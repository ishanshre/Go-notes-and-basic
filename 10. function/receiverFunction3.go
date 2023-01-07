package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string, items map[string]float64) bill {
	b := bill{
		name:  name,
		items: items,
		tip:   0,
	}
	return b
}

func (p bill) format() string {
	fs := "Bill break down \n"
	var total float64 = 0
	for k, v := range p.items {
		fs += fmt.Sprintf("%v: ... $%v \n", k, v)
		total += v
	}
	fs += fmt.Sprintf("%v: ...$%0.2f\n", "total", total)
	fs += fmt.Sprintf("%v: ...$%0.2f\n", "tip", p.tip)
	return fs
}

// while updating or adding in receiver we use pointer
func (p *bill) update(tip float64) {
	p.tip = tip
}

func (p *bill) addItem(name string, price float64) {
	p.items[name] = price
}
