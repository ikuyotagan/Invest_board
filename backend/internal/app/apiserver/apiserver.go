package apiserver

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Artemchikus/api/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
)

func Start(config *Config) error {
	URL := config.DatabaseURL
	if config.CustomDatabseURl != "" {
		URL = config.CustomDatabseURl
	}

	db, err := newDB(URL)
	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)
	sessionStore := newSessionStore(config.SessionKey)
	srv := newServer(store, sessionStore)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connection to db is successfull")
	return db, nil
}

func newSessionStore(sessionKey string) *sessions.CookieStore {
	session := sessions.NewCookieStore([]byte(sessionKey))
	session.Options = &sessions.Options{
		Domain:   "localhost",
		Path:     "/",
		MaxAge:   3600 * 8,
		HttpOnly: true,
	}
	return session
}
