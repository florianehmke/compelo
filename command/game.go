package command

type game struct {
	guid        string
	projectGUID string

	name string

	matches      map[string]match
	competitions map[string]competition
}
