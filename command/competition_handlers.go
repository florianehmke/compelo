package command

import "compelo/event"

func (svc *Service) handleCompetitionCreated(e *event.CompetitionCreated) {

	// Map teams to interal type.
	var teams []competitionTeam
	for _, t := range e.Teams {
		teams = append(teams, competitionTeam{
			playerGUIDs: t.PlayerGUIDs,
		})
	}

	svc.data.projects[e.ProjectGUID].games[e.GameGUID].competitions[e.GUID] = competition{
		guid:        e.GUID,
		gameGUID:    e.GameGUID,
		projectGUID: e.ProjectGUID,
		name:        e.Name,
		rounds:      e.Rounds,
		teams:       teams,
	}
}
