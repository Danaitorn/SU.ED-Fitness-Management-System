package config

import (
	"fmt"
	"log"
	"os"

	"backend/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	_ = godotenv.Load()

	var dsn string

	// ถ้ามี DB_DSN ให้ใช้เลย (production)
	if os.Getenv("DB_DSN") != "" {

		dsn = os.Getenv("DB_DSN")

	} else {

		// ใช้สำหรับ local / docker
		host := "fitness_postgres"

		if os.Getenv("HOSTNAME") == "" {
			host = "localhost"
		}

		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			host,
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}

	err = database.AutoMigrate(
		&models.User{},
		&models.Package{},
		&models.Membership{},
		&models.FitnessClass{},
		&models.News{},
		&models.MembershipInfo{},
		&models.HealthAnswer{},
		&models.FitnessVisit{},
	)

	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	DB = database

	log.Println("✅ Database connected & migrated successfully!")
}

func GetJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}