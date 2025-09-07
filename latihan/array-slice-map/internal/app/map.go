package app

import (
	"fmt"
)

func HandlerMap() {
	var mahasiswa = map[string]string{
		"20231001": "Budi Santoso",
		"20231002": "Siti Aminah",
		"20231003": "Andi Wijaya",
	}

	// Menampilkan seluruh data mahasiswa
	fmt.Println("=== Data Mahasiswa ===")
	for nim, nama := range mahasiswa {
		fmt.Printf("NIM: %s, Nama: %s\n", nim, nama)
	}

	// Mengakses data berdasarkan key
	fmt.Println("\nMahasiswa dengan NIM 20231002:", mahasiswa["20231002"])

	// Menambahkan data baru
	mahasiswa["20231004"] = "Dewi Lestari"

	// Menghapus data
	delete(mahasiswa, "20231001")

	fmt.Println("\n=== Data Mahasiswa Setelah Update ===")
	for nim, nama := range mahasiswa {
		fmt.Printf("NIM: %s, Nama: %s\n", nim, nama)
	}
}
