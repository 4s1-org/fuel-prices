[![Pipeline](https://gitlab.com/4s1/fuel-prices/badges/main/pipeline.svg)](https://gitlab.com/4s1/fuel-prices/pipelines)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=fuel-prices&metric=bugs)](https://sonarcloud.io/summary/overall?id=fuel-prices)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=fuel-prices&metric=code_smells)](https://sonarcloud.io/summary/overall?id=fuel-prices)

# Fuel Prices

Lädt die aktuellen Kraftstoffpreise von [Tankerkönig](https://www.tankerkoenig.de/) herunterladen und extrahiert die Preise von Tankstellen in CSV-Dateien.

## Starten

```bash
# Bauen
make

# Konfiguration erzeugen
go run . --init

# config.json ausfüllen

# Anwendung erneut starten
go run .
```
