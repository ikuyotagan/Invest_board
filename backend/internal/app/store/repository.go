package store

import (
	"time"

	"github.com/Artemchikus/api/internal/app/model"
)

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	SetTinkoffKey(*model.User) error
}

type StockRepository interface {
	Create(*model.Stock) error
	Find(int) (*model.Stock, error)
	FindByName(string) (*model.Stock, error)
	FindByFIGI(string) (*model.Stock, error)
}

type CandelRepository interface {
	Create(*model.Candel) error
	Find(int) (*model.Candel, error)
	FindByTimeAndStockID(time.Time, int) (*model.Candel, error)
	FindbyPeriodAndStokID(time.Time, time.Time, int) ([]*model.Candel, error)
}

type PersonalStockRepository interface {
	Create(*model.PersonalStock) error
	Find(int) (*model.PersonalStock, error)
	FindStocksByUserID(int) ([]*model.PersonalStock, error)
	FindByUserIDAndStockID(int, int) (*model.PersonalStock, error)
}
