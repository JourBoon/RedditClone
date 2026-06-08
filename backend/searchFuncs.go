package fonction_go

import (
	"fmt"
)

func Search(searchQuery string, searchType string) []Mess {
	db, err := dbConnection()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	allMess,err:= getMess(db)
	if err!=nil{
		println(err)
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var result []Mess
	for _, m := range allMess {
		switch searchType {
		case "tag":
			for _, tag := range m.Tag {
				if tag == searchQuery {
					result = append(result, m)
					break
				}
			}
		case "username":
			if m.Username == searchQuery {
				result = append(result, m)
			}
		case "created_at":
			if m.Created_at == searchQuery {
				result = append(result, m)
			}
		}
	}
	return result
}
