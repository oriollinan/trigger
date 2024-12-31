package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"trigger.com/api/src/database"
	"trigger.com/api/src/parser"
	"trigger.com/api/src/server"
)

func main() {
	args, err := parser.CmdArgs()

	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(*args.EnvPath)
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Open(database.ConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.WithValue(context.Background(), database.CtxKey, db)
	server, err := server.Create(*args.Port, ctx)
	if err != nil {
		log.Fatal(err)
	}

	go server.Start()
	defer server.Stop()
}