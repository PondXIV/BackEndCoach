package service

import (
	"backEndGo/models"
	"backEndGo/repository"
	"fmt"

	"gorm.io/gorm"
)

type ShowDataService interface {
	Login(Email string, Password string, Type int) (*[]models.Coach, *[]models.Customer, error)
	LoginNotType(Email string, Password string) (*[]models.Coach, *[]models.Customer, error)
	// Table Coach
	PrintAllCoach()
	LoginCoach(Email string, Password string) (*[]models.Coach, error)

	// Table Customer
	PrintAllCustomer()
	LoginCustomer(Email string, Password string) (*[]models.Customer, error)

	// Table Buying
	PrintAllBuying()

	// Table Chat
	PrintAllChat()

	// Table Course
	PrintAllCourse()
}
type ShowData struct {
	db *gorm.DB
}

// LoginTwo implements ShowDataService
func (s ShowData) LoginNotType(Email string, Password string) (*[]models.Coach, *[]models.Customer, error) {
	repo := repository.NewUserRepository(s.db)
	coach, cus, err := repo.LoginNotType(Email, Password)
	if err != nil {
		panic(err)
	}
	return coach, cus, nil
}

// Login implements ShowDataService
func (s ShowData) Login(Email string, Password string, Type int) (*[]models.Coach, *[]models.Customer, error) {
	repo := repository.NewUserRepository(s.db)
	coach, cus, err := repo.Login(Email, Password, Type)
	if err != nil {
		panic(err)
	}
	return coach, cus, nil
}

// Table Course
// PrintAllCourse implements ShowDataService
func (s ShowData) PrintAllCourse() {
	repo := repository.NewCourseRepository(s.db)
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
	repo := repository.NewCoachRepository(s.db)
	coachs, err := repo.GetCoachAll()

	if err != nil {
		panic(err)
	}
	for _, v := range *coachs {
		fmt.Printf("%v\n", v)
	}
}

// LoginCoach implements ShowDataService
func (s ShowData) LoginCoach(Email string, Password string) (*[]models.Coach, error) {
	repo := repository.NewCoachRepository(s.db)
	coa, err := repo.LoginCoach(Email, Password)

	if err != nil {
		panic(err)
	}
	return coa, nil
}

// Table Customer
// PrintAllCustomer implements ShowDataService
func (s ShowData) PrintAllCustomer() {
	repo := repository.NewCustomerRepository(s.db)
	cus, err := repo.GetCustomerAll()

	if err != nil {
		panic(err)
	}
	for _, v := range *cus {
		fmt.Printf("%v\n", v)
	}
}

// LoginCustomer implements ShowDataService
func (s ShowData) LoginCustomer(Email string, Password string) (*[]models.Customer, error) {
	repo := repository.NewCustomerRepository(s.db)
	cus, err := repo.LoginCustomer(Email, Password)

	if err != nil {
		panic(err)
	}
	return cus, nil
}

// Table Buying
// PrintAllBuying implements ShowDataService
func (s ShowData) PrintAllBuying() {
	repo := repository.NewBuyingRepository(s.db)
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
	repo := repository.NewChatRepository(s.db)
	chats, err := repo.GetChatAll()

	if err != nil {
		panic(err)
	}
	for _, v := range *chats {
		fmt.Printf("%v\n", v)
	}
}
func NewShowDataService(gormdb *gorm.DB) ShowDataService {
	return ShowData{db: gormdb}
}
