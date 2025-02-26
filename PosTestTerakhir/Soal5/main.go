package main

import "fmt"

type player struct {
	name  string
	goals int
	assists int
}

func main() {
	var n int
	fmt.Scan(&n)
	players := make([]player, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&players[i].name, &players[i].goals, &players[i].assists)
	}

	insertionSort(players, n)

	for _, p := range players {
		fmt.Println(p.name, p.goals, p.assists)
	}
}

func insertionSort(players []player, n int) {
	for i := 1; i < n; i++ {
		key := players[i]
		j := i - 1
		for j >= 0 && compare(players[j], key) {
			players[j+1] = players[j]
			j--
		}
		players[j+1] = key
	}
}

func compare(a, b player) bool {
	if a.goals != b.goals {
		return a.goals < b.goals
	}
	if a.assists != b.assists {
		return a.assists < b.assists
	}
	return a.name > b.name
}