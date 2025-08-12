package config

import (
	env "AuthInGo/config/env"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.User = env.GetString("DB_USER", "root")
	cfg.Passwd = env.GetString("DB_PASSWORD", "root")
	cfg.Net = env.GetString("DB_NET", "tcp")
	cfg.Addr = env.GetString("DB_ADDR", "localhost:3306")
	cfg.DBName = env.GetString("DBName", "test_db")

	fmt.Println("configuring database with the following settings:")
	fmt.Printf("User: %s, Net: %s, Addr: %s, DBName: %s\n, formatting the DSN:%s\n",
		cfg.User, cfg.Net, cfg.Addr, cfg.DBName, cfg.FormatDSN())

	// here FormatDSN() is used to format the DSN (Data Source Name) string
	// which is used to connect to the database.
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return nil, err
	}
	fmt.Println("Trying to ping the database...")
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Error pinging the database:", pingErr)
		return nil, pingErr
	}
	fmt.Println("Successfully connected to the database!", cfg.DBName)

	return db, nil
}
