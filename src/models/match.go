package models

import (
	"strings"

	"../utils"
)

const (
	MATCH_NUM_INDEX                 = 6
	SETS_INDEX                      = 27
	BEST_OF_INDEX                   = 28
	ROUND_INDEX                     = 29
	MINUTES_INDEX                   = 30
	WINNER_ID_INDEX                 = 7
	WINNER_SEED_INDEX               = 8
	WINNER_ENTRY_INDEX              = 9
	WINNER_RANK_INDEX               = 15
	WINNER_RANK_POINTS_INDEX        = 16
	WINNER_ACES_INDEX               = 31
	WINNER_DOULBE_FAULTS_INDEX      = 32
	WINNER_SERVES_INDEX             = 33
	WINNER_FIRST_SRV_IN_INDEX       = 34
	WINNER_FISRT_SRV_WON_INDEX      = 35
	WINNER_SEC_SRV_WON_INDEX        = 36
	WINNER_SERVICE_GAMES_INDEX      = 37
	WINNER_BREAK_POINTS_SAVED_INDEX = 38
	WINNER_BREAK_POINTS_FACED_INDEX = 39
	LOSER_ID_INDEX                  = 17
	LOSER_SEED_INDEX                = 18
	LOSER_ENTRY_INDEX               = 19
	LOSER_RANK_INDEX                = 25
	LOSER_RANK_POINTS_INDEX         = 26
	LOSER_ACES_INDEX                = 40
	LOSER_DOUBLE_FAULTS_INDEX       = 41
	LOSER_SERVES_INDEX              = 42
	LOSER_FIRST_SRV_IN_INDEX        = 43
	LOSER_FIRST_SRV_WON_INDEX       = 44
	LOSER_SEC_SRV_WON_INDEX         = 45
	LOSER_SERVICE_GAMES_INDEX       = 46
	LOSER_BREAK_POINTS_SAVED_INDEX  = 47
	LOSER_BREAK_POINTS_FACED_INDEX  = 48
)

type Match struct {
	ID           uint
	TournamentId string     `gorm:"index:tournid"`
	MatchNum     string     `gorm:"index:matchnum"`
	Sets         []MatchSet `gorm: "-"`
	WinnerStats  PlayerStats
	LoserStats   PlayerStats
	BestOf       int
	Round        string
	Minutes      int
}

type MatchSet struct {
	MatchId uint `gorm:"index:matchid"`
	Score   string
}

func NewMatch(csv_line []string) *Match {
	m := new(Match)

	m.TournamentId = csv_line[TOURNAMENT_ID_INDEX]
	m.MatchNum = csv_line[MATCH_NUM_INDEX]
	m.BestOf = utils.GetNumber(csv_line[BEST_OF_INDEX])
	m.Round = csv_line[ROUND_INDEX]
	m.Minutes = utils.GetNumber(csv_line[MINUTES_INDEX])

	/* The winner has passed to next round without game because the loser was injured */
	if csv_line[SETS_INDEX] == "W/O" {
		return m
	} else {
		for _, set := range strings.Split(csv_line[SETS_INDEX], " ") {
			var s MatchSet
			s.Score = set

			m.Sets = append(m.Sets, s)
		}

	}

	m.WinnerStats = NewPlayerStats(csv_line, WINNER_ID_INDEX, WINNER_SEED_INDEX, WINNER_ENTRY_INDEX, WINNER_RANK_INDEX, WINNER_RANK_POINTS_INDEX,
		WINNER_ACES_INDEX, WINNER_DOULBE_FAULTS_INDEX, WINNER_SERVES_INDEX, WINNER_FIRST_SRV_IN_INDEX, WINNER_FISRT_SRV_WON_INDEX,
		WINNER_SEC_SRV_WON_INDEX, WINNER_SERVICE_GAMES_INDEX, WINNER_BREAK_POINTS_SAVED_INDEX, WINNER_BREAK_POINTS_FACED_INDEX)

	m.LoserStats = NewPlayerStats(csv_line, LOSER_ID_INDEX, LOSER_SEED_INDEX, LOSER_ENTRY_INDEX, LOSER_RANK_INDEX, LOSER_RANK_POINTS_INDEX, LOSER_ACES_INDEX,
		LOSER_DOUBLE_FAULTS_INDEX, LOSER_SERVES_INDEX, LOSER_FIRST_SRV_IN_INDEX, LOSER_FIRST_SRV_WON_INDEX, LOSER_SEC_SRV_WON_INDEX, LOSER_SERVICE_GAMES_INDEX,
		LOSER_BREAK_POINTS_SAVED_INDEX, LOSER_BREAK_POINTS_FACED_INDEX)

	return m
}
