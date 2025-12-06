# Temperaturer ETL

Dette prosjektet er et lite datasystem som viser hvordan tre utviklingsmiljøer kan samarbeide gjennom Docker. Systemet henter temperaturdata, behandler dem og viser resultatet i en nettleser.

## Oversikt

Prosjektet består av tre tjenester:

1. **Datainnsamler (Python)**  
   Henter temperaturprognoser fra Meteorologisk institutt og lagrer dem i data/data.csv.

2. **Databehandler (Go)**  
   Leser CSV-filen, beregner gjennomsnittstemperatur per dag og lagrer resultatet i data/data.json. Bruker dummy.csv hvis data.csv ikke finnes.

3. **Web (nginx + HTML + JavaScript)**  
   Viser en graf basert på data.json i nettleseren. Bruker dummy.json hvis data.json ikke finnes.

Alle tre tjenestene deler den samme datamappen via volumes i compose.yaml.

## Kjøring med Docker

Klon repoet:

    git clone https://github.com/smidig-it2/temperaturer-etl.git
    cd temperaturer-etl

Start alle tjenestene:

    docker compose up

Nettsiden er tilgjengelig på:

    http://localhost:8080

Trykk Ctrl+C for å stoppe.

Bygg på nytt hvis kode eller Dockerfile er endret:

    docker compose up --build

Rydd opp:

    docker compose down --remove-orphans

## Lokalt utviklingsmiljø (uten Docker)

Dette kan brukes hvis du vil teste hver del før du bygger Docker-images.

### Python (datainnsamler)

    cd datainnsamler
    python hent_data.py

### Go (databehandler)

    cd databehandler
    go run .

### Web (HTML og JavaScript)

Høyreklikk på <code>index.html</code> og velg Open withLive Server hvis du har denne utvidelsen installert i VS Code. Alternativt, start en enkel webserver i rotmappen med

    python -m http.server 8000

Åpne i nettleseren

    http://localhost:8000/web/index.html
    

## Mappestruktur

    temperaturer-etl/
    │
    ├─ datainnsamler/      # Python-kode og Dockerfile
    ├─ databehandler/      # Go-kode og Dockerfile
    ├─ web/                # HTML, JavaScript og Dockerfile
    ├─ data/               # data.csv, data.json, dummy.csv, dummy.json
    └─ compose.yaml

## Krav

- Docker Desktop
- docker compose
- (Valgfritt) Python 3.13 og Go 1.23 for lokal testing

## Lisens

Dette prosjektet er lisensiert under MIT-lisensen. Se LICENSE-filen for mer informasjon.
