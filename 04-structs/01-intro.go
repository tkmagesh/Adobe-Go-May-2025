package main

import "fmt"

/* product => id, name, cost */

func main() {
	/*
		var p struct {
			Id   int
			Name string
			Cost float64
		}
		p.Id = 100
		p.Name = "Pen"
		p.Cost = 5
	*/

	var p struct {
		Id   int
		Name string
		Cost float64
	} = struct {
		Id   int
		Name string
		Cost float64
	}{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}
	// fmt.Println(p)
	// fmt.Printf("%#v\n", p)
	fmt.Printf("%+v\n", p)

	// struct instances are values
	var p1 = struct {
		Id   int
		Name string
		Cost float64
	}{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}

	var p2 = struct {
		Id   int
		Name string
		Cost float64
	}{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}
	fmt.Println("p1 == p2 :", p1 == p2)

	fmt.Println("Before applying discount")
	fmt.Println(Format(p))
	ApplyDiscount(&p, 10)
	fmt.Println("After applying discount")
	fmt.Println(Format(p))

}

/* Write the following functions
1. Format(product) => "Id = 100, Name = Pen, Cost = 10"
2. ApplyDiscount(product, discountPercentage) => no return result. updates the cost after applying the discount
*/

func Format(p struct {
	Id   int
	Name string
	Cost float64
}) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func ApplyDiscount(p *struct {
	Id   int
	Name string
	Cost float64
}, discountPercentage float64) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
