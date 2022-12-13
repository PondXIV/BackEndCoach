package main

import (
	"backEndGo/service"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// SERVER ******
	StartServer()

	//เข้ารหัส DB
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	//fmt.Println(viper.GetString("mysql.dsn"))
	// Connect to database

	dsn := viper.GetString("mysql.dsn")
	dialector := mysql.Open(dsn)

	db, err := gorm.Open(dialector)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connection Success %v\n", db)

	// getAll Buying tast
	// buying := service.NewShowDataService(db)
	// buying.PrintAllBuying()

	// getAll Chat tast
	// chats := service.NewShowDataService(db)
	// chats.PrintAllChat()

	// getAll Course tast
	// courses := service.NewShowDataService(db)
	// courses.PrintAllCourse()

	// getAll Customer tast
	// cus := service.NewShowDataService(db)
	// cus.PrintAllCustomer()

	// getAll Coach tast
	coa := service.NewShowDataService(db)
	// coa.PrintAllCoach()

	// fmt.Printf("\n")
	//LoginCustomer
	// result, err := cus.LoginCustomer("benzzaa@gmail.com", "1234")
	// if err != nil {
	// 	panic(err)
	// }
	// for _, v := range *result {
	// 	fmt.Printf("\n%v\n", v)
	// }

	//LoginCoach
	// result, err := coa.LoginCoach("babe@gmail.com", "1234")
	// if err != nil {
	// 	panic(err)
	// }
	// for _, v := range *result {
	// 	fmt.Printf("\n%v\n", v)
	// }

	/// Login
	//Coach
	coach, cus, err := coa.Login("babe@gmail.com", "1234", 0)
	if err != nil {
		panic(err)
	}
	for _, v := range *coach {
		fmt.Printf("\n%v\n", v)
	}
	for _, v := range *cus {
		fmt.Printf("\n%v\n", v)
	}
	//Customer
	coach, cus, err = coa.Login("Tpangpond@gmail.com", "15978", 1)
	if err != nil {
		panic(err)
	}
	for _, v := range *coach {
		fmt.Printf("\n%v\n", v)
	}
	for _, v := range *cus {
		fmt.Printf("\n%v\n", v)
	}
	///
	coach, cus, err = coa.Login("asdasdsad", "1213123", 1)
	if err != nil {
		panic(err)
	}
	for _, v := range *coach {
		fmt.Printf("\n%v\n", v)
	}
	for _, v := range *cus {
		fmt.Printf("\n%v\n", v)
	}
}
