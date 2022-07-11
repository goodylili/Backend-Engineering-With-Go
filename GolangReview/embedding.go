package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	// business logic
	return []Employee{}
}

func main() {
	mgtStaff := Manager{
		Employee: Employee{
			Name: "Mike Tyson",
			ID:   "0x23456",
		},
		Reports: []Employee{},
	}
	fmt.Println(mgtStaff.ID)
	fmt.Println(mgtStaff.Description())
}
