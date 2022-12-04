package options

import (
	"github.com/spf13/pflag"
)

type options interface {
	AddFlags(set *pflag.FlagSet)
	Validate() []error
}
