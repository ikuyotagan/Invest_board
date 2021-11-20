package apiserver

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Artemchikus/api/internal/app/api/tinkoff"
	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"net/http"
	"time"

	"github.com/Artemchikus/api/internal/app/model"
	"github.com/Artemchikus/api/internal/app/store"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

const (
	sessionName        = "restapi"
	ctxKeyUser  ctxKey = iota
	ctxKeyRequestID
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errNotAutheticated          = errors.New("not authenticated")
	errEmailAlreadyExists       = errors.New("email is already registered")
	errSmallPassword            = errors.New("password needs at least 8 simbols")
	errNoApiKey                 = errors.New("tinkoff api key needed")
)

type ctxKey int8

type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        store,
		sessionStore: sessionStore,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)

	s.router.HandleFunc("/sessions", s.handleOPTIONS()).Methods("OPTIONS")
	s.router.HandleFunc("/users", s.handleOPTIONS()).Methods("OPTIONS")
	s.router.HandleFunc("/logout", s.handleOPTIONS()).Methods("OPTIONS")

	s.router.Use(s.handleCORS)

	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")

	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.authenticateUser)
	private.HandleFunc("/whoami", s.handleWhoami()).Methods("GET")
	private.HandleFunc("/logout", s.handleLogout()).Methods("POST")
	private.HandleFunc("/set_tinkoff", s.handleSetTinkoff()).Methods("POST")

	withTinkoffKey := private.PathPrefix("/tinkoff").Subrouter()
	withTinkoffKey.Use(s.isTinkoffKeyExist)
	withTinkoffKey.HandleFunc("/personal_stocks", s.handleGetPersonalStocks()).Methods("POST")
	withTinkoffKey.HandleFunc("/last_candle", s.handleGetLastCandle()).Methods("POST")
	withTinkoffKey.HandleFunc("/analytics", s.handleGetAnalytics()).Methods("POST")
}

func (s *server) handleOPTIONS() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		if r.Method == http.MethodOptions {
			return
		}
	}
}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

func (s *server) handleCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
		w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
		next.ServeHTTP(w, r)
	})
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remoute_addr": r.RemoteAddr,
			"request_id":   r.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Infof("complited with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Since(start))
	})
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if _, err := s.store.User().FindByEmail(u.Email); err == nil {
			s.error(w, r, http.StatusUnprocessableEntity, errEmailAlreadyExists)
			return
		}

		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, errSmallPassword)
			return
		}

		/*if err := s.store.User().SetTinkoffKey(u); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}*/

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleSessionsCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		client := sdk.NewRestClient(u.TinkoffAPIKey)
		acc, err := client.Accounts(context.WithValue(r.Context(), ctxKeyRequestID, u.ID))
		stocks, err := client.Portfolio(context.WithValue(r.Context(), ctxKeyRequestID, u.ID), acc[0].ID)

		tinkoff.SetData(&stocks.Positions, s.store, u.ID)
		/*
			for _, stock := range stocks.Positions{
				tinkoff.SubscribeCandle(&stock.FIGI)
			}*/

		session, err := s.sessionStore.New(r, sessionName)
		if err != nil {
			return
		}

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = u.ID
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, sessionName)

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			s.error(w, r, http.StatusUnauthorized, errNotAutheticated)
			return
		}

		u, err := s.store.User().Find(id.(int))
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errNotAutheticated)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

func (s *server) isTinkoffKeyExist(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, sessionName)

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			s.error(w, r, http.StatusUnauthorized, errNotAutheticated)
			return
		}

		u, err := s.store.User().Find(id.(int))
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errNotAutheticated)
			return
		}

		if u.TinkoffAPIKey == "" || &u.TinkoffAPIKey == nil {
			s.error(w, r, http.StatusUnauthorized, errNoApiKey)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

func (s *server) handleWhoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*model.User))
	}
}

func (s *server) handleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		session.Options.MaxAge = -1

		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) handleSetTinkoff() http.HandlerFunc {
	type request struct {
		TinkoffAPIKey string `json:"tinkoffapikey"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := r.Context().Value(ctxKeyUser).(*model.User)

		u.TinkoffAPIKey = req.TinkoffAPIKey

		err := s.store.User().SetTinkoffKey(u)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) handleGetPersonalStocks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := r.Context().Value(ctxKeyUser).(*model.User)

		ps, err := s.store.PersonalStock().FindStocksByUserID(u.ID)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, ps)
	}
}

func (s *server) handleGetLastCandle() http.HandlerFunc {
	type request struct {
		Figi string `json:"figi"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		//u := r.Context().Value(ctxKeyUser).(*model.User)

		lc, err := s.store.Candel().FindLastByStockFIGI(req.Figi)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, lc)
	}
}

func (s *server) handleGetAnalytics() http.HandlerFunc {
	type request struct {
		Figi string `json:"figi"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, nil)
	}
}
