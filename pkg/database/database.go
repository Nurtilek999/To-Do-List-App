package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

//type Db struct {
//	DB *sql.DB
//}

func InitDB() (*sql.DB, error) {
	username := viper.GetString("Database.Username")
	password := viper.GetString("Database.Password")
	//host := viper.GetString("Database.Host")
	//port := viper.GetInt("Database.Port")
	dbname := viper.GetString("Database.DBName")

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, dbname)
	//var err error
	DB, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return DB, nil
}

//func InitDB() (*gorm.DB, error) {
//	//username := "postgres"
//	//password := "admin"
//	//host := "127.0.0.1"
//	//port := ":8080"
//	//dbname := "postgres"
//
//	dsn := "host=localhost user=postgres  password=admin dbname=postgres port=5432 sslmode=disable"
//
//	db, err := gorm.Open("postgres", dsn)
//
//	if err != nil {
//		return nil, err
//	}
//
//	//if err := db.DB().Ping(); err != nil {
//	//	return nil, err
//	//}
//
//	return db, nil
//}
