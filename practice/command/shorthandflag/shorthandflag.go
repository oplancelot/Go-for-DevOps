// flag.[Type]Var()
package main

import (
	"flag"
	"log"
	"os"
)

var (
	useProd = flag.Bool("prod", true, "Use a production endpoint")
	useDev  = flag.Bool("dev", false, "Use a development endpoint")
	help    = new(bool)
)

func init() {
	flag.BoolVar(help, "help", false, "Display help text")
	flag.BoolVar(help, "h", false, "Display help text (shorthand)")
}

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

// [root@tx01 ~/Go-for-DevOps/practice/command/shorthandflag]# go run shorthandflag.go -h
//   -dev
//         Use a development endpoint
//   -h    Display help text (shorthand)
//   -help
//         Display help text
//   -prod
//         Use a production endpoint (default true)
