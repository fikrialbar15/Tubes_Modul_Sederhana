package main

import (
	"fmt"
	"strings"
)

// Maksimal user,soal quiz dan forumdis yang bisa registrasi
const USERMAX int = 100
const QUIZMAX int = 20
const FORUMMAX int = 50

type siswa struct {
	nama, username, password, jawabTugas string
	quiz, attempt, ID                    int
}

type quiz struct {
	pertanyaan, jawaban string
}

type forum struct {
	nama, username, komen string
}

type user [USERMAX]siswa
type quizSet [QUIZMAX]quiz
type forumdis [FORUMMAX]forum

func main() {
	var tugas, judulForum string
	var T user
	var Q quizSet
	var F forumdis
	var userAmount, quizAmount, forumAmount int
	presetUser(&T, &userAmount)
	presetQuiz(&Q, &quizAmount)
	presetForum(&F, &forumAmount, &judulForum)
	menuUtama(&T, &userAmount, &quizAmount, &tugas, &Q, &judulForum, &forumAmount, &F)
}

// Males registrasi mulu pas masukin user mulai
func presetUser(T *user, n *int) {
	T[0].nama = "Raihan"
	T[0].username = "gauhover"
	T[0].password = "Zale"
	T[0].ID = 0
	T[0].quiz = -1
	T[0].attempt = 0
	*n++
	T[1].nama = "Ashley"
	T[1].username = "Dreamer"
	T[1].password = "Sky"
	T[1].jawabTugas = "Benar sekali!"
	T[1].ID = 1
	T[1].quiz = 100
	T[1].attempt = 2
	*n++
	T[2].nama = "Zera"
	T[2].username = "Ziera"
	T[2].password = "Mecha"
	T[2].jawabTugas = "No."
	T[2].ID = 2
	T[2].quiz = 66
	T[2].attempt = 3
	*n++
	T[3].nama = "Fiora"
	T[3].username = "Flower"
	T[3].password = "Catlover123"
	T[3].ID = 3
	T[3].quiz = 33
	T[3].attempt = 1
	*n++
	T[4].nama = "Eva"
	T[4].username = "Eve123"
	T[4].password = "Eva123"
	T[4].jawabTugas = "Tolong!"
	T[4].ID = 4
	T[4].quiz = 0
	T[4].attempt = 5
	*n++
}

// Males bikin quiz mulu
func presetQuiz(Q *quizSet, n *int) {
	Q[0].pertanyaan = "50+50 = ???"
	Q[0].jawaban = "100"
	*n++
	Q[1].pertanyaan = "Pelajaran apa yang mengajar keuangan?"
	Q[1].jawaban = "Ekonomi"
	*n++
	Q[2].pertanyaan = "Pelajaran apa yang mengajar angka?"
	Q[2].jawaban = "Matematika"
	*n++
}

// Males bikin forum mulu
func presetForum(F *forumdis, n *int, judul *string) {
	*judul = "Membahas Ujian Nasional"
	F[0].nama = "Ashley"
	F[0].username = "Dreamer"
	F[0].komen = "Belajar biar dapat nilai baik"
	*n++
	F[1].nama = "Zera"
	F[1].username = "Mecha"
	F[1].komen = "Haduh, hancurlah masa depanku..."
	*n++
	F[2].nama = "Eva"
	F[2].username = "Eve123"
	F[2].komen = "Tidurlah!"
	*n++
}

// Header
func header() {
	fmt.Println("=================================================")
	fmt.Println("    	   APLIKASI MOODLE SEDERHANA   ")
	fmt.Println("       FIKRI ALBAR AL GHIFARI_1303220028   ")
	fmt.Println("    RAIHAN RADITYARAHMAN ARRAFI_1303220172   ")
	fmt.Println("=================================================")
}

// Males ngetik ngulang scan mulu, jadiin fungsi aja
func input() string {
	var masukan string
	fmt.Print("Pilihan : ")
	fmt.Scan(&masukan)
	return masukan
}

// Menu utama dari program
func menuUtama(T *user, n, quizAmount *int, tugas *string, Q *quizSet, judulForum *string, forumAmount *int, F *forumdis) {
	var exit bool
	var masukan string
	for masukan != "0" && masukan != "1" && masukan != "2" && masukan != "3" || exit != true {
		header()
		fmt.Println("Selamat datang!")
		fmt.Println("1. Registration")
		fmt.Println("2. Login")
		fmt.Println("3. Admin")
		fmt.Println("0. Exit")
		masukan = input()
		if masukan == "1" {
			registrasiMenu(T, n)
			exit = false
		} else if masukan == "2" {
			login(T, n, quizAmount, tugas, Q, judulForum, F, forumAmount)
			exit = false
		} else if masukan == "3" {
			adminPanel(T, n, quizAmount, tugas, Q, judulForum, forumAmount, F)
			exit = false
		} else if masukan == "0" {
			exit = true
		}
	}
}

// Menu registrasi untuk menambah user ke dalam Array.
func registrasiMenu(T *user, n *int) {
	var exit bool
	var masukan string
	for masukan != "2" && exit != true {
		header()
		fmt.Println("Siap melakukan registrasi?")
		fmt.Println("1. Mulai")
		fmt.Println("2. Kembali")
		masukan = input()
		if masukan == "1" {
			//Check berapa banyak user terdaftar
			fmt.Println(n)
			if USERMAX > *n {
				registrasi(T, n)
			} else {
				fmt.Println("================================")
				fmt.Println("Maaf, kapasitas user penuh. Mohon akses panel Admin untuk menghapus beberapa user")
				pauseToContinue()
			}
		} else if masukan == "2" {
			exit = true
		}
	}
}

// Fungsi untuk registrasi
func registrasi(T *user, n *int) {
	var masukan, nama, username, password string
	var option bool
	option = false
	header()
	fmt.Println("Peringatan!! Jangan menggunakan spasi!")
	fmt.Print("Masukan Nama anda     : ")
	fmt.Scan(&nama)
	// Check jika nama ada yang sama. Kalo ada, else dibawah akan memberi notifikasi.
	if dupeCheck("nama", *T, *n, nama) {
		fmt.Print("Masukan username anda : ")
		fmt.Scan(&username)
		if dupeCheck("username", *T, *n, username) {
			fmt.Print("Masukan password anda : ")
			fmt.Scan(&password)
			for masukan != "2" && masukan != "1" || option == false {
				option = true
				fmt.Println("================================")
				fmt.Println("Apakah data yang dimasukan telah benar?")
				fmt.Println("Nama     :", nama)
				fmt.Println("Username :", username)
				fmt.Println("Password :", password)
				fmt.Println("1. Benar")
				fmt.Println("2. Kembali")
				masukan = input()
				if masukan == "1" {
					// Data yang diinput dimasukkan ke dalam Array. Jika 2 ditekan, tidak terjadi.
					T[*n].nama = nama
					T[*n].username = username
					T[*n].password = password
					T[*n].ID = *n
					T[*n].quiz = -1
					*n++
					fmt.Println("================================")
					fmt.Println("Registrasi berhasil! Silahkan login dari menu utama")
					pauseToContinue()
					masukan = "1"
					option = true
				} else if masukan == "2" {
					option = true
				} else {
					option = false
				}
			}
		} else {
			fmt.Println("Maaf, username yang dimasukan telah diambil. Mohon registrasi ulang.")
		}
	} else {
		fmt.Println("Maaf, nama yang dimasukan telah diambil. Mohon registrasi ulang.")
	}
}

// Fungsi untuk mengecek duplikasi input dari Array
func dupeCheck(flag string, T user, n int, masukan string) bool {
	var i int
	if flag == "nama" {
		for i = 0; i < n; i++ {
			if T[i].nama == masukan {
				return false
			}
		}
	} else if flag == "username" {
		for i = 0; i < n; i++ {
			if T[i].username == masukan {
				return false
			}
		}
		// Tidak dipakai karena aneh masa password gk boleh sama??
		// Buat jaga jagalah serah :v
	} else if flag == "password" {
		for i = 0; i < n; i++ {
			if T[i].password == masukan {
				return false
			}
		}
	}
	return true
}

// Login
func login(T *user, n, quizAmount *int, tugas *string, Q *quizSet, judulForum *string, F *forumdis, forumAmount *int) {
	var username, password string
	var currentUser int
	header()
	if *n < 1 {
		fmt.Println("Tidak ada akun yang telah dibuat, mohon registrasi dari menu utama")
		pauseToContinue()
	} else {
		fmt.Print("Masukan usernama anda : ")
		fmt.Scan(&username)
		if loginSearch(T, n, username) != -1 {
			fmt.Print("Masukan password anda : ")
			fmt.Scan(&password)
			if T[loginSearch(T, n, username)].password == password {
				currentUser = loginSearch(T, n, username)
				fmt.Println("================================")
				fmt.Println("Login berhasil!")
				pauseToContinue()
				siswaMenu(T, &currentUser, quizAmount, tugas, Q, judulForum, F, forumAmount)
			} else {
				fmt.Println("Maaf, password yang dimasukan tidak benar")
				pauseToContinue()
			}
		} else {
			fmt.Println("Maaf, username yang dimasukan tidak dapat ditemukan")
			pauseToContinue()
		}
	}
}

// Menu Admin
func adminPanel(T *user, n, quizAmount *int, tugas *string, Q *quizSet, judulForum *string, forumAmount *int, F *forumdis) {
	var masukan string
	var exit bool
	for masukan != "0" && masukan != "1" && masukan != "2" && masukan != "3" && masukan != "4" && masukan != "5" || exit != true {
		header()
		fmt.Println("Panel Admin")
		fmt.Println("1. Tugas Configuration")
		fmt.Println("2. Quiz Configuration")
		fmt.Println("3. Forum Configuration")
		fmt.Println("4. Student Database")
		fmt.Println("5. User Editor")
		fmt.Println("0. Kembali")
		masukan = input()
		if masukan == "1" {
			adminTugas(T, n, tugas)
			exit = false
		} else if masukan == "2" {
			adminQuiz(T, Q, n, quizAmount)
			exit = false
		} else if masukan == "3" {
			adminForum(F, judulForum, forumAmount)
			exit = false
		} else if masukan == "4" {
			databaseView(T, n)
			exit = false
		} else if masukan == "5" {
			for masukan != "1" && masukan != "2" && masukan != "0" {
				fmt.Println("================================")
				fmt.Println("Mencari akun dengan cara?")
				fmt.Println("1. ID")
				fmt.Println("2. Username")
				fmt.Println("0. Kembali")
				masukan = input()
				if masukan == "1" {
					userSearchID(T, n)
				} else if masukan == "2" {
					userSearchUsername(T, n)
				}
			}
			exit = false
		} else if masukan == "0" {
			exit = true
		}
	}
}

// Mencari user untuk modifikasi (Sequential search) dengan username
func userSearchUsername(T *user, n *int) {
	var masukan string
	var userFound, i int
	userFound = -1
	header()
	fmt.Println("Masukan username yang ingin dimodifikasi")
	fmt.Scan(&masukan)
	for i = 0; i < *n; i++ {
		if T[i].username == masukan {
			userFound = i
		}
	}
	if userFound >= 0 {
		userEditor(T, n, userFound)
	} else {
		fmt.Println("================================")
		fmt.Println("Maaf, username yang dimasukan tidak dapat ditemukan")
		pauseToContinue()
	}
}

// Mencari user untuk modifikasi (Binary Search) dengan ID
func userSearchID(T *user, n *int) {
	var masukan, userFound, kn, kr, med int
	kr = 0
	kn = *n - 1
	userFound = -1
	SortAsc(T, n, "ID")
	header()
	fmt.Println("Masukan ID yang ingin dimodifikasi")
	fmt.Scan(&masukan)
	for kr <= kn && userFound == -1 {
		med = (kr + kn) / 2
		if masukan < T[med].ID {
			kn = med - 1
		} else if masukan > T[med].ID {
			kr = med + 1
		} else {
			userFound = med
		}
	}
	if userFound >= 0 {
		userEditor(T, n, userFound)
	} else {
		fmt.Println("================================")
		fmt.Println("Maaf, ID yang dimasukan tidak dapat ditemukan")
		pauseToContinue()
	}
}

// Menu edit saat sudah ketemu username
func userEditor(T *user, n *int, i int) {
	var exit, deleteConfirm bool
	var masukan string
	for masukan != "0" && masukan != "1" && masukan != "2" && masukan != "3" && masukan != "4" || exit != true {
		header()
		fmt.Println("Akun ditemukan")
		fmt.Println("ID       :", T[i].ID)
		fmt.Println("Nama     :", T[i].nama)
		fmt.Println("Username :", T[i].username)
		fmt.Println("Password :", T[i].password)
		fmt.Println("================================")
		fmt.Println("1. Ganti Nama")
		fmt.Println("2. Ganti Username")
		fmt.Println("3. Ganti Password")
		fmt.Println("4. Hapus Akun")
		fmt.Println("0. Kembali")
		masukan = input()
		deleteConfirm = false
		if masukan == "1" {
			userConfiguration("nama", T, n, i)
			exit = false
		} else if masukan == "2" {
			userConfiguration("username", T, n, i)
			exit = false
		} else if masukan == "3" {
			userConfiguration("password", T, n, i)
			exit = false
		} else if masukan == "4" {
			userDeletion(T, n, i, &deleteConfirm)
			if deleteConfirm == true {
				exit = true
			} else {
				exit = false
			}
		} else if masukan == "0" {
			exit = true
		}
	}
}

// Ganti Nama, Username atau Password
func userConfiguration(flag string, T *user, n *int, i int) {
	header()
	var ii int
	var exit, dupeFound bool
	var masukan, option string
	dupeFound = false
	exit = false
	if flag == "nama" {
		fmt.Println("Masukan nama yang mau dipakai")
		fmt.Scan(&masukan)
		for ii = 0; ii < *n; ii++ {
			if T[ii].nama == masukan && ii != i {
				dupeFound = true
			}
		}
		if dupeFound != true {
			for option != "1" && option != "2" || exit != true {
				fmt.Println("================================")
				fmt.Println("Apakah ini nama yang mau dipakai?")
				fmt.Println("Nama lama :", T[i].nama)
				fmt.Println("Nama baru :", masukan)
				fmt.Println("================================")
				fmt.Println("1.Iya")
				fmt.Println("2.Batal")
				option = input()
				if option == "1" {
					T[i].nama = masukan
					fmt.Println("================================")
					fmt.Println("Nama berhasil diganti!")
					exit = true
				} else if option == "2" {
					exit = true
				}
			}
		} else {
			println("Maaf, nama yang dimasukan sudah diambil.")
			pauseToContinue()
		}
	}
	if flag == "username" {
		fmt.Println("Masukan usename yang mau dipakai")
		fmt.Scan(&masukan)
		for ii = 0; ii < *n; ii++ {
			if T[ii].username == masukan && ii != i {
				dupeFound = true
			}
		}
		if dupeFound != true {
			for option != "1" && option != "2" || exit != true {
				fmt.Println("Apakah ini Username yang mau dipakai?")
				fmt.Println("Usernama lama :", T[i].username)
				fmt.Println("Username baru :", masukan)
				fmt.Println("================================")
				fmt.Println("1.Iya")
				fmt.Println("2.Batal")
				option = input()
				if option == "1" {
					T[i].username = masukan
					fmt.Println("================================")
					fmt.Println("Username berhasil diganti!")
					exit = true
				} else if option == "2" {
					exit = true
				}
			}
		} else {
			println("Maaf, username yang dimasukan sudah diambil.")
			pauseToContinue()
		}
	}
	if flag == "password" {
		fmt.Println("Masukan password yang mau dipakai")
		fmt.Scan(&masukan)
		for option != "1" && option != "2" || exit != true {
			fmt.Println("Apakah ini password yang mau dipakai?")
			fmt.Println("Password lama :", T[i].password)
			fmt.Println("Password baru :", masukan)
			fmt.Println("================================")
			fmt.Println("1.Iya")
			fmt.Println("2.Batal")
			option = input()
			if option == "1" {
				T[i].password = masukan
				fmt.Println("================================")
				fmt.Println("Password berhasil diganti!")
				exit = true
			} else if option == "2" {
				exit = true
			}
		}
	}
}

// Hapus user dari Array
func userDeletion(T *user, n *int, i int, deleteConfirm *bool) {
	header()
	var ii int
	var masukan string
	for masukan != "1" && masukan != "2" {
		fmt.Println("Yakin ingin menghapus user?")
		fmt.Println("1. Hapus")
		fmt.Println("0. Kembali")
		masukan = input()
		if masukan == "1" {
			// Kalo user cuma 1, hapus aja
			if *n == 1 || i+1 == *n {
				T[i].nama = ""
				T[i].username = ""
				T[i].password = ""
				T[i].jawabTugas = ""
				T[i].ID = 0
				T[i].quiz = -1
				T[i].attempt = 0
				*n--
			} else {
				// Kalo ada user lebih dari 1, replace dengan user setelahnya dan hapus yang sudah diswap
				if i < *n {
					for ii = i; ii < *n-1; ii++ {
						// Swap
						T[ii] = T[ii+1]
						// Hapus yang tadi udh di swap
						T[ii+1].nama = ""
						T[ii+1].username = ""
						T[ii+1].password = ""
						T[ii+1].jawabTugas = ""
						T[ii+1].quiz = -1
						T[ii+1].attempt = 0
						T[ii+1].ID = 0
					}
					*n--
				}
			}
			*deleteConfirm = true
			fmt.Println("================================")
			fmt.Println("Akun berhasil dihapus!")
			pauseToContinue()
		} else {
			*deleteConfirm = false
		}
	}
}

// Menampilkan seluruh user yang ada dalam database
func databaseView(T *user, n *int) {
	var i, cap int
	var exit bool
	var masukan, flag, sort string
	cap = 1
	sort = "Asc"
	flag = "ID"
	for masukan != "0" && masukan != "1" && masukan != "2" && masukan != "3" && masukan != "4" && masukan != "5" && masukan != "6" && masukan != "7" && masukan != "8" || exit != true {
		if sort == "Asc" {
			SortAsc(T, n, flag)
		} else if sort == "Desc" {
			SortDes(T, n, flag)
		}
		header()
		fmt.Println("[ID] Nama siswa (Username@Password) : Tugas || Nilai Quiz || Percobaan Quiz, dari", 1+(10*(cap-1)), "ke", 10*cap, "sorted by", flag, "with", sort+".")
		i = 0 + (10 * (cap - 1))
		for i < *n && i < 10*cap && 10*(cap-1) <= i {
			fmt.Print("[", T[i].ID, "] ", T[i].nama, " (", T[i].username, "@", T[i].password, ")", " : ")
			if T[i].jawabTugas == "" {
				fmt.Print("Belum")
			} else {
				fmt.Print("Sudah")
			}
			fmt.Print(" || ")
			if T[i].quiz == -1 {
				fmt.Print("(Belum Quiz!)")
			} else {
				fmt.Print(T[i].quiz)
			}
			fmt.Print(" || ")
			if T[i].attempt == 0 {
				fmt.Print("(Belum Quiz!)")
			} else {
				fmt.Print(T[i].attempt)
			}
			fmt.Println("")
			i++
		}
		fmt.Println("================================")
		if cap != 1 {
			fmt.Print("1. Previous ")
		}
		if (cap)*10 < *n {
			fmt.Print("2. Next ")
		}
		fmt.Print("3. Sort Ascending ")
		fmt.Print("4. Sort Descending ")
		fmt.Print("5. Sort by Nilai ")
		fmt.Print("6. Sort by Nama ")
		fmt.Print("7. Sort by Username ")
		fmt.Print("8. Sort by ID ")
		fmt.Println("0. Kembali")
		masukan = input()
		if masukan == "1" && cap != 1 {
			cap--
			exit = false
		} else if masukan == "2" && (cap)*10 < *n {
			cap++
			exit = false
		} else if masukan == "3" {
			sort = "Asc"
			exit = false
		} else if masukan == "4" {
			sort = "Desc"
			exit = false
		} else if masukan == "5" {
			flag = "Nilai"
			exit = false
		} else if masukan == "7" {
			flag = "Username"
			exit = false
		} else if masukan == "8" {
			flag = "ID"
			exit = false
		} else if masukan == "0" {
			SortDes(T, n, "ID")
			exit = true
		}
	}
}

// Pause biar bisa kebaca saat ada output
func pauseToContinue() {
	var masukan string
	for masukan != "x" {
		fmt.Println("Ketik 'x' untuk lanjut... ")
		masukan = input()
	}
}

// Cari username saat login
func loginSearch(T *user, n *int, masukan string) int {
	var i int
	for i = 0; i < *n; i++ {
		if T[i].username == masukan {
			return i
		}
	}
	return -1
}

// Buka menu siswa dengan user yang sudah login
func siswaMenu(T *user, currentUser, quizAmount *int, tugas *string, Q *quizSet, judulForum *string, F *forumdis, forumAmount *int) {
	var masukan string
	var exit bool
	for masukan != "0" && masukan != "1" && masukan != "2" && masukan != "3" || exit != true {
		header()
		fmt.Println("Selamat datang,", T[*currentUser].nama)
		fmt.Println("1. Tugas")
		fmt.Println("2. Quiz")
		fmt.Println("3. Forum")
		fmt.Println("0. Log out")
		masukan = input()
		if masukan == "1" {
			siswaTugas(T, currentUser, tugas)
			exit = false
		} else if masukan == "2" {
			siswaQuizMenu(T, currentUser, Q, quizAmount)
			exit = false
		} else if masukan == "3" {
			siswaForum(currentUser, F, forumAmount, judulForum, T)
			exit = false
		} else if masukan == "0" {
			exit = true
		}
	}
}

// Siswa buka tugas
func siswaTugas(T *user, currentUser *int, tugas *string) {
	var masukan, jawab string
	var exit bool
	for masukan != "0" && masukan != "1" && masukan != "2" && masukan != "3" || exit != true {
		header()
		fmt.Println("Tugas Anda sekarang adalah : ")
		if *tugas == "" {
			fmt.Println("(Tugas belum diberikan!)")
		} else {
			fmt.Println(*tugas)
		}
		fmt.Println("Jawaban Anda :")
		if T[*currentUser].jawabTugas == "" {
			fmt.Println("(Anda belum menjawab!)")
		} else {
			fmt.Println(T[*currentUser].jawabTugas)
		}
		fmt.Println("================================")
		fmt.Println("1. Jawab")
		fmt.Println("2. Hapus jawaban")
		fmt.Println("0. Kembali")
		masukan = input()
		if masukan == "1" {
			fmt.Println("================================")
			fmt.Println("Gunakan '_' sebagai spasi!")
			fmt.Print("Jawab : ")
			fmt.Scan(&jawab)
			T[*currentUser].jawabTugas = underscoreDeleter(jawab)
			fmt.Println("Jawaban tersimpan!")
			pauseToContinue()
			exit = false
		} else if masukan == "2" {
			for masukan != "0" && masukan != "1" {
				fmt.Println("================================")
				fmt.Println("Hapus jawaban?")
				fmt.Println("1. Hapus!")
				fmt.Println("0. Kembali")
				masukan = input()
				if masukan == "1" {
					fmt.Println("Jawaban dihapus!")
					T[*currentUser].jawabTugas = ""
				}
			}
			exit = false
		} else if masukan == "0" {
			exit = true
		}
	}
}

// Siswa buka quiz
func siswaQuizMenu(T *user, currentUser *int, Q *quizSet, n *int) {
	var exit bool
	var masukan string
	if *n == 0 {
		fmt.Println("================================")
		fmt.Println("Maaf, tidak ada soal quiz untuk dijawab! Mohon tambahkan soal dari panel admin")
		pauseToContinue()
	} else {
		for masukan != "0" && masukan != "1" || exit != true {
			header()
			fmt.Println("Quiz siap untuk dikerjakan!")
			fmt.Print("Nilai : ")
			if T[*currentUser].quiz == -1 {
				fmt.Println("(Belum dikerjakan!)")
			} else {
				fmt.Println(T[*currentUser].quiz)
			}
			fmt.Print("Percobaan ke : ")
			if T[*currentUser].attempt == 0 {
				fmt.Println("(Belum dikerjakan!)")
			} else {
				fmt.Println(T[*currentUser].attempt)
			}
			fmt.Println("================================")
			fmt.Println("1. Mulai quiz!")
			fmt.Println("0. Kembali")
			masukan = input()
			if masukan == "1" {
				menjawabQuiz(Q, T, currentUser, n)
				exit = false
			} else if masukan == "0" {
				exit = true
			}
		}
	}
}

// Fungsi menjawab quiz dan akhir nilai dari quiz
func menjawabQuiz(Q *quizSet, T *user, currentUser, n *int) {
	var i, benar int
	var jawaban, jawabanQuiz string
	for i = 0; i < *n; i++ {
		header()
		fmt.Print("Soal no.")
		fmt.Println(i + 1)
		fmt.Println(Q[i].pertanyaan)
		fmt.Println("Gunakan '_' sebagai spasi!")
		fmt.Print("Jawab : ")
		fmt.Scan(&jawaban)
		jawabanQuiz = Q[i].jawaban
		jawabanQuiz = underscoreDeleter(jawabanQuiz)
		if strings.ToLower(jawabanQuiz) == strings.ToLower(jawaban) {
			fmt.Println("Benar!")
			pauseToContinue()
			benar++
		} else {
			fmt.Println("Salah!")
			fmt.Println("Jawaban benar :", Q[i].jawaban)
			pauseToContinue()
		}
	}
	header()
	fmt.Println("Quiz selesai!")
	fmt.Println("Jawaban benar :", benar, "dari", *n, "soal")
	fmt.Println("NILAI :", (benar*100) / *n)
	T[*currentUser].quiz = (benar * 100) / *n
	T[*currentUser].attempt++
	pauseToContinue()
}

// Mengatur tugas dari admin
func adminTugas(T *user, n *int, tugas *string) {
	var masukan string
	var exit bool
	for masukan != "0" && masukan != "1" && masukan != "2" && masukan != "3" || exit != true {
		header()
		fmt.Println("Tugas yang sedang diberikan adalah...")
		if *tugas == "" {
			fmt.Println("(Tugas belum diberikan!)")
		} else {
			fmt.Println(*tugas)
		}
		fmt.Println("================================")
		fmt.Println("1. Tambah / Edit")
		fmt.Println("2. Hapus")
		fmt.Println("3. Lihat jawaban siswa")
		fmt.Println("0. Kembali")
		masukan = input()
		if masukan == "1" {
			editTugas(T, n, tugas)
			exit = false
		} else if masukan == "2" {
			masukan = "3"
			for masukan != "1" && masukan != "0" {
				fmt.Println("================================")
				fmt.Println("Hapus tugas?")
				fmt.Println("================================")
				fmt.Println("1. Hapus")
				fmt.Println("0. Batal")
				masukan = input()
				if masukan == "1" {
					*tugas = ""
					fmt.Println("Tugas berhasil dihapus!")
					pauseToContinue()
				}
			}
			exit = false
		} else if masukan == "3" {
			checkTugas(T, n, tugas)
			exit = false
		} else if masukan == "0" {
			exit = true
		}
	}
}

// Menganti pertanyaan tugas dan bisa menghapus semua jawaban siswa
func editTugas(T *user, n *int, tugas *string) {
	var masukan, tugasCheck string
	var i int
	header()
	fmt.Println("Masukan pertanyaan yang ingin diberikan")
	fmt.Println("Gunakan '_' untuk spasi!")
	fmt.Scan(&tugasCheck)
	*tugas = underscoreDeleter(tugasCheck)
	fmt.Println("================================")
	fmt.Println("Pertanyaan berhasil diberikan!")
	pauseToContinue()
	for masukan != "1" && masukan != "2" {
		fmt.Println("================================")
		fmt.Println("Hapus semua jawaban siswa yang sudah masuk?")
		fmt.Println("1. Iya")
		fmt.Println("2. Tidak")
		masukan = input()
		if masukan == "1" {
			for i = 0; i < *n; i++ {
				T[i].jawabTugas = ""
			}
			fmt.Println("Semua jawaban siswa telah di reset!")
			pauseToContinue()
		}
	}
}

// Melihat jawaban tugas siswa dari admin
func checkTugas(T *user, n *int, tugas *string) {
	var masukan string
	var i, cap int
	var exit bool
	cap = 1
	exit = false
	for masukan != "0" && masukan != "1" && masukan != "2" || exit != true {
		header()
		fmt.Println("Jawaban tugas dari semua siswa")
		fmt.Println("Pertanyaan :", *tugas)
		i = 0 + (10 * (cap - 1))
		for i < *n && i < 10*cap && 10*(cap-1) <= i {
			fmt.Print("[", T[i].ID, "] ", T[i].nama, "(", T[i].username, ")", " : ")
			if T[i].jawabTugas == "" {
				fmt.Println("(Siswa belum menjawab)")
			} else {
				fmt.Println(T[i].jawabTugas)
			}
			i++
		}
		fmt.Println("================================")
		if cap != 1 {
			fmt.Println("1. Previous")
		}
		if cap*10 < *n {
			fmt.Println("2. Next")
		}
		fmt.Println("0. Kembali")
		masukan = input()
		if masukan == "1" && cap != 1 {
			cap--
			exit = false
		} else if masukan == "2" && (cap)*10 < *n {
			cap++
			exit = false
		} else if masukan == "0" {
			exit = true
		}
	}
}

// Mengatur quiz dari admin
func adminQuiz(T *user, Q *quizSet, n, quizAmount *int) {
	var i, cap int
	var masukan string
	var exit, unedited bool
	unedited = true
	cap = 1
	for masukan != "0" && masukan != "1" && masukan != "2" || exit != true {
		header()
		fmt.Println("Pertanyaan Quiz dan jawabannya.")
		if *quizAmount == 0 {
			fmt.Println("(Tidak ada soal!)")
		} else {
			i = 0 + (5 * (cap - 1))
			for i < *quizAmount && i < 5*cap && 5*(cap-1) <= i {
				fmt.Print(i+1, ". ", Q[i].pertanyaan)
				fmt.Println("")
				fmt.Print("Jawaban : ", Q[i].jawaban)
				fmt.Println("")
				i++
			}
		}
		fmt.Println("================================")
		if cap != 1 {
			fmt.Print("4. Previous ")
		}
		if (cap)*5 < *quizAmount {
			fmt.Print("5. Next")
		}
		fmt.Println("")
		fmt.Println("1. Tambah")
		fmt.Println("2. Edit / Hapus")
		fmt.Println("3. Lihat nilai siswa")
		fmt.Println("0. Keluar")
		masukan = input()
		if masukan == "1" {
			if unedited == true {
				quizReset(T, n, &unedited)
			} else {
				tambahQuiz(Q, quizAmount)
			}
			exit = false
		} else if masukan == "2" {
			if unedited == true {
				quizReset(T, n, &unedited)
			} else {
				editQuiz(Q, quizAmount)
			}
			exit = false
		} else if masukan == "3" {
			quizReview(T, n)
			exit = false
		} else if masukan == "4" && cap != 1 {
			cap--
			exit = false
		} else if masukan == "5" && (cap)*5 < *quizAmount {
			cap++
			exit = false
		} else if masukan == "0" {
			exit = true
		}
	}
}

// Fungsi untuk menambah soal dan jawaban dari quiz
func tambahQuiz(Q *quizSet, n *int) {
	var question, answer, masukan string
	if QUIZMAX > *n {
		fmt.Println("================================")
		fmt.Println("Masukan pertanyaan quiz")
		fmt.Println("Gunakan '_' sebagai spasi")
		fmt.Scan(&question)
		question = underscoreDeleter(question)
		fmt.Println("================================")
		fmt.Println("Masukan jawaban dari pertanyaan")
		fmt.Println("Gunakan '_' sebagai spasi")
		fmt.Scan(&answer)
		answer = underscoreDeleter(answer)
		for masukan != "0" && masukan != "1" {
			fmt.Println("================================")
			fmt.Println("Pertanyaan :", question)
			fmt.Println("Jawaban :", answer)
			fmt.Println("================================")
			fmt.Println("1. Tambah")
			fmt.Println("0. Kembali")
			masukan = input()
			if masukan == "1" {
				Q[*n].pertanyaan = question
				Q[*n].jawaban = answer
				*n++
			}
		}
	} else {
		fmt.Println("Maaf, kapasitas quiz penuh. Mohon hapus beberapa soal")
		pauseToContinue()
	}
}

// Fungsi untk reset nilai siswa jika ada perubahan dari soal quiz
func quizReset(T *user, n *int, unedited *bool) {
	var masukan string
	var i int
	if *unedited == true {
		for masukan != "0" && masukan != "1" {
			header()
			fmt.Println("Mengubah quiz akan membuat semua nilai siswa kosong!")
			fmt.Println("Lanjut?")
			fmt.Println("================================")
			fmt.Println("1. Reset nilai")
			fmt.Println("0. Kembali")
			fmt.Println("================================")
			masukan = input()
			if masukan == "1" {
				for i = 0; i < *n; i++ {
					T[i].quiz = -1
					T[i].attempt = 0
				}
				fmt.Println("Nilai siswa telah dihapus!")
				pauseToContinue()
				*unedited = false
			}
		}
	}
}

// Fungsi untuk menghapus underscore "_" dari sebuah string
func underscoreDeleter(kata string) string {
	return strings.Replace(kata, "_", " ", -1)
}

// Fungsi untuk ngedit pertanyaan atau jawaban dari soal
func editQuiz(Q *quizSet, n *int) {
	var i int
	var exit, deleteConfirm bool
	var masukan string
	header()
	fmt.Println("Masukan nomor soal yang ingin di edit")
	fmt.Scan(&i)
	i = i - 1
	if i < *n {
		for masukan != "0" && masukan != "1" && masukan != "2" && masukan != "3" || exit != true {
			header()
			fmt.Println("Soal :")
			fmt.Println(Q[i].pertanyaan)
			fmt.Println("Jawaban :")
			fmt.Println(Q[i].jawaban)
			fmt.Println("================================")
			fmt.Println("1. Ganti Soal")
			fmt.Println("2. Ganti Jawaban")
			fmt.Println("3. Hapus soal")
			fmt.Println("0. Kembali")
			masukan = input()
			deleteConfirm = false
			if masukan == "1" {
				quizConfiguration("pertanyaan", Q, i)
				exit = false
			} else if masukan == "2" {
				quizConfiguration("jawaban", Q, i)
				exit = false
			} else if masukan == "3" {
				hapusQuiz(Q, n, i, &deleteConfirm)
				if deleteConfirm == true {
					exit = true
				} else {
					exit = false
				}
			} else if masukan == "0" {
				exit = true
			}
		}
	} else {
		fmt.Println("================================")
		fmt.Println("Maaf, soal tidak dapat ditemukan")
		pauseToContinue()
	}
}

// Fungsi untuk menghapus soal dari quiz
func hapusQuiz(Q *quizSet, n *int, i int, deleteConfirm *bool) {
	header()
	var ii int
	var masukan string
	for masukan != "1" && masukan != "2" {
		fmt.Println("Yakin ingin menghapus soal?")
		fmt.Println("1. Hapus")
		fmt.Println("0. Kembali")
		masukan = input()
		if masukan == "1" {
			// Kalo soal cuma 1, hapus aja
			if *n == 1 || i+1 == *n {
				Q[i].pertanyaan = ""
				Q[i].jawaban = ""
				*n--
			} else {
				// Kalo ada user lebih dari 1, replace dengan user setelahnya dan hapus yang sudah diswap
				if i < *n {
					for ii = i; ii < *n-1; ii++ {
						// Swap
						Q[ii] = Q[ii+1]
						// Hapus yang tadi udh di swap
						Q[ii+1].pertanyaan = ""
						Q[ii+1].jawaban = ""
					}
					*n--
				}
			}
			*deleteConfirm = true
			fmt.Println("================================")
			fmt.Println("Soal berhasil dihapus!")
			pauseToContinue()
		} else {
			*deleteConfirm = false
		}
	}
}

// Selection sort untuk Ascending
func SortAsc(T *user, n *int, flag string) {
	var i, j, idx_min int
	var t siswa
	i = 1
	if flag == "Nilai" {
		for i < *n {
			idx_min = i - 1
			j = i
			for j < *n {
				if T[idx_min].quiz > T[j].quiz {
					idx_min = j
				}
				j++
			}
			t = T[idx_min]
			T[idx_min] = T[i-1]
			T[i-1] = t
			i++
		}
	} else if flag == "Nama" {
		for i < *n {
			idx_min = i - 1
			j = i
			for j < *n {
				if T[idx_min].nama > T[j].nama {
					idx_min = j
				}
				j++
			}
			t = T[idx_min]
			T[idx_min] = T[i-1]
			T[i-1] = t
			i++
		}
	} else if flag == "Username" {
		for i < *n {
			idx_min = i - 1
			j = i
			for j < *n {
				if T[idx_min].username > T[j].username {
					idx_min = j
				}
				j++
			}
			t = T[idx_min]
			T[idx_min] = T[i-1]
			T[i-1] = t
			i++
		}
	} else if flag == "ID" {
		for i < *n {
			idx_min = i - 1
			j = i
			for j < *n {
				if T[idx_min].ID > T[j].ID {
					idx_min = j
				}
				j++
			}
			t = T[idx_min]
			T[idx_min] = T[i-1]
			T[i-1] = t
			i++
		}
	}
}

// Insert sort untuk Descending
func SortDes(T *user, n *int, flag string) {
	var i, j int
	var temp siswa
	i = 1
	if flag == "Nilai" {
		for i < *n {
			j = i
			temp = T[j]
			for j > 0 && temp.quiz > T[j-1].quiz {
				T[j] = T[j-1]
				j = j - 1
			}
			T[j] = temp
			i++
		}
	} else if flag == "Nama" {
		for i < *n {
			j = i
			temp = T[j]
			for j > 0 && temp.nama > T[j-1].nama {
				T[j] = T[j-1]
				j = j - 1
			}
			T[j] = temp
			i++
		}
	} else if flag == "Username" {
		for i < *n {
			j = i
			temp = T[j]
			for j > 0 && temp.username > T[j-1].username {
				T[j] = T[j-1]
				j = j - 1
			}
			T[j] = temp
			i++
		}
	} else if flag == "ID" {
		for i < *n {
			j = i
			temp = T[j]
			for j > 0 && temp.ID > T[j-1].ID {
				T[j] = T[j-1]
				j = j - 1
			}
			T[j] = temp
			i++
		}
	}
}

// Fungsi untuk menampilkan semua nilai dari siswa
func quizReview(T *user, n *int) {
	var i, cap int
	var exit bool
	var masukan, flag, sort string
	cap = 1
	sort = "Desc"
	flag = "Nilai"
	for masukan != "0" && masukan != "1" && masukan != "2" && masukan != "3" && masukan != "4" && masukan != "5" && masukan != "6" && masukan != "7" && masukan != "8" || exit != true {
		if sort == "Asc" {
			SortAsc(T, n, flag)
		}
		if sort == "Desc" {
			SortDes(T, n, flag)
		}
		header()
		fmt.Println("[ID] Nama siswa (Username) : Nilai dari", 1+(10*(cap-1)), "ke", 10*cap, "sorted by", flag, "with", sort+".")
		i = 0 + (10 * (cap - 1))
		for i < *n && i < 10*cap && 10*(cap-1) <= i {
			fmt.Print("[", T[i].ID, "] ", T[i].nama, " (", T[i].username, ")", " : ")
			if T[i].quiz == -1 {
				fmt.Println("(Belum mengerjakan!)")
			} else {
				fmt.Println(T[i].quiz)
			}
			i++
		}
		fmt.Println("================================")
		if cap != 1 {
			fmt.Print("1. Previous ")
		}
		if (cap)*10 < *n {
			fmt.Print("2. Next ")
		}
		fmt.Print("3. Sort Ascending ")
		fmt.Print("4. Sort Descending ")
		fmt.Print("5. Sort by Nilai ")
		fmt.Print("6. Sort by Nama ")
		fmt.Print("7. Sort by Username ")
		fmt.Print("8. Sort by ID ")
		fmt.Println("0. Kembali")
		masukan = input()
		if masukan == "1" && cap != 1 {
			cap--
			exit = false
		} else if masukan == "2" && (cap)*10 < *n {
			cap++
			exit = false
		} else if masukan == "3" {
			sort = "Asc"
			exit = false
		} else if masukan == "4" {
			sort = "Desc"
			exit = false
		} else if masukan == "5" {
			flag = "Nilai"
			exit = false
		} else if masukan == "6" {
			flag = "Nama"
			exit = false
		} else if masukan == "7" {
			flag = "Username"
			exit = false
		} else if masukan == "8" {
			flag = "ID"
			exit = false
		} else if masukan == "0" {
			SortDes(T, n, "ID")
			exit = true
		}
	}
}

// Ganti pertanyaan dan jawaban soal dari quiz
func quizConfiguration(flag string, Q *quizSet, i int) {
	header()
	var exit bool
	var masukan, option string
	exit = false
	if flag == "pertanyaan" {
		fmt.Println("Masukan pertanyaan baru")
		fmt.Println("Gunakan '_' sebagai spasi!")
		fmt.Scan(&masukan)
		masukan = underscoreDeleter(masukan)
		for option != "1" && option != "2" || exit != true {
			fmt.Println("================================")
			fmt.Println("Apakah ini pertanyaan yang mau dipakai?")
			fmt.Println("Pertanyaan lama :", Q[i].pertanyaan)
			fmt.Println("Pertanyaan baru :", masukan)
			fmt.Println("================================")
			fmt.Println("1.Iya")
			fmt.Println("2.Batal")
			option = input()
			if option == "1" {
				Q[i].pertanyaan = masukan
				fmt.Println("================================")
				fmt.Println("Pertanyaan berhasil diganti!")
				pauseToContinue()
				exit = true
			} else if option == "2" {
				exit = true
			}
		}
	}
	if flag == "jawaban" {
		fmt.Println("Masukan jawaban baru")
		fmt.Println("Gunakan '_' sebagai spasi!")
		fmt.Scan(&masukan)
		masukan = underscoreDeleter(masukan)
		for option != "1" && option != "2" || exit != true {
			fmt.Println("Apakah ini jawaban yang mau dipakai?")
			fmt.Println("Jawaban lama :", Q[i].jawaban)
			fmt.Println("Jawaban baru :", masukan)
			fmt.Println("================================")
			fmt.Println("1.Iya")
			fmt.Println("2.Batal")
			option = input()
			if option == "1" {
				Q[i].jawaban = masukan
				fmt.Println("================================")
				fmt.Println("Jawaban berhasil diganti!")
				pauseToContinue()
				exit = true
			} else if option == "2" {
				exit = true
			}
		}
	}
}

// Fungsi untuk mengatur forum dari admin
func adminForum(F *forumdis, judulForum *string, n *int) {
	var masukan string
	var i, cap int
	var exit bool
	cap = 1
	exit = false
	for masukan != "0" && masukan != "1" && masukan != "2" && masukan != "3" && masukan != "4" && masukan != "5" || exit != true {
		header()
		fmt.Print("Judul forum diskusi : ")
		if *judulForum == "" {
			fmt.Println("(Belum diberikan)")
		} else {
			fmt.Println(*judulForum)
		}
		fmt.Println("================================")
		fmt.Println("[ID] Nama Siswa (@username) : Komentar siswa")
		if *n == 0 {
			fmt.Println("Belum ada Komentar")
		} else {
			i = 0 + (10 * (cap - 1))
			for i < *n && i < 10*cap && 10*(cap-1) <= i {
				fmt.Print("[", i, "] ", F[i].nama, "(@", F[i].username, ") : ", F[i].komen)
				fmt.Println("")
				i++
			}
		}
		fmt.Println("================================")
		if cap != 1 {
			fmt.Println("1. Previous")
		}
		if cap*10 < *n {
			fmt.Println("2. Next")
		}
		fmt.Println("3. Edit Judul")
		fmt.Println("4. Hapus Komentar")
		fmt.Println("5. Reset Forum")
		fmt.Println("0. Kembali")
		masukan = input()
		if masukan == "1" && cap != 1 {
			cap--
			exit = false
		} else if masukan == "2" && (cap)*10 < *n {
			cap++
			exit = false
		} else if masukan == "3" {
			forumJudulEdit(judulForum)
			exit = false
		} else if masukan == "4" {
			forumHapus(F, n)
			exit = false
		} else if masukan == "5" {
			forumReset(F, n)
			exit = false
		} else if masukan == "0" {
			exit = true
		}
	}
}

// Fungsi untuk menganti judul forum
func forumJudulEdit(judul *string) {
	var masukan, option string
	header()
	fmt.Println("Masukan judul baru untuk forum diskusi")
	fmt.Println("Gunakan '_' sebagai spasi!")
	fmt.Scan(&masukan)
	masukan = underscoreDeleter(masukan)
	for option != "2" && option != "1" {
		fmt.Println("Apakah ini judul yang mau dipakai?")
		fmt.Println("Judul lama :", *judul)
		fmt.Println("Judul baru :", masukan)
		fmt.Println("================================")
		fmt.Println("1.Iya")
		fmt.Println("2.Batal")
		option = input()
		if option == "1" {
			*judul = masukan
			fmt.Println("================================")
			fmt.Println("Judul forum diskusi berhasil diganti!")
			pauseToContinue()
		}
	}
}

// Fungsi untuk menghapus sebuah komentar dari forum
func forumHapus(F *forumdis, n *int) {
	var option string
	var i, ii int
	header()
	fmt.Println("Masukan ID forum komentar")
	fmt.Scan(&i)
	if i+1 > *n {
		fmt.Println("Komentar tidak dapat ditemukan")
		pauseToContinue()
	} else if i < 0 {
		fmt.Println("ID tidak bisa negative!")
		pauseToContinue()
	} else {
		for option != "2" && option != "1" {
			fmt.Println("Apakah ini komentar yang mau dihapus?")
			fmt.Print("Nama(@username) : ", F[i].nama, "(@", F[i].username, ")")
			fmt.Println(" ")
			fmt.Println("Komentar :", F[i].komen)
			fmt.Println("================================")
			fmt.Println("1.Iya")
			fmt.Println("2.Batal")
			option = input()
			if option == "1" {
				// Kalo user cuma satu atau terakhir, hapus aja
				if *n == 1 || i+1 == *n {
					F[i].nama = ""
					F[i].username = ""
					F[i].komen = ""
					*n--
				} else {
					// Kalo ada user lebih dari 1, replace dengan user setelahnya dan hapus yang sudah diswap
					if i < *n {
						for ii = i; ii < *n-1; ii++ {
							// Swap
							F[ii] = F[ii+1]
							// Hapus yang tadi udh di swap
							F[ii+1].nama = ""
							F[ii+1].username = ""
							F[ii+1].komen = ""

						}
						*n--
					}
				}
				fmt.Println("Komentar berhasil dihapus!")
				pauseToContinue()
			}
		}
	}
}

// Fungsi untuk menghapus semua komentar dari forum
func forumReset(F *forumdis, n *int) {
	var masukan string
	var i int
	header()
	fmt.Println("Hapus SELURUH komentar forum?")
	fmt.Println("Ketik 'HAPUS' dengan semua kapital untuk lanjut")
	fmt.Scan(&masukan)
	if masukan == "HAPUS" {
		for i = 0; i < *n; i++ {
			F[i].nama = ""
			F[i].username = ""
			F[i].komen = ""
			*n = 0
			fmt.Println("Penghapusan berhasil!")
			pauseToContinue()
		}
	} else {
		fmt.Println("Masukan tidak sesuai, membatalkan penghapusan...")
		pauseToContinue()
	}
}

// Fungsi untuk menampilkan forum dari siswa
func siswaForum(currentUser *int, F *forumdis, forumAmount *int, judulForum *string, T *user) {
	var masukan string
	var i, cap int
	var exit bool
	cap = 1
	exit = false
	for masukan != "0" && masukan != "1" && masukan != "2" && masukan != "3" && masukan != "4" && masukan != "5" || exit != true {
		header()
		fmt.Print("Judul forum diskusi : ")
		if *judulForum == "" {
			fmt.Println("(Belum diberikan)")
		} else {
			fmt.Println(*judulForum)
		}
		if *forumAmount == 0 {
			fmt.Println("Belum ada Komentar! Jadilah yang pertama untuk berkomentar!")
		} else {
			i = 0 + (10 * (cap - 1))
			for i < *forumAmount && i < 10*cap && 10*(cap-1) <= i {
				fmt.Print("[", i+1, "] ", F[i].nama, "(@", F[i].username, ") : ", F[i].komen)
				fmt.Println("")
				i++
			}
		}
		fmt.Println("================================")
		fmt.Print("Masuk sebagai : ", T[*currentUser].nama, "(@", T[*currentUser].username, ")")
		fmt.Println(" ")
		fmt.Println("================================")
		if cap != 1 {
			fmt.Println("1. Previous")
		}
		if cap*10 < *forumAmount {
			fmt.Println("2. Next")
		}
		fmt.Println("3. Komentar")
		fmt.Println("0. Kembali")
		masukan = input()
		if masukan == "1" && cap != 1 {
			cap--
			exit = false
		} else if masukan == "2" && (cap)*10 < *forumAmount {
			cap++
			exit = false
		} else if masukan == "3" {
			if FORUMMAX > *forumAmount {
				siswaKomentar(T, currentUser, forumAmount, F)
			} else {
				fmt.Println("Maaf, kapasitas Forum penuh. Mohon hubungi Admin untuk memperluas Forum")
				pauseToContinue()
			}
			exit = false
		} else if masukan == "0" {
			exit = true
		}
	}
}

// Fungsi untuk berkomentar ke forum
func siswaKomentar(T *user, currentUser, forumAmount *int, F *forumdis) {
	var komentar, option string
	header()
	fmt.Println("Masukan komentar baru untuk forum diskusi")
	fmt.Println("Gunakan '_' sebagai spasi!")
	fmt.Scan(&komentar)
	komentar = underscoreDeleter(komentar)
	for option != "2" && option != "1" {
		fmt.Println("================================")
		fmt.Println("Apakah ini komentar yang mau tampilkan?")
		fmt.Println("Komentar :", komentar)
		fmt.Println("================================")
		fmt.Println("1.Iya")
		fmt.Println("2.Batal")
		option = input()
		if option == "1" {
			F[*forumAmount].nama = T[*currentUser].nama
			F[*forumAmount].username = T[*currentUser].username
			F[*forumAmount].komen = komentar
			*forumAmount++
			fmt.Println("================================")
			fmt.Println("Komentar berhasil dikirim!")
			pauseToContinue()
		}
	}
}
