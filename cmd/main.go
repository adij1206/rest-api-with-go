package main

import (
	"database/sql"
	"ecommerce/cmd/api"
	"ecommerce/config"
	"ecommerce/db"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {

	// One way to create the mysql configs but it is not good as in the main method itself we are doing everything
	db, err := db.NewMySqlStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":3000", db)

	err1 := server.Run()

	if err1 != nil {
		log.Fatal(err1)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB : Connected Successfully")
}
