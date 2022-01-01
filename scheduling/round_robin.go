package scheduling

const dummy = "D"

type Pairing struct {
	A string
	B string
}

type RoundRobin struct {
	players []string
}

func NewRoundRobin() *RoundRobin {
	return &RoundRobin{}
}

func (r *RoundRobin) AddPlayer(name string) {
	r.players = append(r.players, name)
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
		schedule[i] = pairings(halfA, halfB)
		players = rotate(players)
	}

	return schedule
}

func (r *RoundRobin) initialize() {
	if len(r.players)%2 == 1 {
		r.AddPlayer(dummy)
	}
}

func pairings(halfA []string, halfB []string) []Pairing {
	var pairings []Pairing
	for i := 0; i < len(halfA); i++ {
		if halfA[i] != dummy && halfB[i] != dummy {
			pairings = append(pairings, Pairing{
				A: halfA[i],
				B: halfB[i],
			})
		}
	}
	return pairings
}

func rotate(slice []string) []string {
	s := copySlice(slice)
	fixed := s[0]
	moves := s[len(s)-1]
	return append([]string([]string{fixed, moves}), s[1:len(s)-1]...)
}

func bisect(slice []string) ([]string, []string) {
	s := copySlice(slice)
	halfA, halfB := divide(s)
	return halfA, reverse(halfB)
}

func reverse(slice []string) []string {
	s := copySlice(slice)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func divide(slice []string) ([]string, []string) {
	s := copySlice(slice)
	half := len(s) / 2
	halfA := s[:half]
	halfB := s[half:]
	return halfA, halfB
}

func copySlice(slice []string) []string {
	s := make([]string, len(slice))
	copy(s, slice)
	return s
}
