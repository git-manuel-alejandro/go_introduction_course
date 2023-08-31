package intro

import (
	"fmt"
)

type Vehicle interface {
	Distance() float64
}

const (
	CarVehicle        = "car"
	MotorcycleVehicle = "motorcycle"
	TruckVehicle      = "truck"
)

type Car struct {
	Time int
}
type Truck struct {
	Time int
}

type MotorcyVehicle struct {
	Time int
}

func NewVehicle(v string, time int) (Vehicle, error) {
	switch v {
	case CarVehicle:
		return &Car{Time: time}, nil

	case TruckVehicle:
		return &Truck{Time: time}, nil

	case MotorcycleVehicle:
		return &MotorcyVehicle{Time: time}, nil
	}

	return nil, fmt.Errorf("vehicle '%s' not exists", v)

}

func (c *Car) Distance() float64 {
	fmt.Println(TruckVehicle)

	return 100 * (float64(c.Time) / 60)
}

func (c *Truck) Distance() float64 {

	return 70 * (float64(c.Time) / 60)
}

func (c *MotorcyVehicle) Distance() float64 {

	return 120 * (float64(c.Time) / 60)
}
