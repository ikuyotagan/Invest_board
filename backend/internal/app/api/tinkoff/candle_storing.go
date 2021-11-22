package tinkoff

import (
	"context"
	"github.com/Artemchikus/api/internal/app/model"
	"github.com/Artemchikus/api/internal/app/store"
	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"github.com/pkg/errors"
	"log"
	"time"
)

func CandleStoring(store store.Store) {
	go func() {
		client := sdk.NewRestClient(*token)
		ctx := context.WithValue(context.Background(), "Candle storing", "Fuck")
		stocks, err := store.Stock().GetAll()
		if err != nil {
			log.Println(errors.Errorf("Oh shit, ", err))
		}
		for _, stock := range stocks {
			lastCandle, _ := store.Candel().FindLastByStockID(stock.ID)
			candles := make([]sdk.Candle, 0)
			if lastCandle != nil && time.Since(lastCandle.Time) > time.Hour*24 {
				candles, err = client.Candles(ctx, lastCandle.Time, time.Now(), sdk.CandleInterval1Day, stock.FIGI)
			} else {
				lastCandl := time.Unix(time.Now().Unix() - 31556926, 0)
				candles, err = client.Candles(ctx, lastCandl, time.Now(), sdk.CandleInterval1Day, stock.FIGI)
			}
			if err != nil {
				log.Println(errors.Errorf("Oh shit, ", err))
			}
			for _, candle := range candles {
				err := store.Candel().Create(CandleConverter(&candle, stock.ID))
				if err != nil {
					log.Fatal(errors.Errorf("Oh shit, ", err))
				}
			}
		}
		time.Sleep(time.Hour * 24)
		for {
			stocks, err = store.Stock().GetAll()
			if err != nil {
				log.Fatal(errors.Errorf("Oh shit, ", err))
			}
			for _, stock := range stocks {
				candles, err := client.Candles(ctx, time.Unix(time.Now().Unix()-int64(time.Hour.Seconds()), 0), time.Now(), sdk.CandleInterval1Day, stock.FIGI)
				if err != nil {
					log.Fatal(errors.Errorf("Oh shit, ", err))
				}
				for _, candle := range candles {
					err := store.Candel().Create(CandleConverter(&candle, stock.ID))
					if err != nil {
						log.Fatal(errors.Errorf("Oh shit, ", err))
					}
				}
			}
			time.Sleep(time.Hour * 24)
		}
	}()
}

func CandleConverter(candle *sdk.Candle, stockId int) *model.Candel {
	return &model.Candel{
		Time:          candle.TS,
		OpenPrice:     float32(candle.OpenPrice),
		ClosePrice:    float32(candle.ClosePrice),
		HighestPrice:  float32(candle.HighPrice),
		LowestPrice:   float32(candle.LowPrice),
		TradingVolume: float32(candle.Volume),
		StockID:       stockId,
	}
}
