package main

import "fmt"

func main() {
	ak47, _ := getGun("ak47")
	maverick, _ := getGun("maverick")
	printDetail(ak47)
	printDetail(maverick)
}

func printDetail(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
