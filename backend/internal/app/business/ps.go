package business

import "github.com/Artemchikus/api/internal/app/model"

func GetPs(cost []*model.Candel) []*model.Coeff {
	var (
		vyr   float32 = 1000000000 // Выручка компании
		kolAc float32 = 1000000    // Число акций
	)
	n := cap(cost) // количество свечай нв входе
	result := make([]*model.Coeff, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, &model.Coeff{
			Name:    `P/S`,
			Value:   cost[i].ClosePrice * kolAc / vyr,
			Time:    cost[i].Time,
			StockId: cost[i].StockID})
	}
	return result
}
