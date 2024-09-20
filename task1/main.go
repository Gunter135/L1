package main

import "fmt"

// struct embedding
type Human struct {
	id int
}

func (h *Human) describe() string {
	return fmt.Sprintf("human with id=%v", h.id)
}

func (h *Human) dosmth() {
	fmt.Println(h.id)
}


type Action struct {
	Human
	actionstuff string
}

func main() {
	h := Human{11}
	h.dosmth()
	action := Action{
		Human: h,
		actionstuff: "stuff",
	}
	action.dosmth()
	type ds interface {
		describe() string
	}
	var d ds = &action
	fmt.Println(d.describe())
}
