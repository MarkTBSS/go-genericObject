package main

import "fmt"

type GameOrMovie interface {
	getPrice() int
}

type Game struct {
	Name     string
	Platform string
	Price    int
}

type Movie struct {
	Name  string
	Price int
}

func (g Game) getPrice() int {
	return g.Price
}

func (m Movie) getPrice() int {
	return m.Price
}

func sum[V GameOrMovie](objs []V) int {
	var total int
	for _, obj := range objs {
		total += obj.getPrice()
	}
	return total
}

func main() {
	games := []Game{
		{Name: "League of Legends", Platform: "PC", Price: 1000},
		{Name: "League of Legends", Platform: "PS4", Price: 2000},
		{Name: "League of Legends", Platform: "Xbox", Price: 3000},
	}

	movies := []Movie{
		{Name: "Star Wars", Price: 1000},
		{Name: "Star Wars", Price: 2000},
		{Name: "Star Wars", Price: 3000},
	}

	fmt.Println(sum(games))
	fmt.Println(sum(movies))
}
