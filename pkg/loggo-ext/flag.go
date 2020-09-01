package loggoext

import (
	"flag"
	"strings"
)

var (
	debug string
)

func init(){
	usage := "--debug=module1,module2,... to set debug level for specific modules.\n --debug=<root> to set all loggers to debug.\n"
	flag.StringVar(&debug, "debug", "", usage)
	flag.StringVar(&debug, "d", "", usage+" (shorthand)")
}

func FlagConfigure(){
	if !flag.Parsed() {
		flag.Parse()
	}
	DebugLoggers(strings.Split(debug, ","))
}