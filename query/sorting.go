package query

import "sort"

func sortProjectsByCreatedDate(values []*Project) {
	sort.Slice(values, func(i, j int) bool {
		return values[i].getCreatedDate().Before(values[j].getCreatedDate())
	})
}

func sortGamesByCreatedDate(values []*Game) {
	sort.Slice(values, func(i, j int) bool {
		return values[i].getCreatedDate().Before(values[j].getCreatedDate())
	})
}

func sortPlayersByCreatedDate(values []*Player) {
	sort.Slice(values, func(i, j int) bool {
		return values[i].getCreatedDate().Before(values[j].getCreatedDate())
	})
}

func sortMatchesByCreatedDate(values []*Match) {
	sort.Slice(values, func(i, j int) bool {
		return values[i].getCreatedDate().After(values[j].getCreatedDate())
	})
}
