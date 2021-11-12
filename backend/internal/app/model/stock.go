package model

type Stock struct {
	ID            int    `json:"-"`
	Name          string `json:"name"`
	FIGI          string `json:"figi"`
}
