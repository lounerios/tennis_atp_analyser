package models

type Player struct  {
     Id string
     Firstname string
     Lastname string
     Status string
     Birthdate string
     Country string
}

func NewPlayer(csv_line []string) *Player {
    return &Player{csv_line[0], csv_line[1], csv_line[2], csv_line[3], csv_line[4], csv_line[5]}
}
