package main

import (
	"fmt"
	"log"

	"github.com/brianvoe/gofakeit"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sandeepgoutele/golang-coding/airline_checkin/airline"
)

func main() {
	dbInstance := airline.GetInstance()
	fmt.Println("Cleaned up users table before populating new data!")
	query := `TRUNCATE TABLE users`
	_, err := dbInstance.DbObj.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	// Seed the random generator
	gofakeit.Seed(0)
	for idx := 0; idx < 120; idx++ {
		// Generate a random user
		user := airline.User{
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Email:     gofakeit.Email(),
			Phone:     gofakeit.Phone(),
			Address:   gofakeit.Address().Address,
		}

		// Insert the user into the user table
		query = `INSERT INTO users (first_name, last_name, email, phone, address) VALUES (?, ?, ?, ?, ?)`
		_, err = dbInstance.DbObj.Exec(query, user.FirstName, user.LastName, user.Email, user.Phone, user.Address)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("User data inserted successfully!")

	fmt.Println("Cleaned up seats table before populating new data!")
	query = `TRUNCATE TABLE seats`
	_, err = dbInstance.DbObj.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	// Populate seats table
	seatNum := 1
	seatLabel := 'A'
	for idx := 0; idx < 120; idx++ {
		// Insert the user into the user table
		query = `INSERT INTO seats (name, trip_id) VALUES (?, ?)`
		if seatLabel == 'G' {
			seatNum += 1
			seatLabel = 'A'
		}
		_, err = dbInstance.DbObj.Exec(query, fmt.Sprintf("%d-%c", seatNum, seatLabel), 1)
		if err != nil {
			log.Fatal(err)
		}
		seatLabel += 1
	}
	fmt.Println("Updated seats data!")
}
