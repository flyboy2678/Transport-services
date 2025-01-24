package main

import (
	"log"
	"transportService/internal/db"
	"transportService/internal/env"
	"transportService/internal/store"

	_ "github.com/lib/pq"
)

func main() {

	env.Init()
	cfg := config{
		addr: env.GetString("ADDR", ":3000"),
		db: dbConfig{
			addr: env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost:5432/transportService?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}
	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store: store,
	}


	mux:= app.mount()

	log.Fatal(app.run(mux))
}