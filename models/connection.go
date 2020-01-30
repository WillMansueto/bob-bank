package models

import(
	"database/sql"
	"fmt"
	"log"

	"bob-bank/config"

	_ "github.com/lib/pq"
)

var configs = config.LoadConfigs()

func Connect() *sql.DB {
	URL := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", configs.Database.User, 
	configs.Database.Pass, configs.Database.Name, "disable")
	db, err := sql.Open("postgres", URL)
	if err != nil{
		log.Fatal(err)
		return nil
	}
	return db
}

func TestConnection() {
	con := Connect()
	defer con.Close()
	err := con.Ping()
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	}
	fmt.Println("Database connected!")
}