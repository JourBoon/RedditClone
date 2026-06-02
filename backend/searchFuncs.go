package fonction_go

import (
	"fmt"
)

func Search(searchQuery string, searchType string) []mess {
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

	var result []mess
	for _, m := range allMess {
		switch searchType {
		case "tag":
			for _, tag := range m.tag {
				if tag == searchQuery {
					result = append(result, m)
					break
				}
			}
		case "username":
			if m.username == searchQuery {
				result = append(result, m)
			}
		case "created_at":
			if m.created_at == searchQuery {
				result = append(result, m)
			}
		}
	}
	return result
}
