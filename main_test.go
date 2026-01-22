package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("processTruck", func(t *testing.T) {
		t.Run("Should load and unload the cargo from the truck", func(t *testing.T) {
			nt := &NormalTruck{id: "1"}
			et := &ElectricTruck{id: "2"}

			err := processTruck(nt)
			if err != nil {
				t.Fatalf("Error processing truck: %s", err)
			}

			err = processTruck(et)
			if err != nil {
				t.Fatalf("Error processing truck: %s", err)
			}

			//Assertion
			if nt.cargo != 0 {
				t.Fatal("The cargo should be 0")
			}

			if nt.Diesel != -2 {
				t.Fatal("The diesel should be -2")
			}

			if et.battery != -2 {
				t.Fatal("The battery should be -2")
			}
		})
	})
}
