package main

import (
	"context"
	"log"

	"github.com/viktorbeznosov/job4j_go_lang_base/internal/config"
	"github.com/viktorbeznosov/job4j_go_lang_base/internal/db"
	"github.com/viktorbeznosov/job4j_go_lang_base/internal/repository"
	"github.com/viktorbeznosov/job4j_go_lang_base/internal/tracker"
)

func main() {
	ctx := context.Background()

	cfg := db.Config{
		Host:     config.Env("DB_HOST", "localhost"),
		Port:     config.EnvInt("DB_PORT", 6543),
		User:     config.Env("DB_USER", "postgres"),
		Password: config.Env("DB_PASSWORD", "password"),
		DBName:   config.Env("DB_NAME", "tracker"),
		SSLMode:  config.Env("DB_SSLMODE", "disable"),
	}

	pool, err := db.NewPool(ctx, cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	repo := repository.NewRepoPg(pool)

	ui := tracker.UI{
		In:    tracker.ConsoleInput{},
		Out:   tracker.ConsoleOutput{},
		Store: repo,
	}

	if err := ui.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
