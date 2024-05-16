package main

import "fmt"

// Definisi tipe Ban
type Ban string

// Definisi struct Pintu
type Pintu struct {
	ketukan string
	buka    string
}

// Metode untuk mengetuk pintu
func (p *Pintu) Ketuk() string {
	return p.ketukan
}

// Metode untuk membuka pintu
func (p *Pintu) Buka() string {
	return p.buka
}

// Definisi struct Mobil
type Mobil struct {
	roda  []Ban
	pintu [2]Pintu
}

// Metode untuk menginisialisasi Mobil
func NewMobil() *Mobil {
	pintuKanan := Pintu{ketukan: "Knock", buka: "beep"}
	pintuKiri := Pintu{ketukan: "beep", buka: "Knock"}
	return &Mobil{
		roda:  []Ban{"ban karet", "ban karet", "ban karet", "ban karet"},
		pintu: [2]Pintu{pintuKanan, pintuKiri},
	}
}

// Metode untuk mengganti roda
func (m *Mobil) GantiBan(ban Ban, indeks int) {
	if indeks >= 0 && indeks < len(m.roda) {
		m.roda[indeks] = ban
	} else {
		fmt.Println("Indeks roda tidak valid")
	}
}

func main() {
	// Membuat objek mobil baru
	mobilBaru := NewMobil()

	// Menampilkan roda mobil
	fmt.Println(mobilBaru.roda)

	// Mengganti roda ke-3 dengan ban kayu
	mobilBaru.GantiBan("ban kayu", 2)

	// Mengganti roda ke-1 dengan ban kayu
	mobilBaru.GantiBan("ban besi", 0)

	// Menampilkan roda mobil setelah ganti
	fmt.Println(mobilBaru.roda)

	// Mengetuk dan membuka pintu kanan
	fmt.Println(mobilBaru.pintu[0].Ketuk())
	fmt.Println(mobilBaru.pintu[0].Buka())

	// Mengetuk dan membuka pintu kiri
	fmt.Println(mobilBaru.pintu[1].Ketuk())
	fmt.Println(mobilBaru.pintu[1].Buka())
}
