## Requirements
-   go 1.8
-   posgres 9.6

or
-   docker-compose 1.13

## Before run
-   `GOPATH=<project folder>` - set **GOPATH** to project folder
-   `go get -u github.com/kardianos/govendor` - install [govendor](https://github.com/kardianos/govendor) for dependencies managment
-   `cd ./src/tournament_server/ && govendor sync` - go to source folder and download dependencies
-   edit settings `./config.json` if you are going to run app as usual
-   or create file `./tournament.env` with enviroment settings using `./tournament.example.env` as template if you are going to run app as containers
-   depends on way you are running app, create database and user
-   use `./db.sql` to create database tables

## Structure

- config.json
- db.sql
- docker-compose.yml
- tournament.example.env
-   src/
    -   tournament_server/
        - Dockerfile
        - server.go
        - config/
            - config.go
        - handlers/
            - messages.go
            - playerHandlers.go
            - playerStructs.go
            - tournamentHandlers.go
            - tournamentStructs.go
            - utils.go
        - models
            - joinEventModel.go
            - moneyTransactionModel.go
            - tournamentModel.go
        - vendor
            - vendor.json 

## Usage
- **/fund** - GET with *playerId* - string, required; *points* - uint64, required
- **/take** - GET with *playerId* - string, required; *points* - uint64, required
- **/balance** - GET with *playerId* - string, required
- **/announceTournament** - GET with *tournamentId* - string, required; *deposit* - uint64, required
- **/joinTournament** - GET with *tournamentId* - string, required; *playerId* - string, required; *backerId* list of strings
- **/resultTournament** - POST with JSON: 
```json
{
    "tournamentId": <string>,
    "winners": [
        {
            "playerId": <string>,
            "prize": <uint64>
        }
    ]
}
```
### Examples

