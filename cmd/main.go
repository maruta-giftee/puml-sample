package main

import (
	"fmt"
)

type Member struct {
	Name string
	Age  int
}

type Team struct {
	Name    string
	Members []*Member
}

func (t *Team) AddMember(m *Member) {
	t.Members = append(t.Members, m)
}

func (t *Team) RemoveMember(member *Member) {
	members := []*Member{}
	for _, m := range t.Members {
		if m != member {
			members = append(members, m)
		}
	}
	t.Members = members
}

func (t *Team) ShowMembers() {
	fmt.Println(t.Name)
	for _, m := range t.Members {
		fmt.Println(" - " + m.Name)
	}
	fmt.Println()
}

func main() {
	jedi := &Team{Name: "Jedi"}
	sith := &Team{Name: "Sith"}

	// create member
	yoda := &Member{
		Name: "Master Yoda",
		Age:  800,
	}
	luke := &Member{
		Name: "Luke Skywalker",
		Age:  19,
	}
	// create member
	anakin := &Member{
		Name: "Anakin Skywalker",
		Age:  19,
	}
	obiwan := &Member{
		Name: "Obi-Wan Kenobi",
		Age:  25,
	}

	// edit member
	fmt.Println("episode 1")
	jedi.AddMember(yoda)
	jedi.AddMember(obiwan)
	jedi.AddMember(anakin)
	jedi.ShowMembers()

	fmt.Println("episode 3")
	jedi.RemoveMember(anakin)
	sith.AddMember(anakin)
	jedi.ShowMembers()
	sith.ShowMembers()

	fmt.Println("episode 4")
	jedi.AddMember(luke)
	jedi.RemoveMember(obiwan)
	jedi.ShowMembers()
	sith.ShowMembers()
}
