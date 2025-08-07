package models

import "time"

type VehicleType int

const (
	Small VehicleType = iota
	Medium
	Large
)

type SlotSize int

const (
	SP SlotSize = iota // Small Parking
	MP                 // Medium Parking
	LP                 // Large Parking
)

type Vehicle struct {
	PlateNumber string      `json:"plate_number"`
	Type        VehicleType `json:"type"`
	EntryPoint  int         `json:"entry_point"`
	LastExit    time.Time   `json:"last_exit"`
}

type ParkingSlot struct {
	ID         int
	Size       SlotSize
	Distances  []int // distances to each entry
	Occupied   bool
	Occupant   *Vehicle
	ParkedTime time.Time
}

type ParkingSystem struct {
	EntryPoints int
	Slots       []*ParkingSlot
	VehicleMap  map[string]*ParkingSlot // vehicle plate -> slot
}

type ParkingAssignment struct {
	PlateNumber string    `json:"plate_number"`
	BestSlot    int       `json:"best_slot"`
	Size        SlotSize  `json:"size"`
	ParkedTime  time.Time `json:"parked_time"`
}

type ParkingRelease struct {
	PlateNumber string  `json:"plate_number"`
	Fee         float64 `json:"fee"`
}

// Compatibility check
func (v VehicleType) CanFit(s SlotSize) bool {
	switch v {
	case Small:
		return true
	case Medium:
		return s == MP || s == LP
	case Large:
		return s == LP
	}
	return false
}
