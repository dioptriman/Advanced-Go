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
	id    string
	cargo int
}

func (t NormalTruck) LoadCargo() error {
	return nil
}

func (t NormalTruck) UnloadCargo() error {
	return nil
}

func (e ElectricTruck) LoadCargo() error {
	return nil
}

func (e ElectricTruck) UnloadCargo() error {
	return nil
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

// processTruck handles the loading and unloading of a truck
func processTruck(truck Truck) error {
	fmt.Printf("Processing truck : %v", truck)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error Loading Cargo : %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("Error Unloading Cargo : %w", err)
	}

	return nil
}

func main() {
	err := processTruck(NormalTruck{id: "1"})
	if err != nil {
		log.Fatalf("Error processing truck: %s", err)
	}

	err = processTruck(ElectricTruck{id: "2"})
	if err != nil {
		log.Fatalf("Error processing truck: %s", err)
	}
}
