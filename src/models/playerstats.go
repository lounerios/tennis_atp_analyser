package models

import (
  "utils"
)

type PlayerStats struct {
     Id string
     Seed int
     Entry string
     Rank int
     RankPoints int
     Aces int
     DoubleFaults int
     Serves int
     FirstSrvIn int
     FirstSrvWon int
     SecSrvWon int
     ServiceGames int
     BreakPointsSaved int
     BreakPointsFaced int
}

func NewPlayerStats(csv_line []string, id_idx int, seed_idx int, entry_idx int, rnk_idx int, rnkpts_idx int, aces_idx int, db_flts_idx int, serves_idx int, f_srv_in_idx int, f_srv_won_idx int, s_srv_won_idx int, srv_gms_idx int, brk_pts_saved_idx int, brk_pts_faced_idx int) *PlayerStats{
    ps := new(PlayerStats)

    ps.Id = csv_line[id_idx]
    ps.Seed = utils.GetNumber(csv_line[seed_idx])
    ps.Entry = csv_line[entry_idx]
    ps.Rank = utils.GetNumber(csv_line[rnk_idx])
    ps.RankPoints = utils.GetNumber(csv_line[rnkpts_idx])
    ps.Aces = utils.GetNumber(csv_line[aces_idx])
    ps.DoubleFaults = utils.GetNumber(csv_line[db_flts_idx])
    ps.Serves = utils.GetNumber(csv_line[serves_idx])
    ps.FirstSrvIn = utils.GetNumber(csv_line[f_srv_in_idx])
    ps.FirstSrvWon = utils.GetNumber(csv_line[f_srv_won_idx])
    ps.SecSrvWon = utils.GetNumber(csv_line[s_srv_won_idx])
    ps.ServiceGames = utils.GetNumber(csv_line[srv_gms_idx])
    ps.BreakPointsSaved = utils.GetNumber(csv_line[brk_pts_saved_idx])
    ps.BreakPointsFaced = utils.GetNumber(csv_line[brk_pts_faced_idx])

   return ps
}
