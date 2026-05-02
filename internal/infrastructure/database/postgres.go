package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() (*sql.DB, *gorm.DB, error) {
	// err := godotenv.Load()
	// if err != nil {
	// 	return nil, nil, fmt.Errorf("error loading .env file: %v", err)
	// }
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	schema := os.Getenv("DB_SCHEMA")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=%s",
		host, port, user, password, dbname, schema, sslmode)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	sqlDB, err := gormDB.DB()
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	sqlDB.SetConnMaxLifetime(10)
	sqlDB.SetMaxOpenConns(5)
	if err := sqlDB.Ping(); err != nil {
		fmt.Printf("failed to ping PostgreSQL: %v", err)
		return nil, nil, err
	}

	fmt.Println("Successfully connected to PostgreSQL")
	return sqlDB, gormDB, nil
}
