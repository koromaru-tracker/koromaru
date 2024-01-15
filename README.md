# Koromaru

Koromaru is an open source BitTorrent Indexer written in Go and uses chihaya as tracker.

Features include:

- User registration
- User login
- User Torrent Pass
- Torrent Upload

## Why Koromaru?

koromaru is a private BitTorrent tracker for the chihaya Tracker. It is written in Go and uses SQLite as a database. It is designed to be light, fast, secure and easy to set up.

## Development

### Getting Started

In order to compile the project, the latest stable version of Go and knowledge of a working Go environment are required.

```bash
git clone git@github.com:koromaru-tracker/koromaru.git
cd koromaru
go build ./cmd/koromaru
./koromaru --help
```
