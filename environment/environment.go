package environment

type Provisioner interface {
	Configure() error
	New() Instance
}

type Instance interface{}
