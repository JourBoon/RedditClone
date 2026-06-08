package fonction_go
import(
	"unicode/utf8"
	"sort"
)

func Search(list []Mess, searchQuery string, searchType string) []Mess {
	var result []Mess;
	var noSort [][]any;
	for i:=0;i<len(list);i++{
		indice := matchExact(list[i].Subject,searchQuery)+matchLetter(list[i].Subject,searchQuery);
		noSort = append(noSort, []any{indice, list[i]})
	}
	result = sortMess(noSort);
	return result
}

func matchExact(s1 string, s2 string) int {
	var indice int
	for i := 0; i < utf8.RuneCountInString(s2); i++ {
		if i>utf8.RuneCountInString(s2){
			break
		}
		if s2[i]==s1[i]{
			indice+=1
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
        index, ok1 := item[0].(float64)
        mess, ok2 := item[1].(Mess)
        if !ok1 || !ok2 {
            continue
        }
        indexed = append(indexed, indexedMess{index, mess})
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