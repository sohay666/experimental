package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"parking-system/internal/models"
	"parking-system/internal/services"
	"time"
)

type ParkingHandler struct {
	allocator *services.Allocator
	releaser  *services.Releaser
}

func NewParkingHandler(entryPoints int, distances [][]int, sizes []models.SlotSize) (*ParkingHandler, error) {

	if entryPoints < 3 {
		return nil, errors.New("must have at least 3 entry points")
	}
	if len(distances) != len(sizes) {
		return nil, errors.New("distances and sizes must match")
	}

	slots := []*models.ParkingSlot{}
	for i, dist := range distances {
		slots = append(slots, &models.ParkingSlot{
			ID:        i,
			Size:      sizes[i],
			Distances: dist,
		})
	}

	parkingSystem := models.ParkingSystem{
		EntryPoints: entryPoints,
		Slots:       slots,
		VehicleMap:  make(map[string]*models.ParkingSlot),
	}

	return &ParkingHandler{
		allocator: services.NewAllocator(parkingSystem),
		releaser:  services.NewReleaser(parkingSystem),
	}, nil
}

func (h *ParkingHandler) Park(w http.ResponseWriter, r *http.Request) {
	var request struct {
		EntryPoint  int                `json:"entry_point"`
		VehicleType models.VehicleType `json:"vehicle_type"`
		PlateNumber string             `json:"plate_number"`
		StartTime   *time.Time         `json:"start_time"`
	}
	var errMsg struct {
		Error      string `json:"error"`
		StatusCode int    `json:"status_code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errMsg.Error = err.Error()
		errMsg.StatusCode = http.StatusBadRequest
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	assignment, err := h.allocator.Allocate(request.EntryPoint, request.VehicleType, request.PlateNumber, request.StartTime)
	if err != nil {
		errMsg.Error = err.Error()
		errMsg.StatusCode = http.StatusUnprocessableEntity
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	var response struct {
		Data       models.ParkingAssignment `json:"data"`
		StatusCode int                      `json:"status_code"`
	}

	response.Data = *assignment
	response.StatusCode = http.StatusOK

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *ParkingHandler) Unpark(w http.ResponseWriter, r *http.Request) {

	var request struct {
		PlateNumber string    `json:"plate_number"`
		EndTime     time.Time `json:"end_time"`
	}
	var errMsg struct {
		Error      string `json:"error"`
		StatusCode int    `json:"status_code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errMsg.Error = err.Error()
		errMsg.StatusCode = http.StatusBadRequest
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	if request.PlateNumber == "" {
		errMsg.Error = "Plate number is required"
		errMsg.StatusCode = http.StatusBadRequest
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	release, err := h.releaser.Release(request.PlateNumber, request.EndTime)
	if err != nil {
		errMsg.Error = err.Error()
		errMsg.StatusCode = http.StatusUnprocessableEntity
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	var response struct {
		Data       models.ParkingRelease `json:"data"`
		StatusCode int                   `json:"status_code"`
	}

	response.Data = *release
	response.StatusCode = http.StatusOK

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
