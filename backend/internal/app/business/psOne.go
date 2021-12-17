package business

import "github.com/Artemchikus/api/internal/app/model"

func GetPsOne(cost *model.Candel) *model.Coeff {
	var (
		vyr   float32 = 1000000000 // Выручка компании
		kolAc float32 = 1000000    // Число акций
	)
	//n:= cap(cost) // количество свечай нв входе
	//result:=make([]*model.Coeff,0,n)
	//for i:= 0; i< n; i++{
	var result *model.Coeff = &model.Coeff{
		Name:    `P/S`,
		Value:   cost.ClosePrice * kolAc / vyr,
		Time:    cost.Time,
		StockId: cost.StockID}
	//}
	return result
}
