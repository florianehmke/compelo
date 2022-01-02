package scheduling

const dummy = -1

type Pairing struct {
	A int
	B int
}

type RoundRobin struct {
	players []int
}

func NewRoundRobin() *RoundRobin {
	return &RoundRobin{}
}

func (r *RoundRobin) AddPlayer(id int) {
	r.players = append(r.players, id)
}

func (r *RoundRobin) Count() int {
	return len(r.players)
}

func (r *RoundRobin) Schedule() map[int][]Pairing {
	schedule := make(map[int][]Pairing)

	r.initialize()
	players := r.players
	for i := 0; i < len(r.players)-1; i++ {
		halfA, halfB := bisect(players)
		schedule[i] = pairings(i, halfA, halfB)
		players = rotate(players)
	}

	return schedule
}

func (r *RoundRobin) ScheduleRounds(rounds int) map[int]map[int][]Pairing {
	schedules := make(map[int]map[int][]Pairing)

	schedule := r.Schedule()
	for i := 0; i < rounds; i++ {
		schedules[i] = schedule
		schedule = swapSides(schedule)
	}

	return schedules
}

func (r *RoundRobin) initialize() {
	if len(r.players)%2 == 1 {
		r.AddPlayer(dummy)
	}
}

func pairings(day int, halfA []int, halfB []int) []Pairing {
	var pairings []Pairing
	for i := 0; i < len(halfA); i++ {
		if halfA[i] != dummy && halfB[i] != dummy {
			var a int
			var b int

			// swap side based on day
			if day%2 == 0 {
				a = halfA[i]
				b = halfB[i]
			} else {
				a = halfB[i]
				b = halfA[i]
			}

			pairings = append(pairings, Pairing{A: a, B: b})
		}
	}
	return pairings
}

func rotate(slice []int) []int {
	s := copySlice(slice)
	fixed := s[0]
	moves := s[len(s)-1]
	return append([]int([]int{fixed, moves}), s[1:len(s)-1]...)
}

func bisect(slice []int) ([]int, []int) {
	s := copySlice(slice)
	halfA, halfB := divide(s)
	return halfA, reverse(halfB)
}

func reverse(slice []int) []int {
	s := copySlice(slice)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func divide(slice []int) ([]int, []int) {
	s := copySlice(slice)
	half := len(s) / 2
	halfA := s[:half]
	halfB := s[half:]
	return halfA, halfB
}

func copySlice(slice []int) []int {
	s := make([]int, len(slice))
	copy(s, slice)
	return s
}

func swapSides(schedule map[int][]Pairing) map[int][]Pairing {
	swapped := make(map[int][]Pairing)
	for i, pairings := range schedule {
		swapped[i] = make([]Pairing, 0)
		for _, pairing := range pairings {
			swapped[i] = append(swapped[i], Pairing{A: pairing.B, B: pairing.A})
		}
	}
	return swapped
}
