#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
hent_data.py
Henter værvarsel for de neste ni dagene for en gitt bredde- og lengdegrad fra Meterologisk institutt
Lagrer temperaturene for hver time i data.csv i formatet tid,temperatur.
Se dokumentasjon: https://api.met.no/weatherapi/locationforecast/2.0/documentation
"""

import csv
from pathlib import Path

import requests

# ---------------- Konfigurasjon ----------------
LAT = 59.55   # Breddegrad (Mysen)
LON = 11.33   # Lengdegrad
OUTPUT_DIR = Path(__file__).resolve().parent.parent / "data"
OUTFILE = OUTPUT_DIR / "data.csv"
HEADERS = {"User-Agent": "smidig-it2-etl-eksempel"} # MET ber om User-Agent som identifiserer programmet høflig

def lag_url(lat, lon):
    """Lager URL til MET sitt Locationforecast-API."""
    return (
        "https://api.met.no/weatherapi/locationforecast/2.0/compact"
        f"?lat={lat}&lon={lon}"
    )


def hent_data(lat, lon):
    """
    Henter JSON-data fra MET.
    raise_for_status() sørger for at vi får feilmelding ved 4xx/5xx.
    """
    url = lag_url(lat, lon)
    print(f"Henter {url}")
    r = requests.get(url, headers=HEADERS, timeout=30)
    r.raise_for_status()
    return r.json()


def main():
    # Opprett data-mappen hvis den ikke finnes
    OUTPUT_DIR.mkdir(parents=True, exist_ok=True)

    # Hent JSON fra MET
    data = hent_data(LAT, LON)

    # Timeserien ligger i properties.timeseries
    timeserier = data["properties"]["timeseries"]

    rader = []  # liste med (tid, temperatur)

    for punkt in timeserier:
        tid = punkt["time"]  # ISO-tidspunkt, f.eks. "2025-11-20T20:00:00Z"

        detaljer = punkt["data"]["instant"]["details"]
        temp = detaljer.get("air_temperature")

        # Hopper over timer som mangler data
        if temp is None:
            continue

        rader.append((tid, temp))

    # Skriv CSV-fil
    # newline="" skrur av automatisk linjeskift-konvertering i Python
    # lineterminator="\n" tvinger LF på alle plattformer for konsistens
    with OUTFILE.open("w", newline="", encoding="utf-8") as f:
        writer = csv.writer(f, lineterminator="\n")
        writer.writerow(["tid", "temperatur"])
        writer.writerows(rader)

    print(f"Skrev {len(rader)} rader til {OUTFILE}")
    print(f"Koordinater: lat={LAT}, lon={LON}")


if __name__ == "__main__":
    main()

# Smidig IT-2 © TIP AS, 2025
