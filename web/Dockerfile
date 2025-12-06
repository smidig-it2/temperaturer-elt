# Enkel Nginx for statiske filer
FROM nginx:stable-alpine

# Arbeidsmappe for HTML og JavaScript
WORKDIR /usr/share/nginx/html

# Kopier nettsiden inn i containeren
COPY index.html .
COPY app.js .

# Nginx sitt standardoppsett holder til dette
# Containeren lytter på port 80

# Smidig IT-2 © TIP AS, 2025
