package query

import (
	"compelo/rating"
	"sort"
	"time"
)

type Result string

const (
	Win  Result = "Win"
	Draw Result = "Draw"
	Loss Result = "Loss"
)

type Match struct {
	MetaData
	GUID        string `json:"guid"`
	GameGUID    string `json:"gameGuid"`
	ProjectGUID string `json:"projectGuid"`

	Date  time.Time    `json:"date" ts_type:"string"`
	Teams []*MatchTeam `json:"teams"`

	next *Match
	prev *Match
}

type MatchTeam struct {
	Players     []*Player `json:"players"`
	Score       int       `json:"score"`
	Result      Result    `json:"result"`
	RatingDelta int       `json:"ratingDelta"`
}

func (m *Match) determineResult() {
	highScore := 0
	highScoreCount := 0
	for _, t := range m.Teams {
		if t.Score > highScore {
			highScore = t.Score
			highScoreCount = 1
		} else if t.Score == highScore {
			highScoreCount++
		}
	}
	if highScoreCount < len(m.Teams) {
		for i := range m.Teams {
			if m.Teams[i].Score == highScore {
				m.Teams[i].Result = Win
			} else {
				m.Teams[i].Result = Loss
			}
		}
	} else {
		for i := range m.Teams {
			m.Teams[i].Result = Draw
		}
	}
}

func (m *Match) calculateTeamElo(ratings map[string]*Rating) {
	rm := rating.NewRatedMatch()
	for i, t := range m.Teams {
		sum := 0
		for _, p := range t.Players {
			sum += ratings[p.GUID].Current
		}
		avg := sum / len(t.Players)

		// The rating service expects a "rank" to sort players.
		// Here we just use the negative score instead, should
		// result in the same thing for most games..
		rm.AddPlayer(i, -t.Score, avg)
	}
	rm.Calculate()

	for i := range m.Teams {
		m.Teams[i].RatingDelta = rm.GetRatingDelta(i)
	}
}

// FIXME don't use this, update via MatchList
func (m *Match) updatePlayerRatings() {
	for _, team := range m.Teams {
		for _, player := range team.Players {
			player.ratings[m.GameGUID].Current += team.RatingDelta
		}
	}
}

func (m *Match) scoreDifference() int {
	lowestScore := 0
	highestScore := 0
	for i, t := range m.Teams {
		if t.Score < lowestScore || i == 0 {
			lowestScore = t.Score
		}
		if t.Score > highestScore || i == 0 {
			highestScore = t.Score
		}
	}
	return highestScore - lowestScore
}

func (m *Match) scoreSum() int {
	sum := 0
	for _, t := range m.Teams {
		sum += t.Score
	}
	return sum
}

func sortMatchesByCreatedDate(values []*Match) {
	sort.Slice(values, func(i, j int) bool {
		return values[i].ID > values[j].ID
	})
}

// eloMatchList contains the basic linked matchList
// as well as player ratings of all players that
// were part of any its matches.
type eloMatchList struct {
	matchList

	playerRatings map[string]int
}

func (ml *eloMatchList) addEloMatch(m *Match) {
	m.determineResult()
	ml.calculateTeamElo(m)
	ml.updatePlayerRatings(m)
	ml.addMatch(m)
}

func (ml *eloMatchList) removeEloMatch(m *Match) {
	ml.removeMatch(m)
	ml.recalculateElo()
}

func (ml *eloMatchList) recalculateElo() {
	ml.playerRatings = make(map[string]int)

	current := ml.head
	for current != nil {
		current.determineResult()
		ml.calculateTeamElo(current)
		ml.updatePlayerRatings(current)
		current = current.next
	}
}

func (ml *eloMatchList) playerRating(playerGUID string) int {
	if r, ok := ml.playerRatings[playerGUID]; ok {
		return r
	}
	ml.playerRatings[playerGUID] = rating.InitialRating
	return rating.InitialRating
}

func (ml *eloMatchList) calculateTeamElo(m *Match) {
	rm := rating.NewRatedMatch()
	for i, t := range m.Teams {
		sum := 0
		for _, p := range t.Players {
			sum += ml.playerRating(p.GUID)
		}
		avg := sum / len(t.Players)

		// The rating service expects a "rank" to sort players.
		// Here we just use the negative score instead, should
		// result in the same thing for most games..
		rm.AddPlayer(i, -t.Score, avg)
	}
	rm.Calculate()

	for i := range m.Teams {
		m.Teams[i].RatingDelta = rm.GetRatingDelta(i)
	}
}

func (ml *eloMatchList) updatePlayerRatings(m *Match) {
	for _, team := range m.Teams {
		for _, player := range team.Players {
			ml.playerRatings[player.GUID] += team.RatingDelta
		}
	}
}

type matchList struct {
	entries map[string]*Match
	head    *Match
	tail    *Match
}

func (ml *matchList) addMatch(m *Match) {
	if ml.head == nil {
		ml.head = m
	}

	if ml.tail != nil {
		ml.tail.next = m
		m.prev = ml.tail
	}

	ml.tail = m
	ml.entries[m.GUID] = m
}

func (ml *matchList) removeMatch(m *Match) {
	if ml.head == m {
		ml.head = m.next
	}
	if ml.tail == m {
		ml.tail = m.prev
	}

	if m.prev != nil && m.next != nil {
		m.prev.next = m.next
		m.next.prev = m.prev
	} else if m.prev != nil {
		m.prev.next = nil
	} else if m.next != nil {
		m.next.prev = nil
	}
}
