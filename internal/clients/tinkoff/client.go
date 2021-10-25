package tinkoff

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
)

//TODO:: Выводить токены из бд
var token = flag.String("token", "", "your token")

func newStreamingClient() (*sdk.StreamingClient, *log.Logger) {
	rand.Seed(time.Now().UnixNano()) // инициируем Seed рандома для функции requestID
	flag.Parse()
	logger := log.New(os.Stdout, "[invest-openapi-go-sdk]", log.LstdFlags)

	client, err := sdk.NewStreamingClient(logger, *token)
	if err != nil {
		log.Fatalln(err)
	}
	//stream(client, logger)
	return client, logger
}

func closeStream(client *sdk.StreamingClient) error {
	return client.Close()
}

func NewRestClient() *sdk.RestClient {
	client := sdk.NewRestClient(*token)
	//rest(client)
	return client
}

/*func stream(client *sdk.StreamingClient, logger *log.Logger) {

	// Запускаем цикл обработки входящих событий. Запускаем асинхронно
	// Сюда будут приходить сообщения по подпискам после вызова соответствующих методов
	// SubscribeInstrumentInfo, SubscribeCandle, SubscribeOrderbook
	go func() {
		err := client.RunReadLoop(func(event interface{}) error {
			logger.Printf("Got event %+v", event)
			return nil
		})
		if err != nil {
			log.Fatalln(err)
		}
	}()
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Генерируем уникальный ID для запроса
func requestID() string {
	b := make([]rune, 12)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}*/

func errorHandle(err error) error {
	if err == nil {
		return nil
	}

	if tradingErr, ok := err.(sdk.TradingError); ok {
		if tradingErr.InvalidTokenSpace() {
			tradingErr.Hint = "Do you use sandbox token in production environment or vise verse?"
			return tradingErr
		}
	}

	return err
}
