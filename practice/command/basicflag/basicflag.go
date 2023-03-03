package main

import (
	"flag"
	"log"
	"os"
)

// Bool defines a bool flag with specified name, default value, and usage string. The return value is the address of a bool variable that stores the value of the flag.
var (
	useProd = flag.Bool("prod", true, "Use a production endpoint")
	useDev  = flag.Bool("dev", false, "Use a development endpoint")
	help    = flag.Bool("help", false, "Display help text")
)

func main() {
	flag.Parse()
	if *help {
		flag.PrintDefaults()
		return
	}
	switch {
	case *useProd && *useDev:
		log.Println("Error: --prod and --dev cannot both be set")
		flag.PrintDefaults()
		os.Exit(1)
	case !(*useProd || *useDev):
		log.Println("Error: either --prod or --dev must be set")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
