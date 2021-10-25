package tinkoff

import (
	"context"
	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"log"
	"time"
)

//TODO:: Добавить коммент к каждой ручке

func GetAccounts(client *sdk.RestClient) []sdk.Account {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Получение всех брокерских счетов")
	accounts, err := client.Accounts(ctx)
	if err != nil {
		log.Fatalln(errorHandle(err))
	}
	log.Printf("%+v\n", accounts)
	return accounts
}

func GetCurrencies(client *sdk.RestClient) []sdk.Instrument {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Получение валютных инструментов")
	// Например: USD000UTSTOM - USD, EUR_RUB__TOM - EUR
	currencies, err := client.Currencies(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", currencies)
	return currencies
}

func GetETFs(client *sdk.RestClient) []sdk.Instrument {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Получение фондовых инструментов")
	// Например: FXMM - Казначейские облигации США, FXGD - золото
	etfs, err := client.ETFs(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", etfs)
	return etfs
}

func GetBonds(client *sdk.RestClient) []sdk.Instrument {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Получение облигационных инструментов")
	// Например: SU24019RMFS0 - ОФЗ 24019
	bonds, err := client.Bonds(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", bonds)
	return bonds
}

func GetStocks(client *sdk.RestClient) []sdk.Instrument {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Получение акционных инструментов")
	// Например: SBUX - Starbucks Corporation
	stocks, err := client.Stocks(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", stocks)
	return stocks
}

func GetByTicker(client *sdk.RestClient, ticker string) []sdk.Instrument {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Получение инструменов по тикеру TCS")
	// Получение инструмента по тикеру, возвращает массив инструментов потому что тикер уникален только в рамках одной биржи
	// но может совпадать на разных биржах у разных кампаний
	// Например: https://www.moex.com/ru/issue.aspx?code=FIVE и https://www.nasdaq.com/market-activity/stocks/FIVE
	// В этом примере получить нужную компанию можно проверив поле Currency
	instruments, err := client.InstrumentByTicker(ctx, ticker)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", instruments)
	return instruments
}

func GetByFigi(client *sdk.RestClient, figi string) sdk.Instrument {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Получение инструмента по FIGI BBG005DXJS36 (TCS)")
	// Получение инструмента по FIGI(https://en.wikipedia.org/wiki/Financial_Instrument_Global_Identifier)
	// Узнать FIGI нужного инструмента можно методами указанными выше
	// Например: BBG000B9XRY4 - Apple, BBG005DXJS36 - Tinkoff
	instrument, err := client.InstrumentByFIGI(ctx, figi)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", instrument)
	return instrument
}

func GetOperationsByDate(client *sdk.RestClient, startDate *time.Time, endDate *time.Time, figi string) []sdk.Operation {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Получение списка операций для счета по-умолчанию за последнюю неделю по инструменту(FIGI) BBG000BJSBJ0")
	// Получение списка операций за период по конкретному инструменту(FIGI)
	// Например: ниже запрашиваются операции за последнюю неделю по инструменту NEE
	operations, err := client.Operations(ctx, sdk.DefaultAccount, *startDate, *endDate, figi)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", operations)
	return operations
}

func GetPositionsPortfolio(client *sdk.RestClient) []sdk.PositionBalance {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Получение списка НЕ валютных активов портфеля для счета по-умолчанию")
	positions, err := client.PositionsPortfolio(ctx, sdk.DefaultAccount)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", positions)
	return positions
}

func GetCurrenciesPortfolio(client *sdk.RestClient) []sdk.CurrencyBalance {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Получение списка валютных активов портфеля для счета по-умолчанию")
	positionCurrencies, err := client.CurrenciesPortfolio(ctx, sdk.DefaultAccount)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", positionCurrencies)
	return positionCurrencies
}

func GetPortfolio(client *sdk.RestClient) sdk.Portfolio {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Получение списка валютных и НЕ валютных активов портфеля для счета по-умолчанию")
	// Метод является совмещеним PositionsPortfolio и CurrenciesPortfolio
	portfolio, err := client.Portfolio(ctx, sdk.DefaultAccount)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", portfolio)
	return portfolio
}

func GetOrders(client *sdk.RestClient) []sdk.Order {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Получение списка выставленных заявок(ордеров) для счета по-умолчанию")
	orders, err := client.Orders(ctx, sdk.DefaultAccount)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", orders)
	return orders
}

func GetCandles(client *sdk.RestClient, startDate *time.Time, endDate *time.Time, figi string) []sdk.Candle {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Получение часовых свечей за последние 24 часа по инструменту BBG005DXJS36 (TCS)")
	// Получение свечей(ордеров)
	// Внимание! Действуют ограничения на промежуток и доступный размер свечей за него
	// Интервал свечи и допустимый промежуток запроса:
	// - 1min [1 minute, 1 day]
	// - 2min [2 minutes, 1 day]
	// - 3min [3 minutes, 1 day]
	// - 5min [5 minutes, 1 day]
	// - 10min [10 minutes, 1 day]
	// - 15min [15 minutes, 1 day]
	// - 30min [30 minutes, 1 day]
	// - hour [1 hour, 7 days]
	// - day [1 day, 1 year]
	// - week [7 days, 2 years]
	// - month [1 month, 10 years]
	// Например получение часовых свечей за последние 24 часа по инструменту BBG005DXJS36 (TCS)
	candles, err := client.Candles(ctx, *startDate, *endDate, sdk.CandleInterval1Hour, figi)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", candles)
	return candles
}

func GetOrderbook(client *sdk.RestClient, depth int, figi string) sdk.RestOrderBook {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Получение ордербука(он же стакан) глубиной 10 по инструменту BBG005DXJS36")
	// Получение ордербука(он же стакан) по инструменту
	orderbook, err := client.Orderbook(ctx, depth, figi)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", orderbook)
	return orderbook
}

func SetMarketOrder(client *sdk.RestClient, figi string, lots int, operationType sdk.OperationType) sdk.PlacedOrder {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Выставление рыночной заявки для счета по-умолчанию на покупку ОДНОЙ акции BBG005DXJS36 (TCS)")
	// Выставление рыночной заявки для счета по-умолчанию
	// В примере ниже выставляется заявка на покупку ОДНОЙ акции BBG005DXJS36 (TCS)
	placedOrder, err := client.MarketOrder(ctx, sdk.DefaultAccount, figi, lots, operationType)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", placedOrder)
	return placedOrder
}

func SetLimitOrder(
	client *sdk.RestClient,
	figi string,
	lots int,
	operationType sdk.OperationType,
	price float64) sdk.PlacedOrder {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Выставление лимитной заявки для счета по-умолчанию на покупку ОДНОЙ акции BBG005DXJS36 (TCS) по цене не выше 20$")
	// Выставление лимитной заявки для счета по-умолчанию
	// В примере ниже выставляется заявка на покупку ОДНОЙ акции BBG005DXJS36 (TCS) по цене не выше 20$
	placedOrder, err := client.LimitOrder(ctx, sdk.DefaultAccount, figi, lots, operationType, price)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", placedOrder)
	return placedOrder
}

func OrderCancel(client *sdk.RestClient, placedOrder *sdk.PlacedOrder) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Printf("Отмена ранее выставленной заявки для счета по-умолчанию. %+v\n", placedOrder)
	// Отмена ранее выставленной заявки для счета по-умолчанию.
	// ID заявки возвращается в структуре PlacedLimitOrder в поле ID в запросе выставления заявки client.LimitOrder
	// или в структуре Order в поле ID в запросе получения заявок client.Orders
	err := client.OrderCancel(ctx, sdk.DefaultAccount, placedOrder.ID)
	if err != nil {
		log.Fatalln(err)
	}
}
