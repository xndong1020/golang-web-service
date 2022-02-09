package main

import (
	"fmt"

	db "example.com/api/db"
	models "example.com/api/models"
	utils "example.com/api/utils"

	_ "github.com/lib/pq"
)

func main() {

	// get db connection
	connection := db.GetDbConnection()

	rows, err := connection.Query(`select * from samurais`)
	utils.CheckError(err)

	// create an empty Samurai array
	samurais := make([]models.Samurai, 0)

	for rows.Next() {
		var samurai models.Samurai
		rows.Scan(&samurai.Id, &samurai.Name)
		samurais = append(samurais, samurai)
	}

	defer rows.Close()

	fmt.Println("Connected!", samurais)
}
