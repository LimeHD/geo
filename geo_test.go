package geo

import (
	"path/filepath"
	"testing"
)

func TestGeoIp(t *testing.T) {
	CountryIpList := []struct {
		ip string
		is int
	}{
		{"87.226.37.80", BALTIA_ID},
		{"82.200.232.150", SNG_ID},
		{"217.195.59.22", BALTIA_ID},
		{"84.15.143.111", BALTIA_ID},
		{"89.218.110.222", SNG_ID},
		{"193.193.240.36", SNG_ID},
		{"95.141.132.138", SNG_ID},
		{"89.218.5.109", SNG_ID},
		{"115.42.64.226", WORLD_ID},
		{"190.109.167.9", WORLD_ID},
		{"195.46.20.146", WORLD_ID},
		{"91.197.174.108", RU_ID},
		{"78.140.7.239", RU_ID},
		{"91.187.75.48", WORLD_ID},
		{"104.248.200.184", WORLD_ID},
		{"80.245.117.131", RU_ID},
		{"80.245.117.130", RU_ID},
		{"91.243.36.246", RU_ID},
		{"80.94.229.172", SNG_ID},
		{"86.57.219.183", SNG_ID},
		{"91.215.176.237", SNG_ID},
	}

	abs, err := filepath.Abs("geoip/GeoLite2-City.mmdb")

	if err != nil {
		t.Errorf("Error, can't find database file in %s", abs)
	}

	reader := &Reader{}

	if err = reader.OpenDatabase(abs); err != nil {
		t.Errorf("Error, can't open database from %s, error is: %v", abs, err)
	}

	for _, country := range CountryIpList {
		id, name := reader.GetRegion(country.ip)

		if id != country.is {
			t.Errorf(
				"Crash happened, the country is not who it claims to be. Region (%s): %s [%d] is not [%d] \n",
				country.ip, name, id, country.is)
		}
	}
}
