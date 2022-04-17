package gamble

import "github.com/ImTheTom/OtherProjects/race-track/pkg/race"

type Status int

const (
	Pending Status = iota
	Paid
)

type Gamble struct {
	UserID     int
	RaceID     int
	EntrantNum int
	Bet        float64
	Odds       float64
	Status     Status
}

var gambles []*Gamble

func CreateNewGamble(user, race int, entrantNum int, bet, odds float64) *Gamble {
	gmb := &Gamble{
		UserID:     user,
		RaceID:     race,
		EntrantNum: entrantNum,
		Bet:        bet,
		Odds:       odds,
		Status:     Pending,
	}
	gambles = append(gambles, gmb)

	AdjustUserBalance(user, -bet)

	return gmb
}

func PayoutWinnings(rc *race.Race) {
	for _, v := range gambles {
		if v.RaceID == rc.ID {
			entrantPos := rc.Entrants[v.EntrantNum-1].Position

			if entrantPos == 1 {
				payoutAmt := v.Bet * v.Odds
				AdjustUserBalance(v.UserID, payoutAmt)
			}

			v.Status = Paid
		}
	}
}
