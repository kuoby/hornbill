package environment

type Provisioner interface {
	Configure() error
	NewProcess() Process
}
