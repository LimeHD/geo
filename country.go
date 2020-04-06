package geo

var (
	networks = []string{
		"192.168.0.0/24",
		"127.0.0.1/32",
		"85.234.0.52/32",
		"172.20.0.2/32",
	}
)

type Country struct {
	Id          int
	Ip          string
	City        string `json:"city"`
	Subdivision interface{}
	Country     string `json:"country"`
	Timezone    string
	IsoCode     string
	Coords      []float64

	Net string `json:"net"`
}
