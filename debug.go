package loggoext

import (	
	"fmt"
	"github.com/juju/loggo"
)

func DebugLoggers(debugModules []string) {
	conf := loggo.Config{
		"": loggo.INFO,
	}
	for _, mod := range debugModules {
		if mod == "<root>":
			mod = ""
		conf[mod] = loggo.DEBUG
	}

	fmt.Printf("Setting logging level to debug for modules: %v", debugModules)
	err := loggo.ConfigureLoggers(conf.String())
	if err != nil {
		panic(err)
	}
}