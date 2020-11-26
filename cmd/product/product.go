package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"github.com/jmoiron/sqlx" 
	"github.com/seminarioGo/internal/config"
	"github.com/seminarioGo/internal/database"
	"github.com/seminarioGo/internal/service/product"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := readConfig()	
	db, err := database.NewDatabase(cfg)
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}	
	if err := createSchema(db); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	service, _ := product.New(db, cfg)
	httpService := product.NewHTTPTransport(service)
	r := gin.Default()
	httpService.Register(r)
	r.Run()
}

func readConfig() *config.Config{
	configFile := flag.String("config","./config.yaml","this is the service config")
	flag.Parse() 
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {	
		fmt.Println(err.Error())
		os.Exit(1) 
	}
	return cfg
}

func createSchema(db *sqlx.DB) error {
	schema := "CREATE TABLE IF NOT EXISTS product (id integer primary key autoincrement, name varchar, price real);"
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}	
	query := "INSERT INTO product (name, price) VALUES (?,?)"
	_, _ = db.Exec(query, "Prueba "+string(time.Now().Second()), 10) 	
	return nil	
} 