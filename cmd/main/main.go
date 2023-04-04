package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/paridhimaheshwari/book-management-system/pkg/config"
	"github.com/paridhimaheshwari/book-management-system/pkg/routes"
)
var db *gorm.DB

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}


func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book(
	db.NewRecord(b)
	db.Create(&b)
)