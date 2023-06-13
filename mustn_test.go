package must_test

import (
	"fmt"
	"testing"

	"github.com/tebeka/must"
)

type User struct {
	ID   int
	Name string
}

func NewUser(id int, name string) (*User, error) {
	if id <= 0 || name == "" {
		return nil, fmt.Errorf("bad id or name")
	}

	u := User{id, name}
	return &u, nil
}

func TestWrap2OK(t *testing.T) {
	fn := must.Wrap2(NewUser)
	id, name := 7, "bond"
	u := fn(id, name)
	if u.ID != id || u.Name != name {
		t.Fatal(u)
	}
}

func TestWrap2Panic(t *testing.T) {
	fn := must.Wrap2(NewUser)
	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("no panic")
		}
	}()

	fn(-1, "")
}
