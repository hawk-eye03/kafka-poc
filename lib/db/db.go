package db

import (
	"fmt"
	"log"

	"github.com/hawk-eye03/kafka-poc/lib/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
)

type MySQLConnection struct {
	DB *gorm.DB
}

func NewMainDBConnection(config *config.ConfigMap) *MySQLConnection {
	return &MySQLConnection{
		DB: connectDB(&config.MainDB),
	}
}

func connectDB(dbConfig *config.DBConfig) *gorm.DB {
	if !dbConfig.ValidateDBCreds() {
		zap.L().Info("Invalid DB Creds")
		panic("Cannot make DB Connection!!")
	}
	// Open a database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	zap.L().Info("Connected to the MySQL database")
	return db
}
