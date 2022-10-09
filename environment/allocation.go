package environment

type Allocation struct {
	Default   bool   `json:"default"`
	IpAddress string `json:"ip_address"`
	Port      int    `json:"port"`
}
