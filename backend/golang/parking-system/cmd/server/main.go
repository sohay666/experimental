package main

import (
	"log"
	"net/http"
	"parking-system/internal/handlers"
	"parking-system/internal/models"
)

func DefaultParkingSystem() *handlers.ParkingHandler {
	entryPoint := 3
	distances := [][]int{
		{1, 3, 5}, // slot 0 distances from entry 0, 1, 2   | SP
		{2, 1, 4}, // slot 1 distances                      | MP
		{3, 2, 3}, // slot 2 distances                      | LP
	}
	sizes := []models.SlotSize{models.SP, models.MP, models.LP}
	ph, err := handlers.NewParkingHandler(entryPoint, distances, sizes)
	if err != nil {
		log.Fatal(err.Error())
	}
	return ph
}

func main() {
	parkingHandler := DefaultParkingSystem()

	http.HandleFunc("/park", parkingHandler.Park)
	http.HandleFunc("/unpark", parkingHandler.Unpark)

	log.Printf("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
