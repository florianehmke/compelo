package command

type competition struct {
	guid        string
	gameGUID    string
	projectGUID string

	teams []competitionTeam
}

type competitionTeam struct {
	playerGUIDs []string
}
