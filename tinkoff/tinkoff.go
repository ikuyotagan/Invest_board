package tinkoff

import (
	"context"
	"flag"
	"fmt"
	"jwt/database"
	"log"
	"time"

	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
)

var token = flag.String("token",
	"t.zi-GzKc-R6kOQF3MqGFBCSCSKX-B4aed_OwhJbjqQZCL9va5cPlmiMRNRfCDMW2_cB7RhSylrxHLvX_AOlW10g",
	"your token")

// func main() {
// 	rand.Seed(time.Now().UnixNano())
// 	flag.Parse()
// 	rest()

// 	// stream()
// }

// func stream() {
// 	logger := log.New(os.Stdout, "[invest-openapi-go-sdk]", log.LstdFlags)

// 	client, err := sdk.NewStreamingClient(logger, *token)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer client.Close()

// 	go func() {
// 		err := client.RunReadLoop(func(event interface{}) error {
// 			logger.Printf("Got event: %v", event)
// 			return nil
// 		})
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 	}()

// 	err = client.SubscribeInstrumentInfo("BBG005DXJS36", requestID())
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	time.Sleep(10 * time.Second)
// }

// type Request struct {
// 	FIGI string `json:"figi"`
// 	Name string `json:"name"`
// }

type Candle struct {
	Id        string    `json:"-"`
	HighPrice float64   `json:"h"`
	LowPrice  float64   `json:"l"`
	TS        time.Time `json:"time"`
}

func Rest() {
	client := sdk.NewRestClient(*token)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// portfolio, err := client.Portfolio(ctx, sdk.DefaultAccount)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// var r Request
	// for i := range portfolio.Positions {
	// 	r = Request{
	// 		Name: portfolio.Positions[i].Name,
	// 		FIGI: portfolio.Positions[i].FIGI,
	// 	}
	// 	log.Printf("Name - %s; FIGI - %s \n", r.Name, r.FIGI)
	// }

	// ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	var c Candle
	candles, err := client.Candles(ctx, time.Now().AddDate(-1, 0, 0), time.Now(), sdk.CandleInterval1Day, "BBG005DXJS36")
	if err != nil {
		log.Fatalln("no candels")
	}
	for i := range candles {
		c = Candle{
			HighPrice: candles[i].HighPrice,
			LowPrice:  candles[i].LowPrice,
			TS:        candles[i].TS,
		}
		if err := database.DB.QueryRow("SELECT id FROM candels WHERE time = $1", 
		candles[i].TS).Scan(&c.Id); err != nil {
			fmt.Println(err)
			if err := database.DB.QueryRow(
				"INSERT INTO candels (price, time, name) VALUES ($1, $2, $3) RETURNING id",
				((c.HighPrice+c.LowPrice)/2),
				c.TS,
				"tinkoff",
			).Scan(&c.Id); err != nil {
				log.Fatalln("not in db")
			}
		}
	}
}

// var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// func requestID() string {
// 	b := make([]rune, 12)
// 	for i := range b {
// 		b[i] = letterRunes[rand.Intn(len(letterRunes))]
// 	}
// 	return string(b)
// }
