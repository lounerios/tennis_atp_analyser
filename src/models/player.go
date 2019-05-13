package models

type Player struct {
	ID        string
	Firstname string
	Lastname  string
	Status    string `gorm:"index:status"`
	Birthdate string
	Country   string `gorm:"index:country"`
}

func NewPlayer(csv_line []string) *Player {
	return &Player{csv_line[0], csv_line[1], csv_line[2], csv_line[3], csv_line[4], csv_line[5]}
}
