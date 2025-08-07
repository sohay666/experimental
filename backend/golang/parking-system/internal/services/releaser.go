package services

import (
	"errors"
	"parking-system/internal/models"
	"time"
)

type Releaser struct {
	parkingSystem models.ParkingSystem
}

func NewReleaser(parkingSystem models.ParkingSystem) *Releaser {
	return &Releaser{
		parkingSystem: parkingSystem,
	}
}

func (r *Releaser) Release(plateNumber string, endTime time.Time) (*models.ParkingRelease, error) {

	slot, exists := r.parkingSystem.VehicleMap[plateNumber]
	if !exists {
		return nil, errors.New("no active parking assignment found for this vehicle")
	}

	calculator := NewCalculator()
	fee := calculator.Calculate(slot, endTime)

	// delete Vehicle by plateNumber from VehicleMap
	slot.Occupied = false
	slot.Occupant.LastExit = time.Now()
	delete(r.parkingSystem.VehicleMap, plateNumber)

	releaser := models.ParkingRelease{
		PlateNumber: plateNumber,
		Fee:         fee,
	}
	return &releaser, nil
}
