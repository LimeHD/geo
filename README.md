# Geo service package

### Installation

`go get github.com/LimeHD/geo`

### Example usage

1. Как модуль в проекте:

```go
...

abs, err := filepath.Abs("geoip/GeoLite2-City.mmdb")
reader := &Reader{}

id, name := reader.GetRegion(ip)
...

```

2. Как консольная утилита:

in command line:

```
> geo -l /usr/local/var/geolite/GeoLite2-City.mmdb 192.123.43.12
20
```

### Tests

`go test`

### TODO

- [x] Возможность определение региона по IP
- [ ] Возможность определение региона по координатам

### Maintainers

<table>
<tr>
<td align="center">
<img src="https://avatars1.githubusercontent.com/u/23422968?s=460&u=668229465690637b50f6581df0fa9918d7fb6c1e&v=4" width="100px;" alt=""/>
<br /><sub><b>zikwall</b></sub></a><br />
</td>
</tr>
</table>
