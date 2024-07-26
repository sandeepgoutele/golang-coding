package main

import (
	"sync"

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
		LIMIT 1`)
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
}
