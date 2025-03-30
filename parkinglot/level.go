package parkinglot

import (
	"parkinglotgo/vehicle"
)

type Level struct {
	floorNo      int
	parkingSpots []*ParkingSpot
	full         bool
	active       bool
}

var floor *Level

func CreateFloor(floorNo int, slots int) *Level {
	floor = &Level{
		floorNo:      floorNo,
		parkingSpots: []*ParkingSpot{},
		full:         false,
		active:       true,
	}
	for i := 0; i < slots; i++ {
		if floorNo == 1 {
			floor.parkingSpots = append(floor.parkingSpots, CreateParkingSpot(i+1, vehicle.Bus))
		} else {
			var carsPortion float64 = 0.4
			if float64(i) > float64(slots)*carsPortion {
				floor.parkingSpots = append(floor.parkingSpots, CreateParkingSpot(i+1, vehicle.Car))
			} else {
				floor.parkingSpots = append(floor.parkingSpots, CreateParkingSpot(i+1, vehicle.TwoWheeler))
			}
		}
	}
	return floor
}

func (l *Level) AddParkingSpot(spot *ParkingSpot) {
	l.parkingSpots = append(l.parkingSpots, spot)
}

func (l *Level) ParkVehicle(vehicle *vehicle.Vehicle) bool {
	for _, spot := range l.parkingSpots {
		if spot.isSpotAvailableForParking() == true && spot.GetVehicleType() == vehicle.GetVehicleType() {
			spot.ParkVehicle(vehicle)
			return true
		}
	}
	return false
}

func (l *Level) UnparkVehicle(spot *ParkingSpot, vehicle *vehicle.Vehicle) {
	for _, spot := range l.parkingSpots {
		if spot.closed == false && spot.GetVehicleId() == vehicle.GetId() {
			spot.UnparkVehicle(vehicle)
		}
	}
}

func (l *Level) isFloorFull() bool {
	if l.full == true {
		return true
	}
	return false
}

func (l *Level) isFloorActive() bool {
	if l.active == true {
		return true
	}
	return false
}
