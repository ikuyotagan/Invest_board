package business

import "github.com/Artemchikus/api/internal/app/model"

func GetPbv(cost []*model.Candel) []*model.Coeff {
	var (
		kap   float32 = 1000000000 // Капитал компании
		kolAc float32 = 1000000    // Число акций
	)
	n := cap(cost) // количество свечай нв входе
	/*for i:= 0; i< n; i++ {
		cost = append(cost, &model.Candel{
			ClosePrice: float32((i+1)*100),
		})
	}*/
	result := make([]*model.Coeff, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, &model.Coeff{
			Name:    `P/BV`,
			Value:   cost[i].ClosePrice * kolAc / kap,
			Time:    cost[i].Time,
			StockId: cost[i].StockID})
	}
	return result
}
