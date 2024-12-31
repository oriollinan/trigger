package main

import (
	"context"
	"log"
	"os"

	"trigger.com/trigger/internal/sync"
	"trigger.com/trigger/pkg/arguments"
	"trigger.com/trigger/pkg/middleware"
	"trigger.com/trigger/pkg/mongodb"
	"trigger.com/trigger/pkg/router"
	"trigger.com/trigger/pkg/server"
)

func main() {
	args, err := arguments.Command()
	if err != nil {
		log.Fatal(err)
	}

	err = sync.Env(*args.EnvPath)
	if err != nil {
		log.Fatal(err)
	}

	mongoClient, _, err := mongodb.Open(mongodb.ConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	syncCollection := mongoClient.Database(
		os.Getenv("MONGO_DB"),
	).Collection("sync")

	router, err := router.Create(
		context.WithValue(
			context.TODO(),
			mongodb.CtxKey,
			syncCollection,
		),
		sync.Router,
	)
	if err != nil {
		log.Fatal(err)
	}

	server, err := server.Create(
		router,
		middleware.Create(
			middleware.Logging,
			middleware.Cors,
		),
		*args.Port,
	)
	if err != nil {
		log.Fatal(err)
	}

	go server.Start()
	server.Stop()
}
