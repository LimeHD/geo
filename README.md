# Geo service package

### Installation

`go get github.com/LimeHD/geo`

### Example usage

```go
...

abs, err := filepath.Abs("geoip/GeoLite2-City.mmdb")
reader := &Reader{}

id, name := reader.GetRegion(ip)
...

```

### Tests

`go test`