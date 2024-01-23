# How To Use
## Frontend
- Keine Grosse Konfiguration nötig; Läuft auf Port `3000`.
- Zum Starten: 
  - Node.js installieren, Repo clonen
  - npm - Packages via `npm i` installieren.
  - Frontend - Server mit `npm run dev` starten
  - (Unter http://localhost:3000 im Web aufrufbar)

## Backend
- Läuft auf Port `8080`

# Bürokratisches
## Loggingkonzept

## Speicherung der Passwörter

## Verwendete Bibliotheken
### Frontend
- `nuxt v3`: Da in diesem Projekt mit dynamischen Daten (Posts, Comments) gearbeitet wird, eignet sich ein Webframework sehr - Da die Daten ansonsten "von Hand" generiert werden müssten ;-) -> Beispiel: `v-for` - Loop, welcher Daten dynamisch generiert
- `nuxt-security`: Ein Security-Package, welches die meisten Security-Schwachstellen abdeckt (XSS, iFrame Options (Clickjacking), etc)

### Backend

## Schutz vor XSS
XSS kann relativ einfach vorgebeugt werden; Nuxt "sanitized" sämtliche Inputs, bevor sie für eine Request verwendet werden.
`nuxt-security` scannt via `X-XSS` Header die Seite und verhindert das Laden, wenn schädliche Inhalte erkennt werden. Ein alternativer Modus des Headers "sanitized" (Entfernt alle schädlichen Inhalte) die Seite, bevor sie geladen wird.