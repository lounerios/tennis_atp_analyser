package models

import (
  "utils"
)

const (
     TOURNAMENT_ID_INDEX        = 0
     TOURNAMENT_NAME_INDEX      = 1
     TOURNAMENT_SURFACE_INDEX   = 2
     TOURNAMENT_DRAW_SIZE_INDEX = 3
     TOURNAMENT_LEVEL_INDEX     = 4
     TOURNAMENT_DATE_INDEX      = 5
)

type Tournament struct {
  Id string
  Name string
  Surface string
  DrawSize int
  Level string
  Date string
}

func NewTournament(csv_line []string) *Tournament{
  t := new(Tournament)

  t.Id = csv_line[TOURNAMENT_ID_INDEX]
  t.Name = csv_line[TOURNAMENT_NAME_INDEX]
  t.Surface = csv_line[TOURNAMENT_SURFACE_INDEX]
  t.DrawSize = utils.GetNumber(csv_line[TOURNAMENT_DRAW_SIZE_INDEX])
  t.Level = csv_line[TOURNAMENT_LEVEL_INDEX]
  t.Date = csv_line[TOURNAMENT_DATE_INDEX]

  return t
}
