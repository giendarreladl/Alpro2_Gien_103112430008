<h1 align="center">Laporan Praktikum Modul 2 <br> REVIEW PENGENALAN PEMROGRAMAN </h1> 

___
<h4 align="center">Gien Darrel Adli - 103112430008 </h4>

### Unguided
#### 2A-01.
Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program.
Silakan masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran
yang diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?
```go
package main

import "fmt"

func main() {

    var (
        satu, dua, tiga string
        temp            string
    )

    fmt.Print("Masukan input string: ")
    fmt.Scanln(&satu)
    fmt.Print("Masukan input string: ")
    fmt.Scanln(&dua)
    fmt.Print("Masukan input string: ")
    fmt.Scanln(&tiga)
    fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

    temp = satu
    satu = dua
    dua = tiga
    tiga = temp

    fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
![[2A-01.png]]
Program ini menerima tiga input string dari pengguna, lalu mencetak urutan awalnya. Setelah itu, program menukar posisi string dengan pergeseran ke kiri, di mana nilai pertama berpindah ke posisi kedua, nilai kedua ke posisi ketiga, dan nilai ketiga kembali ke posisi pertama menggunakan variabel sementara. Terakhir, program mencetak urutan string setelah pertukaran.

#### 2A-02
Tahun kabisat adalah tahun yang habis dibagi 400 atau habis dibagi 4 tetapi tidak habis
dibagi 100. Buatlah sebuah program yang menerima input sebuah bilangan bulat dan
memeriksa apakah bilangan tersebut merupakan tahun kabisat (true) atau bukan (false)
```go
package main

import "fmt"  

func main() {

    var tahun int
    fmt.Print("Masukkan tahun: ")
    fmt.Scan(&tahun)

    kabisat := (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)

    fmt.Println("Kabisat:", kabisat)

}
```
![[2A-02.png]]
Program ini meminta pengguna untuk memasukkan sebuah tahun, lalu menentukan apakah tahun tersebut merupakan tahun kabisat berdasarkan aturan kalender Gregorian. Tahun kabisat adalah tahun yang habis dibagi 400 atau tahun yang habis dibagi 4 tetapi tidak habis dibagi 100. Program menghitung berdasarkan aturan ini dan menyimpan hasilnya dalam variabel `kabisat`, yang akan bernilai `true` jika tahun tersebut kabisat dan `false` jika bukan. Hasil akhirnya ditampilkan di layar dalam format "Kabisat: true" atau "Kabisat: false".

#### 2A-03
Buat program Bola yang menerima input jari-jari suatu bola (bilangan bulat). Tampilkan
Volume dan Luas kulit bola. 𝑣𝑜𝑙𝑢𝑚𝑒𝑏𝑜𝑙𝑎 = 4/3 𝜋𝑟3 dan 𝑙𝑢𝑎𝑠𝑏𝑜𝑙𝑎 = 4𝜋𝑟2 (π ≈3.1415926535).
```go
package main 

import "fmt"  

func main() {

    var r float64
    fmt.Print("Jejari = ")
    fmt.Scan(&r)

    pi := 3.1415926535
    v := (4.0 / 3.0) * pi * (r * r * r)
    l := 4 * pi * (r * r)

    fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", r, v, l)
    
}
```
![[2A-03.png]]
Program ini menghitung volume dan luas permukaan bola berdasarkan jejari yang dimasukkan pengguna. Nilai π digunakan untuk akurasi, lalu perhitungan dilakukan dengan rumus matematika yang sesuai. Hasil akhirnya ditampilkan dalam format desimal yang rapi.

#### 2A-04
Dibaca nilai temperatur dalam derajat Celsius. Nyatakan temperatur tersebut dalam
Fahrenheit
𝐶𝑒𝑙𝑠𝑖𝑢𝑠 = (𝐹𝑎ℎ𝑟𝑒𝑛ℎ𝑒𝑖𝑡 − 32) × 5/9 𝑅𝑒𝑎𝑚𝑢𝑟 = 𝐶𝑒𝑙𝑐𝑖𝑢𝑠 × 4/5 𝐾𝑒𝑙𝑣𝑖𝑛 = (𝐹𝑎ℎ𝑟𝑒𝑛ℎ𝑒𝑖𝑡 + 459.67) × 5/9
```go
package main

import "fmt"

func main() {

    var c float64
    fmt.Print("Temperatur Celsius: ")
    fmt.Scan(&c)

    r := (4.0 / 5.0) * c
    f := (9.0/5.0)*c + 32
    k := c + 273

    fmt.Printf("Derajat Reamur: %.0f\n", r)
    fmt.Printf("Derajat Fahrenheit: %.0f\n", f)
    fmt.Printf("Derajat Kelvin: %.0f\n", k)

}
```
![[2A-04.png]]
Program ini mengonversi suhu dari Celsius ke Reamur, Fahrenheit, dan Kelvin berdasarkan nilai yang dimasukkan oleh pengguna. Perhitungan dilakukan menggunakan rumus konversi standar, di mana Reamur diperoleh dengan mengalikan Celsius dengan 4/5, Fahrenheit dihitung dengan (9/5 × Celsius) + 32, dan Kelvin didapat dengan menambahkan 273 ke Celsius. Hasil konversi kemudian ditampilkan dalam format tanpa angka desimal.

### 2A-05
Tipe karakter sebenarnya hanya apa yang tampak dalam tampilan. Di dalamnya
tersimpan dalam bentuk biner 8 bit (byte) atau 32 bit (rune) saja. Buat program ASCII yang akan membaca 5 buat data integer dan mencetaknya dalam format karakter. Kemudian membaca 3 buah data karakter dan mencetak 3 buah karakter setelah karakter tersebut (menurut tabel ASCII)
Masukan terdiri dari dua baris. Baris pertama berisi 5 buah data integer. Data integer mempunyai nilai antara 32 s.d. 127. Baris kedua berisi 3 buah karakter yang berdampingan satu dengan yang lain (tanpa dipisahkan spasi). Keluaran juga terdiri dari dua baris. Baris pertama berisi 5 buah representasi karakter dari data yang diberikan, yang berdampingan satu dengan lain, tanpa dipisahkan spasi. Baris kedua berisi 3 buah karakter (juga tidak dipisahkan oleh spasi).
```go
package main

import "fmt"

func main() {

	var a1, a2, a3, a4, a5 int
	var y1, y2, y3 rune

	fmt.Scan(&a1, &a2, &a3, &a4, &a5)
	fmt.Scanln()
	fmt.Scanf("%c%c%c\n", &y1, &y2, &y3)

	y1 += 1
	y2 += 1
	y3 += 1

	fmt.Printf("%c%c%c%c%c\n", a1, a2, a3, a4, a5)
	fmt.Printf("%c%c%c\n", y1, y2, y3)
}
```
![[2A-05.png]]
Program ini membaca lima bilangan bulat dan tiga karakter sebagai input. Lima bilangan pertama dikonversi ke karakter ASCII dan ditampilkan sebagai teks. Tiga karakter berikutnya diubah ke karakter selanjutnya dalam urutan ASCII sebelum ditampilkan.


### 2B-01
Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan
praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana
susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta
untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan
warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’,
‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang.
Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi
sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan
warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false
untuk urutan warna lainnya.
Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah
adalah input/read):
```go
package main

import "fmt"

func main() {

    var warna1, warna2, warna3, warna4 string
    berhasil := true
    
    for i := 1; i <= 5; i++ {
        fmt.Scan(&warna1, &warna2, &warna3, &warna4)

        if warna1 != "merah" || warna2 != "kuning" || warna3 != "hijau" || warna4 != "ungu" {
            berhasil = false

        }

    }
    fmt.Println(berhasil)
}
```
![[2B-01.png]]Program ini membaca empat warna dari input sebanyak lima kali dan mengecek apakah setiap set warna selalu berurutan **"merah", "kuning", "hijau", "ungu"**. Jika ada set yang tidak sesuai, variabel berhasil diubah menjadi false. Setelah perulangan selesai, hasil akhir (`true` atau `false`) ditampilkan.

### 2B-02
Suatu pita (string) berisi kumpulan nama-nama bunga yang dipisahkan oleh spasi dan ‘–‘, contoh pita diilustrasikan seperti berikut ini. Pita: mawar – melati – tulip – teratai – kamboja – anggrek. Buatlah sebuah program yang menerima input sebuah bilangan bulat positif (dan tidak nol) N, kemudian program akan meminta input berupa nama bunga secara berulang sebanyak N kali dan nama tersebut disimpan ke dalam pita. (Petunjuk: gunakan operasi penggabungan string dengan operator “+” ). Tampilkan isi pita setelah proses input selesai.
```go
package main

import (
"fmt"
)

func main() {

	var n int
    var bunga, pita string
    jumlah := 0

    fmt.Print("N: ")
    fmt.Scan(&n)
    
    for i := 1; i <= n; i++ {
        fmt.Printf("Bunga %d: ", i)
        fmt.Scan(&bunga)
        if bunga == "SELESAI" {
            break
        }

        if jumlah > 0 {
            pita += " - "
        }

        pita += bunga
        jumlah++
    }
    fmt.Println("Pita:", pita)
    fmt.Println("Bunga:", jumlah)
}
```
![[2B-02.png]]Program ini membaca nama bunga sebanyak `n` kali dan menyusunnya dalam format berpisah dengan tanda " - ". Jika input "SELESAI" dimasukkan, perulangan berhenti lebih awal. Setelah itu, program menampilkan rangkaian bunga dalam bentuk pita serta jumlah bunga yang telah dimasukkan.

### 2B-03
Setiap hari Pak Andi membawa banyak barang belanjaan dari pasar dengan mengendarai
sepeda motor. Barang belanjaan tersebut dibawa dalam kantong terpal di kiri-kanan motor. Sepeda motor tidak akan oleng jika selisih berat barang di kedua kantong sisi tidak lebih dari 9 kg. Buatlah program Pak Andi yang menerima input dua buah bilangan real positif yang menyatakan berat total masing-masing isi kantong terpal. Program akan terus meminta input bilangan tersebut hingga salah satu kantong terpal berisi 9 kg atau lebih.
```go
package main

import "fmt"

func main() {

    var beratKiri, beratKanan, bedaBerat float32
    var isOleng bool

    for {

        fmt.Print("Masukan berat belanjaan di kedua kantong: ")
        fmt.Scan(&beratKiri, &beratKanan)

        if beratKiri < 0 || beratKanan < 0 {
            fmt.Println("Proses selesai")
            break
        }

        if beratKiri+beratKanan > 150 {
            fmt.Println("Proses selesai")
            break
        }

        if beratKiri > beratKanan {
            bedaBerat = beratKiri - beratKanan
        } else {
            bedaBerat = beratKanan - beratKiri
        }

        isOleng = bedaBerat >= 9
        fmt.Println("Sepeda motor pak Andi akan oleng:", isOleng)
    }
}
```
![[2B-03.png]]
Program ini membaca berat belanjaan di dua kantong dan menentukan apakah sepeda motor oleng berdasarkan selisih berat (≥9). Program berhenti jika berat total >150 atau ada input negatif.

### 2B-04-a
Diberikan sebuah persamaan sebagai berikut ini.
𝑓(𝑘) = (4𝑘 + 2)^2 / (4𝑘 + 1)(4𝑘 + 3)
Buatlah sebuah program yang menerima input sebuah bilangan sebagai K, kemudian
menghitung dan menampilkan nilai f(K) sesuai persamaan di atas.
```go
package main

import "fmt" 

func main() {
    var K float64
    fmt.Print("Nilai K = ")
    fmt.Scan(&K)

    fk := ((4*K + 2) * (4*K + 2)) / ((4*K + 1) * (4*K + 3))

    fmt.Printf("Nilai f(K) = %.10f\n", fk)

}
```
![[2B-04-a.png]]Program ini menghitung f(K) menggunakan rumus matematika dan menampilkan hasil dengan 10 desimal.

### 2B-04-b
√2 merupakan bilangan irasional. Meskipun demikian, nilai tersebut dapat dihampiri
dengan rumus berikut:
√2 = ∏ (4𝑘 + 2)^2 / (4𝑘 + 1)(4𝑘 + 3) ∞ 𝑘=0
Modifikasi program sebelumnya yang menerima input integer 𝐾 dan menghitung √2
untuk 𝐾 tersebut. Hampiran √2 dituliskan dalam ketelitian 10 angka di belakang koma.
```go
package main

import "fmt"

func main() {

    var K float64
    fmt.Print("Nilai K = ")
    fmt.Scan(&K)

    hasil := 1.0

    for i := 0.0; i < K; i++ {
        hasil *= ((4*i + 2) * (4*i + 2)) / ((4*i + 1) * (4*i + 3))
    }
    fmt.Printf("Nilai akar 2 = %.10f\n", hasil)
}
```
![[2B-04-b.png]]rogram ini menghitung nilai perkiraan akar 2 menggunakan rumus iteratif berdasarkan nilai K yang dimasukkan pengguna. Program mengalikan hasil dengan rasio tertentu dalam loop hingga K iterasi, lalu mencetak hasilnya dengan presisi 10 desimal.
### 2C-01
PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka,
buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan
ketentuan sebagai berikut!
Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam
gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari
500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari
500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat
(yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg
```go
package main

import "fmt"

func main() {

    var berat int
    fmt.Print("Berat parsel (gram): ")
    fmt.Scan(&berat)

    kg := berat / 1000
    gram := berat % 1000

    biayaKg := kg * 10000
    biayaGram := 0

    if gram > 0 {
        if gram >= 500 {
            biayaGram = gram * 5
        } else {
            biayaGram = gram * 15
        }
    }
    
    if kg >= 10 {
        biayaGram = 0
    }
    
    total := biayaKg + biayaGram
    fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gram)
    fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaGram)
    fmt.Printf("Total biaya: Rp. %d\n", total)
}
```
![[2C-01.png]]Program ini menghitung biaya pengiriman parsel berdasarkan berat yang dimasukkan pengguna dalam gram. Berat dikonversi ke kilogram dan gram, lalu biaya dihitung sesuai tarif yang ditentukan. Jika berat mencapai 10 kg atau lebih, biaya tambahan untuk gram diabaikan. Program menampilkan detail berat, rincian biaya, dan total biaya pengiriman.

### 2C-02
a. Jika nam diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesuai spesifikasi soal?

Jawab : Eksekusi ini **sesuai** dengan spesifikasi soal, karena 80.1 memang lebih dari 80 dan seharusnya mendapat nilai "A".

b.Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program seharusnya!

Jawab: 
1. Kesalahan penggunaan operator perbandingan
    - Dalam soal, beberapa kondisi menggunakan bentuk X < NAM <= Y, tetapi dalam program, kondisi ini tidak ditulis dengan benar.
    - Misalnya, untuk rentang 65 ≤ NAM < 72.5, seharusnya nam >= 65 && nam < 72.5, tetapi dalam program hanya nam > 65.
2. Tidak menggunakan `else if` dengan benar
    - Program menggunakan if secara terpisah, sehingga beberapa kondisi bisa dievaluasi secara bersamaan dan menghasilkan kesalahan logika.
3. Kesalahan penulisan kondisi untuk beberapa rentang nilai
    - Rentang 72.5 ≤ NAM ≤ 80 tidak dimasukkan dalam kondisi yang benar.
    - 50 ≤ NAM < 57.5 seharusnya nam >= 50 && nam < 57.5, bukan nam > 50.

Alur program yang benar:
1. Jika NAM > 80, maka NMK = 'A'.
2. Jika 72.5 ≤ NAM ≤ 80, maka NMK = 'AB'.
3. Jika 65 ≤ NAM < 72.5, maka NMK = 'B'.
4. Jika 57.5 ≤ NAM < 65, maka NMK = 'BC'.
5. Jika 50 ≤ NAM < 57.5, maka NMK = 'C'.
6. Jika 40 ≤ NAM < 50, maka NMK = 'D'.
7. Jika NAM < 40, maka NMK = 'E'.

c. Perbaiki program tersebut! Ujilah dengan masukan: 93.5; 70.6; dan 49.5. Seharusnya keluaran yang diperoleh adalah ‘A’, ‘B’, dan ‘D’.

Jawab : 
```go
package main

import "fmt"

func main() {
    var nam float64
    var nmk string

    fmt.Print("Nilai akhir mata kuliah: ")
    fmt.Scan(&nam)

    if nam > 80 {
        nmk = "A"
    } else if nam >= 72.5 {
        nmk = "AB"
    } else if nam >= 65 {
        nmk = "B"
    } else if nam >= 57.5 {
        nmk = "BC"
    } else if nam >= 50 {
        nmk = "C"
    } else if nam >= 40 {
        nmk = "D"
    } else {
        nmk = "E"
    }

    fmt.Println("Nilai mata kuliah:", nmk)
}
```
![[2C-02.png]]Program ini menerima input berupa nilai akhir mata kuliah (NAM) dalam bentuk bilangan riil, lalu menentukan nilai mata kuliah (NMK) berdasarkan rentang yang telah ditentukan. Program menggunakan struktur if-else if untuk mengklasifikasikan NAM ke dalam kategori NMK, kemudian mencetak hasilnya.

### 2C-03
```go
package main

import "fmt"

func main() {

    var b int
    fmt.Print("Bilangan: ")
    fmt.Scan(&b)

    fmt.Print("Faktor: ")
    count := 0

    for i := 1; i <= b; i++ {
        if b%i == 0 {
            fmt.Print(i, " ")
            count++
        }
    }
    fmt.Println()

    isPrime := count == 2

    fmt.Println("Prima:", isPrime)
}
```
![[2C-03.png]]
Program ini menerima input bilangan bulat b > 1, mencari dan mencetak semua faktor bilangan tersebut. Program juga menentukan apakah bilangan tersebut merupakan bilangan prima dengan menghitung jumlah faktornya. Jika bilangan memiliki tepat dua faktor (1 dan dirinya sendiri), maka program mencetak `true` untuk bilangan prima, selain itu mencetak false.