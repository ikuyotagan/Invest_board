package tinkoff

import (
	"context"
	"github.com/Artemchikus/api/internal/app/model"
	"github.com/Artemchikus/api/internal/app/store"
	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"
)

func CandleStoring(store store.Store) {
	go func() {
		client := sdk.NewRestClient(*token)
		ctx := context.WithValue(context.Background(), "Candle storing", "Fuck")
		stocks, err := store.Stock().GetAll()
		if err != nil {
			stocksRaw, err := fuckUp(&ctx, client)
			if err != nil {
				log.Println(http.StatusInternalServerError, err)
				return
			}
			for _, stockRaw := range stocksRaw.Positions {
				stocks = append(stocks, StockConverter(&stockRaw))
			}
			for _, stock := range stocks {
				_, err := store.Stock().FindByFIGI(stock.FIGI)
				if err != nil {
					err = store.Stock().Create(stock)
					if err != nil {
						log.Println(http.StatusInternalServerError, err)
						return
					}
				}
			}
		}
		for _, stock := range stocks {
			lastCandle, _ := store.Candel().FindLastByStockID(stock.ID)

			candles := make([]sdk.Candle, 0)
			if lastCandle != nil && time.Since(lastCandle.Time) > time.Hour*24 {
				candles, err = client.Candles(ctx, lastCandle.Time, time.Now(), sdk.CandleInterval1Day, stock.FIGI)
			} else if lastCandle == nil {
				lastCandl := time.Unix(time.Now().Unix()-31556926, 0)
				candles, err = client.Candles(ctx, lastCandl, time.Now(), sdk.CandleInterval1Day, stock.FIGI)
			}
			if err != nil {
				log.Println(errors.Errorf("Oh shit, ", err))
			}
			for _, candle := range candles {
				err := store.Candel().Create(CandleConverter(&candle, stock.ID))
				if err != nil {
					log.Println(errors.Errorf("Oh shit, ", err))
				}
			}
			log.Printf("Загружено %d свечей по акции \"%s\"", len(candles), stock.Name)
		}
		time.Sleep(time.Hour * 24)
		for {
			stocks, err = store.Stock().GetAll()
			if err != nil {
				log.Println(errors.Errorf("Oh shit, ", err))
			}
			for _, stock := range stocks {
				candles, err := client.Candles(ctx, time.Unix(time.Now().Unix()-int64(time.Hour.Seconds())*25, 0), time.Now(), sdk.CandleInterval1Day, stock.FIGI)
				if err != nil {
					log.Println(errors.Errorf("Oh shit, ", err))
				}
				for _, candle := range candles {
					err := store.Candel().Create(CandleConverter(&candle, stock.ID))
					if err != nil {
						log.Println(errors.Errorf("Oh shit, ", err))
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

func fuckUp(ctx *context.Context, client *sdk.RestClient) (*sdk.Portfolio, error) {
	acc, err := client.Accounts(*ctx)
	if err != nil {
		log.Println(http.StatusInternalServerError, err)
		return nil, err
	}

	stocks, err := client.Portfolio(*ctx, acc[0].ID)
	if err != nil {
		log.Println(http.StatusInternalServerError, err)
		return nil, err
	}
	return &stocks, nil
}
