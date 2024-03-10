package app

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func App() {

	app := &cli.App{
		Name:   "multiplyall",
		Usage:  "make the multiplication between all numbers",
		Action: calculateMathPow,
		Flags: []cli.Flag{
			&cli.IntSliceFlag{
				Name:     "numbers",
				Usage:    "Send numbers to be multiplied at --numbers 10 --numbers 5 format //output 50",
				Required: true,
			},
		},
	}
	// this command will pass the arguments of command line to the package.
	// may return a error
	if error := app.Run(os.Args); error != nil {
		log.Fatal(error)
	}

}

// Operate a math calculation
func calculateMathPow(cCtx *cli.Context) error {
	numbers := cCtx.IntSlice("numbers")
	fmt.Println(numbers)

	var result int = numbers[0]

	for _, number := range numbers[1:] {
		result *= number
	}
	fmt.Println(result)

	return nil
}
