# Parking lot System
I'll implement a parking allocation system for the Object-Oriented Mall in Golang.

# Project Structure
```
/parking-system
  /cmd
    /server
      main.go
  /integration_tests
     parking_test.go
  /internal
    /handlers
      parking.go
    /models
      models.go
    /services
      allocator.go
      calculator.go
      releaser.go
  go.mod
  go.sum
```

# Key Features Implemented

1. Parking Allocation:
   - Finds closest available slot based on vehicle type
   - Handles three entry points (extensible to more)
   - Validates vehicle types and slot compatibility
2. Pricing Calculation:
    - Flat rate for first 3 hours
    - Different hourly rates based on slot size
    - Daily rate for stays over 24 hours
    - Continuous rate for vehicles returning within 1 hour
3. Data Models:
    - Parking System with entry points and distances to slots
    - Parking slots with types and distances
    - Vehicles with types and plate numbers
    - Parking assignments with timestamps

# Run Server
```
go run cmd/server/main.go                         # Running the server
```

# Run Integrated Test
```
go test -timeout 30s -run ^TestParkingFlow$ parking-system/integration_tests -gcflags=all=-N -gcflags=all=-l -count=1 -v
```

# Test Case

1. Park Small car and Unpark after 2 hours (Flat rate only)
```
curl -X POST http://localhost:8080/park \
  -H "Content-Type: application/json" \
  -d '{"entry_point": 0, "vehicle_type": 0, "plate_number": "S001", "start_time": "2025-06-06T08:00:00Z"}'

curl -X DELETE http://localhost:8080/unpark \
  -H "Content-Type: application/json" \
  -d '{"plate_number": "S001", "end_time": "2025-06-06T10:00:00Z"}'
```

2. Park Medium car and Unpark after 5 hours (Flat + hourly MP rate)
```
curl -X POST http://localhost:8080/park \
  -H "Content-Type: application/json" \
  -d '{"entry_point": 1, "vehicle_type": 1, "plate_number": "M001", "start_time": "2025-06-06T08:00:00Z"}'

curl -X DELETE http://localhost:8080/unpark \
  -H "Content-Type: application/json" \
  -d '{"plate_number": "M001", "end_time": "2025-06-06T13:00:00Z"}'
```

3. Park Large car and Unpark after 7.5 hours (rounded to 8h total)
```
curl -X POST http://localhost:8080/park \
  -H "Content-Type: application/json" \
  -d '{"entry_point": 2, "vehicle_type": 2, "plate_number": "L001", "start_time": "2025-06-06T09:00:00Z"}'

curl -X DELETE http://localhost:8080/unpark \
  -H "Content-Type: application/json" \
  -d '{"plate_number": "L001", "end_time": "2025-06-06T16:30:00Z"}'
```

4. Unpark after 30 hours (24h chunk + hourly remainder)
```
curl -X POST http://localhost:8080/park \
  -H "Content-Type: application/json" \
  -d '{"entry_point": 0, "vehicle_type": 0, "plate_number": "S002", "start_time": "2025-06-06T01:00:00Z"}'

curl -X DELETE http://localhost:8080/unpark \
  -H "Content-Type: application/json" \
  -d '{"plate_number": "S002", "end_time": "2025-06-07T07:00:00Z"}'
```

5. Exit and Re-enter Within 1 Hour (Continuous rate)
```
# First park
curl -X POST http://localhost:8080/park \
  -H "Content-Type: application/json" \
  -d '{"entry_point": 0, "vehicle_type": 0, "plate_number": "S003", "start_time": "2025-06-06T01:00:00Z"}'

# Unpark at 03:00
curl -X DELETE http://localhost:8080/unpark \
  -H "Content-Type: application/json" \
  -d '{"plate_number": "S003", "end_time": "2025-06-06T03:00:00Z"}'

# Re-enter at 03:30 (within 1h)
curl -X POST http://localhost:8080/park \
  -H "Content-Type: application/json" \
  -d '{"entry_point": 0, "vehicle_type": 0, "plate_number": "S003", "start_time": "2025-06-06T03:30:00Z"}'

# Unpark again later
curl -X DELETE http://localhost:8080/unpark \
  -H "Content-Type: application/json" \
  -d '{"plate_number": "S003", "end_time": "2025-06-06T06:30:00Z"}'
```

6. No Available Slot (Expected error message)
```
Fill up all slots with 3 vehicles first:

curl -X POST http://localhost:8080/park \
  -H "Content-Type: application/json" \
  -d '{"entry_point": 0, "vehicle_type": 0, "plate_number": "X1", "start_time": "2025-06-06T01:00:00Z"}'

curl -X POST http://localhost:8080/park \
  -H "Content-Type: application/json" \
  -d '{"entry_point": 0, "vehicle_type": 0, "plate_number": "X2", "start_time": "2025-06-06T01:00:00Z"}'

curl -X POST http://localhost:8080/park \
  -H "Content-Type: application/json" \
  -d '{"entry_point": 0, "vehicle_type": 0, "plate_number": "X3", "start_time": "2025-06-06T01:00:00Z"}'

# Try to park one more
curl -X POST http://localhost:8080/park \
  -H "Content-Type: application/json" \
  -d '{"entry_point": 0, "vehicle_type": 0, "plate_number": "X4", "start_time": "2025-06-06T01:00:00Z"}'
```

7. No Active Parking assignment found
```
curl -X POST http://localhost:8080/park \
  -H "Content-Type: application/json" \
  -d '{"entry_point": 0, "vehicle_type": 0, "plate_number": "S004", "start_time": "2025-06-06T01:00:00Z"}'
```

8. Vehicle already parked
```
curl -X POST http://localhost:8080/park \
  -H "Content-Type: application/json" \
  -d '{"entry_point": 0, "vehicle_type": 0, "plate_number": "S004", "start_time": "2025-06-06T01:00:00Z"}'

curl -X POST http://localhost:8080/park \
  -H "Content-Type: application/json" \
  -d '{"entry_point": 0, "vehicle_type": 0, "plate_number": "S004", "start_time": "2025-06-06T01:10:00Z"}'
```
