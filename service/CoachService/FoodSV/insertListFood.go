package foodsv

type InsertListFoodDataService interface {
}
type InsertListFoodData struct {
}

func NewInsertListFoodDataService() InsertListFoodDataService {
	return InsertListFoodData{}
}
