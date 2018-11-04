package cod

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// API returns stored user data for use in API calls.
type API struct {
	BaseURL  *url.URL
	Game     string
	Platform string
	UserName string
}

// Leaderboard returns a list of best users by platform and scope.
type Leaderboard struct {
	Rows     int    `json:"rows"`
	Platform string `json:"platform"`
	Scope    string `json:"scope"`
	Entries  []struct {
		UserName string `json:"username"`
		Platform string `json:"platform"`
		Level    struct {
			ID    int    `json:"id"`
			Image string `json:"image"`
		} `json:"level"`
		Prestige struct {
			ID    int    `json:"id"`
			Image string `json:"image"`
		} `json:"prestige"`
		Kills       int `json:"kills"`
		Deaths      int `json:"deaths"`
		EKIA        int `json:"ekia"`
		Wins        int `json:"wins"`
		Losses      int `json:"losses"`
		GamesPlayed int `json:"gamesplayed"`
		TimePlayed  int `json:"timeplayed"`
	} `json:"entries"`
}

// RecentMatches returns a list of recently-played matches.
type RecentMatches struct {
	Success  bool   `json:"success"`
	Rows     int    `json:"rows"`
	Game     string `json:"game"`
	Platform string `json:"platform"`
	Entries  []struct {
		MID       string `json:"mid"`
		UTCStart  int    `json:"utcStart"`
		UTCEnd    int    `json:"utcEnd"`
		MatchInfo struct {
			MatchDuration int    `json:"matchDuration"`
			MatchType     string `json:"matchType"`
			MatchMapID    string `json:"matchMapId"`
			MatchMode     string `json:"matchMode"`
		} `json:"matchInfo"`
		Teams struct {
			TeamScore struct {
				Team1 int `json:"team1"`
				Team2 int `json:"team2"`
			}
			WinningTeam int `json:"winningTeam"`
		} `json:"teams"`
		PlayerEntries []struct {
			UID               int `json:"uid"`
			Prestige          int `json:"prestige"`
			Rank              int `json:"rank"`
			Team              int `json:"team"`
			Position          int `json:"position"`
			Kills             int `json:"kills"`
			Deaths            int `json:"deaths"`
			EKIA              int `json:"ekia"`
			HighestKillStreak int `json:"highestkillstreak"`
			Assists           int `json:"assists"`
			Headshots         int `json:"headshots"`
			ShotsFired        int `json:"shotsfired"`
			ShotsLanded       int `json:"shotslanded"`
			ShotsMissed       int `json:"shotsmissed"`
		} `json:"playerEntries"`
	}
}

// SpecificMatch returns a single recently-played match.
type SpecificMatch struct {
	Success  bool   `json:"success"`
	Rows     int    `json:"rows"`
	Game     string `json:"game"`
	Platform string `json:"platform"`
	Entry    struct {
		MID       string `json:"mid"`
		UTCStart  int    `json:"utcStart"`
		UTCEnd    int    `json:"utcEnd"`
		MatchInfo struct {
			MatchDuration int    `json:"matchDuration"`
			MatchType     string `json:"matchType"`
			MatchMapID    string `json:"matchMapId"`
			MatchMode     string `json:"matchMode"`
		} `json:"matchInfo"`
		Teams struct {
			TeamScore struct {
				Team1 int `json:"team1"`
				Team2 int `json:"team2"`
			}
			WinningTeam int `json:"winningTeam"`
		} `json:"teams"`
		PlayerEntries []struct {
			UID               int `json:"uid"`
			Prestige          int `json:"prestige"`
			Rank              int `json:"rank"`
			Team              int `json:"team"`
			Position          int `json:"position"`
			Kills             int `json:"kills"`
			Deaths            int `json:"deaths"`
			EKIA              int `json:"ekia"`
			HighestKillStreak int `json:"highestkillstreak"`
			Assists           int `json:"assists"`
			Headshots         int `json:"headshots"`
			ShotsFired        int `json:"shotsfired"`
			ShotsLanded       int `json:"shotslanded"`
			ShotsMissed       int `json:"shotsmissed"`
		} `json:"playerEntries"`
	}
}

// UserIDToUserName returns a list of users.
type UserIDToUserName []struct {
	UID      int    `json:"uid"`
	UserName string `json:"username"`
	Platform string `json:"platform"`
	Game     string `json:"game"`
}

// UserStats returns a bunch of stats for a specific user.
type UserStats struct {
	Identifier string `json:"identifier"`
	Type       string `json:"type"`
	User       struct {
		ID       int    `json:"id"`
		UserName string `json:"username"`
		Platform string `json:"platform"`
		Title    string `json:"title"`
		Avatar   string `json:"avatar"`
	} `json:"user"`
	Cache struct {
		Time     int `json:"time"`
		Expire   int `json:"expire"`
		Interval int `json:"interval"`
	} `json:"cache"`
	Stats struct {
		Level             int    `json:"level"`
		MaxLevel          int    `json:"maxlevel"`
		Prestige          int    `json:"prestige"`
		PrestigeID        int    `json:"prestigeid"`
		MaxPrestige       int    `json:"maxprestige"`
		Kills             int    `json:"kills"`
		KillsConfirmed    int    `json:"killsconfirmed"`
		Deaths            int    `json:"deaths"`
		GamesPlayed       int    `json:"gamesplayed"`
		Wins              int    `json:"wins"`
		Losses            int    `json:"losses"`
		Melee             int    `json:"melee"`
		Hits              int    `json:"hits"`
		Misses            int    `json:"misses"`
		RankXP            int    `json:"rankxp"`
		CareerScore       int    `json:"careerscore"`
		TotalHeals        int    `json:"totalheals"`
		EKIA              int    `json:"ekia"`
		LongestKillStreak int    `json:"longestkillstreak"`
		CurWinStreak      int    `json:"curwinstreak"`
		TotalShots        int    `json:"totalshots"`
		TeamKills         int    `json:"teamkills"`
		Suicides          int    `json:"suicides"`
		OffEnds           int    `json:"offends"`
		KillsDenied       int    `json:"killsdenied"`
		Captures          int    `json:"captures"`
		Defends           int    `json:"defends"`
		TimePlayed        int    `json:"timeplayed"`
		WeaponData        string `json:"weapondata"`
	} `json:"stats"`
	Matches []struct {
		Identifier  string `json:"identifier"`
		Kills       int    `json:"kills"`
		Deaths      int    `json:"deaths"`
		EKIA        int    `json:"ekia"`
		GamesPlayed int    `json:"gamesplayed"`
		Wins        int    `json:"wins"`
		Losses      int    `json:"losses"`
		TotalShots  int    `json:"totalshots"`
		Captures    int    `json:"captures"`
		Defends     int    `json:"defends"`
		CareerScore int    `json:"careerscore"`
		TimePlayed  int    `json:"timeplayed"`
		RankXP      int    `json:"rankxp"`
		Time        int    `json:"time"`
		Format      string `json:"format"`
	} `json:"matches"`
	LastMatch struct {
		Identifier  string `json:"identifier"`
		Kills       int    `json:"kills"`
		Deaths      int    `json:"deaths"`
		EKIA        int    `json:"ekia"`
		GamesPlayed int    `json:"gamesplayed"`
		Wins        int    `json:"wins"`
		Losses      int    `json:"losses"`
		TotalShots  int    `json:"totalshots"`
		Captures    int    `json:"captures"`
		Defends     int    `json:"defends"`
		CareerScore int    `json:"careerscore"`
		TimesPlayed int    `json:"timesplayed"`
		Time        int    `json:"time"`
		Format      string `json:"format"`
	} `json:"lastmatch"`
	WeaponData []struct {
		Identifier      string `json:"identifier"`
		Name            string `json:"name"`
		Kills           int    `json:"kills"`
		BackstabberKill int    `json:"backstabberkills"`
		Deaths          int    `json:"deaths"`
		TimesUsed       int    `json:"timesused"`
		Used            int    `json:"used"`
		DeathsDuringUse int    `json:"deathsduringuse"`
		Hits            int    `json:"hits"`
		EKIA            int    `json:"ekia"`
		Destroyed       int    `json:"destroyed"`
		Headshots       int    `json:"headshots"`
		Shots           int    `json:"shots"`
		Assists         int    `json:"assists"`
		DamageDone      int    `json:"damagedone"`
	} `json:"weapondata"`
}

// Validation returns information about a single user.
type Validation struct {
	ID       int    `json:"id"`
	Success  bool   `json:"success"`
	UserName string `json:"username"`
}

// New initializes the API object and sets some variables.
func New(game string, platform string, username string) (*API, error) {
	base, err := url.Parse("https://cod-api.theapinetwork.com/api/")
	if err != nil {
		return &API{}, err
	}

	return &API{
		BaseURL:  base,
		Game:     game,
		Platform: platform,
		UserName: username,
	}, nil
}

// NewRequest builds the GET request and URL for the API call.
func (a *API) NewRequest(endpoint string) (*http.Request, error) {
	end, err := url.Parse(endpoint)
	if err != nil {
		return &http.Request{}, err
	}
	urlStr := a.BaseURL.ResolveReference(end)

	req, err := http.NewRequest("GET", urlStr.String(), nil)
	if err != nil {
		return req, err
	}

	req.Header.Set("Accept", "application/json")

	return req, nil
}

// Do runs the request against the API and returns JSON data.
func (a *API) Do(req *http.Request, i interface{}) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &i)
}

// GetLeaderboard queries the API for the best users in a specific scope.
func (a *API) GetLeaderboard(scope string, rows int) (*Leaderboard, error) {
	endpoint := "leaderboard/" + a.Game + "/" + a.Platform + "/" + scope + "?rows=" + strconv.Itoa(rows)
	req, err := a.NewRequest(endpoint)

	if err != nil {
		return &Leaderboard{}, err
	}

	var board Leaderboard
	err = a.Do(req, &board)

	return &board, err
}

// GetRecentMatches queries the API for recently-played matches.
func (a *API) GetRecentMatches(rows int) (*RecentMatches, error) {
	endpoint := "matches/recent?rows=" + strconv.Itoa(rows)
	req, err := a.NewRequest(endpoint)

	if err != nil {
		return &RecentMatches{}, err
	}

	var matches RecentMatches
	err = a.Do(req, &matches)

	return &matches, err
}

// GetSpecificMatch queries the API for a single recently-played match.
func (a *API) GetSpecificMatch(mid string) (*SpecificMatch, error) {
	endpoint := "matches/get?mid=" + mid
	req, err := a.NewRequest(endpoint)

	if err != nil {
		return &SpecificMatch{}, err
	}

	var match SpecificMatch
	err = a.Do(req, &match)

	return &match, err
}

// GetUserNames queries the API for users matching specific ID(s).
func (a *API) GetUserNames(uids ...int) (*UserIDToUserName, error) {
	endpoint := "users/ids?"
	num := 1
	for _, uid := range uids {
		if num > 1 {
			endpoint += "&"
		}
		endpoint += "id=" + strconv.Itoa(uid)
		num++
	}
	req, err := a.NewRequest(endpoint)

	if err != nil {
		return &UserIDToUserName{}, err
	}

	var users UserIDToUserName
	err = a.Do(req, &users)

	return &users, err
}

// GetUserStats queries the API for all stats for a specific user.
func (a *API) GetUserStats(matchType string) (*UserStats, error) {
	endpoint := "stats/" + a.Game + "/" + url.QueryEscape(a.UserName) + "/" + a.Platform + "?type=" + matchType
	req, err := a.NewRequest(endpoint)

	if err != nil {
		return &UserStats{}, err
	}

	var stats UserStats
	err = a.Do(req, &stats)

	return &stats, err
}

// ValidateUser queries the API to check if a user exists.
func (a *API) ValidateUser() (*Validation, error) {
	endpoint := "validate/" + a.Game + "/" + url.QueryEscape(a.UserName) + "/" + a.Platform
	req, err := a.NewRequest(endpoint)

	if err != nil {
		return &Validation{}, err
	}

	var validUser Validation
	err = a.Do(req, &validUser)

	return &validUser, err
}
