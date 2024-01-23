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
- Zum Starten:
  -Go installieren, MongoDB Compass installieren (sicherstellen dass der Service gestartet ist)
  -Packages im Serverfolder via `go install` installieren
  -Starten mit `go run main.go`

## Benutzer
-Beim ersten Mal ausführen des Backends werden zwei Nutzer der Datenbank hinzugefügt. Weitere können mit der UI hinzugefügt werden. (Aus sicherheitsgründen kann man nur Nicht-Adminuser via. der UI erstellen)
-Credentials: 
  -Username: `Manfred`
  -Passwort: `Manfred123.`
  -Admin: `true`

  -Username: `Michael`
  -Passwort: `Admin123.`
  -Admin: `false`

Da bei diesen Usern keine Telefonnummer hinterlegt wurde, kann man sich Ohne SMS-Authentifizierung anmelden.

# Bürokratisches
## Loggingkonzept
Unter `server\log\apilog.log` wird ein Logfile stetig aktualisiert. Darin halten wir Datum und Uhrzeit fest, wenn sich ein User anmeldet, abmeldet oder sich Einträge in der Datenbank verändern oder neue erstellt werden.

## Speicherung der Passwörter
Bevor die Passwörter in der Datenbank gespeichert werden, werden diese gehasht. Dafür wurde die bcrypt Funktion der `crypto` Bibliothek verwendet. Um zusätzlich vor einem Rainbowtable Angriff zu schützen wurde ein salt hinzugefügt.

## Verwendete Bibliotheken
### Frontend
- `nuxt v3`: Da in diesem Projekt mit dynamischen Daten (Posts, Comments) gearbeitet wird, eignet sich ein Webframework sehr - Da die Daten ansonsten "von Hand" generiert werden müssten ;-) -> Beispiel: `v-for` - Loop, welcher Daten dynamisch generiert
- `nuxt-security`: Ein Security-Package, welches die meisten Security-Schwachstellen abdeckt (XSS, iFrame Options (Clickjacking), etc)

### Backend
- `gin-gonic`: Dieses http web framework wurde verwendet um den Webserver des Backends aufzusetzen. Die Routes des Backends laufen über gin.
- `mongo-driver`: Durch diesen Treiber wurde die Verbindung zwischen dem Backend und der MongoDB-Datenbank sichergestellt. Jeglicher Datenaustausch zwischen Backend und Datenbank funktionierte über dies Bibliothek.
- `crypto`: Mit der crypto Bibliothek wurden die Passwörte gehasht, bevor sie der Datenbank hinzugefügt wurden.

## Schutz vor XSS
XSS kann relativ einfach vorgebeugt werden; Nuxt "sanitized" sämtliche Inputs, bevor sie für eine Request verwendet werden.
`nuxt-security` scannt via `X-XSS` Header die Seite und verhindert das Laden, wenn schädliche Inhalte erkennt werden. Ein alternativer Modus des Headers "sanitized" (Entfernt alle schädlichen Inhalte) die Seite, bevor sie geladen wird.

## .env
Ursprünglich wollten wir aus sicherheitsgründen strings wie die `MONGO_URL` in einem .env speichern. Jedoch aufgrund von Zeitmangel und dadurch, dass jegliche verwendeten ENV's nur localhost strings waren, haben wir auf diesen Aufwand verzichtet.
