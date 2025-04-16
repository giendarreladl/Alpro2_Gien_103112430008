<h1 align="center">Laporan Praktikum Modul 10 <br>  PENCARIAN NILAI EKSTRIM PADA HIMPUNAN DATA</h1> 

___
<h4 align="center">Gien Darrel Adli - 103112430008 </h4>

### Unguided

### Soal-1. 
Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar.
Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak
kelinci yang akan dijual.
Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan
bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya
N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual.
Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan
terbesar.
```go
package main

import (
	"fmt"
)

type Kelinci struct {
	nama  string
	berat float64
}

func main() {
	var n int
	var kelinci [1000]Kelinci

	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Nama kelinci ke-%d: ", i+1)
		fmt.Scan(&kelinci[i].nama)
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&kelinci[i].berat)
	}

	min := kelinci[0]
	max := kelinci[0]

	for i := 1; i < n; i++ {
		if kelinci[i].berat < min.berat {
			min = kelinci[i]
		}
		if kelinci[i].berat > max.berat {
			max = kelinci[i]
		}
	}

	fmt.Printf("Kelinci paling ringan: %s dengan berat %.2f\n", min.nama, min.berat)
	fmt.Printf("Kelinci paling berat : %s dengan berat %.2f\n", max.nama, max.berat)
}
```
![](s(1).png)

>Program ini digunakan untuk mencatat data sejumlah anak kelinci berupa nama dan berat badannya. Data disimpan dalam array bertipe struct Kelinci yang memiliki dua field: nama dan berat. Setelah pengguna memasukkan jumlah kelinci dan mengisi data masing-masing kelinci, program mencari kelinci dengan berat paling ringan dan paling berat dengan membandingkan nilai berat dari setiap elemen array. Hasil akhir menampilkan nama dan berat kelinci paling ringan dan paling berat.

### Soal-2
Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program
ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan
dijual.
Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan
y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya
ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil
yang menyatakan banyaknya ikan yang akan dijual.
Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan
total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang
dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah
sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.
```go
package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah ikan dan jumlah ikan per wadah (x y): ")
	fmt.Scan(&x, &y)

	beratIkan := make([]float64, x)
	fmt.Println("Masukkan berat ikan satu per satu:")
	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
	}

	jumlahWadah := (x + y - 1) / y
	totalPerWadah := make([]float64, jumlahWadah)

	for i := 0; i < x; i++ {
		indexWadah := i / y
		totalPerWadah[indexWadah] += beratIkan[i]
	}

	var totalSemua float64
	for _, berat := range beratIkan {
		totalSemua += berat
	}
	rataRata := totalSemua / float64(jumlahWadah)

	fmt.Println("\nTotal berat per wadah:")
	for i, total := range totalPerWadah {
		fmt.Printf("Wadah %d: %.2f\n", i+1, total)
	}
	fmt.Printf("Rata-rata berat ikan per wadah: %.2f\n", rataRata)
}

```
![](s(2).png)

>Program ini digunakan untuk menghitung total dan rata-rata berat ikan dalam beberapa wadah. Pengguna diminta memasukkan jumlah total ikan dan jumlah ikan per wadah. Berat masing-masing ikan kemudian diinput satu per satu dan disimpan dalam slice. Program menghitung jumlah wadah yang dibutuhkan, lalu menjumlahkan berat ikan untuk setiap wadah sesuai urutan. Setelah itu, program menghitung rata-rata total berat ikan terhadap jumlah wadah. Hasil akhirnya adalah daftar total berat per wadah dan rata-rata berat ikan per wadah.

### Soal-3
Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data
berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data
yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.
Buatlah program dengan spesifikasi subprogram sebagai berikut:
```go
package main

import (
	"fmt"
)

const maxData = 100

type arrBalita [maxData]float64

func hitungMinMax(arr arrBalita, n int, bMin *float64, bMax *float64) {
	*bMin = arr[0]
	*bMax = arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < *bMin {
			*bMin = arr[i]
		}
		if arr[i] > *bMax {
			*bMax = arr[i]
		}
	}
}

func hitungRata(arr arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	return total / float64(n)
}

func main() {
	var data arrBalita
	var n int
	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	var min, max float64
	hitungMinMax(data, n, &min, &max)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)

	rata := hitungRata(data, n)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rata)
}
```
![](s(3).png)

>Program ini mencatat data berat beberapa balita, lalu menghitung berat minimum, maksimum, dan rata-ratanya. Berat balita disimpan dalam array bertipe khusus arrBalita yang berisi maksimal 100 elemen bertipe float64. Fungsi hitungMinMax digunakan untuk mencari berat balita paling ringan dan paling berat, sedangkan fungsi hitungRata menghitung rata-rata dari seluruh data berat. Di fungsi main, pengguna diminta memasukkan jumlah data dan berat tiap balita, kemudian hasil perhitungan ditampilkan ke layar.
