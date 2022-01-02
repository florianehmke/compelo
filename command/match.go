package command

type match struct {
	guid        string
	gameGUID    string
	projectGUID string

	teams []matchTeam
}

type matchTeam struct {
	playerGUIDs []string
	score       int
}
