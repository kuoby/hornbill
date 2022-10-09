package environment

type Configuration struct {
	envVars     []string
	allocations []Allocation
	resource    Resource
}
