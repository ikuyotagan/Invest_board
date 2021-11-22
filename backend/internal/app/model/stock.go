package model

type Stock struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	FIGI          string `json:"figi"`
}
