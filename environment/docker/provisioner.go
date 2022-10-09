package docker

import (
	"github.com/kuoby/hornbill/environment"
)

type provisioner struct{}

func NewProvisioner() environment.Provisioner {
	return &provisioner{}
}

func (p *provisioner) Configure() error {
	return nil
}

func (p *provisioner) NewProcess() environment.Process {
	return &process{}
}
