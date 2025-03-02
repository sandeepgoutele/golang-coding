package main

import (
	"database/sql"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func newConn() *sql.DB {
	_db, err := sql.Open("mysql", "sandeep:sandeep@tcp(localhost:3306)/practise")
	if err != nil {
		panic(err)
	}
	return _db
}

type conn struct {
	db *sql.DB
}

type cpool struct {
	channel chan interface{}
	maxConn int
	conns   []*conn
	mu      *sync.Mutex
}

func NewPool(maxConn int) (*cpool, error) {
	var mu = sync.Mutex{}
	pool := &cpool{
		mu:      &mu,
		conns:   make([]*conn, 0, maxConn),
		channel: make(chan interface{}, maxConn),
		maxConn: maxConn,
	}

	for idx := 0; idx < maxConn; idx++ {
		pool.conns = append(pool.conns, &conn{newConn()})
		pool.channel <- nil
	}
	return pool, nil
}

func (pool *cpool) Get() (*conn, error) {
	<-pool.channel

	pool.mu.Lock()
	cn := pool.conns[0]
	pool.conns = pool.conns[1:]
	pool.mu.Unlock()
	return cn, nil
}

func (pool *cpool) Put(cn *conn) {
	pool.mu.Lock()
	pool.conns = append(pool.conns, cn)
	pool.mu.Unlock()
	pool.channel <- nil
}

func benchmarkPool() {
	startTime := time.Now()
	pool, err := NewPool(10)
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	for idx := 0; idx < 500; idx++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			con, err := pool.Get()
			if err != nil {
				panic(err)
			}
			_, err = con.db.Exec("SELECT SLEEP(0.1);")
			if err != nil {
				panic(err)
			}
			pool.Put(con)
		}()
	}
	wg.Wait()
	log.Print("Benchmark connection pool ", time.Since(startTime))
}

func benchmarkNonPool() {
	startTime := time.Now()
	var wg sync.WaitGroup
	for idx := 0; idx < 50; idx++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			db := newConn()
			_, err := db.Exec("SELECT SLEEP(0.1);")
			if err != nil {
				panic(err)
			}
			db.Close()
		}()
	}
	wg.Wait()
	log.Print("Benchmark non connection pool ", time.Since(startTime))
}

func main() {
	// benchmarkNonPool()
	benchmarkPool()
}
