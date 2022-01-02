package command

type competition struct {
	guid        string
	gameGUID    string
	projectGUID string

	name   string
	rounds int

	teams []competitionTeam
}

type competitionTeam struct {
	playerGUIDs []string
}
