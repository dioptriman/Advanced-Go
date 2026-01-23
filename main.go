package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("Not Implemented")
	ErrTruckNotFound  = errors.New("Truck Not Found")
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id     string
	cargo  int
	Diesel float64
}

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 1
	t.Diesel -= 1
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	t.Diesel -= 1
	return nil
}

func (e *ElectricTruck) LoadCargo() error {
	e.cargo += 1
	e.battery -= 1
	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	e.cargo = 0
	e.battery -= 1
	return nil
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

// processTruck handles the loading and unloading of a truck
func processTruck(truck Truck) error {
	fmt.Printf("Processing truck : %+v \n", truck)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error Loading Cargo : %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("Error Unloading Cargo : %w", err)
	}

	return nil
}

func main() {
	truckID := 42
	anotherTruckID := &truckID

	log.Println(truckID)
	log.Printf("The value of the truckID : %v \n", *anotherTruckID)
	log.Printf("The address of the truckID : %v \n", &anotherTruckID)
}
