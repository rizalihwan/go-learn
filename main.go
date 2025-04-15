package main

import "fmt"

func main() {
	fmt.Println("increment result: ", incremet(50, 10))
	fmt.Println(defineStatus("Rizal Ihwan", "21"))
}

func incremet(a int, b int) int {
	return a + b
}

func defineStatus(nama string, umur string) string {
	return fmt.Sprintf("Nama saya %s, umur saya %s", nama, umur)
}
