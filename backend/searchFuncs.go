package fonction_go

import (
	"sort"
	"unicode/utf8"
)

func Search(list []Mess, searchQuery string, searchType string) []Mess {

	if searchQuery == "" {
		return list
	}

	var result []Mess
	var noSort [][]any
	switch searchType {
	case "subject":
		for i := 0; i < len(list); i++ {
			indice := matchExact(list[i].Subject, searchQuery) + matchLetter(list[i].Subject, searchQuery)
			noSort = append(noSort, []any{indice, list[i]})
		}
	case "user":
		for i := 0; i < len(list); i++ {
			indice := matchExact(list[i].Username, searchQuery) + matchLetter(list[i].Username, searchQuery)
			noSort = append(noSort, []any{indice, list[i]})
		}
	default:
		for i := 0; i < len(list); i++ {
			for j := 0; j < len(list[i].Tag); j++ {
				indice := matchExact(list[i].Tag[j], searchQuery) + matchLetter(list[i].Tag[j], searchQuery)
				noSort = append(noSort, []any{indice, list[i]})
			}
		}
	}
	result = sortMess(noSort)
	return result
}

func matchExact(s1 string, s2 string) int {
	var indice int
	for i := 0; i < utf8.RuneCountInString(s2); i++ {
		if i > utf8.RuneCountInString(s2) {
			break
		}
		if s2[i] == s1[i] {
			indice += 1
		}
	}
	return indice
}

func matchLetter(s1 string, s2 string) int {
	var indice int
	runes1 := []rune(s1)
	runes2 := []rune(s2)

	for i := 0; i < len(runes2); i++ {
		for j := 0; j < len(runes1); j++ {
			if runes2[i] == runes1[j] {
				indice++
				runes1 = append(runes1[:j], runes1[j+1:]...)
				break
			}
		}
	}
	return indice
}

func sortMess(list [][]any) []Mess {
	type indexedMess struct {
		index float64
		mess  Mess
	}

	var indexed []indexedMess
	for _, item := range list {
		if len(item) < 2 {
			continue
		}
		index, ok1 := item[0].(int)
		mess, ok2 := item[1].(Mess)
		if !ok1 || !ok2 {
			continue
		}
		indexed = append(indexed, indexedMess{float64(index), mess})
	}

	sort.Slice(indexed, func(i, j int) bool {
		return indexed[i].index < indexed[j].index
	})

	result := make([]Mess, len(indexed))
	for i, v := range indexed {
		result[i] = v.mess
	}
	return result
}
