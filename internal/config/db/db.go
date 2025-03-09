package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("⚠️ Warning: No .env file found, relying on system env variables")
	}

	dbURL := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	fmt.Println("📌 Connecting to DB:", dbURL)

	var err error
	DB, err = sql.Open("postgres", dbURL) // ✅ Global DB Fix

	if err != nil {
		log.Fatalf("❌ Failed to connect to the database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("❌ Database not responding: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL!")
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}
func createTables() {
	createTableHandler(CreateUserTableQuery, "User")
	createTableHandler(CreateChanelTableQuery, "Chanel")
	createTableHandler(CreateChannelUserTableQuery, "Chanel Members")
	createTableHandler(CreateMessageTableQuery, "Message")
	createTableHandler(CreateArchivedMessageTableQuery, "Archived Message")
}

func createTableHandler(query string, tableName string) error {
	fmt.Printf("Creating %s table...\n", tableName)
	_, err := DB.Exec(query)
	if err != nil {
		panic(fmt.Sprintf("Failed to create %s table \n %v", tableName, err))
	}

	fmt.Printf("✅ Created %s table!", tableName)
	return nil
}
