# Gator
gator is an RSS aggregator and program built in Go.

## Installation
You first need to have Postgres and Go installed in order to run
this program.
Install with:
```bash
go install github.com/DavidMWeaver4/rssProject@latest
```
`go install` installs `gator` to `~/go/bin` (or `$GOBIN`). Ensure that directory is on your `PATH`.

## Config
Create the config file at ~/.gatorconfig.json
example:
```json
{
  "db_url": "postgres://username@localhost:5432/gator?sslmode=disable"
}
```
Depending on your Postgres setup, you may need to include a password in the URL.
example:
```json
{
  "db_url": "postgres://username:password123@localhost:5432/gator?sslmode=disable"
}
```

## Commands
```bash
gator register <name>       #registers a new username and sets as current user

gator login <name>          #logs in to a username

gator addfeed <name> <url>  #adds an RSS feed to the database

gator agg 30s               #aggregates the list of feeds at a specific time

gator browse 10             #browses the most recent published posts, can        
                            #be changed to as many as you like
```
