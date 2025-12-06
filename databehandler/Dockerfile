# Go-miljø for å bygge programmet
FROM golang:1.23-alpine

# Arbeidsmappe i containeren
WORKDIR /app

# Kopier mod-filer først for bedre cache
COPY go.mod ./
RUN go mod download

# Kopier resten av koden
COPY . .

# Bygg binæren
RUN go build -o databehandler

# Start programmet når containeren kjører
CMD ["./databehandler"]

# Smidig IT-2 © TIP AS, 2025
