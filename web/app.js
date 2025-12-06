// app.js
// Leser data/data.json og viser daglige gjennomsnitt i en stolpediagram

async function hentDataOgPlott() {
  try {
    // Prøv først data.json, fall tilbake til dummy.json
    let respons = await fetch('/data/data.json');
    if (!respons.ok) {
      console.warn('data.json ikke funnet, prøver dummy.json');
      respons = await fetch('/data/dummy.json');
      if (!respons.ok) {
        throw new Error('Klarte ikke å hente verken data.json eller dummy.json');
      }
    }

    const dager = await respons.json();
    // Forventer format [{ "dato": "2025-11-01", "gjennomsnitt": 3.4 }, ...]

    const etiketter = dager.map(d => d.dato);
    const verdier = dager.map(d => d.gjennomsnitt);

    const ctx = document.getElementById('tempChart').getContext('2d');

    new Chart(ctx, {
      type: 'bar',
      data: {
        labels: etiketter,
        datasets: [{
          label: 'Gjennomsnittstemperatur (°C)',
          data: verdier
        }]
      },
      options: {
        responsive: true,
        scales: {
          y: {
            beginAtZero: false,
            title: {
              display: true,
              text: 'Temperatur (°C)'
            }
          },
          x: {
            title: {
              display: true,
              text: 'Dato'
            }
          }
        }
      }
    });
  } catch (err) {
    console.error(err);
    alert('Kunne ikke laste temperaturdata. Sjekk at data.json finnes.');
  }
}

document.addEventListener('DOMContentLoaded', hentDataOgPlott);
