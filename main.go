package main

import (
	"fmt" //Package Untuk Memformat Input dan Output
	"math"//Package Untuk Menyediakan Fungsi-Fungsi MTK Dasar dan Lanjutan
)

//Inner Function Soal 1
func soal_1() {
	var (
		bj  = 54  //BJ = BEBEK JANTAN		//SATU ANGGKA SETELAH BB/BJ ADALAH BULAN KE-N
		bb  = 168 //BB = BEBEK BETINA
		bb1 = bb + bj
		bj1 = bj * 1 / 3
		bj2 = bj1 + bb1
		bb2 = bb1 * 5/6
		bb3 = bj2
		bj3 = bj2 * 1/5

	)
//Code Untuk Menampilkan Output
	fmt.Println(bb1, bj1)
	fmt.Println(bj2, bb2)
	fmt.Println(bb3, bj3)
}

//Inner Function Soal 2
func soal_2() {

//Menggunakan Variable Secara Langsung Karena Variable r2&T Tidak Memiliki Value Pasti
	var r2, T float64	//Menggunakan Float64 Karena Value dari r2&T Harus Berupa Angka

	//Input Nilai Diameter
	fmt.Print("masukan diameter(cm):")
	fmt.Scan(&r2)

	//Input Nilai Tinggi Tabung
	fmt.Print("masukan tinggi tabung(cm):")
	fmt.Scan(&T)
	fmt.Println()

	//Code Jari" Lingkaran
	var r = r2 / 2

	var (
		//Variable & Code Luas Lingkaran
		L  = math.Pi * r * r
		//Variable & Code Keliling Lingkaran
		K  = 2 * math.Pi * r
		//Variable & Code Volume Tabung
		V  = math.Pi * r * r * T
		//Variable & Code Luas Permukaan Tabung
		Lp = 2 * math.Pi * r * (r + T)
	)
	//Print Hasil Luas Lingkaran
	fmt.Printf("Luas lingkaran(cm2) = %.2f\n", L)
	//Print Hasi Keliling Lingkaran
	fmt.Printf("Keliling lingkaran(cm) = %.2f\n", K)
	//Print Hasil Volume Tabung
	fmt.Printf("Volume Tabunng(cm3) = %.2f\n", V)
	//Print Luas Permukaan
	fmt.Printf("Luas Permukaan(cm2) = %.2f\n", Lp)

}

//Core Function
func main() {
	fmt.Println("Jawaban soal 1")
	soal_1()

	fmt.Println()

	fmt.Println("Jawaban soal 2")
	soal_2()
}
