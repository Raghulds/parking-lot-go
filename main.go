package main

import (
	"bufio"
	"fmt"
	"os"
	"parkinglotgo/parkinglot"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter number of floors for Parking lot: ")
	scanner.Scan()
	floors, err := strconv.Atoi(scanner.Text())
	for err != nil {
		fmt.Print("Not a valid number")
		fmt.Print("Enter number of floors for Parking lot: ")
		scanner.Scan()
		floors, err = strconv.Atoi(scanner.Text())
	}

	fmt.Print("Enter number of spots in each floor")
	scanner.Scan()
	slots, err := strconv.Atoi(scanner.Text())
	for slots < 3 {
		fmt.Print("Spots in each floor cannot be less than 3")
		fmt.Print("Enter number of spots in each floor: ")
		scanner.Scan()
		slots, err = strconv.Atoi(scanner.Text())
	}

	var lot = parkinglot.CreateParkingLot(floors, slots)
	fmt.Println(lot)
}
