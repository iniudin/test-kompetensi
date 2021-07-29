package main

import "fmt"

func main() {
	var jam, menit, detik int
	var suffix string

	fmt.Println("Masukkan Jam: (AM/PM)")
	fmt.Scanf("%d:%d:%d%s", &jam, &menit, &detik, &suffix)

	// Jika suffix AM ubah jam menjadi 0,
	// sebaliknya jika suffix PM dan jam != 12, maka tambahkan 12 jam
	if suffix == "AM" && jam == 12 {
		jam = 0

	} else if suffix == "PM" && jam != 12 {
		jam += 12
	}

	fmt.Println("Hasil konversi Waktu: ")
	fmt.Printf("%02d:%02d:%02d", jam, menit, detik)
}
