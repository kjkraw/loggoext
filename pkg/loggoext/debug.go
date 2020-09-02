package loggoext

import (	
	"github.com/juju/loggo"
)

func DebugLoggers(debugModules []string) {
	conf := loggo.Config{
		"<root>": loggo.INFO,
	}
	for _, mod := range debugModules {
		conf[mod] = loggo.DEBUG
	}

	err := loggo.ConfigureLoggers(conf.String())
	if err != nil {
		panic(err)
	}
}