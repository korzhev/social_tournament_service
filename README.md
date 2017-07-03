# Notes 03.07.2017
-   New database schema: money_transaction table should be updated. See `db.sql`.
-   If multiple parallel requests were made to update the same player balance, only first will be processed, rest ones will be rejected.

## Requirements
-   go 1.8
-   posgres 9.6

or
-   docker-compose 1.13

## Before run
-   `GOPATH=<project folder>` - add project folder to **GOPATH**
-   `go get -u github.com/kardianos/govendor` - install [govendor](https://github.com/kardianos/govendor) for dependencies managment
-   `cd ./src/tournament_server/ && govendor sync` - go to source folder and download dependencies
-   edit settings `./config.json` if you are going to run app as usual
-   or create file `./tournament.env` with enviroment settings using `./tournament.example.env` as template if you are going to run app as containers
-   depends on way you are running app, create database and user
-   use `./db.sql` to create database tables

## Structure

- config.json  - default config file for development
- db.sql - init sql for postgres
- docker-compose.yml - docker-compose file
- tournament.example.env - template for env_file in `docker-compose.yml`
-   src/ - source code
    -   tournament_server/ - package
        - Dockerfile - dockerfile for [hub.docker](https://hub.docker.com/r/korzhev/social_tournament_service/)
        - Dockerfile-alternative - dockerfile for docker-compose
        - server.go - app entry point
        - config/ - config package
            - config.go - get config from file or env
        - handlers/ - http handlers
            - messages.go - error messages
            - playerHandlers.go - handlers for player
            - playerStructs.go - structs for player
            - tournamentHandlers.go - handlers for tournament
            - tournamentStructs.go - structs for tournament
            - utils.go - common and util functions
        - models - models for orm
            - joinEventModel.go
            - moneyTransactionModel.go
            - tournamentModel.go
        - vendor - folder for vendor packages
            - vendor.json - vendor info, do not edit

## Usage

All responses are JSON

- **/fund** - GET with *playerId* - string, required; *points* - uint64, required
- **/take** - GET with *playerId* - string, required; *points* - uint64, required
- **/balance** - GET with *playerId* - string, required
- **/announceTournament** - GET with *tournamentId* - string, required; *deposit* - uint64, required
- **/joinTournament** - GET with *tournamentId* - string, required; *playerId* - string, required; *backerId* list of strings
- **/resultTournament** - POST with JSON: 
```json
{
    "tournamentId": "<string>",
    "winners": [
        {
            "playerId": "<string>",
            "prize": "<uint64>"
        }
    ]
}
```

## Notes

- Using [pg](https://github.com/go-pg/pg) as ORM
- Using [echo](https://echo.labstack.com/) as web framework
- May be there are better ORM and Web packages.
- There isn't any info about **PlayerID** or **TournamentID**, so they are strings.
- Also there isnt't any info about relationship between *Items* like **Tournament**, **Player**, etc., so there aren't player table or tournament description table.
- There isn't any info about how to store announced tournament and how to stop it or is it possible to have several tournaments at the same time. You should set **tournamentId** when getting result at /resultTournament
- No tests
- https://hub.docker.com/r/korzhev/social_tournament_service/