package database

import (
	"testing"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "github.com/DarkLord017/athena/athena/database/models"
)
func TestMigrations(t *testing.T) {
	dsn := "root:MySQLDatabase$24@tcp(127.0.0.1:3306)/athena?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Migrate models
	MigrateUp(db)
}