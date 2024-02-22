package main

import (
	"fmt"
	"slices"
)

type Team struct {
	name    string
	players []string
}

type League struct {
	teams []string
	wins  map[string]int
}

func NewLeague(teams []string) League {
	winMap := map[string]int{}
	for _, v := range teams {
		winMap[v] = 0
	}

	return League{
		teams: append(make([]string, 0, len(teams)), teams...),
		wins:  winMap,
	}
}

func (l *League) MatchResult(team1 string, score1 int, team2 string, score2 int) {
	if _, ok := l.wins[team1]; !ok {
		return
	}
	if _, ok := l.wins[team2]; !ok {
		return
	}

	switch {
	case score1 > score2:
		l.wins[team1]++
	case score1 < score2:
		l.wins[team2]++
	default:
		// Tie: do nothing
	}
}

func (l League) Ranking() []string {
	result := make([]string, 0, len(l.teams))
	result = append(result, l.teams...)

	slices.SortFunc(result, func(team1, team2 string) int {
		return l.wins[team2] - l.wins[team1]
	})

	return result
}

func main() {
	league := NewLeague([]string{"team1", "team2", "team3"})
	fmt.Println(league.Ranking())
	league.MatchResult("team1", 5, "team3", 10)
	league.MatchResult("team2", 5, "team3", 10)
	fmt.Println(league.Ranking())
	league.MatchResult("team1", 5, "team2", 10)
	league.MatchResult("team3", 5, "team2", 10)
	league.MatchResult("team3", 5, "team2", 10)
	league.MatchResult("team3", 5, "team2", 10)
	fmt.Println(league.Ranking())
}
