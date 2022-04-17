package race

import "log"

func (r *Race) StepRace() {
	r.Step++
	r.stepEntrantsForward()
	r.checkRaceIsFinished()
	r.updatePositions()

	log.Println(r)
}

func (r *Race) stepEntrantsForward() {
	for _, v := range r.Entrants {
		horseSpeed := v.Horse.Speed
		v.Travelled += horseSpeed
	}
}

func (r *Race) checkRaceIsFinished() {
	for _, v := range r.Entrants {
		if v.Travelled > r.Distance {
			r.Finished = true

			return
		}
	}
}

// FIXME - Works but spaghetti.
func (r *Race) updatePositions() {
	correctlyPositioned := []int{}
	pos := 1

	for i := 0; i < len(r.Entrants); i++ {
		max := 0.0
		index := 0

		for j, v := range r.Entrants {
			can := true

			for _, v := range correctlyPositioned {
				if v == j {
					can = false
				}
			}

			if !can {
				continue
			}

			if v.Travelled > max {
				max = v.Travelled
				index = j
			}
		}

		correctlyPositioned = append(correctlyPositioned, index)
		r.Entrants[index].Position = pos
		pos++
	}
}
