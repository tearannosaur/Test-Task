package main

import (
	"log"
	db "task/database"
	h "task/handlers"
	r "task/repository"
	s "task/server"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	database, err := db.DBInit()
	if err != nil {
		log.Fatalln("Database init error", err)
	}
	log.Println("Successfully database connection")
	repo := r.RepositoryModuleInit(database)
	handler := h.HandlerModuleInit(repo)
	s.ServerInit(handler)
}
