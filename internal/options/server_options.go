package options

import (
	"crypto/tls"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"

	"github.com/gaulzhw/go-server/internal/server"
)

type ServerOptions struct {
	Address string
	CertDir string
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{}
}

func (o *ServerOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Address, "address", ":8080",
		"Address of go-server to bind on. Default to :8080.")
	fs.StringVar(&o.CertDir, "cert-dir", "",
		"The directory of cert to use, use tls.crt and tls.key as certificates. Default to disable https.")
}

func (o *ServerOptions) Validate() []error {
	var errs []error

	return errs
}

func (o *ServerOptions) NewServer() (*server.Server, error) {
	var tlsConfig *tls.Config = nil
	if o.CertDir != "" {
		cert, err := tls.LoadX509KeyPair(filepath.Join(o.CertDir, "tls.crt"), filepath.Join(o.CertDir, "tls.key"))
		if err != nil {
			return nil, errors.WithMessage(err, "unable to load tls certificate")
		}
		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
	}

	return server.NewServer(o.Address, tlsConfig), nil
}
