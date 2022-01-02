package query

import "compelo/event"

func (h *defaultHandler) handleCompetitionCreated(e *event.CompetitionCreated) {
	project := h.data.projects[e.ProjectGUID]

	teams := []*CompetitionTeam{}

	for _, t := range e.Teams {
		var players []*Player
		for _, guid := range t.PlayerGUIDs {
			players = append(players, project.players[guid])
		}
		sortPlayersByCreatedDate(players)
		teams = append(teams, &CompetitionTeam{Players: players})
	}

	competition := Competition{
		MetaData: MetaData{
			ID:          e.GetID(),
			CreatedDate: e.Date,
			UpdatedDate: e.Date,
		},
		GUID:        e.GUID,
		GameGUID:    e.GameGUID,
		ProjectGUID: e.ProjectGUID,
		Rounds:      e.Rounds,
		Name:        e.Name,
		Teams:       teams,
	}

	h.data.projects[e.ProjectGUID].games[e.GameGUID].competitions[e.GUID] = &competition
}
