package main

import (
	"fmt"
	"github.com/danryan/go-group/os/group"
	"strconv"
)

type Group struct {
	g    *group.Group
	gid  int
	name string
}

func GroupByName(name string) (*Group, error) {
	g, err := group.LookupGroup(name)
	if err != nil {
		return nil, err
	}

	gid, err := strconv.Atoi(g.Gid)
	if err != nil {
		return nil, err
	}

	return &Group{
		name: g.Name,
		gid:  gid,
	}, nil
}

func main() {
	g := GroupByName("administrator")
	fmt.Println(g.name, g.gid)
}
