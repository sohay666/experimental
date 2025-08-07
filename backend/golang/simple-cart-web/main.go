package main

import (
	"fmt"
	"log"
	"net/http"
	"shop-cart/config"
	ctr "shop-cart/controller"
)

func routes() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong!")
	})

	http.HandleFunc("/products", ctr.Products)
	http.HandleFunc("/add-to-cart", ctr.AddToCart)
	http.HandleFunc("/list-cart", ctr.ListCart)
}

func main() {
	routes()
	fmt.Printf("Starting server at port %s\n", config.PORT)
	if err := http.ListenAndServe(":"+config.PORT, nil); err != nil {
		log.Fatal(err)
	}
}
