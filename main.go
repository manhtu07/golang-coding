package main

import "fmt"

type observer interface {
	update(prodName string)
}

type item struct {
	observers []observer
	id        uint32
	name      string
}

func newItem(id uint32, name string) *item {
	return &item{
		id:   id,
		name: name,
	}
}

func (i *item) register(usr observer) {
	i.observers = append(i.observers, usr)
}

func (i *item) notiAll() {
	for _, usr := range i.observers {
		usr.update(i.name)
	}
}

type customer struct {
	id   uint32
	name string
}

func (c *customer) update(prodName string) {
	fmt.Printf("user subscribed: id: %d, name: %s, product: %s\n", c.id, c.name, prodName)
}

func main() {
	newItem := newItem(1, "Iphone 12")
	newItem.register(&customer{id: 155, name: "phonghm"})
	newItem.register(&customer{id: 159, name: "satan"})
	newItem.notiAll()
}
