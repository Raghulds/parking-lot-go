package parkinglot

import (
	"fmt"
)

var lot *ParkingLot

type ParkingLot struct {
	levels []*Level
}

func GetParkingLotInstance() *ParkingLot {
	if lot == nil {
		lot = &ParkingLot{
			levels: []*Level{},
		}
	}
	return lot
}

func (p *ParkingLot) AddLevel(number int, slots int) *Level {
	floor := CreateFloor(number, slots)
	return floor
}

func CreateParkingLot(floors int, slots int) *ParkingLot {
	fmt.Printf("Creating Parking lot with %d floors and %d slots in each floors...\n", floors, slots)
	lot := &ParkingLot{}
	for i := 0; i < floors; i++ {
		floor := lot.AddLevel(i+1, slots)
		lot.levels = append(lot.levels, floor)
	}
	fmt.Println("Parking lot created successfully")
	lot.PrintParkingLot()
	return lot
}

func (p *ParkingLot) PrintParkingLot() {
	fmt.Println("\n=== Parking Lot Details ===")
	for i, level := range p.levels {
		fmt.Printf("\nFloor %d:\n", i+1)
		fmt.Printf("Status: %s\n", getStatus(level))
		fmt.Println("Parking Spots:")
		for _, spot := range level.parkingSpots {
			status := "Available"
			if spot.closed {
				status = "Closed"
			} else if spot.vehicle != nil {
				status = fmt.Sprintf("Occupied by %s (ID: %s)", spot.GetVehicleType(), spot.GetVehicleId())
			}
			fmt.Printf("  Spot %d: %s for %s\n", spot.spot, status, spot.GetVehicleType())
		}
	}
	fmt.Println("\n=== End of Parking Lot Details ===")
}

func getStatus(level *Level) string {
	if !level.active {
		return "Inactive"
	}
	if level.full {
		return "Full"
	}
	return "Active"
}
