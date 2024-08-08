package main

import (
	"github.com/heldercavalcante/api-bank/internal/configs"
	"github.com/heldercavalcante/api-bank/internal/database"
	"github.com/heldercavalcante/api-bank/internal/http"
)



func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	database.InitDBConnection()
	conn, err := database.GetDB()

	if err != nil {
		panic(err)
	}

	defer conn.Close()
	
	server := http.NewAPIServer(configs.GetServerPort())
	server.Run()
}