package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sandeepgoutele/golang-coding/airline_checkin/airline"
	log "github.com/sirupsen/logrus"
)

func book(dbInstance *airline.DBSingleton, user *airline.User) (*airline.Seat, error) {
	txn, err := dbInstance.DbObj.Begin()
	if err != nil {
		log.Fatal(err)
	}

	row := txn.QueryRow(`SELECT id, name, trip_id, user_id 
		FROM seats 
		WHERE trip_id = 1 and user_id is null 
		ORDER BY id 
		LIMIT 1 
		FOR UPDATE`)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var seat airline.Seat
	row.Scan(&seat.ID, &seat.Name, &seat.TripID, &seat.UserID)

	_, err = txn.Exec("UPDATE seats set user_id = ? WHERE id = ?", user.ID, seat.ID)
	if err != nil {
		return nil, err
	}

	err = txn.Commit()

	if err != nil {
		return nil, err
	}
	return &seat, nil
}

func main() {
	now := time.Now()
	log.SetLevel(log.DebugLevel)
	dbInstance := airline.GetInstance()
	defer dbInstance.DbObj.Close()
	users := airline.GetUsers(dbInstance)
	log.Debugf("Simulating airline booking for %d users", len(users))
	var wg sync.WaitGroup
	wg.Add(len(users))
	for idx := range users {
		go func(user *airline.User) {
			seat, err := book(dbInstance, user)
			userName := user.FirstName + " " + user.LastName
			if err != nil {
				log.Errorf("We couldn't assign seat %s to user %s", seat.Name, userName)
				log.Error(err)
			} else {
				log.Infof("%s was assigned the seat %s", userName, seat.Name)
			}
			defer wg.Done()
		}(&users[idx])
	}
	wg.Wait()

	seats := airline.GetSeats(dbInstance)
	seatPlot := make([]string, 6)
	for _, seat := range seats {
		label := []rune(strings.Split(seat.Name, "-")[1])
		occupied := "."
		if seat.UserID.Valid {
			occupied = "x"
		}
		seatPlot[label[0]-'A'] = fmt.Sprintf("%s %s ", seatPlot[label[0]-'A'], occupied)
	}

	for idx, plot := range seatPlot {
		if idx%3 == 0 {
			log.Println()
		}
		log.Println(plot)
	}

	fmt.Println()
	airline.ResetSeats(dbInstance)

	fmt.Println()
	newNow := time.Now()
	log.Println("Time taken for full execution: ", newNow.Sub(now))
}
