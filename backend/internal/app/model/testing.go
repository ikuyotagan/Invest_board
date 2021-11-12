package model

import (
	"testing"
	"time"
)

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@gmail.com",
		Password: "password",
	}
}

func TinkoffAPITestUser(t *testing.T) *User {
	return &User{
		Email:         "user@gmail.com",
		Password:      "password",
		TinkoffAPIKey: "t.zi-GzKc-R6kOQF3MqGFBCSCSKX-B4aed_OwhJbjqQZCL9va5cPlmiMRNRfCDMW2_cB7RhSylrxHLvX_AOlW10g",
	}
}

func TestCandel1(t *testing.T) *Candel {
	date := "2019-08-07T15:35:00Z"
	formatDate, _ := time.Parse(time.RFC3339Nano, date)
	location, _ := time.LoadLocation("Europe/Moscow")
	return &Candel{
		Time:          formatDate.In(location),
		OpenPrice:     1.1,
		ClosePrice:    1.1,
		HighestPrice:  1.1,
		LowestPrice:   1.1,
		TradingVolume: 1.1,
		StockID:       1,
	}
}

func TestCandel2(t *testing.T) *Candel {
	date := "2019-08-08T15:35:00Z"
	formatDate, _ := time.Parse(time.RFC3339Nano, date)
	location, _ := time.LoadLocation("Europe/Moscow")
	return &Candel{
		Time:          formatDate.In(location),
		OpenPrice:     2.2,
		ClosePrice:    2.2,
		HighestPrice:  2.2,
		LowestPrice:   2.2,
		TradingVolume: 2.2,
		StockID:       1,
	}
}

func TestStock(t *testing.T) *Stock {
	return &Stock{
		Name: "HP",
		FIGI: "BBG000BLNNH6",
	}
}

func TestPersonalStock1(t *testing.T) *PersonalStock {
	return &PersonalStock{
		UserID:         1,
		StockID:        1,
		UserStockValue: 1.1,
	}
}

func TestPersonalStock2(t *testing.T) *PersonalStock {
	return &PersonalStock{
		UserID:         1,
		StockID:        2,
		UserStockValue: 2.2,
	}
}
