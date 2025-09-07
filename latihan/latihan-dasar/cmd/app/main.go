package main

import "fmt"

// fungsi untuk perkalian
func kali(a, b int) int {
	return a * b
}

func main() {
	// --- Biodata ---
	nama := "Afes Oktavianus"
	umur := 38
	hobi := "Ngoding & Membaca"

	fmt.Println("=== Biodata ===")
	fmt.Println("Nama :", nama)
	fmt.Println("Umur :", umur)
	fmt.Println("Hobi :", hobi)
	fmt.Println()

	// --- Perkalian ---
	fmt.Println("=== Perkalian ===")
	x, y := 7, 5
	hasil := kali(x, y)
	fmt.Printf("%d x %d = %d\n", x, y, hasil)
	fmt.Println()

	// --- Angka Genap 1-10 ---
	fmt.Println("=== Angka Genap 1-10 ===")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
