package airline

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Address   string
	ID        int
}

type Seat struct {
	ID     int
	Name   string
	TripID int
	UserID int
}

type DBSingleton struct {
	DbObj *sql.DB
}

var instance *DBSingleton
var once sync.Once

func GetInstance() *DBSingleton {
	once.Do(func() {
		// Connect to the database
		dsn := "sandeep:sandeep@tcp(127.0.0.1:3306)/practise"
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal(err)
		}

		// Test the connection
		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connected to the database successfully!")
		instance = &DBSingleton{DbObj: db}
	})

	return instance
}

func GetUsers(dbInstance *DBSingleton) []User {
	query := `SELECT * from users`
	rows, err := dbInstance.DbObj.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.Address, &user.ID)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	/*
		// Print the users
		for _, user := range users {
			fmt.Printf("ID: %d, First Name: %s, Last Name: %s, Email: %s, Phone: %s, Address: %s\n",
				user.ID, user.FirstName, user.LastName, user.Email, user.Phone, user.Address)
		}
	*/

	return users
}
