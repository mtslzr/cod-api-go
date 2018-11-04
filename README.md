# Call of Duty Tracker API (Go)

Wrapper for [Call of Duty Tracker's API](https://callofdutytracker.com/site-api), written in Go.

## Installation

```bash
go get github.com/mtslzr/cod-api/go
```

## Usage
```go
import cod "cod-api-go"

client, err := cod.New("GAME", "PLATFORM", "USERNAME")
```

Replace the following placeholders:
* GAME: bo3, bo4, wwii, iw
* PLATFORM: bnet, psn, steam, xbl
* USERNAME: Platform-specific username (include #1111 tag for Battle.net)

See [API Documentation](https://callofdutytracker.com/site-api) for more information on variables.

## Endpoints

### Validate User

```go
data, err := client.ValidateUser()
```

### User Stats

```go
data, err := client.GetUserStats("MATCHTYPE")
```

_MATCHTYPE can be "mp" or "blackout" for specific match data. Or pass "" for default (mp)._

### Recent Matches

```go
data, err := client.GetRecentMatches(NUMMATCHES)
```

_NUMMATCHES is a number of matches between 1 and 100._

### Specific Match

```go
data, err := client.GetSpecificMatch(MATCHID)
```

_MATCHID is a specific Match ID to search._

### Users

```go
data, err := client.GetUserNames("USERID1", "USERID2", "USERID3")
```

_USERID can be one (more more) User IDs to search; results will show user for each ID._

### Leaderboard

```go
data, err := client.GetLeaderboard("SCOPE", NUMPLAYERS)
```

_SCOPE is which leaderboard you'd like to return (kills, deaths, ekia, wins, losses, gamesplayed, timeplayed). NUMPLAYERS is the number of players between 1 and 100._

## To Do

* Add Unit Tests for each endpoint.
* Add validation for variables being passed to API struct and each endpoint.

## Credits

Based heavily on code from [pubg-go](https://github.com/albshin/go-pubg) and converted for this API.