package cli

import (
	"flag"
	"os"
)

type stockAnalysisFlags struct {
	ConfigPath string
}

func New() *stockAnalysisFlags {
	f := &stockAnalysisFlags{}
	flagSet := flag.NewFlagSet("stock-analysis", flag.ExitOnError)

	flagSet.StringVar(
		&f.ConfigPath, "config-path", "config/config.yaml", "config file path")

	argsWithoutExecutes := os.Args[1:]

	if err := flagSet.Parse(argsWithoutExecutes); err != nil {
		panic(err)
	}

	return f
}
