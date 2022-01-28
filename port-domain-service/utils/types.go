package utils

type JsonStatusResponse struct {
	Status string `json:"status"`
}

type JsonPortsResponse struct {
	Status string `json:"status"`
	Ports  []Port `json:"ports"`
}

type Port struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float32 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code,omitempty"`
}

type Entry struct {
	PortName string
	Port     Port
}