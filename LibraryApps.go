package main

import "fmt"

const NMAX int = 100

type Book struct {
	info    [NMAX]detailBook
	nElemen int
}

type User struct {
	info    [NMAX]detailUser
	nElemen int
}

type detailUser struct {
	name          string
	denda         int
	tgl, bln, thn int
}

type detailBook struct {
	name        string
	catID       int
	bookcode    string
	borrowed    bool
	nborrow     int
	currentUser int
}

func main() {
	menu()
}

func menu() {
	var A Book
	var B User
	var x int
	x = -1

	for x != 0 {
		header()
		fmt.Println("||===============================||")
		fmt.Println("||          Pilih opsi:          ||")
		fmt.Println("||1. Add Book                    ||")
		fmt.Println("||2. Edit Book                   ||")
		fmt.Println("||3. Delete Book                 ||")
		fmt.Println("||4. View Book                   ||")
		fmt.Println("||5. Borrow a Book               ||")
		fmt.Println("||6. Find a Book                 ||")
		fmt.Println("||7. Favorite Book               ||")
		fmt.Println("||8. Return Book                 ||")
		fmt.Println("||0. Exit Menu                   ||")
		fmt.Println("||===============================||")
		fmt.Print("Masukan nomor sesuai opsi yang anda ingin pilih: ")
		fmt.Scan(&x)
		fmt.Println()

		if x == 1 {
			addBook(&A)
		} else if x == 2 {
			viewBook(A)
			editBook(&A)
		} else if x == 3 {
			viewBook(A)
			deleteBook(&A)
		} else if x == 4 {
			viewBook(A)
		} else if x == 5 {
			borrowBook(&A, &B)
		} else if x == 6 {
			findBook(A)
		} else if x == 7 {
			sortFav(&A)
		} else if x == 8 {
			returnBook(&A, &B)
		} else if x < 0 && x > 5 {
			x = -1
		}
	}

}

//A opsi untuk librarian
func header() {
	fmt.Println("* --------------------------------------- *")
	fmt.Println("* -------- Aplikasi Perpustakaan -------- *")
	fmt.Println("* -- Tugas Besar Algoritma Pemrograman -- *")
	fmt.Println("* --------------------------------------- *")
}

func addBook(set *Book) {
	var input, code string
	var tempCat int
	fmt.Println("Berikan simbol _ sebagai spasi untuk judul buku yang mengandung lebih dari 1 kata")
	fmt.Println("Masukkan Judul:")
	fmt.Scan(&input)
	fmt.Println("Masukkan Code Book:")
	fmt.Scan(&code)

	//STOP untuk berhenti melakukan inputan data
	//judul bisa sama asalkan codebook-nya berbeda
	for input != "STOP" && code != "STOP" {

		if set.nElemen < NMAX && search(*set, input) == -1 || searchCode(set, code) == -1 {
			set.info[set.nElemen].name = input
			set.info[set.nElemen].bookcode = code
			fmt.Println("Pilih Kategori:")
			fmt.Println("1. Fiksi")
			fmt.Println("2. Non Fiksi")

			fmt.Println("Masukan nomor sesuai kategori yang dipilih:")
			fmt.Scan(&tempCat)
			for tempCat < 1 && tempCat > 2 {
				fmt.Println("Kategori tidak sesuai")
				fmt.Scan(&tempCat)
			}
			set.info[set.nElemen].catID = tempCat
			set.info[set.nElemen].borrowed = false
			set.nElemen++
		} else if set.nElemen < NMAX && search(*set, input) != -1 && searchCode(set, code) != -1 {
			fmt.Println("Buku", input, "sudah terdaftar")
		}
		fmt.Println("Masukkan Judul:")
		fmt.Scan(&input)
		fmt.Println("Masukkan Code Book:")
		fmt.Scan(&code)
	}
}

func editBook(set *Book) {
	var bef string
	var aft string
	var code string
	var idx int
	var tempCat int

	fmt.Println()
	fmt.Println("Pilih judul buku yang ingin anda ubah:")
	fmt.Scan(&bef)
	idx = search(*set, bef)
	if idx != -1 {
		//i jadi suatu idx yg udh ditemukan
		fmt.Println("Masukan Judul Buku baru:")
		fmt.Scan(&aft)
		set.info[idx].name = aft
		fmt.Println("Masukan Kode Buku baru:")
		fmt.Scan(&code)
		set.info[idx].bookcode = code

		fmt.Println("Pilih Kategori:")
		fmt.Println("1. Fiksi")
		fmt.Println("2. Non Fiksi")

		fmt.Println("Masukan nomor sesuai kategori yang dipilih:")
		fmt.Scan(&tempCat)
		for tempCat < 1 && tempCat > 2 {
			fmt.Println("Kategori tidak sesuai")
			fmt.Scan(&tempCat)
		}
		set.info[idx].catID = tempCat
	} else {
		fmt.Println("Buku tidak ditemukan")
	}
}

func deleteBook(set *Book) {
	var input string
	var idx int

	fmt.Println("Pilih buku yang ingin anda hapus")
	fmt.Scan(&input)
	idx = search(*set, input)
	if idx != -1 {
		//i jadi suatu idx yg udh ditemukan
		for i := idx; i < set.nElemen; i++ {
			set.info[i] = set.info[i+1]
		}
		//nElemen diluar perulangan, karna jika di dalam akan --
		//ikut mengurangi jumlah nElemen dan hanya terprint dari jumlah nElemen yang baru
		set.info[set.nElemen].name = ""
		set.info[set.nElemen].bookcode = ""
		set.info[set.nElemen].catID = 0
		set.nElemen--

		fmt.Println("Buku yang dipilih sudah terhapus")
	} else {
		fmt.Println("Buku tidak ditemukan")
	}
}

func viewBook(set Book) {
	var pilih, pilihCat int
	var category string

	fmt.Println("Pilih:")
	fmt.Println("1. Lihat semua")
	fmt.Println("2. Cari berdasarkan kategori")
	sortJudul(&set)
	fmt.Scan(&pilih)
	for pilih != 1 && pilih != 2 {
		fmt.Println("Pilihan tidak valid")
		fmt.Scan(&pilih)
	}
	if pilih == 1 {
		for i := 0; i < set.nElemen; i++ {
			category = categories(set.info[i].catID)
			fmt.Printf("Judul: %s, Kategori: %s\n", set.info[i].name, category)
		}
	} else if pilih == 2 {
		fmt.Println("---------------------")
		fmt.Println("1. Fiksi")
		fmt.Println("2. Non Fiksi")
		fmt.Println("Masukkan ID kategori:")
		fmt.Scan(&pilihCat)
		for pilihCat < 1 && pilihCat > 2 {
			fmt.Println("Pilihan tidak valid")
			fmt.Scan(&pilihCat)
		}
		printCategories(set, pilihCat)
	}
}

//search di gunakan untuk mengecek apakah buku yang dicari ada dan berisi pengeluaran berupa indeks

func search(set Book, input string) int {
	var idx, i int
	idx = -1
	for i = 0; i < set.nElemen; i++ {
		if set.info[i].name == input {
			idx = i
		}
	}
	return idx
}

func printCategories(set Book, catID int) {
	var category string
	for i := 0; i < set.nElemen; i++ {
		if set.info[i].catID == catID {
			category = categories(set.info[i].catID)
			fmt.Printf("Judul: %s, Kategori: %s\n", set.info[i].name, category)
		}
	}
}

func sortJudul(A *Book) {
	var i, j int
	var temp detailBook
	for i = 1; i < A.nElemen; i++ {
		j = i
		temp = A.info[j]
		for j > 0 && temp.name < A.info[j-1].name {
			A.info[j] = A.info[j-1]
			j--
		}
		A.info[j] = temp
	}
}

func categories(catID int) string {
	if catID == 1 {
		return "Fiksi"
	} else if catID == 2 {
		return "Non Fiksi"
	} else {
		return "UNKNOWN"
	}
}

func huntBook(set Book, p User, input string, ava *int) {
	var idx int
	var borrow bool
	idx = search(set, input)
	borrow = cekStatus(set, idx)

	if idx != -1 {
		if borrow {
			fmt.Println("Buku ", input, "dipinjam oleh", p.info[p.nElemen].name)
		} else {
			fmt.Println("Buku ", input, "Tersedia")
			*ava = idx
		}
	} else {
		fmt.Println("Buku tidak tersedia")
	}
}

func searchCode(set *Book, code string) int {
	var idx, i int
	idx = -1
	for i = 0; i < set.nElemen; i++ {
		if set.info[i].bookcode == code {
			idx = i
		}
	}
	return idx
}

func findBook(set Book) {
	var idx int
	var code string

	fmt.Println("Masukan Code")
	fmt.Scan(&code)

	idx = searchCode(&set, code)
	if idx != -1 {
		fmt.Println("Buku,", set.info[idx].name, "tersedia")
	} else {
		fmt.Println("Buku tidak tersedia")
	}
}

func cekStatus(set Book, idx int) bool {
	// true = dipinjam, false == tersedia
	return set.info[idx].borrowed
}

func borrowBook(set1 *Book, set2 *User) {
	var judul, opt string
	var dBorrow, mBorrow, yBorrow int
	var dReturn, mReturn, yReturn int
	var durasi, ava int

	fmt.Println("Masukan Judul Buku yang ingin dipinjam:")
	fmt.Scan(&judul)
	ava = -1
	huntBook(*set1, *set2, judul, &ava)
	if ava != -1 {
		fmt.Println("1. Pinjam buku")
		fmt.Println("2. Kembali")
		fmt.Print("Select option:")
		fmt.Scan(&opt)
		if opt == "1" {
			fmt.Print("Nama Peminjam: ")
			fmt.Scan(&set2.info[set2.nElemen].name)
			fmt.Print("Tanggal Peminjaman (dd mm yy): ")
			fmt.Scan(&dBorrow, &mBorrow, &yBorrow)
			for !validDate(dBorrow, mBorrow, yBorrow) {
				fmt.Println("Tanggal tidak valid")
				fmt.Println("Tanggal Peminjaman (dd mm yy): ")
				fmt.Scan(&dBorrow, &mBorrow, &yBorrow)
				fmt.Println()
			}

			fmt.Print("Durasi Peminjaman (hari): ")
			fmt.Scan(&durasi)
			countDate(dBorrow, mBorrow, yBorrow, durasi, &dReturn, &mReturn, &yReturn)
			fmt.Printf("Biaya peminjaman buku: Rp%d\n", durasi*1000)
			fmt.Print("Batas pengembalian buku: ")
			fmt.Printf("%2.d/%2.d/%2.d\n", dReturn, mReturn, yReturn)
			set2.info[set2.nElemen].tgl = dReturn
			set2.info[set2.nElemen].bln = mReturn
			set2.info[set2.nElemen].thn = yReturn
			set1.info[ava].currentUser = set2.nElemen
			set1.info[ava].borrowed = true
			fmt.Println(set1.info[ava].borrowed)
			set1.info[ava].nborrow++

			fmt.Println()
		} else if opt == "2" {

		}

	}

}

func returnBook(set *Book, set2 *User) {
	var judul string
	var idx, dReturn, mReturn, yReturn, currentUser int
	var reset detailUser

	fmt.Println("Masukan Judul Buku yang ingin dikembalikan:")
	fmt.Scan(&judul)
	idx = search(*set, judul)
	currentUser = set.info[idx].currentUser
	if cekStatus(*set, idx) {
		fmt.Print("Tanggal Pengembalian : ")
		fmt.Scan(&dReturn, &mReturn, &yReturn)
		if !(dReturn >= set2.info[currentUser].tgl && mReturn >= set2.info[currentUser].bln && yReturn >= set2.info[currentUser].thn) {
			fmt.Println("Anda telah melewati tanggal pengembalian")
			fmt.Println("Anda terkena denda sejumlah Rp100.000")
			fmt.Println("Tolong bayar : gopay : 082146828525")
			set.info[idx].borrowed = false
		} else {
			fmt.Println(dReturn > set2.info[currentUser].tgl, mReturn > set2.info[currentUser].bln, yReturn > set2.info[currentUser].thn)
			fmt.Println(set2.info[currentUser].tgl, set2.info[currentUser].bln, set2.info[currentUser].thn)
			fmt.Println("Buku anda telah disimpan")
			set2.info[idx] = reset
			set.info[idx].borrowed = false
		}
	} else {
		fmt.Println("Piye to")
	}

}

func sortFav(A *Book) {
	var i, pass int
	var temp detailBook
	pass = 1
	for pass = 1; pass < 5 && pass < A.nElemen; pass++ {
		i = pass
		temp = A.info[i]
		for i > 0 && temp.nborrow > A.info[i-1].nborrow {
			A.info[i] = A.info[i-1]
			i--
		}
		A.info[i] = temp
	}
	// for i = 1; i < 5; i++ {
	// 	j = i
	// 	temp = A.info[j]
	// 	for j > 0 && temp.nborrow < A.info[j-1].nborrow {
	// 		A.info[j] = A.info[j-1]
	// 		j--
	// 	}
	// 	A.info[j] = temp
	// }
	for i := 0; i < 5 && i < A.nElemen; i++ {
		fmt.Println("Judul :", A.info[i].name)
	}
}

func validDate(day, month, year int) bool {
	var validD, validM, validY bool
	var temp int
	getday(month, year, &temp)
	validD = day > 0 && day <= temp
	validM = month > 0 && month <= 12
	validY = year > 0
	return validD && validM && validY
}

func countDate(d1, m1, y1, durasi int, d2, m2, y2 *int) {
	var temp int
	getday(m1, y1, &temp)
	*d2 = d1 + durasi
	*m2 = m1
	*y2 = y1
	for *d2 > temp {
		*d2 -= temp
		*m2++
		getday(m1, y1, &temp)
	}
	for *m2 > 12 {
		*m2 -= 12
		*y2++
	}
}

func leapyear(year int) bool {
	return year%400 != 0 || year%100 != 0 && year%4 == 0
}

func getday(month, year int, day *int) {
	if month == 2 {
		if leapyear(year) {
			*day = 29
		} else {
			*day = 28
		}
	} else if month <= 7 {
		if month%2 != 0 {
			*day = 31
		} else {
			*day = 30
		}
	} else if month >= 8 {
		if month%2 == 0 {
			*day = 31
		} else {
			*day = 30
		}
	}
}

