package main

import "fmt"

func main() {
	var n, sumNum, sumActual int

	//input untuk nilai n
	fmt.Print("Masukan bilangan untuk baris pertama : ")
	fmt.Scan(&n)

	//disini karena nilai n itu adalah total panjang jadi saya buat ke dalam array
	myArray := make([]int, n)

	//memasukan nilai untuk di baris ke 2 dan disimpan ke array tadi
	fmt.Print("Masukan nilai untuk baris ke dua : ")
	for i := range myArray {
		fmt.Scan(&myArray[i])

		// menambahkan logic if, jika nilai yang kita kasih di baris ke 2 itu lebih dari yang kita inginkan maka error
		// if myArray[i] > n-1 {
		// 	fmt.Println("Nilai yang diberikan lebih dari nilai yang di tetapkan di awal")
		// 	i--
		// }
	}

	//untuk mencari angka yang hilang perlu beberapa cara
	// pertama kita cari tahu dulu jumlah dari rentang angka 1 sampai ke n
	sumNum = n * (n + 1) / 2

	// selanjutnya kita menjumlahkan nilai yang kita masukan untuk baris ke 2
	for i := 0; i < n; i++ {
		sumActual += myArray[i]
	}

	// yang ke 3 mencari nilai yang hilang tersebut dengan cara jumlah semestinya di kurang dengan jumlah data yang ada
	missingNumber := sumNum - sumActual

	// terakhir menampilkan nilai yang hilang
	fmt.Printf("Angka yang hilang dari array di atas adalah: %d\n", missingNumber)
}
