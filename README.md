# 🚀 Commandarr

**Commandarr** is an open-source Telegram Bot built in Go to seamlessly manage and automate media downloads with Sonarr (TV series) and Radarr (movies). It's secure, extensible, and easy to deploy using Docker.

---

## 🌟 Current Features

- ✅ **Telegram Integration**: Bot setup and basic commands.
- ✅ **Sonarr Integration**:
  - Search and add TV series.
  - Immediate search and download trigger upon addition.
- ✅ **Radarr Integration**:
  - Search and add movies.
  - Immediate search and download trigger upon addition.

---

## 🚩 Available Commands

- `/search_movie <movie-name>`: Search for a movie.
- `/add_movie <movie-name>`: Add a movie to Radarr.
- `/search_series <series-name>`: Search for a series.
- `/add_series <series-name>`: Add a series to Sonarr.

---

## 🛠️ Quick Start

### Docker Compose

```yaml
version: '3.8'

services:
  commandarr:
    image: yourdockerhubusername/commandarr:latest
    container_name: commandarr
    environment:
      TELEGRAM_TOKEN: "<your-telegram-token>"
      SONARR_URL: "http://your-sonarr-url:8989"
      SONARR_API_KEY: "<sonarr-api-key>"
      RADARR_URL: "http://your-radarr-url:7878"
      RADARR_API_KEY: "<radarr-api-key>"
```

Run Commandarr:

```bash
docker-compose up -d
```

### Docker Run

```bash
docker run -d \
  --name commandarr \
  -e TELEGRAM_TOKEN="<your-telegram-token>" \
  -e SONARR_URL="http://your-sonarr-url:8989" \
  -e SONARR_API_KEY="<sonarr-api-key>" \
  -e RADARR_URL="http://your-radarr-url:7878" \
  -e RADARR_API_KEY="<radarr-api-key>" \
  yourdockerhubusername/commandarr:latest
```

---

## 🔐 Configuration

Set the following environment variables:

| Variable          | Description                          | Example                                     |
|-------------------|--------------------------------------|---------------------------------------------|
| `TELEGRAM_TOKEN`  | Telegram Bot API Token               | `123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11` |
| `SONARR_API_KEY`  | Sonarr API key                      | `abcdef123456789`                           |
| `RADARR_API_KEY`  | Radarr API key                      | `abcdef123456789`                           |
| `SONARR_URL`      | URL of your Sonarr instance          | `http://localhost:8989`                     |
| `RADARR_URL`      | Radarr URL                          | `http://localhost:7878`                     |

---

## 🚩 Next Steps

- 🔔 Telegram notifications for downloads.
- 📊 Structured logging and error handling.
- 📚 Complete documentation with PlantUML diagrams.
- 🐳 Docker Hub publication for community use.

---

## 🤝 Contributing

We welcome contributions! Check [CONTRIBUTING.md](CONTRIBUTING.md) for guidance.

---

## 📜 License

Commandarr is licensed under the [Apache-2.0 License](LICENSE).