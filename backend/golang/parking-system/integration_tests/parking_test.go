package integration_tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"parking-system/internal/handlers"
	"parking-system/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Unpark struct {
	PlateNumber string `json:"plate_number"`
	Duration    string `json:"duration"`
	Fee         int    `json:"fee"`
}
type UnparkResponse struct {
	Data       Unpark `json:"data"`
	StatusCode int    `json:"status_code"`
}

type ErrorResponse struct {
	Error      string `json:"error"`
	StatusCode int    `json:"status_code"`
}

func TestParkingFlow(t *testing.T) {
	entryPoint := 3
	distances := [][]int{
		{1, 3, 5}, // slot 0 | SP(0)
		{2, 1, 4}, // slot 1 | MP(1)
		{3, 2, 3}, // slot 2 | LP(2)
	}
	sizes := []models.SlotSize{models.SP, models.MP, models.LP}

	handler, _ := handlers.NewParkingHandler(entryPoint, distances, sizes)

	mux := http.NewServeMux()
	mux.HandleFunc("/park", handler.Park)
	mux.HandleFunc("/unpark", handler.Unpark)

	server := httptest.NewServer(mux)
	defer server.Close()

	t.Run("Park Small Car 2 Hours (Flat rate)", func(t *testing.T) {
		plate := "S001"
		// Park
		parkPayload := map[string]interface{}{
			"entry_point":  0,
			"vehicle_type": 0,
			"plate_number": plate,
			"start_time":   "2025-06-06T08:00:00Z",
		}
		assertHTTPPost(t, server.URL+"/park", parkPayload)

		// Unpark
		unparkPayload := map[string]interface{}{
			"plate_number": plate,
			"end_time":     "2025-06-06T10:00:00Z",
		}
		resp := assertHTTPDelete(t, server.URL+"/unpark", unparkPayload)

		// Check fee
		var result UnparkResponse
		err := json.Unmarshal(resp, &result)
		assert.NoError(t, err)
		assert.Equal(t, plate, result.Data.PlateNumber)
		assert.Equal(t, 40, result.Data.Fee) // adjust based on your actual flat rate
	})

	t.Run("Park Medium car and Unpark after 5 hours (Flat + hourly MP rate)", func(t *testing.T) {
		plate := "M001"
		// Park
		parkPayload := map[string]interface{}{
			"entry_point":  1,
			"vehicle_type": 1,
			"plate_number": plate,
			"start_time":   "2025-06-06T08:00:00Z",
		}
		assertHTTPPost(t, server.URL+"/park", parkPayload)

		// Unpark
		unparkPayload := map[string]interface{}{
			"plate_number": plate,
			"end_time":     "2025-06-06T13:00:00Z",
		}
		resp := assertHTTPDelete(t, server.URL+"/unpark", unparkPayload)

		// // Check fee
		var result UnparkResponse
		err := json.Unmarshal(resp, &result)
		assert.NoError(t, err)
		assert.Equal(t, plate, result.Data.PlateNumber)
		assert.Equal(t, 160, result.Data.Fee) // adjust based on your actual flat rate
	})

	t.Run("Park Large car and Unpark after 7.5 hours (rounded to 8h total)", func(t *testing.T) {
		plate := "L001"
		// Park
		parkPayload := map[string]interface{}{
			"entry_point":  2,
			"vehicle_type": 2,
			"plate_number": plate,
			"start_time":   "2025-06-06T09:00:00Z",
		}
		assertHTTPPost(t, server.URL+"/park", parkPayload)

		// Unpark
		unparkPayload := map[string]interface{}{
			"plate_number": plate,
			"end_time":     "2025-06-06T16:30:00Z",
		}
		resp := assertHTTPDelete(t, server.URL+"/unpark", unparkPayload)

		// // Check fee
		var result UnparkResponse
		err := json.Unmarshal(resp, &result)
		assert.NoError(t, err)
		assert.Equal(t, plate, result.Data.PlateNumber)
		assert.Equal(t, 540, result.Data.Fee) // adjust based on your actual flat rate
	})

	t.Run("Unpark after 30 hours (24h chunk + hourly remainder)", func(t *testing.T) {
		plate := "S002"
		// Park
		parkPayload := map[string]interface{}{
			"entry_point":  0,
			"vehicle_type": 0,
			"plate_number": plate,
			"start_time":   "2025-06-06T01:00:00Z",
		}
		assertHTTPPost(t, server.URL+"/park", parkPayload)

		// Unpark
		unparkPayload := map[string]interface{}{
			"plate_number": plate,
			"end_time":     "2025-06-07T07:00:00Z",
		}
		resp := assertHTTPDelete(t, server.URL+"/unpark", unparkPayload)

		// Check fee
		var result UnparkResponse
		err := json.Unmarshal(resp, &result)
		assert.NoError(t, err)
		assert.Equal(t, plate, result.Data.PlateNumber)
		assert.Equal(t, 5100, result.Data.Fee) // adjust based on your actual flat rate
	})

	t.Run("Exit and Re-enter Within 1 Hour (Continuous rate)", func(t *testing.T) {
		plate := "S003"
		// Park
		parkPayload := map[string]interface{}{
			"entry_point":  0,
			"vehicle_type": 0,
			"plate_number": plate,
			"start_time":   "2025-06-06T01:00:00Z",
		}
		assertHTTPPost(t, server.URL+"/park", parkPayload)

		// Unpark
		unparkPayload := map[string]interface{}{
			"plate_number": plate,
			"end_time":     "2025-06-06T03:00:00Z",
		}
		resp := assertHTTPDelete(t, server.URL+"/unpark", unparkPayload)
		// Check fee
		var result UnparkResponse
		json.Unmarshal(resp, &result)

		// Re-enter at 03:30 (within 1h)
		parkPayload["start_time"] = "2025-06-06T03:30:00Z"
		assertHTTPPost(t, server.URL+"/park", parkPayload)

		//Unpark again later
		unparkPayload["end_time"] = "2025-06-06T06:30:00Z"
		resp2 := assertHTTPDelete(t, server.URL+"/unpark", unparkPayload)

		var result2 UnparkResponse
		err := json.Unmarshal(resp2, &result2)

		assert.NoError(t, err)
		assert.Equal(t, plate, result.Data.PlateNumber)
		assert.Equal(t, 40, result.Data.Fee)  // adjust based on your actual flat rate
		assert.Equal(t, 40, result2.Data.Fee) // adjust based on your actual flat rate
	})

	t.Run("Error while No Available Slot (Expected error message)", func(t *testing.T) {
		plates := []string{
			"S001", "S002", "S003",
		}

		// Fill up all slots with 3 vehicles first:
		for _, plate := range plates {
			parkPayload := map[string]interface{}{
				"entry_point":  0,
				"vehicle_type": 0,
				"plate_number": plate,
				"start_time":   "2025-06-06T01:00:00Z",
			}
			assertHTTPPost(t, server.URL+"/park", parkPayload)
		}

		parkPayload := map[string]interface{}{
			"entry_point":  0,
			"vehicle_type": 0,
			"plate_number": "S004",
			"start_time":   "2025-06-06T01:00:00Z",
		}
		resp := assertHTTPPost(t, server.URL+"/park", parkPayload)

		var result ErrorResponse
		json.Unmarshal(resp, &result)
		assert.Equal(t, "no available parking slot", result.Error) // adjust based on your actual flat rate
		assert.Equal(t, http.StatusUnprocessableEntity, result.StatusCode)
	})

	t.Run("Error while No Active Parking assignment found", func(t *testing.T) {
		unparkPayload := map[string]interface{}{
			"plate_number": "S004",
			"end_time":     "2025-06-06T03:00:00Z",
		}
		resp := assertHTTPDelete(t, server.URL+"/unpark", unparkPayload)
		var result ErrorResponse
		json.Unmarshal(resp, &result)
		assert.Equal(t, "no active parking assignment found for this vehicle", result.Error)
		assert.Equal(t, http.StatusUnprocessableEntity, result.StatusCode)
	})

	t.Run("Error while Vehicle already parked", func(t *testing.T) {
		plate := "S003"
		parkPayload := map[string]interface{}{
			"entry_point":  0,
			"vehicle_type": 0,
			"plate_number": plate,
			"start_time":   "2025-06-06T01:00:00Z",
		}
		assertHTTPPost(t, server.URL+"/park", parkPayload)

		// Re-enter at 03:30 (within 1h)
		parkPayload["start_time"] = "2025-06-06T03:30:00Z"
		resp := assertHTTPPost(t, server.URL+"/park", parkPayload)
		var result ErrorResponse
		json.Unmarshal(resp, &result)
		assert.Equal(t, "vehicle already parked", result.Error)
		assert.Equal(t, http.StatusUnprocessableEntity, result.StatusCode)
	})
}

func assertHTTPPost(t *testing.T, url string, payload interface{}) []byte {
	body, _ := json.Marshal(payload)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	assert.NoError(t, err)
	defer resp.Body.Close()
	result, _ := io.ReadAll(resp.Body)
	return result
}

func assertHTTPDelete(t *testing.T, url string, payload interface{}) []byte {
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()
	result, _ := io.ReadAll(resp.Body)
	return result
}
