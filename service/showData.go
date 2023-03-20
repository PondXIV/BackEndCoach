package service

import (
	// "backEndGo/models"
	"backEndGo/repository"
	"fmt"
)

type ShowDataService interface {

	// Table Coach
	PrintAllCoach()

	// Table Customer
	PrintAllCustomer()

	// Table Buying
	PrintAllBuying()

	// Table Chat
	PrintAllChat()

	// Table Course
	PrintAllCourse()
}
type ShowData struct {
}

// Table Course
// PrintAllCourse implements ShowDataService
func (s ShowData) PrintAllCourse() {
	repo := repository.NewCourseRepository()
	courses, err := repo.GetCourseAll()

	if err != nil {
		panic(err)
	}
	for _, v := range *courses {
		fmt.Printf("%v\n", v)
	}
}

// Table Coach
// PrintAllCoach implements ShowDataService
func (s ShowData) PrintAllCoach() {
	repo := repository.NewCoachRepository()
	coachs, err := repo.GetCoachAll()

	if err != nil {
		panic(err)
	}
	for _, v := range *coachs {
		fmt.Printf("%v\n", v)
	}
}

// Table Customer
// PrintAllCustomer implements ShowDataService
func (s ShowData) PrintAllCustomer() {
	repo := repository.NewCustomerRepository()
	cus, err := repo.GetCustomerAll()

	if err != nil {
		panic(err)
	}
	for _, v := range *cus {
		fmt.Printf("%v\n", v)
	}
}

// Table Buying
// PrintAllBuying implements ShowDataService
func (s ShowData) PrintAllBuying() {
	repo := repository.NewBuyingRepository()
	buying, err := repo.GetBuyingrAll()

	if err != nil {
		panic(err)
	}
	for _, v := range *buying {
		fmt.Printf("%v\n", v)
	}
}

// Table Chat
// PrintAllChat implements ShowDataService
func (s ShowData) PrintAllChat() {
	repo := repository.NewChatRepository()
	chats, err := repo.GetChatAll()

	if err != nil {
		panic(err)
	}
	for _, v := range *chats {
		fmt.Printf("%v\n", v)
	}
}
func NewShowDataService() ShowDataService {
	return ShowData{}
}
