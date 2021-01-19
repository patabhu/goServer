package adapters

import (
	"fmt"
	"goServer/config"
	"goServer/models"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
)

// NewDBAdapterRepository - Repository layer for database connection
func NewDBAdapterRepository(config *config.Config) *gorm.DB {

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.DBName, config.Database.Pass)

	dbConn, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Println("connection string ", connectionString)
		log.Fatal("unable to connect ", err)
		return nil
	}
	log.Println("connected to database")
	dbConn.AutoMigrate(models.ToDo{})
	return dbConn
}
