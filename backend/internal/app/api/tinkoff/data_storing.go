package tinkoff

import (
	"github.com/Artemchikus/api/internal/app/model"
	"github.com/Artemchikus/api/internal/app/store"
	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"github.com/pkg/errors"
)

func SetData(stocks *[]sdk.PositionBalance, store store.Store, userId int) error {
	for _, stock := range *stocks {
		stk, err := store.Stock().FindByFIGI(stock.FIGI)
		if err != nil {
			stk = StockConverter(&stock)
			err = store.Stock().Create(stk)
			if err != nil {
				return errors.Errorf("Oh shit, error when setting stock: " + stock.Name)
			}
		}

		_, err = store.PersonalStock().FindByUserIDAndStockID(userId, stk.ID)
		if err != nil {
			err = store.PersonalStock().Create(PersonalStockConverter(&stock, stk, userId))
			if err != nil {
				return errors.Errorf("Oh shit, error when setting personal stock: " + stock.Name)
			}
		}
	}
	return nil
}

func StockConverter(balance *sdk.PositionBalance) *model.Stock {
	return &model.Stock{
		Name: balance.Name,
		FIGI: balance.FIGI,
	}
}

func PersonalStockConverter(persStock *sdk.PositionBalance, stk *model.Stock, uid int) *model.PersonalStock {
	return &model.PersonalStock{
		UserID:         uid,
		StockID:        stk.ID,
		UserStockValue: float32(persStock.Balance),
	}
}
