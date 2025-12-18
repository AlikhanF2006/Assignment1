package company

import "fmt"

type Employee interface {
	GetDetails() string
}

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary float64
}

func (f FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("Full-time: %s, Salary: %.2f", f.Name, f.Salary)
}

type PartTimeEmployee struct {
	ID     uint64
	Name   string
	Salary float64
}

func (p PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("Part-time: %s, Salary: %.2f", p.Name, p.Salary)
}

type Company struct {
	Employees map[uint64]Employee
}

func NewCompany() Company {
	return Company{
		Employees: make(map[uint64]Employee),
	}
}

func (c Company) AddEmployee(id uint64, e Employee) Company {
	c.Employees[id] = e
	return c
}

func (c Company) ListEmployees() {
	for _, e := range c.Employees {
		fmt.Println(e.GetDetails())
	}
}
