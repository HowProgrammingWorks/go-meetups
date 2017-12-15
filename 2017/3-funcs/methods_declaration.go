package main

import "fmt"

type Student struct {
	Name  string
	Group string
	Age   int
}

func (s Student) String() string {
	return fmt.Sprintf("%s %d years old from group %s", s.Name, s.Age, s.Group)
}

func (s *Student) GoToArmy() error {
	if s.Group == "" {
		return fmt.Errorf("he's already there")
	}
	s.Group = ""
	return nil
}
