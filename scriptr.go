package main

import "flag"
import "fmt"

const (
	def_script = ""
	usage      = "path to script file"
)

var script string

func init() {

	flag.StringVar(&script, "script", def_script, usage)
	flag.StringVar(&script, "s", def_script, usage+" (shorthand)")

}

func main() {

	flag.Parse()

	if (def_script == "script") && (flag.Args() == nil) {
		fmt.Println("Please provide a stript")
	} else if def_script != "" {
		fmt.Println("script:", script)
	} else {
		fmt.Println("script", flag.Arg(0))
	}

}
