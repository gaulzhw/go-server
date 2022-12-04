package options

import (
	"github.com/spf13/pflag"
)

type MySQLOptions struct {
}

var _ options = (*MySQLOptions)(nil)

func NewMySQLOptions() *MySQLOptions {
	return &MySQLOptions{}
}

func (o *MySQLOptions) AddFlags(fs *pflag.FlagSet) {
}

func (o *MySQLOptions) Validate() []error {
	var errs []error
	return errs
}
