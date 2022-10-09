package environment

type Resource struct {
	Memory      int64 `json:"memory"`
	Swap        int64 `json:"swap"`
	IoWeight    int64 `json:"io_weight"`
	CpuLimit    int64 `json:"cpu_limit"`
	Threads     int64 `json:"threads"`
	OomDisabled bool  `json:"oom_disabled"`
}
