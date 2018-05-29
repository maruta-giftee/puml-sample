package main

import (
	"fmt"
	"strconv"
)

type Skill struct {
	Name      string
	Rank      string
	Reference string
}

type Member struct {
	Name   string
	Age    int
	Skills []Skill
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
		fmt.Println(" - " + m.Name + "(" + strconv.Itoa(m.Age) + ")")
		for _, s := range m.Skills {
			fmt.Println("  -> " + s.Name + ": " + s.Rank)
		}
		fmt.Println()
	}
}

func main() {
	jedi := &Team{Name: "Jedi"}
	sith := &Team{Name: "Sith"}

	// create member
	yoda := &Member{
		Name: "Master Yoda",
		Age:  800,
		Skills: []Skill{
			{
				Name:      "Rails",
				Rank:      "contributor",
				Reference: "https://rubyonrails.org/",
			},
			{
				Name:      "AWS",
				Rank:      "solution architect",
				Reference: "https://aws.amazon.com/jp/certification/certified-solutions-architect-associate/",
			},
			{
				Name:      "Mysql",
				Rank:      "database administrator",
				Reference: "https://education.oracle.com/pls/web_prod-plq-dad/db_pages.getpage?page_id=654&get_params=p_id:260&p_org_id=70&lang=JA#tabs-2-2",
			},
		},
	}
	luke := &Member{
		Name: "Luke Skywalker",
		Age:  19,
		Skills: []Skill{
			{
				Name:      "Rails",
				Rank:      "developer",
				Reference: "https://rubyonrails.org/",
			},
			{
				Name:      "AWS",
				Rank:      "rookie",
				Reference: "https://aws.amazon.com/",
			},
		},
	}
	// create member
	anakin := &Member{
		Name: "Anakin Skywalker",
		Age:  19,
		Skills: []Skill{
			{
				Name:      "Rails",
				Rank:      "developer",
				Reference: "https://rubyonrails.org/",
			},
			{
				Name:      "Go",
				Rank:      "developer",
				Reference: "https://golang.org/",
			},
			{
				Name:      "AWS",
				Rank:      "developer",
				Reference: "https://aws.amazon.com/",
			},
			{
				Name:      "Mysql",
				Rank:      "database administrator",
				Reference: "https://education.oracle.com/pls/web_prod-plq-dad/db_pages.getpage?page_id=654&get_params=p_id:260&p_org_id=70&lang=JA#tabs-2-2",
			},
		},
	}

	// edit member
	fmt.Println("episode 1")
	jedi.AddMember(yoda)
	jedi.AddMember(anakin)
	jedi.ShowMembers()

	fmt.Println("episode 3")
	jedi.RemoveMember(anakin)
	sith.AddMember(anakin)
	jedi.ShowMembers()
	sith.ShowMembers()

	fmt.Println("episode 4")
	jedi.AddMember(luke)
	jedi.ShowMembers()
	sith.ShowMembers()
}
