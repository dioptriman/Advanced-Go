package main

import (
	"fmt"
	"log"
)

type Truck struct {
	id string
}

// processTruck handles the loading and unloading of a truck
func processTruck(truck Truck) error {
	fmt.Printf("Processing truck: %s\n", truck.id)

	return nil
}

func main() {
	trucks := []Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived \n", truck.id)
		err := processTruck(truck)
		if err != nil {
			log.Fatalf("Error processing truck: %s", err)
		}
	}
}
