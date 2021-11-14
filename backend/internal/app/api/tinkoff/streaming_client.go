package tinkoff

import (
	"flag"
	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"log"
	"math/rand"
	"os"
)

var (
	client *sdk.StreamingClient
	token  = flag.String("token", "", "default token")
)

func NewClient() {
	logger := log.New(os.Stdout, "[invest-openapi-go-sdk]", log.LstdFlags)
	var err error
	client, err = sdk.NewStreamingClient(logger, *token)
	if err != nil {
		logger.Fatalln(err)
	}
	defer client.Close()
}

func StreamingListener() {
	go func() {
		err := client.RunReadLoop(func(event interface{}) error {
			log.Printf("Got event %+v", event)

			return nil
		})
		if err != nil {
			log.Fatalln(err)
		}
	}()
}

func SubscribeCandle(figi *string) {
	err := client.SubscribeCandle(*figi, sdk.CandleInterval1Min, requestID())
	if err != nil {
		log.Fatalln(err)
	}
}

func CloseClient() {
	client.Close()
}

//func StreamingManager ()

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Генерируем уникальный ID для запроса
func requestID() string {
	b := make([]rune, 12)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
