package main

import (
	"log"
	"net/http"
	"pharma"
	"pharma/db"
	"pharma/web"
)

func main() {
	product, err := db.NewStore("postgres://postgres:12345@localhost/pharmacy?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}
	defer product.Close()

	var emptyCategory pharma.Category
	var emptyCategory1 pharma.Supplies
	var emptyCategory2 pharma.PharmacyInfo
	var emptyCategory3 pharma.Order

	h := web.NewHandler(product, emptyCategory, emptyCategory1, emptyCategory2, emptyCategory3)
	http.ListenAndServe(":8080", h)
}
