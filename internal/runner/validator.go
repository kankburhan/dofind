package runner

import (
	"os"

	"github.com/kankburhan/dofind/pkg/errors"
)

func (o *Options) validate() {
	if o.Target == "" {
		errors.Exit("No target input provided.")
	}

	if o.Saveto != "" {
		f, e := os.OpenFile(o.Saveto,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if e != nil {
			errors.Exit(e.Error())
		}
		o.Output = f
	}
}
