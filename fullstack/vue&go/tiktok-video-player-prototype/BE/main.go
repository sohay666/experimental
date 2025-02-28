package main

import (
	"fmt"
	"log"
	"net/http"
	ctr "toktok/controller"
	"toktok/utils"
)

func routes() {
	prefix := utils.Config.App.Prefix
	http.HandleFunc(prefix+"/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong!")
	})
	http.HandleFunc(prefix+"/search/videos", ctr.VideosHandler)
	http.HandleFunc(prefix+"/search/suggestions", ctr.SuggestionsHandler)
}

func main() {
	utils.LoadConfig()
	routes()
	port := utils.Config.App.Port
	fmt.Printf("Starting server at port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
