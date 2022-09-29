package initialize

import (
	flag "github.com/spf13/pflag"
)

func RetrieveContextIfChanged(contextFlag *flag.Flag) string {
	if contextFlag != nil && contextFlag.Changed {
		return contextFlag.Value.String()
	}
	return ""
}
