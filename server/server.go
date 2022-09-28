package server

type Instance struct {
	fs *Filesystem
}

// NewInstance returns a new Instance.
func NewInstance() *Instance {
	return &Instance{}
}
