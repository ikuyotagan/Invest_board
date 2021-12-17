package business

import "github.com/Artemchikus/api/internal/app/model"

func GetPeOne(cost *model.Candel) *model.Coeff {
	var (
		prof  float32 = 1000000000 // Чистая прибыль компании
		kolAc float32 = 1000000    // Число акций
	)
	//n:= cap(cost) // количество свечай нв входе
	/*for i:= 0; i< n; i++ {
		cost = append(cost, &model.Candel{
			ClosePrice: float32((i+1)*100),
		})
	}*/
	//result:=make([]*model.Coeff,0,n)
	//for i:= 0; i< n; i++{
	var result *model.Coeff = &model.Coeff{
		Name:    `P/E`,
		Value:   cost.ClosePrice / (prof / kolAc),
		Time:    cost.Time,
		StockId: cost.StockID}
	//}
	return result
}
