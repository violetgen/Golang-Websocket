package main

import (
	"fmt"

	r "github.com/dancannon/gorethink"
)

func main() {

	type User struct {
		Id   string `gorethink:"id,omitempty"` //leave if it is empty
		Name string `gorethink:"name"`
	}

	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "golang_react",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// user := User{
	// 	Name: "anonymous",
	// }
	// response, err := r.Table("user").Insert(user).RunWrite(session)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// response, _ := r.Table("user").Get("8ff3c230-6d7e-4c92-983a-1bd3372c51e3").Delete().RunWrite(session)
	// fmt.Printf("%#v\n", response)

	cursor, _ := r.Table("user").Changes(r.ChangesOpts{IncludeInitial: true}).Run(session)
	var changeResponse r.ChangeResponse

	for cursor.Next(&changeResponse) {
		fmt.Printf("%#v\n", changeResponse)
	}
}
