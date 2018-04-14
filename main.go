package main

import (
	"fmt"
	"github.com/Pattern-Recognition1/gomnist"
	"os"
)

func main() {

	//taking user input
	fmt.Println("Part 1 (type 1) or Part 2 (type 2)")
	var part int
	fmt.Scanf("%d", &part)

	switch part {
	case 1:
		{
			fmt.Println("executing part 1")
			fmt.Println("Enter digit followed by type (0 for mean, 1 for standard deviation) :")

			var digit int
			var typeOF int
			fmt.Scanf("%d %d", &digit,&typeOF)

			if digit < 0 || digit > 9 {
				fmt.Println("You entered invalid digit !")
				os.Exit(0)
			}

			gomnist.GetMeanAndSD(digit,typeOF)
		}
	case 2:
		{
			fmt.Println("executing part 2")
			gomnist.BayesianDecisionClassification()
		}
	default:
		fmt.Println("Invalid input")
	}

}
