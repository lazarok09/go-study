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
		Commands: []cli.Command{
			{
				Action: searchIPbyHost,
				Name:   "ip",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "host",
						Usage: "A DNS server to be lookedup ex: amazon.com",
						Value: "github.com",
					},
				},
				Usage: "Make network ip's search using the host",
			},
			{
				Action: searchServers,
				Name:   "server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "host",
						Usage:    "Send the DNS to be looked the servers ex: amazon.com",
						Value:    "github.com",
						Required: true,
					},
				},
				Usage: "Make server network search using the DNS",
			},
		},
	}

	if error := app.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}

func searchIPbyHost(cCtx *cli.Context) error {
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

func searchServers(cCtx *cli.Context) error {
	host := cCtx.String("host")
	results, error := net.LookupNS(host)

	if error != nil {
		fmt.Println(host)
		log.Fatal(error)
	}

	fmt.Printf("Success search of %d DNS counts \n", len(results))

	// print all ips
	for _, ip := range results {

		fmt.Println(ip.Host)
	}
	// print the host looked
	fmt.Printf("At host %s", host)

	return nil
}
