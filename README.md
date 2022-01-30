# Fuel Prices

Lädt die aktuellen Kraftstoffpreise von [Tankerkönig](https://www.tankerkoenig.de/) herunterladen und extrahiert die Preise von Tankstellen in CSV-Dateien.

## Starten

```bash
# Bauen
go build fuel-prices.go

# Konfiguration erzeugen
./fuel-prices --init

# config.json ausfüllen

# Anwendung erneut starten
./fuel-prices
```

## Go Hilfen

```bash
# Projekt initalisieren
go mod init gitlab.com/4s1/fuel-prices

# Starten
go run fuel-prices.go

# Bauen
go build fuel-prices.go
#   oder
make build
```
