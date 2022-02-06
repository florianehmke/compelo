package query

import (
	"compelo/event"
)

func (h *defaultHandler) handleMatchCreated(e *event.MatchCreated) {
	project := h.data.projects[e.ProjectGUID]
	teams := []*MatchTeam{}

	for _, t := range e.Teams {
		var players []*Player
		for _, guid := range t.PlayerGUIDs {
			players = append(players, project.players[guid])
		}
		sortPlayersByCreatedDate(players)

		teams = append(teams, &MatchTeam{
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

	h.data.projects[e.ProjectGUID].games[e.GameGUID].eloMatchList.addEloMatch(&match)
}

func (h *defaultHandler) handleMatchDeleted(e *event.MatchDeleted) {
	h.data.projects[e.ProjectGUID].games[e.GameGUID].eloMatchList.removeEloMatch(e.GUID)
}
