package environment

var _ Provisioner = (*Docker)(nil)

type Docker struct{}

func NewDocker() *Docker {
	return &Docker{}
}

func (d *Docker) Configure() error {
	return nil
}

func (d *Docker) New() Instance {
	return nil
}
