package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/claerhead/go_blockchain/explorer"
	"github.com/claerhead/go_blockchain/rest"
)

func usage() {
	fmt.Printf("Welcome to 지선 코인\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("-port: Set the PORT of the server\n")
	fmt.Printf("-mode: Choose between 'html' and 'rest'\n\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")
	flag.Parse()
	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	case "all":
		diffPort := *port + 1000
		go explorer.Start(diffPort)
		rest.Start(*port)
	default:
		usage()
	}
}
