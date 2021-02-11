package main

import "errors"

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"mail"` // => могу поля структуры привязывать к полям json
}

var users = make(map[string]*user)
var tokens = make(map[string]string)

func (u user) getToken() (string, error) {
	for t, un := range tokens {
		if u.Username == un {
			return t, nil
		}
	}
	return "", errors.New("Token Not Found")
}

func startDataInit() {
	// Заполняем пародию на базу данных
	/*newId, _ := uuid.NewUUID()
	users[newId.String()] = &user{}
	users[newId.String()].Email = "test@gmail.com"*/
	// Сначала создаем структуру, только потом изменяем ее поля

	users["1"] = &user{
		"qwe",
		"123",
		"test@gmail.com",
	}
}
