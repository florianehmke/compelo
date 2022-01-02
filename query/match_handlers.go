package query

import (
	"compelo/event"
)

func (h *defaultHandler) handleMatchCreated(e *event.MatchCreated) {
	project := h.data.projects[e.ProjectGUID]
	ratings := make(map[string]*Rating)
	teams := []*Team{}

	for _, t := range e.Teams {
		var players []*Player
		for _, guid := range t.PlayerGUIDs {
			players = append(players, project.players[guid])

			if rating, ok := project.players[guid].ratings[e.GameGUID]; ok {
				ratings[guid] = rating
			} else {
				rating := initialRatingFor(guid, e.GameGUID)
				ratings[guid] = rating
				project.players[guid].ratings = make(map[string]*Rating)
				project.players[guid].ratings[e.GameGUID] = rating
			}
		}
		sortPlayersByCreatedDate(players)

		teams = append(teams, &Team{
			Score:   t.Score,
			Players: players,
		})
	}

	match := Match{
		MetaData: MetaData{
			ID:          e.GetID(),
			CreatedDate: e.Date,
			UpdatedDate: e.Date,
		},
		GUID:        e.GUID,
		GameGUID:    e.GameGUID,
		ProjectGUID: e.ProjectGUID,
		Date:        e.Date,
		Teams:       teams,
	}

	match.determineResult()
	match.calculateTeamElo(ratings)
	match.updatePlayerRatings(ratings)

	h.data.projects[e.ProjectGUID].games[e.GameGUID].matches[e.GUID] = &match
}
