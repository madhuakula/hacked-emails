package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/madhuakula/hacked-emails/api"
	"github.com/madhuakula/hacked-emails/version"
)

const (
	// BANNER is what is printed for help/info output.
	BANNER = `
	██╗  ██╗ █████╗  ██████╗██╗  ██╗███████╗██████╗       ███████╗███╗   ███╗ █████╗ ██╗██╗     ███████╗
	██║  ██║██╔══██╗██╔════╝██║ ██╔╝██╔════╝██╔══██╗      ██╔════╝████╗ ████║██╔══██╗██║██║     ██╔════╝
	███████║███████║██║     █████╔╝ █████╗  ██║  ██║█████╗█████╗  ██╔████╔██║███████║██║██║     ███████╗
	██╔══██║██╔══██║██║     ██╔═██╗ ██╔══╝  ██║  ██║╚════╝██╔══╝  ██║╚██╔╝██║██╔══██║██║██║     ╚════██║
	██║  ██║██║  ██║╚██████╗██║  ██╗███████╗██████╔╝      ███████╗██║ ╚═╝ ██║██║  ██║██║███████╗███████║
	╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚══════╝╚═════╝       ╚══════╝╚═╝     ╚═╝╚═╝  ╚═╝╚═╝╚══════╝╚══════╝
	 
Hacked Emails Command Line Tool
Version: %s
`
)

var (
	vrsn bool
)

func init() {
	// parse flags
	flag.BoolVar(&vrsn, "version", false, "print version and exit")
	flag.BoolVar(&vrsn, "v", false, "print version and exit (shorthand)")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(BANNER, version.VERSION))
		flag.PrintDefaults()
	}

	flag.Parse()

	if vrsn {
		fmt.Printf("hacked-emails version %s", version.VERSION)
		os.Exit(0)
	}

	if flag.NArg() < 1 {
		usageAndExit("Pass an email address: email@domain.com", 1)
	}

	// parse the arg
	arg := strings.Join(flag.Args(), " ")

	if arg == "help" {
		usageAndExit("", 0)
	}

	if arg == "version" {
		fmt.Printf("hacked-emails version %s", version.VERSION)
		os.Exit(0)
	}
}

func main() {
	email := strings.Join(flag.Args(), " ")

	response, err := api.Check(email)
	if err != nil {
		fmt.Printf("Decoding api response as JSON failed: %v", err)
		return
	}

	defResponse := fmt.Sprintf("%s email %s in %d results\n", response.Query, response.Status, response.Results)

	for _, def := range response.Data {
		defResponse += fmt.Sprintf("\n%s : Leaked Date (%s)\n <%s>\n", def.Title, def.Date_leaked, def.Details)
	}

	fmt.Println(defResponse)
}

func usageAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(exitCode)
}
