package model

type Stock struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	FIGI          string  `json:"figi"`
	// Amount        int     `json:"amount"`
	// AnnualRevemue float32 `json:"annual-revenue"`
	// Capital       string  `json:"capital"`
}
