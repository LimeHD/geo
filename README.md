# Geo service package

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