package parkinglot

import (
	"parkinglotgo/vehicle"
)

type ParkingSpot struct {
	spot        int
	vehicleType vehicle.VehicleType
	vehicle     *vehicle.Vehicle
	closed      bool
}

var spot *ParkingSpot

func CreateParkingSpot(spotNumber int, vehicleType vehicle.VehicleType) *ParkingSpot {
	spot = &ParkingSpot{
		spot:        spotNumber,
		vehicleType: vehicle.Bus,
		closed:      false,
	}
	return spot
}
func (s *ParkingSpot) ParkVehicle(vehicle *vehicle.Vehicle) {
	s.vehicle = vehicle
}

func (s *ParkingSpot) UnparkVehicle(vehicle *vehicle.Vehicle) {
	s.vehicle = nil
}

func (s *ParkingSpot) isSpotAvailableForParking() bool {
	if s.closed == false && s.vehicle == nil {
		return true
	}
	return false
}

func (s *ParkingSpot) GetVehicleType() string {
	return s.vehicleType.GetVehicleTypeName()
}

func (s *ParkingSpot) GetVehicleId() string {
	return s.vehicle.GetId()
}
