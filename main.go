package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
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

	time.Sleep(time.Second)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error Loading Cargo : %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("Error Unloading Cargo : %w", err)
	}

	return nil
}

func processFleet(trucks []Truck) error {
	var wg sync.WaitGroup
	for _, t := range trucks {
		wg.Add(1)

		go func(t Truck) {
			if err := processTruck(t); err != nil {
				log.Println(err)
			}
			wg.Done()
		}(t)
	}

	wg.Wait()

	return nil
}

func main() {
	fleet := []Truck{
		&ElectricTruck{id: "ET-1", cargo: 10, battery: 100},
		&NormalTruck{id: "NT-1", cargo: 0, Diesel: 50},
		&ElectricTruck{id: "ET-2", cargo: 30, battery: 30},
		&NormalTruck{id: "NT-3", cargo: 40, Diesel: 20},
	}

	if err := processFleet(fleet); err != nil {
		fmt.Printf("Error Processing fleet : %v \n", err)

		return
	}

	fmt.Println("All trucks processed successfully")
}
