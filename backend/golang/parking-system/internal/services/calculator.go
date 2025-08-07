package services

import (
	"math"
	"parking-system/internal/models"
	"time"
)

type Calculator struct{}

func NewCalculator() *Calculator {
	return &Calculator{}
}

/*
Calculate: to handle the calculation of fees
1. All types of cars pay the flat rate of 40 pesos for the first three (3) hours
2. The exceeding hourly rate beyond the initial three (3) hours will be charged as follows:
  - 20/hour for vehicles parked in SP;
  - 60/hour for vehicles parked in MP;
  - 100/hour for vehicles parked in LP
  - exceeds 24 hours, every full 24-hour chunk is charged 5,000 pesos regardless of the parking slot.
  - Parking fees are calculated using the rounding up e.g. 6.4 hours must be rounded to 7
*/
func (c *Calculator) Calculate(slot *models.ParkingSlot, endTime time.Time) float64 {

	start := slot.ParkedTime
	duration := endTime.Sub(start)
	hours := int(math.Ceil(duration.Hours()))

	fullDays := hours / 24
	remainderHours := hours % 24

	fee := float64(fullDays) * 5000
	if remainderHours <= 3 {
		fee += 40
	} else {
		exceed := float64(remainderHours - 3)
		rate := c.getHourlyRate(slot.Size)
		fee += 40 + (rate * exceed)
	}
	return fee
}

func (c *Calculator) getHourlyRate(slotType models.SlotSize) float64 {
	switch slotType {
	case models.SP:
		return 20
	case models.MP:
		return 60
	case models.LP:
		return 100
	default:
		return 0
	}
}
