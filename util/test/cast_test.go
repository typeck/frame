package test

import (
	"fmt"
	"github.com/typeck/frame/util"
	"testing"
	"time"
)

type User struct {
	Name 	string
	Id 		int
}

func (u User) changeName() {
	u.Name = "tt"
}

func (u *User) changeId() {
	u.Id = -1
}

func (u *User) String() string{
	return u.Name
}

func TestPoint(t *testing.T) {
	user := User{Name: "tt"}
	fmt.Println(util.String(&user))
	go interf(user)
	user.Name = "ii"
	fmt.Println(user.String())
	time.Sleep(5 * time.Second)
}

func interf(i User)  {
	time.Sleep(time.Second)
	fmt.Println(i.String())
}