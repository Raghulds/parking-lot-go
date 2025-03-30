package vehicle

type VehicleType int

const (
	Bus VehicleType = iota
	Car
	TwoWheeler
)

type Vehicle struct {
	id string
	vehicleType VehicleType
}

func (v *Vehicle) GetId() string {
	return v.id
}

func (v *Vehicle) GetVehicleType() string {
	return v.vehicleType.GetVehicleTypeName()
}	

func (v VehicleType) GetVehicleTypeName() string {
	switch v {
	case Bus:
		return "Bus"
	case Car:
		return "Car"
	case TwoWheeler:
		return "TwoWheeler"
	default:
		return "Unknown"
	}
}

