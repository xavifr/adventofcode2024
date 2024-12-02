package Domain

type Day2Report struct {
	Levels []int
}

func (r *Day2Report) Safe(withErrors bool) bool {
	differences := calcDifferences(r.Levels)
	if compareDifferences(differences) {
		return true
	}

	if withErrors {
		for i := 0; i < len(r.Levels); i++ {
			newLevels := append([]int{}, r.Levels[:i]...)
			newLevels = append(newLevels, r.Levels[i+1:]...)
			
			differences := calcDifferences(newLevels)
			if compareDifferences(differences) {
				return true
			}
		}
	}

	return false
}

func calcDifferences(in []int) (out []int) {
	out = make([]int, 0)

	for i := 1; i < len(in); i++ {
		out = append(out, in[i-1]-in[i])
	}

	return
}

func compareDifferences(in []int) bool {
	direction := in[0]
	for i := 0; i < len(in); i++ {
		if i > 0 && direction*in[i] < 0 { // different sign, not safe
			return false
		}

		if in[i] == 0 || AbsInt(in[i]) > 3 {
			return false // difference < 1 or > 3
		}
	}

	return true
}

func AbsInt(x int) int {
	if x > 0 {
		return x
	}

	return x * -1
}
