[![Pipeline](https://gitlab.com/4s1/fuel-prices/badges/main/pipeline.svg)](https://gitlab.com/4s1/fuel-prices/pipelines)
[![Coverage](https://gitlab.com/4s1/fuel-prices/badges/main/coverage.svg)](https://gitlab.com/4s1/fuel-prices/commits/main)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=4s1_fuel-prices&metric=bugs)](https://sonarcloud.io/project/issues?id=4s1_fuel-prices&resolved=false&types=BUG)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=4s1_fuel-prices&metric=vulnerabilities)](https://sonarcloud.io/project/issues?id=4s1_fuel-prices&resolved=false&types=VULNERABILITY)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=4s1_fuel-prices&metric=code_smells)](https://sonarcloud.io/project/issues?id=4s1_fuel-prices&resolved=false&types=CODE_SMELL)

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
