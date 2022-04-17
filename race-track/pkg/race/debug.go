package race

import "fmt"

func (r *Race) String() string {
	header := fmt.Sprintf(
		"RACE %s IS A %.2f METRE RACE ON STEP %d IT STATE IS %v\n",
		r.Name,
		r.Distance,
		r.Step,
		r.Finished,
	)

	var entrantsString string
	for _, v := range r.Entrants {
		entrantsString = fmt.Sprintf(
			"%sENTRANT NUMBER %d IS HORSE %s JOCKEY %s HAS TRAVELLED %.2f AND IS IN %d\n",
			entrantsString,
			v.Number,
			v.Horse.Name,
			v.Jockey.Name,
			v.Travelled,
			v.Position,
		)
	}

	return header + entrantsString
}
