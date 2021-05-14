package main

import "fmt"

type department interface {
	execute(patient *patient)
	setNext(department department)
}

type patient struct {
	name        string
	isReception bool
	isDoctor    bool
	isMedicine  bool
	isCashier   bool
}

type reception struct {
	next department
}

func (r *reception) execute(p *patient) {
	if p.isReception {
		fmt.Printf("%s has gone reception\n", p.name)
		return
	}
	fmt.Printf("%s is staying reception\n", p.name)
	p.isReception = true
	r.next.execute(p)
}

func (r *reception) setNext(d department) {
	r.next = d
}

type doctor struct {
	next department
}

func (r *doctor) execute(p *patient) {
	if p.isDoctor {
		fmt.Printf("%s has gone doctor\n", p.name)
		return
	}
	fmt.Printf("%s is staying doctor\n", p.name)
	p.isDoctor = true
	r.next.execute(p)
}

func (r *doctor) setNext(d department) {
	r.next = d
}

type medicine struct {
	next department
}

func (r *medicine) execute(p *patient) {
	if p.isMedicine {
		fmt.Printf("%s has gone medicine\n", p.name)
		return
	}
	fmt.Printf("%s is staying medicine\n", p.name)
	p.isMedicine = true
	r.next.execute(p)
}

func (r *medicine) setNext(d department) {
	r.next = d
}

type cashier struct{
	next department
}

func (c *cashier) execute(p *patient) {
	if p.isCashier {
		fmt.Printf("%s has gone cashier\n", p.name)
		return
	}
	fmt.Printf("%s is staying cashier\n", p.name)
	p.isCashier = true
}

func (c *cashier) setNext(d department) {
	c.next = d
}

func main() {
	c := cashier{}

	m := medicine{}
	m.setNext(&c)

	d := doctor{}
	d.setNext(&m)

	r := reception{}
	r.setNext(&d)

	p := patient{name: "xavi"}

	r.execute(&p)
}
