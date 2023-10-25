package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/knadh/koanf/v2"
	_ "github.com/lib/pq"
)

func ConnectTodDb(config *koanf.Koanf) *sql.DB {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.String("user"), config.String("password"),
		config.String("database"))
	// dsn := config.String("url")
	fmt.Println(dsn)
	dB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("error starting database", err)
	}
	return dB
}
