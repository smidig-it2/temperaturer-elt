# Bruk et lite Python-image
FROM python:3.13-slim

# Sett arbeidsmappe i containeren
WORKDIR /app

# Installer Python-avhengigheter
# Vi kopierer først requirements.txt for bedre caching
COPY requirements.txt .
RUN pip install --no-cache-dir --upgrade pip \
    && pip install --no-cache-dir -r requirements.txt

# Kopier resten av prosjektet inn i containeren
COPY . .

# Start Python-programmet når containeren kjøres
CMD ["python", "hent_data.py"]

# Smidig IT-2 © TIP AS, 2025