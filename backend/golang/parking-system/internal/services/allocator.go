package services

import (
	"errors"
	"math"
	"parking-system/internal/models"
	"time"
)

type Allocator struct {
	parkingSystem models.ParkingSystem
}

func NewAllocator(parkingSystem models.ParkingSystem) *Allocator {
	return &Allocator{
		parkingSystem: parkingSystem,
	}
}

func (a *Allocator) Allocate(entryPoint int, vehicleType models.VehicleType, plateNumber string, startTime *time.Time) (*models.ParkingAssignment, error) {
	// Check if already parked
	if _, exists := a.parkingSystem.VehicleMap[plateNumber]; exists {
		return nil, errors.New("vehicle already parked")
	}

	// Find best slot
	bestSlot, err := a.findBestSlot(entryPoint, vehicleType, plateNumber, startTime)
	if err != nil {
		return nil, err
	}

	// Create new assignment
	newAssignment := models.ParkingAssignment{
		PlateNumber: plateNumber,
		BestSlot:    bestSlot,
		Size:        a.parkingSystem.VehicleMap[plateNumber].Size,
		ParkedTime:  a.parkingSystem.VehicleMap[plateNumber].ParkedTime,
	}
	return &newAssignment, nil
}

func (a *Allocator) findBestSlot(entryPoint int, vehicleType models.VehicleType, plateNumber string, startTime *time.Time) (int, error) {
	bestSlot := -1
	bestDist := math.MaxInt
	for _, slot := range a.parkingSystem.Slots {
		if slot.Occupied {
			continue
		}
		if !vehicleType.CanFit(slot.Size) {
			continue
		}
		distance := slot.Distances[entryPoint]
		if distance < bestDist {
			bestDist = distance
			bestSlot = slot.ID
		}
	}

	if bestSlot == -1 {
		return bestSlot, errors.New("no available parking slot")
	}

	slot := a.parkingSystem.Slots[bestSlot]
	slot.Occupied = true

	v := models.Vehicle{
		PlateNumber: plateNumber,
		Type:        vehicleType,
		EntryPoint:  entryPoint,
	}
	slot.Occupant = &v
	now := time.Now()
	if startTime != nil {
		now = *startTime
	}

	if now.Sub(v.LastExit) <= time.Hour {
		// Use old time
		slot.ParkedTime = v.LastExit
	} else {
		slot.ParkedTime = now
	}
	a.parkingSystem.VehicleMap[v.PlateNumber] = slot
	return bestSlot, nil
}
