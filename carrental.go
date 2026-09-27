package carrental

import (
	"errors"
	"fmt"
)

// Vehicle represents the car rental structure
type Vehicle struct {
	PlateNumber string
	Model       string
	DailyRate   float64
	IsAvailable bool
	Mileage     int
}

// FleetManager manages the collection of vehicles
type FleetManager struct {
	vehicles map[string]Vehicle
}

// NewFleetManager initializes a new FleetManager
func NewFleetManager() *FleetManager {
	return &FleetManager{
		vehicles: make(map[string]Vehicle),
	}
}

// AddVehicle adds a new car with a unique plate number
func (fm *FleetManager) AddVehicle(v Vehicle) error {
	if _, exists := fm.vehicles[v.PlateNumber]; exists {
		return fmt.Errorf("vehicle with plate number %s already exists", v.PlateNumber)
	}
	fm.vehicles[v.PlateNumber] = v
	return nil
}

// LookupVehicle searches for a vehicle by license plate
func (fm *FleetManager) LookupVehicle(plate string) (*Vehicle, error) {
	vehicle, exists := fm.vehicles[plate]
	if !exists {
		return nil, errors.New("vehicle not found")
	}
	return &vehicle, nil
}

// UpdateMileage increments vehicle mileage after a trip
func (fm *FleetManager) UpdateMileage(plate string, addedKm int) error {
	vehicle, exists := fm.vehicles[plate]
	if !exists {
		return errors.New("vehicle not found")
	}
	if addedKm < 0 {
		return errors.New("added mileage cannot be negative")
	}
	vehicle.Mileage += addedKm
	fm.vehicles[plate] = vehicle
	return nil
}

// DecommissionVehicle deletes a vehicle from the rental pool
func (fm *FleetManager) DecommissionVehicle(plate string) error {
	if _, exists := fm.vehicles[plate]; !exists {
		return errors.New("vehicle not found to decommission")
	}
	delete(fm.vehicles, plate)
	return nil
}