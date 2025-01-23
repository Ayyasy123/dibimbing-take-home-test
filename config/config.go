package config

import (
	"log"

	"github.com/Ayyasy123/dibimbing-take-home-test/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// ambil variable dari .env
	// dbUser := os.Getenv("DB_USER")
	// dbPass := os.Getenv("DB_PASS")
	// dbName := os.Getenv("DB_NAME")
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")

	// buat connection
	// dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	// Konfigurasi koneksi database with xampp
	dsn := "root:@tcp(127.0.0.1:3306)/dibimbing_takehometest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	err = db.AutoMigrate(
		&entity.User{},
		&entity.Event{},
		&entity.Ticket{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}

	DB = db
	log.Println("Database connected and migrated successfully")
}
