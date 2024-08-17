package main

import (
	"go-api/cmd/api"
	"go-api/cmd/config"
	"go-api/cmd/db"

	"github.com/go-sql-driver/mysql"
)

func main() {

	database, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		panic(err)
	}

	db.InitStorage(database)

	api.HandleOpenConnections()

	server := api.NewAPIServer(database, ":8080")

	if error := server.Run(); error != nil {
		panic(error)
	}
}
