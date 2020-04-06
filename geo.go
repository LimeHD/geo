package geo

import (
	"github.com/oschwald/geoip2-golang"
	"github.com/yl2chen/cidranger"
	"net"
	"strings"
)

const (
	DEF_ID = 40

	LIME_ID   = 10
	RU_ID     = 20
	UA_ID     = 30
	SNG_ID    = 35
	WORLD_ID  = 40
	BALTIA_ID = 50
)

const (
	DEF = "WORLD"

	LIME   = "LIME"
	RU     = "RU"
	UA     = "UA"
	AZ     = "AZ"
	AM     = "AM"
	BY     = "BY"
	KZ     = "KZ"
	KG     = "KG"
	MD     = "MD"
	TJ     = "TJ"
	UZ     = "UZ"
	TM     = "TM"
	GE     = "GE"
	LV     = "LV"
	LT     = "LT"
	EE     = "EE"
	SNG    = "СНГ"
	BALTIA = "Страны Балтии"

	// крым наш!
	crimeaISO  = "40"
	crimeaISO2 = "43"
)

type Reader struct {
	GeoReader *geoip2.Reader
}

func (reader *Reader) OpenDatabase(db string) error {
	var err error
	if reader.GeoReader, err = geoip2.Open(db); err != nil {
		return err
	}

	return nil
}

func (reader *Reader) CloseDatabase() error {
	return reader.GeoReader.Close()
}

func (reader *Reader) GetRegion(ipAddress string) (int, string) {
	if ipAddress == "" || strings.Index(ipAddress, "127.0.0.1") != -1 {
		return LIME_ID, LIME
	}

	ip := net.ParseIP(ipAddress)
	record, err := reader.GeoReader.City(ip)
	ranger := cidranger.NewPCTrieRanger()

	if err != nil {
		for _, network := range networks {
			if ipCIDRCheck(ranger, network, ipAddress) {
				return LIME_ID, LIME
			}
		}

		return RU_ID, RU
	}

	country := &Country{
		Id:          0,
		Ip:          "",
		City:        "",
		Subdivision: nil,
		Country:     "",
		Timezone:    "",
		IsoCode:     "",
		Coords:      nil,
		Net:         "",
	}

	country.Country = record.Country.Names["ru"]
	country.City = record.City.Names["ru"]
	country.Timezone = record.Location.TimeZone
	country.IsoCode = record.Country.IsoCode
	country.Coords = []float64{record.Location.Latitude, record.Location.Longitude}
	country.Subdivision = record.Subdivisions

	if country.IsoCode == RU {
		for _, network := range networks {
			if ipCIDRCheck(ranger, network, ipAddress) {
				return LIME_ID, LIME
			}
		}

		return RU_ID, RU
	}

	if country.IsoCode == UA {
		if record.Subdivisions[0].IsoCode == crimeaISO || record.Subdivisions[0].IsoCode == crimeaISO2 {
			return RU_ID, RU
		}

		return UA_ID, UA
	}

	if _, find := In(getSNGCodes(), record.Country.IsoCode); find == true {
		return SNG_ID, SNG
	}

	if _, find := In(getBaltiaCodes(), record.Country.IsoCode); find == true {
		return BALTIA_ID, BALTIA
	}

	return DEF_ID, DEF
}

func ipCIDRCheck(ranger cidranger.Ranger, network string, ipAddress string) bool {
	_, network1, _ := net.ParseCIDR(network)

	if err := ranger.Insert(cidranger.NewBasicRangerEntry(*network1)); err != nil {
		panic(err)
	}

	contains, err := ranger.Contains(net.ParseIP(ipAddress))

	if err != nil {
		panic(err)
	}

	return contains
}

func getSNGCodes() []string {
	return []string{AZ, AM, BY, KZ, KG, MD, TJ, UZ, TM, GE}
}

func getBaltiaCodes() []string {
	return []string{LT, EE, LV}
}
