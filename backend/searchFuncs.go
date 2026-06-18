package fonction_go

import (
	"sort"
)

type scoredMess struct {
	score int
	mess  Mess
}

func Search(list []Mess, searchQuery string, searchType string) []Mess {
	if searchQuery == "" {
		return list
	}

	var scoredList []scoredMess

	switch searchType {
	case "subject":
		for _, m := range list {
			score := matchExact(m.Subject, searchQuery) + matchLetter(m.Subject, searchQuery)
			scoredList = append(scoredList, scoredMess{score: score, mess: m})
		}
	case "user":
		for _, m := range list {
			score := matchExact(m.Username, searchQuery) + matchLetter(m.Username, searchQuery)
			scoredList = append(scoredList, scoredMess{score: score, mess: m})
		}
	default:
		for _, m := range list {
			bestScore := 0
			for _, tag := range m.Tag {
				score := matchExact(tag, searchQuery) + matchLetter(tag, searchQuery)
				if score > bestScore {
					bestScore = score
				}
			}
			scoredList = append(scoredList, scoredMess{score: bestScore, mess: m})
		}
	}

	sort.Slice(scoredList, func(i, j int) bool {
		return scoredList[i].score > scoredList[j].score
	})

	result := make([]Mess, len(scoredList))
	for i, sm := range scoredList {
		result[i] = sm.mess
	}

	return result
}

func matchExact(s1 string, s2 string) int {
	runes1 := []rune(s1)
	runes2 := []rune(s2)
	var indice int

	limit := len(runes1)
	if len(runes2) < limit {
		limit = len(runes2)
	}

	for i := 0; i < limit; i++ {
		if runes1[i] == runes2[i] {
			indice++
		}
	}
	return indice
}

func matchLetter(s1 string, s2 string) int {
	var indice int
	runes1 := []rune(s1)
	runes2 := []rune(s2)

	counts := make(map[rune]int)
	for _, r := range runes1 {
		counts[r]++
	}

	for _, r := range runes2 {
		if counts[r] > 0 {
			indice++
			counts[r]-- 
		}
	}
	return indice
}