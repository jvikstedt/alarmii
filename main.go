package main

import (
	"github.com/boltdb/bolt"
	"github.com/jvikstedt/alarmii/server"
	"github.com/jvikstedt/alarmii/server/repository"
)

func main() {
	db, err := bolt.Open("development.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	jobRepository := repository.NewJobRepository(db)
	server := server.NewServer(3000, jobRepository)
	err = server.Start()
	if err != nil {
		panic(err)
	}
}
