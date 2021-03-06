package models

type AtpRanking struct {
	Date     string
	Number   string `gorm:"index:number"`
	PlayerId string `gorm:"index:playerid"`
	Points   string
}

func NewAtpRanking(csv_line []string) *AtpRanking {
	return &AtpRanking{csv_line[0], csv_line[1], csv_line[2], csv_line[3]}
}
