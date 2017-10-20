package resources

import "fmt"

type UserInfo struct {
	Name string
	Count int
}

var userMap = map [string]UserInfo {
	"simpleUser": UserInfo{"Ivanov invan ivanovich", 4},
	"eventOwner": UserInfo{"Kakoi-to chuvak", 3},
}

func GetUserInfo(id string) (*UserInfo, error){
	inf, ok := userMap[id]
	if !ok{
		return nil, fmt.Errorf("Can't find user by id %s", id)
	}
	return &inf, nil
}
