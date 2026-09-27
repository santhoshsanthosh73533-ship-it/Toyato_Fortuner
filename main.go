package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"cat--1/carrental"
)

func main() {
	manager := carrental.NewFleetManager()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=========================================")
		fmt.Println("   CAR RENTAL FLEET MANAGEMENT SYSTEM   ")
		fmt.Println("=========================================")
		fmt.Println("1. Add Vehicle")
		fmt.Println("2. Lookup Vehicle")
		fmt.Println("3. Update Vehicle Mileage")
		fmt.Println("4. Decommission Vehicle")
		fmt.Println("5. Exit")
		fmt.Println("=========================================")
		fmt.Print("Enter your choice (1-5): ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Println("\n--- Add New Vehicle ---")
			fmt.Print("Enter License Plate Number: ")
			scanner.Scan()
			plate := strings.TrimSpace(scanner.Text())

			fmt.Print("Enter Vehicle Model: ")
			scanner.Scan()
			model := strings.TrimSpace(scanner.Text())

			fmt.Print("Enter Daily Rate (e.g., 2500.50): ")
			scanner.Scan()
			rateStr := strings.TrimSpace(scanner.Text())
			rate, err := strconv.ParseFloat(rateStr, 64)
			if err != nil {
				fmt.Println("Invalid rate format!")
				continue
			}

			fmt.Print("Is Available? (true/false): ")
			scanner.Scan()
			availStr := strings.TrimSpace(scanner.Text())
			available, err := strconv.ParseBool(availStr)
			if err != nil {
				fmt.Println("Invalid boolean input! Enter 'true' or 'false'.")
				continue
			}

			fmt.Print("Enter Initial Mileage (km): ")
			scanner.Scan()
			mileageStr := strings.TrimSpace(scanner.Text())
			mileage, err := strconv.Atoi(mileageStr)
			if err != nil {
				fmt.Println("Invalid mileage format!")
				continue
			}

			newVehicle := carrental.Vehicle{
				PlateNumber: plate,
				Model:       model,
				DailyRate:   rate,
				IsAvailable: available,
				Mileage:     mileage,
			}

			err = manager.AddVehicle(newVehicle)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("SUCCESS: Vehicle added to the fleet!")
			}

		case "2":
			fmt.Println("\n--- Lookup Vehicle ---")
			fmt.Print("Enter License Plate Number: ")
			scanner.Scan()
			plate := strings.TrimSpace(scanner.Text())

			vehicle, err := manager.LookupVehicle(plate)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("\n--- Vehicle Details ---")
				fmt.Printf("Plate Number : %s\n", vehicle.PlateNumber)
				fmt.Printf("Model        : %s\n", vehicle.Model)
				fmt.Printf("Daily Rate   : ₹%.2f\n", vehicle.DailyRate)
				fmt.Printf("Available    : %t\n", vehicle.IsAvailable)
				fmt.Printf("Mileage      : %d km\n", vehicle.Mileage)
			}

		case "3":
			fmt.Println("\n--- Update Mileage ---")
			fmt.Print("Enter License Plate Number: ")
			scanner.Scan()
			plate := strings.TrimSpace(scanner.Text())

			fmt.Print("Enter Additional Kilometers Driven: ")
			scanner.Scan()
			addedKmStr := strings.TrimSpace(scanner.Text())
			addedKm, err := strconv.Atoi(addedKmStr)
			if err != nil {
				fmt.Println("Invalid kilometer format!")
				continue
			}

			err = manager.UpdateMileage(plate, addedKm)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("SUCCESS: Mileage updated successfully!")
			}

		case "4":
			fmt.Println("\n--- Decommission Vehicle ---")
			fmt.Print("Enter License Plate Number to Remove: ")
			scanner.Scan()
			plate := strings.TrimSpace(scanner.Text())

			err := manager.DecommissionVehicle(plate)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("SUCCESS: Vehicle removed from rental pool!")
			}

		case "5":
			fmt.Println("\nExiting program... Goodbye!")
			return

		default:
			fmt.Println("\nInvalid choice! Please select an option between 1 and 5.")
		}
	}
}