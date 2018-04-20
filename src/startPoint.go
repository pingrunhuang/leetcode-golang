package main

import (
	"flag"
	"fmt"

	"./golang_solution"
)

var solutionID = flag.Int("n", 1, "Please specify the number of the solution that you want to check [1-731]")

// calling the solution with flag -n [number]
func main() {
	flag.Parse()

	switch *solutionID {
	case 1:
		fmt.Printf("Solving the %dst challenge\n", *solutionID)
	case 2:
		fmt.Printf("Solving the %dnd challenge\n", *solutionID)
	case 3:
		fmt.Printf("Solving the %drd challenge\n", *solutionID)
	default:
		fmt.Printf("Solving the %dth challenge\n", *solutionID)
	}

	if *solutionID <= 0 || *solutionID > 741 {
		panic(fmt.Errorf("there is no existing solution for the specified number"))
	}

	if *solutionID == 1 {
		solutions.Launch1()
	}
	if *solutionID == 2 {
		solutions.Launch2()
	}

	if *solutionID == 3 {
		solutions.Launch3()
	}

	if *solutionID == 64 {
		solutions.Launch64()
	}

	if *solutionID == 741 {
		solutions.Launch741()
	}
}
