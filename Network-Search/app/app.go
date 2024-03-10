package app

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func App() {
	app := &cli.App{
		Action: SearchIPbyHost,
		Name:   "ip",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Usage: "A DNS server to be lookedup ex: amazon.com",
				Value: "github.com",
			},
		},
		Usage: "Make network searchs using net package",
	}

	if error := app.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}

func SearchIPbyHost(cCtx *cli.Context) error {
	host := cCtx.String("host")
	results, error := net.LookupHost(host)

	if error != nil {
		fmt.Println(host)
		log.Fatal(error)
	}

	fmt.Printf("Success search of %d ip counts \n", len(results))

	// print all ips
	for _, ip := range results {
		fmt.Println(ip)
		fmt.Println(ip)
	}
	// print the host looked
	fmt.Printf("At host %s", host)

	return nil
}
