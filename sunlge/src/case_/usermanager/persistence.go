package usermanager

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Persistence interface {
	load() (map[int]User, error)
	Store(map[int]User) error //另一种写法   Store(users map[int]User) error
}

type JSONFile struct {
	name string
}

func (j JSONFile) load() (map[int]User, error) {
	bytes, err := ioutil.ReadFile(j.name)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[int]User), err // return map[int]User{}, nil
		}
		return nil, err
	}
	var users map[int]User
	err = json.Unmarshal(bytes, &users)
	return users  , err
}


func (j JSONFile) Store(users map[int]User) error {
	bytes, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(j.name, bytes, os.ModePerm)
}


var t1 Persistence = JSONFile{"test"}

func main()  {
	t1.load()
}