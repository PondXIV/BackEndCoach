package main

func main() {

	//เข้ารหัส DB
	// viper.SetConfigName("config")
	// viper.AddConfigPath(".")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	panic(err)
	// }
	// //fmt.Println(viper.GetString("mysql.dsn"))
	// // Connect to database

	// dsn := viper.GetString("mysql.dsn")
	// dialector := mysql.Open(dsn)

	// db, err := gorm.Open(dialector)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Connection Success %v\n", db)

	// SERVER ******
	StartServer()

}
