package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const NMAX int = 100

type text struct {
	textID      int
	authorUname string
	message     string
}

type chat struct {
	receiverUname string
	textData      [NMAX]text
	nText, chatID int
}

type group struct {
	memberUname [30]string
	grupName    string
	groupText   [NMAX]text
	nMember, nGroupText, grupID int
}

type account struct {
	name, uname, pass, gender string
	umur int
	chatData          [NMAX]chat
	joinedGroupID     [NMAX]int
	nChat, nGroup     int
}

type accounts [NMAX]account
type groups [NMAX]group

var dataGrup groups
var dataAkun, dataAkunPending accounts
var nDataAkunPending, nDataAkun, nDataGrup int

func main() {
	menu_welcome()
}

// Tampilan Menu
func menu_welcome_views() {
	fmt.Println("|~~~|-------------------------------------------|~~~|")
	fmt.Println("|~~~|             Aplikasi Chatting             |~~~|")
	fmt.Println("|~~~|    Created by Muhammad Hendika Putra &    |~~~|")
	fmt.Println("|~~~|          Luthfi Iriawan Fadhilah          |~~~|")
	fmt.Println("|~~~|       Algoritma Pemgrograman 2024         |~~~|")
	fmt.Println("|~~~|-------------------------------------------|~~~|")
}

func menu_utama_views() {
	fmt.Println("|~~~|----------------Menu Utama-----------------|~~~|")
	fmt.Println("|~~~| 1. Registrasi                             |~~~|")
	fmt.Println("|~~~| 2. Login                                  |~~~|")
	fmt.Println("|~~~| 3. Admin                                  |~~~|")
	fmt.Println("|~~~| 4. Keluar                                 |~~~|")
	fmt.Println("|~~~|-------------------------------------------|~~~|")
}

func menu_registrasi_views() {
	var name, uname, pass, gender string
	var umur int
	var input string
	var createdAccount account
	
	fmt.Println("|~~~|--------------Menu Registrasi--------------|~~~|")
	fmt.Printf("%15s %6s", "Nama", ": ")
	fmt.Scan(&name)
	fmt.Printf("%15s %6s", "Username", ": ")
	fmt.Scan(&uname)
	fmt.Printf("%15s %6s", "Umur", ": ")
	fmt.Scan(&umur)
	for umur <= 0 {
		fmt.Println("Masukkan umur yang valid.")
		fmt.Printf("%15s %6s", "Umur", ": ")
		fmt.Scan(&umur)
	}
	fmt.Printf("%15s %6s", "Gender (L/P)", ": ")
	fmt.Scan(&gender)
	for gender != "L" && gender != "P" {
		fmt.Println("Masukkan gender yang valid.")
		fmt.Printf("%15s %6s", "Gender (L/P)", ": ")
		fmt.Scan(&gender)
	}
	fmt.Printf("%15s %6s", "Password", ": ")
	fmt.Scan(&pass)
	fmt.Println("|~~~|-------------------------------------------|~~~|")
	fmt.Println("Pilihan: 1. Simpan, 2. Batal")
	fmt.Print("Pilih (1/2): ")
	fmt.Scan(&input)
	for input != "1" && input != "2" {
		fmt.Println("Masukkan salah, silahkan input kembali.")
		fmt.Print("Pilih (1/2): ")
		fmt.Scan(&input)
	}
	if input == "1" {
		createdAccount.name = name
		createdAccount.uname = uname
		createdAccount.pass = pass
		createdAccount.gender = gender
		createdAccount.umur = umur

		if searchAkunIdx(dataAkun, nDataAkun, uname) == -1 && searchAkunIdx(dataAkunPending, nDataAkunPending, uname) == -1 {
			dataAkunPending[nDataAkunPending] = createdAccount
			nDataAkunPending++
			clearScreen()
			fmt.Println("Registrasi berhasil, menunggu persetujuan admin.")
			menu_utama()
		} else {
			clearScreen()
			fmt.Println("Registrasi gagal, Username telah digunakan.")
			menu_utama()
		}
	} else {
		clearScreen()
		menu_utama()
	}
}

func menu_login_views() {
	var uname, pass string
	var input string
	var idxAkunDipakai int

	fmt.Println("|~~~|-----------------Menu Login----------------|~~~|")
	fmt.Printf("%15s %6s", "Username", ": ")
	fmt.Scan(&uname)
	fmt.Printf("%15s %6s", "Password", ": ")
	fmt.Scan(&pass)
	fmt.Println("|~~~|-------------------------------------------|~~~|")
	fmt.Println("Pilihan: 1. Login, 2. Batal")
	fmt.Print("Pilih (1/2): ")
	fmt.Scan(&input)
	for input != "1" && input != "2" {
		fmt.Println("Masukkan salah, silahkan input kembali.")
		fmt.Print("Pilih (1/2): ")
		fmt.Scan(&input)
	}
	if input == "1" {
		idxAkunDipakai = searchAkunIdx(dataAkun, nDataAkun, uname)
		if idxAkunDipakai != -1 {
			if pass == dataAkun[idxAkunDipakai].pass {
				clearScreen()
				menu_home(idxAkunDipakai)
			} else {
				clearScreen()
				fmt.Println("Login gagal, Password salah.")
				menu_utama()
			}
		} else {
			clearScreen()
			fmt.Println("Login gagal, Username tidak ditemukan.")
			menu_utama()
		}
	} else {
		clearScreen()
		menu_utama()
	}
}

func menu_admin_views() {
	fmt.Println("|~~~|-----------------Menu Admin----------------|~~~|")
	fmt.Println("|~~~| 1. Daftar Akun                            |~~~|")
	fmt.Println("|~~~| 2. Daftar Registrasi Akun                 |~~~|")
	fmt.Println("|~~~| 3. Kembali                                |~~~|")
	fmt.Println("|~~~|-------------------------------------------|~~~|")
}

func menu_admin_views_daftar_akun() {
	fmt.Println("|~~~|-----------------Menu Admin----------------|~~~|")
	fmt.Println("|~~~| Daftar Akun                               |~~~|")
	if nDataAkun == 0 {
		fmt.Println("|~~~| Belum ada akun.                           |~~~|")
	} else {
		for i := 0; i < nDataAkun; i++ {
			fmt.Printf("%15s %d\n", "~~ Akun ", i+1)
			fmt.Printf("%15s %6s %s\n", "Nama", ":", dataAkun[i].name)
			fmt.Printf("%15s %6s %s\n", "Username", ":", dataAkun[i].uname)
			fmt.Printf("%15s %6s %s\n", "Gender", ":", dataAkun[i].gender)
			fmt.Printf("%15s %6s %d\n", "Umur", ":", dataAkun[i].umur)
			fmt.Printf("%15s %6s %s\n", "Password", ":", dataAkun[i].pass)
		}
	}
	fmt.Println("|~~~|-------------------------------------------|~~~|")
}

func menu_admin_views_daftar_akun_sort(dataAkun accounts) {
	fmt.Println("|~~~|-----------------Menu Admin----------------|~~~|")
	fmt.Println("|~~~| Daftar Akun                               |~~~|")
	if nDataAkun == 0 {
		fmt.Println("|~~~| Belum ada akun.                           |~~~|")
	} else {
		for i := 0; i < nDataAkun; i++ {
			fmt.Printf("%15s %d\n", "~~ Akun ", i+1)
			fmt.Printf("%15s %6s %s\n", "Nama", ":", dataAkun[i].name)
			fmt.Printf("%15s %6s %s\n", "Username", ":", dataAkun[i].uname)
			fmt.Printf("%15s %6s %s\n", "Password", ":", dataAkun[i].pass)
		}
	}
	fmt.Println("|~~~|-------------------------------------------|~~~|")
}

func menu_admin_views_daftar_akun_pending() {
	fmt.Println("|~~~|-----------------Menu Admin----------------|~~~|")
	fmt.Println("|~~~| Daftar Registrasi Akun                    |~~~|")
	if nDataAkunPending == 0 {
		fmt.Println("|~~~| Belum ada registrasi akun.                |~~~|")
	} else {
		for i := 0; i < nDataAkunPending; i++ {
			fmt.Printf("%15s %d\n", "~~ Akun ", i+1)
			fmt.Printf("%15s %6s %s\n", "Nama", ":", dataAkunPending[i].name)
			fmt.Printf("%15s %6s %s\n", "Username", ":", dataAkunPending[i].uname)
			fmt.Printf("%15s %6s %s\n", "Gender", ":", dataAkunPending[i].gender)
			fmt.Printf("%15s %6s %d\n", "Umur", ":", dataAkunPending[i].umur)
			fmt.Printf("%15s %6s %s\n", "Password", ":", dataAkunPending[i].pass)
		}
	}
	fmt.Println("|~~~|-------------------------------------------|~~~|")
}

func menu_home_views() {
	fmt.Println("|~~~|-----------------Menu Home-----------------|~~~|")
	fmt.Println("|~~~| 1. Chat                                   |~~~|")
	fmt.Println("|~~~| 2. Grup                                   |~~~|")
	fmt.Println("|~~~| 3. Pengaturan Akun                        |~~~|")
	fmt.Println("|~~~| 4. Logout                                 |~~~|")
	fmt.Println("|~~~|-------------------------------------------|~~~|")
}

func menu_chat_views(idxAkunDipakai int) {
	var tampilAkunUname string

	fmt.Println("|~~~|-----------------Menu Chat-----------------|~~~|")
	fmt.Println("|~~~| Daftar Chat                               |~~~|")
	if dataAkun[idxAkunDipakai].nChat == 0 {
		fmt.Println("|~~~| Belum ada chat.                           |~~~|")
	} else {
		for i := 0; i < dataAkun[idxAkunDipakai].nChat; i++ {
			tampilAkunUname = dataAkun[idxAkunDipakai].chatData[i].receiverUname
			fmt.Printf("%6d", i+1)
			fmt.Print(". ", tampilAkunUname, "\n")
		}
	}
	fmt.Println("|~~~|-------------------------------------------|~~~|")
}

func menu_pesan_views(onChat chat) {
	var onChatReceiverIdx int
	onChatReceiverIdx = searchAkunIdx(dataAkun, nDataAkun, onChat.receiverUname)

	fmt.Println("|~~~|----------------Menu Pesan-----------------|~~~|")
	if onChatReceiverIdx != -1 {
		fmt.Printf("%15s Chat\n", dataAkun[onChatReceiverIdx].name)
	} else {
		fmt.Printf("%15s Chat\n", onChat.receiverUname)
	}
	for i := 0; i < onChat.nText; i++ {
		if onChat.textData[i].authorUname == onChat.receiverUname {
			fmt.Printf("%3d %10s\n", onChat.textData[i].textID, onChat.textData[i].authorUname)
			fmt.Printf("%3s %10s\n", "", onChat.textData[i].message)
			fmt.Println()
		} else {
			fmt.Printf("%3d %50s\n", onChat.textData[i].textID, onChat.textData[i].authorUname)
			fmt.Printf("%3s %50s\n", "", onChat.textData[i].message)
			fmt.Println()
		}
	}
}

func menu_grup_views(idxAkunDipakai int) {
	fmt.Println("|~~~|-----------------Menu Grup-----------------|~~~|")
	fmt.Println("|~~~| Daftar Grup                               |~~~|")
	if dataAkun[idxAkunDipakai].nGroup == 0 {
		fmt.Println("|~~~| Belum ada grup.                           |~~~|")
	} else {
		for i := 0; i < dataAkun[idxAkunDipakai].nGroup; i++ {
			fmt.Printf("%6d", i+1)
			fmt.Print(". ", dataGrup[dataAkun[idxAkunDipakai].joinedGroupID[i]].grupName, "\n")
		}
	}
	fmt.Println("|~~~|-------------------------------------------|~~~|")
}

func menu_grup_chat_views(onGrup group, thisAccUname string) {

	fmt.Println("|~~~|--------------Menu Grup Pesan--------------|~~~|")
	fmt.Printf("%15s Chat\n", onGrup.grupName)
	for i := 0; i < onGrup.nGroupText; i++ {
		if onGrup.groupText[i].authorUname == thisAccUname {
			fmt.Printf("%3d %50s\n", onGrup.groupText[i].textID, onGrup.groupText[i].authorUname)
			fmt.Printf("%3s %50s\n", "", onGrup.groupText[i].message)
			fmt.Println()
		} else {
			fmt.Printf("%3d %10s\n", onGrup.groupText[i].textID, onGrup.groupText[i].authorUname)
			fmt.Printf("%3s %10s\n", "", onGrup.groupText[i].message)
			fmt.Println()
		}
	}
}

func menu_grup_chat_setting_views(grupIdx int) {
	fmt.Printf("%15s %s\n", "~~ Grup ", dataGrup[grupIdx].grupName)
	fmt.Println("|~~~|--------Menu Info & Pengaturan Grup--------|~~~|")
	fmt.Println("|~~~| 1. Ubah Nama Grup                         |~~~|")
	fmt.Println("|~~~| 2. Lihat Anggota                          |~~~|")
	fmt.Println("|~~~| 3. Keluar Grup                            |~~~|")
	fmt.Println("|~~~| 4. Kembali                                |~~~|")
	fmt.Println("|~~~|-------------------------------------------|~~~|")
}

func menu_grup_chat_setting_member_views(grupIdx int) {
	fmt.Println("|~~~|----------------Anggota Grup---------------|~~~|")
	for i := 0; i < dataGrup[grupIdx].nMember; i++ {
		fmt.Printf("%15d %s\n", i+1, dataAkun[searchAkunIdx(dataAkun, nDataAkun, dataGrup[grupIdx].memberUname[i])].name)
	}
	fmt.Println("|~~~|-------------------------------------------|~~~|")

}

func menu_home_setting_akun_views(idxAkunDipakai int) {
	var passHidden string
	fmt.Println("|~~~|---------------Pengaturan Akun-------------|~~~|")
	fmt.Println("|~~~|---------------Informasi Akun--------------|~~~|")
	fmt.Printf("%15s %6s %s\n", "Nama", ":", dataAkun[idxAkunDipakai].name)
	fmt.Printf("%15s %6s %s\n", "Username", ":", dataAkun[idxAkunDipakai].uname)
	for i := 0; i < len(dataAkun[idxAkunDipakai].pass); i++ {
		passHidden += "*"
	}
	fmt.Printf("%15s %6s %s\n", "Password", ":", passHidden)
	fmt.Println("|~~~|-------------------------------------------|~~~|")
	fmt.Println("|~~~| 1. Ubah nama                              |~~~|")
	fmt.Println("|~~~| 2. Ubah username                          |~~~|")
	fmt.Println("|~~~| 3. Ubah password                          |~~~|")
	fmt.Println("|~~~| 4. Hapus akun                             |~~~|")
	fmt.Println("|~~~| 5. Kembali                                |~~~|")
	fmt.Println("|~~~|-------------------------------------------|~~~|")
}

// Controller Menu
func menu_welcome() {
	var input string
	menu_welcome_views()
	fmt.Print("\n\nPress (Y/N) to continue...")
	fmt.Scan(&input)
	for input != "Y" && input != "N" {
		fmt.Println("Masukkan salah, silahkan input kembali.")
		fmt.Scan(&input)
	}
	if input == "Y" {
		clearScreen()
		menu_utama()
	}
}

func menu_utama() {
	var input string
	
	menu_welcome_views()
	menu_utama_views()
	fmt.Print("Pilih Menu (1/2/3/4): ")
	fmt.Scan(&input)
	for input != "1" && input != "2" && input != "3" && input != "4" {
		fmt.Println("Masukkan salah, silahkan input kembali.")
		fmt.Print("Pilih Menu (1/2/3/4): ")
		fmt.Scan(&input)
	}
	if input == "1" {
		clearScreen()
		menu_registrasi()
	} else if input == "2" {
		clearScreen()
		menu_login()
	} else if input == "3" {
		clearScreen()
		menu_admin()
	}
}

func menu_registrasi() {
	menu_registrasi_views()
}

func menu_login() {
	menu_login_views()
}

func menu_admin() {
	var input string
	
	menu_admin_views()
	fmt.Print("Pilih Menu (1/2/3): ")
	fmt.Scan(&input)

	for input != "1" && input != "2" && input != "3" {
		fmt.Println("Masukkan salah, silahkan input kembali.")
		fmt.Print("Pilih Menu (1/2/3): ")
		fmt.Scan(&input)
	}

	if input == "1" {
		clearScreen()
		menu_admin_daftar_akun()
	} else if input == "2" {
		clearScreen()
		menu_admin_daftar_akun_pending()
	} else {
		clearScreen()
		menu_utama()
	}
}

func menu_admin_daftar_akun() {
	var input, uname string
	var idxAkun int
	var dataAkunCopy accounts

	menu_admin_views_daftar_akun()
	fmt.Println("Pilihan: 1. Hapus Akun, 2. Sort Akun (nama), 3. Sort Akun (umur), 4. Sort Akun (gender), 5. Kembali")
	fmt.Print("Pilih (1/2/3/4/5): ")
	fmt.Scan(&input)

	for input != "1" && input != "2" && input != "3" && input != "4" && input != "5" {
		fmt.Println("Masukkan salah, silahkan input kembali.")
		fmt.Print("Pilih (1/2/3/4/5): ")
		fmt.Scan(&input)
	}

	if input == "1" {
		fmt.Print("Masukkan username akun yang akan dihapus: ")
		fmt.Scan(&uname)
		idxAkun = searchAkunIdx(dataAkun, nDataAkun, uname)
		if idxAkun != -1 {
			clearScreen()
			fmt.Println("Akun", dataAkun[idxAkun].name, "berhasil dihapus.")
			hapusAkun(&dataAkun, &nDataAkun, dataAkunPending[idxAkun].uname)
			menu_admin_daftar_akun()
		} else {
			clearScreen()
			fmt.Println("Username tidak ditemukan, silahkan input kembali.")
			menu_admin_daftar_akun()
		}
	} else if input == "2" {
		clearScreen()
		dataAkunCopy = dataAkun
		sortAkunByUname(&dataAkunCopy, nDataAkun)
		menu_admin_daftar_akun_sort(dataAkunCopy)
	} else if input == "3" {
		// Sorting akun berdasarkan umur
	} else if input == "4" {
		// Sorting akun berdasarkan gender
	} else {
		clearScreen()
		menu_admin()
	}
}

func menu_admin_daftar_akun_sort(dataAkunCopy accounts) {
	var input string

	menu_admin_views_daftar_akun_sort(dataAkunCopy)
	fmt.Println("Pilihan: 1. Kembali")
	fmt.Print("Pilih (1): ")
	fmt.Scan(&input)

	for input != "1"  {
		fmt.Println("Masukkan salah, silahkan input kembali.")
		fmt.Print("Pilih (1): ")
		fmt.Scan(&input)
	}
	if input == "1" {
		clearScreen()
		menu_admin_daftar_akun()
	}
}

func menu_admin_daftar_akun_pending() {
	var input, uname string
	var idxAkun int
	
	menu_admin_views_daftar_akun_pending()

	fmt.Println("Pilihan: 1. Setujui Akun, 2. Tolak Akun, 3. Kembali")
	fmt.Print("Pilih (1/2/3): ")
	fmt.Scan(&input)

	for input != "1" && input != "2" && input != "3" {
		fmt.Println("Masukkan salah, silahkan input kembali.")
		fmt.Print("Pilih (1/2/3): ")
		fmt.Scan(&input)
	}

	if input == "1" {
		fmt.Print("Masukkan username akun yang akan disetujui: ")
		fmt.Scan(&uname)
		idxAkun = searchAkunIdx(dataAkunPending, nDataAkunPending, uname)
		if idxAkun != -1 {
			clearScreen()
			fmt.Println("Akun", dataAkunPending[idxAkun].name, "berhasil disetujui.")
			dataAkun[nDataAkun] = dataAkunPending[idxAkun]
			nDataAkun++
			hapusAkun(&dataAkunPending, &nDataAkunPending, dataAkunPending[idxAkun].uname)
			menu_admin_daftar_akun_pending()
		} else {
			clearScreen()
			fmt.Println("Username tidak ditemukan, silahkan input kembali.")
			menu_admin_daftar_akun_pending()
		}
	} else if input == "2" {
		fmt.Print("Masukkan username akun yang akan ditolak: ")
		fmt.Scan(&uname)
		idxAkun = searchAkunIdx(dataAkunPending, nDataAkunPending, uname)
		if idxAkun != -1 {
			clearScreen()
			fmt.Println("Akun", dataAkunPending[idxAkun].name, "berhasil ditolak.")
			hapusAkun(&dataAkunPending, &nDataAkunPending, dataAkunPending[idxAkun].uname)
			menu_admin_daftar_akun_pending()
		} else {
			clearScreen()
			fmt.Println("Username tidak ditemukan, silahkan input kembali.")
			menu_admin_daftar_akun_pending()
		}
	} else {
		clearScreen()
		menu_admin()
	}
}

func menu_home(idxAkunDipakai int) {
	var input string
	
	fmt.Println("Selamat Datang,", dataAkun[idxAkunDipakai].name)
	menu_home_views()

	fmt.Print("Pilih Menu (1/2/3/4): ")
	fmt.Scan(&input)

	for input != "1" && input != "2" && input != "3" && input != "4" {
		fmt.Println("Masukkan salah, silahkan input kembali.")
		fmt.Print("Pilih Menu (1/2/3/4): ")
		fmt.Scan(&input)
	}

	if input == "1" {
		clearScreen()
		menu_chat(idxAkunDipakai)
	} else if input == "2" {
		clearScreen()
		menu_grup(idxAkunDipakai)
	} else if input == "3" {
		clearScreen()
		menu_home_setting_akun(idxAkunDipakai)
	} else {
		clearScreen()
		menu_utama()
	}
}

func menu_home_setting_akun(idxAkunDipakai int) {
	var input, inputNama, inputPass, inputUname string
	var grupIdx, grupUnameIdx, chatIdx, idxAkunReceiver, chatIdxReceiver int
	
	menu_home_setting_akun_views(idxAkunDipakai)
	fmt.Print("Pilih Menu (1/2/3/4/5): ")
	fmt.Scan(&input)

	for input != "1" && input != "2" && input != "3" && input != "4" && input != "5" {
		fmt.Println("Masukkan salah, silahkan input kembali.")
		fmt.Print("Pilih Menu (1/2/3/4/5): ")
		fmt.Scan(&input)
	}

	if input == "1" {
		fmt.Print("Massukkan nama baru: ")
		fmt.Scan(&inputNama)

		fmt.Println("Pilihan: 1. Ubah, 2. Batal")
		fmt.Print("Pilih (1/2): ")
		fmt.Scan(&input)
		for input != "1" && input != "2" {
			fmt.Println("Masukkan salah, silahkan input kembali.")
			fmt.Print("Pilih (1/2): ")
			fmt.Scan(&input)
		}
		if input == "1" {
			dataAkun[idxAkunDipakai].name = inputNama
			clearScreen()
			fmt.Println("Nama berhasil diubah.")
			menu_home_setting_akun(idxAkunDipakai)
		} else {
			clearScreen()
			menu_home_setting_akun(idxAkunDipakai)
		}
	} else if input == "2" {
		fmt.Print("Massukkan username baru: ")
		fmt.Scan(&inputUname)

		if searchAkunIdx(dataAkun, nDataAkun, inputUname) == -1 && searchAkunIdx(dataAkunPending, nDataAkunPending, inputUname) == -1 {
			fmt.Println("Pilihan: 1. Ubah, 2. Batal")
			fmt.Print("Pilih (1/2): ")
			fmt.Scan(&input)
			for input != "1" && input != "2" {
				fmt.Println("Masukkan salah, silahkan input kembali.")
				fmt.Print("Pilih (1/2): ")
				fmt.Scan(&input)
			}

			if input == "1" {
				for i := 0; i < dataAkun[idxAkunDipakai].nGroup; i++ {
					grupIdx = dataAkun[idxAkunDipakai].joinedGroupID[i]
					grupUnameIdx = searchUnameInGrup(dataGrup[grupIdx], dataAkun[idxAkunDipakai].uname)
					dataGrup[grupIdx].memberUname[grupUnameIdx] = inputUname
				}
				// Ubah yang di chat
				for i := 0; i < dataAkun[idxAkunDipakai].nChat; i++ {
					idxAkunReceiver = searchAkunIdx(dataAkun, nDataAkun, dataAkun[idxAkunDipakai].chatData[i].receiverUname)
					chatIdxReceiver = searchChatIdx(dataAkun[idxAkunReceiver], dataAkun[idxAkunDipakai].uname)
					chatIdx = searchChatIdx(dataAkun[idxAkunDipakai], dataAkun[idxAkunReceiver].uname)
					for j := 0; j < dataAkun[idxAkunReceiver].chatData[chatIdxReceiver].nText; j++ {
						if dataAkun[idxAkunDipakai].uname == dataAkun[idxAkunReceiver].chatData[chatIdxReceiver].textData[j].authorUname {
							dataAkun[idxAkunReceiver].chatData[chatIdxReceiver].textData[j].authorUname = inputUname
						}
					}
					for j := 0; j < dataAkun[idxAkunDipakai].chatData[chatIdx].nText; j++ {
						if dataAkun[idxAkunDipakai].uname == dataAkun[idxAkunDipakai].chatData[chatIdx].textData[j].authorUname {
							dataAkun[idxAkunDipakai].chatData[chatIdx].textData[j].authorUname = inputUname
						}
					}
					dataAkun[idxAkunReceiver].chatData[chatIdx].receiverUname = inputUname
				}
				// Ubah yang di akun
				dataAkun[idxAkunDipakai].uname = inputUname
				clearScreen()
				menu_home_setting_akun(idxAkunDipakai)
			} else {
				clearScreen()
				menu_home_setting_akun(idxAkunDipakai)
			}
		} else {
			clearScreen()
			fmt.Println("Username telah digunakan.")
			menu_home_setting_akun(idxAkunDipakai)
		}
	} else if input == "3" {
		fmt.Print("Masukkan password lama untuk mengubah: ")
		fmt.Scan(&input)

		if input == dataAkun[idxAkunDipakai].pass {
			fmt.Print("Masukkan password baru: ")
			fmt.Scan(&inputPass)
			fmt.Println("Apakah anda yakin ingin mengubah password.")
			fmt.Println("Pilihan: 1. Yakin, 2. Batal")
			fmt.Print("Pilih (1/2): ")
	
			fmt.Scan(&input)
			for input != "1" && input != "2" {
				fmt.Println("Masukkan salah, silahkan input kembali.")
				fmt.Print("Pilih (1/2): ")
				fmt.Scan(&input)
			}

			if input == "1" {
				dataAkun[idxAkunDipakai].pass = inputPass
				clearScreen()
				fmt.Println("Password berhasil diubah.")
				menu_home_setting_akun(idxAkunDipakai)
			} else {
				clearScreen()
				menu_home_setting_akun(idxAkunDipakai)
			}

		} else {
			clearScreen()
			fmt.Println("Password salah.")
			menu_home_setting_akun(idxAkunDipakai)
		}
	} else if input == "4" {
		fmt.Println("Apakah anda yakin ingin menghapus akun.")
		fmt.Println("Pilihan: 1. Yakin, 2. Batal")
		fmt.Print("Pilih (1/2): ")

		fmt.Scan(&input)
		for input != "1" && input != "2" {
			fmt.Println("Masukkan salah, silahkan input kembali.")
			fmt.Print("Pilih (1/2): ")
			fmt.Scan(&input)
		}

		if input == "1" {
			clearScreen()
			fmt.Println("Akun", dataAkun[idxAkunDipakai].name, "berhasil dihapus.")
			hapusAkun(&dataAkun, &nDataAkun, dataAkunPending[idxAkunDipakai].uname)
			menu_utama()
		} else {
			clearScreen()
			menu_home_setting_akun(idxAkunDipakai)
		}
	} else {
		clearScreen()
		menu_home(idxAkunDipakai)
	}
}

func menu_chat(idxAkunDipakai int) {
	var input, uname string
	var idxAkun, idxChat, idxChatReceiver int
	var newChat, newChatReceiver chat

	menu_chat_views(idxAkunDipakai)

	fmt.Println("Pilihan: 1. Pilih Chat, 2. Tambah Chat, 3. Hapus Chat, 4. Kembali")
	fmt.Print("Pilih (1/2/3/4): ")
	fmt.Scan(&input)

	for input != "1" && input != "2" && input != "3" && input != "4" {
		fmt.Println("Masukkan salah, silahkan input kembali.")
		fmt.Print("Pilih (1/2/3/4): ")
		fmt.Scan(&input)
	}
	if input == "1" {
		fmt.Print("Masukkan username akun yang akan Anda chat: ")
		fmt.Scan(&uname)
		idxChat = searchChatIdx(dataAkun[idxAkunDipakai], uname)
		if idxChat != -1 {
			clearScreen()
			idxAkun = searchAkunIdx(dataAkun, nDataAkun, dataAkun[idxAkunDipakai].chatData[idxChat].receiverUname)
			menu_pesan(idxAkunDipakai, uname, idxAkun)
		} else {
			clearScreen()
			fmt.Println("Username tidak ditemukan, silahkan input kembali.")
			menu_chat(idxAkunDipakai)
		}
	} else if input == "2" {
		fmt.Print("Masukkan username akun yang akan Anda chat: ")
		fmt.Scan(&uname)
		for uname == dataAkun[idxAkunDipakai].uname {
			fmt.Println("Username tidak valid, silahkan input kembali.")
			fmt.Print("Masukkan username akun yang akan Anda chat: ")
			fmt.Scan(&uname)
		}
		idxAkun = searchAkunIdx(dataAkun, nDataAkun, uname)
		if idxAkun != -1 {
			if searchChatIdx(dataAkun[idxAkunDipakai], uname) == -1 {
				newChat.receiverUname = dataAkun[idxAkun].uname
				newChatReceiver.receiverUname = dataAkun[idxAkunDipakai].uname
				dataAkun[idxAkunDipakai].chatData[dataAkun[idxAkunDipakai].nChat] = newChat
				dataAkun[idxAkun].chatData[dataAkun[idxAkun].nChat] = newChatReceiver
				dataAkun[idxAkunDipakai].nChat++
				dataAkun[idxAkun].nChat++
				clearScreen()
				menu_pesan(idxAkunDipakai, uname, idxAkun)
			} else {
				clearScreen()
				fmt.Println("Chat dengan", uname, "sudah ada.")
				menu_pesan(idxAkunDipakai, uname, idxAkun)
			}
		} else {
			clearScreen()
			fmt.Println("Username tidak ditemukan, silahkan input kembali.")
			menu_chat(idxAkunDipakai)
		}
	} else if input == "3" {
		fmt.Print("Masukkan username akun yang akan Anda hapus chatnya: ")
		fmt.Scan(&uname)
		idxChat = searchChatIdx(dataAkun[idxAkunDipakai], uname)
		if idxChat != -1 {
			idxAkun = searchAkunIdx(dataAkun, nDataAkun, uname)
			idxChatReceiver = searchChatIdx(dataAkun[idxAkun], dataAkun[idxAkunDipakai].uname)
			hapusChat(&dataAkun[idxAkunDipakai], idxChat)
			hapusChat(&dataAkun[idxAkun], idxChatReceiver)
			clearScreen()
			fmt.Println("Chat berhasil dihapus.")
			menu_chat(idxAkunDipakai)
		} else {
			clearScreen()
			fmt.Println("Username tidak ditemukan, silahkan input kembali.")
			menu_chat(idxAkunDipakai)
		}
	} else {
		clearScreen()
		menu_home(idxAkunDipakai)
	}
}

func menu_pesan(idxAkunDipakai int, unameReceiver string, idxAkunReceiver int) {
	var onChat chat
	var onChatIdx, onChatReceiverIdx int
	var input string
	var pesanBaru text
	var inputHapusPesan int

	onChatIdx = searchChatIdx(dataAkun[idxAkunDipakai], unameReceiver)
	onChat = dataAkun[idxAkunDipakai].chatData[onChatIdx]
	
	if searchAkunIdx(dataAkun, nDataAkun, unameReceiver) != -1 {
		onChatReceiverIdx = searchChatIdx(dataAkun[idxAkunReceiver], dataAkun[idxAkunDipakai].uname)
		menu_pesan_views(onChat)
		fmt.Println("Pilihan: 1. Tambah Pesan, 2. Hapus Pesan, 3. Kembali")
		fmt.Print("Pilih (1/2/3): ")
		fmt.Scan(&input)
	
		for input != "1" && input != "2" && input != "3" {
			fmt.Println("Masukkan salah, silahkan input kembali.")
			fmt.Print("Pilih (1/2/3): ")
			fmt.Scan(&input)
		}
	
		if input == "1" {
			fmt.Print("Masukkan Pesan: ")
			fmt.Scan(&input)
			pesanBaru.message = input
			pesanBaru.textID = dataAkun[idxAkunDipakai].chatData[onChatIdx].nText+1
			pesanBaru.authorUname = dataAkun[idxAkunDipakai].uname
	
			fmt.Println("Pilihan: 1. Kirim, 2. Batal")
			fmt.Print("Pilih (1/2): ")
			fmt.Scan(&input)
			for input != "1" && input != "2" {
				fmt.Println("Masukkan salah, silahkan input kembali.")
				fmt.Print("Pilih (1/2): ")
				fmt.Scan(&input)
			}
	
			if input == "1" {
				dataAkun[idxAkunDipakai].chatData[onChatIdx].textData[dataAkun[idxAkunDipakai].chatData[onChatIdx].nText] = pesanBaru
				dataAkun[idxAkunDipakai].chatData[onChatIdx].nText++
	
				dataAkun[idxAkunReceiver].chatData[onChatReceiverIdx].textData[dataAkun[idxAkunReceiver].chatData[onChatReceiverIdx].nText] = pesanBaru
				dataAkun[idxAkunReceiver].chatData[onChatReceiverIdx].nText++
				
				clearScreen()
				menu_pesan(idxAkunDipakai, unameReceiver, idxAkunReceiver)
			} else {
				clearScreen()
				menu_pesan(idxAkunDipakai, unameReceiver, idxAkunReceiver)
			}
		} else if input == "2" {
			fmt.Print("Masukkan nomor pesan yang ingin dihapus: ")
			fmt.Scan(&inputHapusPesan)
	
			fmt.Println("Pilihan: 1. Hapus, 2. Batal")
			fmt.Print("Pilih (1/2): ")
			fmt.Scan(&input)
			for input != "1" && input != "2" {
				fmt.Println("Masukkan salah, silahkan input kembali.")
				fmt.Print("Pilih (1/2): ")
				fmt.Scan(&input)
			}
	
			if input == "1" {
				if searchPesanID(inputHapusPesan, dataAkun[idxAkunDipakai], onChatIdx) != -1 {
					if dataAkun[idxAkunDipakai].uname == dataAkun[idxAkunDipakai].chatData[onChatIdx].textData[searchPesanID(inputHapusPesan, dataAkun[idxAkunDipakai], onChatIdx)].authorUname {
						hapusPesan(inputHapusPesan, &dataAkun[idxAkunDipakai], &dataAkun[idxAkunReceiver], onChatIdx, onChatReceiverIdx)
						clearScreen()
						menu_pesan(idxAkunDipakai, unameReceiver, idxAkunReceiver)
					} else {
						clearScreen()
						fmt.Println("Nomor pesan yang dipilih tidak valid.")
						menu_pesan(idxAkunDipakai, unameReceiver, idxAkunReceiver)
					}
				} else {
					clearScreen()
					fmt.Println("Nomor pesan tidak ada.")
					menu_pesan(idxAkunDipakai, unameReceiver, idxAkunReceiver)
				}
			} else {
				clearScreen()
				menu_pesan(idxAkunDipakai, unameReceiver, idxAkunReceiver)
			}
		} else {
			clearScreen()
			menu_chat(idxAkunDipakai)
		}
	} else {
		fmt.Print("Akun ", unameReceiver, " tidak terdaftar.\n")
		menu_pesan_views(onChat)
		fmt.Println("Pilihan: 1. Hapus Chat, 2. Kembali")
		fmt.Print("Pilih (1/2): ")
		fmt.Scan(&input)
	
		for input != "1" && input != "2" {
			fmt.Println("Masukkan salah, silahkan input kembali.")
			fmt.Print("Pilih (1/2): ")
			fmt.Scan(&input)
		}
		if input == "1" {
			fmt.Println("Apakah anda yakin ingin menghapus chat? ")
			fmt.Println("Pilihan: 1. Yakin, 2. Batal")
			fmt.Scan(&input)
			if input == "1" {
				hapusChat(&dataAkun[idxAkunDipakai], onChatIdx)
				clearScreen()
				fmt.Println("Chat berhasil dihapus.")
				menu_chat(idxAkunDipakai)
			} else {
				clearScreen()
				menu_pesan(idxAkunDipakai, unameReceiver, idxAkunReceiver)
			}
		} else {
			clearScreen()
			menu_chat(idxAkunDipakai)
		}
	}
}

func menu_grup(idxAkunDipakai int) {
	var input string
	var noPilihGrup int
	var grupIdx int

	menu_grup_views(idxAkunDipakai)

	fmt.Println("Pilihan: 1. Pilih Grup, 2. Buat Grup, 3. Kembali")
	fmt.Print("Pilih (1/2/3): ")
	fmt.Scan(&input)

	for input != "1" && input != "2" && input != "3" {
		fmt.Println("Masukkan salah, silahkan input kembali.")
		fmt.Print("Pilih (1/2/3): ")
		fmt.Scan(&input)
	}
	if input == "1" {
		fmt.Print("Masukkan nomor grup yang ingin di pilih: ")
		fmt.Scan(&noPilihGrup)
		if noPilihGrup >= 1 && noPilihGrup <= dataAkun[idxAkunDipakai].nGroup {
			grupIdx = searchGrupIdxFromAcc(dataAkun[idxAkunDipakai], noPilihGrup)
			if grupIdx != -1 {
				clearScreen()
				menu_grup_chat(idxAkunDipakai, grupIdx)
			} else {
				clearScreen()
				fmt.Println("Grup tidak ada.")
				menu_grup(idxAkunDipakai)
			}
		} else {
			clearScreen()
			fmt.Println("Nomor Grup tidak valid.")
			menu_grup(idxAkunDipakai)
		}
	} else if input == "2" {
		clearScreen()
		menu_grup_buat_grup(idxAkunDipakai)
	} else {
		clearScreen()
		menu_home(idxAkunDipakai)
	}
}

func menu_grup_buat_grup(idxAkunDipakai int) {
	var stopInput bool
	var uname, inputLanjut, input string
	var grupBaru group
	grupBaru.grupID = nDataGrup+1

	fmt.Println("|~~~|------------------Buat Grup----------------|~~~|")
	fmt.Print("Masukkan nama grup: ")
	fmt.Scan(&grupBaru.grupName)
	fmt.Print("Tambahkan anggota grup (username): ")
	fmt.Scan(&uname)
	for searchAkunIdx(dataAkun, nDataAkun, uname) == -1 || uname == dataAkun[idxAkunDipakai].uname {
		fmt.Println("Username tidak valid, silahkan input kembali.")
		fmt.Print("Tambahkan anggota grup (username): ")
		fmt.Scan(&uname)
	}
	grupBaru.memberUname[grupBaru.nMember] = uname
	grupBaru.nMember++

	for !stopInput {
		fmt.Print("Apakah ingin menambahkan anggota lain? (Y/N): ")
		fmt.Scan(&inputLanjut)
		if inputLanjut == "N" {
			stopInput = true
		} else {
			fmt.Print("Tambahkan anggota grup (username): ")
			fmt.Scan(&uname)
			if searchAkunIdx(dataAkun, nDataAkun, uname) != -1 {
				grupBaru.memberUname[grupBaru.nMember] = uname
				grupBaru.nMember++
			} else {
				fmt.Println("Username tidak ditemukan, silahkan input kembali.")
			}
		}
	}

	grupBaru.memberUname[grupBaru.nMember] = dataAkun[idxAkunDipakai].uname
	grupBaru.nMember++
	
	clearScreen()
	fmt.Println("|~~~|---------------Data Grup Baru--------------|~~~|")
	fmt.Printf("%15s : %s\n", "Nama Grup", grupBaru.grupName)
	fmt.Printf("%15s :\n", "Anggota Grup")
	for i := 0; i < grupBaru.nMember; i++ {
		fmt.Printf("%15d. %s\n", i+1, grupBaru.memberUname[i])
	}
	fmt.Println("|~~~|-------------------------------------------|~~~|")

	fmt.Println("Pilihan: 1. Buat Grup, 2. Batal")
	fmt.Print("Pilih (1/2): ")
	fmt.Scan(&input)

	for input != "1" && input != "2" {
		fmt.Println("Masukkan salah, silahkan input kembali.")
		fmt.Print("Pilih (1/2): ")
		fmt.Scan(&input)
	}

	if input == "1" {
		dataGrup[nDataGrup] = grupBaru
		for i := 0; i < grupBaru.nMember; i++ {
			dataAkun[searchAkunIdx(dataAkun, nDataAkun, grupBaru.memberUname[i])].joinedGroupID[dataAkun[searchAkunIdx(dataAkun, nDataAkun, grupBaru.memberUname[i])].nGroup] = nDataGrup
			dataAkun[searchAkunIdx(dataAkun, nDataAkun, grupBaru.memberUname[i])].nGroup++
		}
		nDataGrup++
		clearScreen()
		fmt.Println("Grup berhasil dibuat.")

		menu_grup_chat(idxAkunDipakai, nDataGrup-1)
	} else {
		clearScreen()
		menu_grup(idxAkunDipakai)
	}
}

func menu_grup_chat(idxAkunDipakai, grupIdx int) {
	var input, inputPesan string
	var pesanBaru text
	var inputNoHapusPesan int

	menu_grup_chat_views(dataGrup[grupIdx], dataAkun[idxAkunDipakai].uname)

	fmt.Println("Pilihan: 1. Tambah Pesan, 2. Hapus Pesan, 3. Info & Pengaturan, 4. Kembali")
	fmt.Print("Pilih (1/2/3/4): ")
	fmt.Scan(&input)

	for input != "1" && input != "2" && input != "3" && input != "4" {
		fmt.Println("Masukkan salah, silahkan input kembali.")
		fmt.Print("Pilih (1/2/3/4): ")
		fmt.Scan(&input)
	}

	if input == "1" {
		fmt.Print("Masukkan Pesan: ")
		fmt.Scan(&inputPesan)

		fmt.Println("Pilihan: 1. Kirim, 2. Batal")
		fmt.Print("Pilih (1/2): ")
		fmt.Scan(&input)

		for input != "1" && input != "2" {
			fmt.Println("Masukkan salah, silahkan input kembali.")
			fmt.Print("Pilih (1/2): ")
			fmt.Scan(&input)
		}

		if input == "1" {
			pesanBaru.message = inputPesan
			pesanBaru.authorUname = dataAkun[idxAkunDipakai].uname
			pesanBaru.textID = dataGrup[grupIdx].nGroupText+1

			dataGrup[grupIdx].groupText[dataGrup[grupIdx].nGroupText] = pesanBaru
			dataGrup[grupIdx].nGroupText++

			clearScreen()
			menu_grup_chat(idxAkunDipakai, grupIdx)
		} else {
			clearScreen()
			menu_grup_chat(idxAkunDipakai, grupIdx)
		}
	} else if input == "2" {
		fmt.Print("Masukkan nomor pesan yang ingin dihapus: ")
		fmt.Scan(&inputNoHapusPesan)

		fmt.Println("Pilihan: 1. Hapus, 2. Batal")
		fmt.Print("Pilih (1/2): ")
		fmt.Scan(&input)
		for input != "1" && input != "2" {
			fmt.Println("Masukkan salah, silahkan input kembali.")
			fmt.Print("Pilih (1/2): ")
			fmt.Scan(&input)
		}

		if input == "1" {
			if searchPesanGrupID(dataGrup[grupIdx], inputNoHapusPesan) != -1 {
				hapusPesanGrup(grupIdx, searchPesanGrupID(dataGrup[grupIdx], inputNoHapusPesan))
				clearScreen()
				menu_grup_chat(idxAkunDipakai, grupIdx)
			} else {
				clearScreen()
				fmt.Println("Nomor pesan yang dipilih tidak valid.")
				menu_grup_chat(idxAkunDipakai, grupIdx)
			}
		} else {
			clearScreen()
			menu_grup_chat(idxAkunDipakai, grupIdx)
		}
	} else if input == "3" {
		clearScreen()
		menu_grup_chat_setting(idxAkunDipakai, grupIdx)
	} else {
		clearScreen()
		menu_grup(idxAkunDipakai)
	}
}

func menu_grup_chat_setting(idxAkunDipakai, grupIdx int) {
	var input, inputNamaGrup string
	
	menu_grup_chat_setting_views(grupIdx)

	fmt.Print("Pilih Menu (1/2/3/4): ")
	fmt.Scan(&input)

	for input != "1" && input != "2" && input != "3" && input != "4" {
		fmt.Println("Masukkan salah, silahkan input kembali.")
		fmt.Print("Pilih Menu (1/2/3/4): ")
		fmt.Scan(&input)
	}
	if input == "1" {
		fmt.Print("Masukkan nama grup baru: ")
		fmt.Scan(&inputNamaGrup)

		fmt.Println("Apakah anda yakin ingin mengubah nama grup.")
		fmt.Println("Pilihan: 1. Yakin, 2. Batal")
		fmt.Print("Pilih (1/2): ")

		fmt.Scan(&input)
		for input != "1" && input != "2" {
			fmt.Println("Masukkan salah, silahkan input kembali.")
			fmt.Print("Pilih (1/2): ")
			fmt.Scan(&input)
		}
		if input == "1" {
			dataGrup[grupIdx].grupName = inputNamaGrup
			clearScreen()
			fmt.Println("Nama grup berhasil diubah.")
			menu_grup_chat_setting(idxAkunDipakai, grupIdx)
		} else {
			clearScreen()
			menu_grup_chat_setting(idxAkunDipakai, grupIdx)
		}
	} else if input == "2" {
		clearScreen()
		menu_grup_chat_setting_member(idxAkunDipakai, grupIdx)
	} else if input == "3" {
		fmt.Println("Apakah anda yakin ingin keluar dari grup ini (Y/N): ")
		fmt.Scan(&input)
		if input == "Y" {
			keluarGrup(&dataAkun[idxAkunDipakai], grupIdx)
			clearScreen()
			menu_grup(idxAkunDipakai)
		} else {
			clearScreen()
			menu_grup_chat_setting(idxAkunDipakai, grupIdx)
		}
	} else {
		clearScreen()
		menu_grup_chat(idxAkunDipakai, grupIdx)
	}
}

func menu_grup_chat_setting_member(idxAkunDipakai, grupIdx int) {
	var input string

	menu_grup_chat_setting_member_views(grupIdx)

	fmt.Println("Pilihan: 1. Tambah Anggota, 2. Kembali")
	fmt.Print("Pilih (1/2): ")
	fmt.Scan(&input)

	for input != "1" && input != "2" {
		fmt.Println("Masukkan salah, silahkan input kembali.")
		fmt.Print("Pilih (1/2): ")
		fmt.Scan(&input)
	}

	if input == "1" {
		fmt.Print("Masukkan username akun yang ingin ditambahkan: ")
		fmt.Scan(&input)
		
		if searchAkunIdx(dataAkun, nDataAkun, input) != -1 {
			if searchUnameInGrup(dataGrup[grupIdx], input) == -1 {
				tambahGrupMember(&dataGrup[grupIdx], input, grupIdx)
				clearScreen()
				fmt.Println("Username berhasil ditambahkan.")
				menu_grup_chat_setting_member(idxAkunDipakai, grupIdx)
			} else {
				clearScreen()
				fmt.Println("Username tersebut sudah ada di dalam grup.")
				menu_grup_chat_setting_member(idxAkunDipakai, grupIdx)
			}
		} else {
			clearScreen()
			fmt.Println("Username tidak ditemukan, silahkan input kembali.")
			menu_grup_chat_setting_member(idxAkunDipakai, grupIdx)
		}
	} else {
		clearScreen()
		menu_grup_chat_setting(idxAkunDipakai, grupIdx)
	}
}

// Fungsi / Prosedur Tambahan
func hapusAkun(data *accounts, nData *int, uname string) {
	var idx int = searchAkunIdx(*data, *nData, uname)
	
	// Tambain kondisi kalo dia ada di grup berarti keluar grup dulu
	for i := 0; i < data[idx].nGroup; i++ {
		keluarGrup(&data[idx], data[idx].joinedGroupID[i])
	}
	// Tambain juga kondisi buat yang chat sama dia
	for i := idx; i < *nData-1; i++ {
		data[i] = data[i+1]
	}
	*nData--
}

func hapusChat(acc *account, idxChat int) {
	for i := idxChat; i < acc.nChat-1; i++ {
		acc.chatData[i] = acc.chatData[i+1]
	}
	acc.nChat--
}

func searchAkunIdx(data accounts, nData int, uname string) int {
	var found int = -1
	var i int
	for i < nData && found == -1 {
		if uname == data[i].uname {
			found = i
		}
		i++
	}
	return found
}

func searchChatIdx(acc account, unameReceiver string) int {
	var found int = -1
	var i int
	for i < acc.nChat && found == -1 {
		if unameReceiver == acc.chatData[i].receiverUname {
			found = i
		}
		i++
	}
	return found
}

func searchPesanID(nomorPesan int, onChatAcc account, idxChat int) int {
	var le, ri, mid, idx int
	idx = -1
	le = 0
	ri = onChatAcc.chatData[idxChat].nText-1
	mid = (le+ri)/2
	for le <= ri && onChatAcc.chatData[idxChat].textData[mid].textID != nomorPesan {
		if nomorPesan < onChatAcc.chatData[idxChat].textData[mid].textID {
			ri = mid - 1
		} else {
			le = mid + 1
		}
		mid = (le+ri) / 2
	}
	if onChatAcc.chatData[idxChat].textData[mid].textID == nomorPesan && mid >= 0 {
		idx = mid
	}
	return idx
}

func hapusPesan(pesanID int, acc *account, accReceiver *account, idxChat, idxChatReceiver int) {
	var idxPesan, tempIdxPesanLama int
	idxPesan = searchPesanID(pesanID, *acc, idxChat)
	for i := idxPesan; i < acc.chatData[idxChat].nText-1; i++ {
		tempIdxPesanLama = acc.chatData[idxChat].textData[i].textID
		acc.chatData[idxChat].textData[i] = acc.chatData[idxChat].textData[i+1]
		acc.chatData[idxChat].textData[i].textID = tempIdxPesanLama
	}
	for i := idxPesan; i < accReceiver.chatData[idxChatReceiver].nText-1; i++ {
		tempIdxPesanLama = accReceiver.chatData[i].textData[idxPesan].textID
		accReceiver.chatData[i].textData[idxPesan] = accReceiver.chatData[i].textData[idxPesan+1]
		accReceiver.chatData[i].textData[idxPesan].textID = tempIdxPesanLama
	}
	acc.chatData[idxChat].nText--
	accReceiver.chatData[idxChatReceiver].nText--
}

func searchGrupIdx(inputIdGrup int) int {
	var found int = -1
	var i int
	for i < nDataGrup && found == -1 {
		if inputIdGrup == dataGrup[i].grupID {
			found = i
		}
		i++
	}
	return found
}

func searchGrupIdxFromAcc(acc account, inputNo int) int {
	return acc.joinedGroupID[inputNo-1]
}

func searchUnameInGrup(onGrup group, uname string) int {
	var found int = -1
	var i int
	for i < onGrup.nMember && found == -1 {
		if uname == onGrup.memberUname[i] {
			found = i
		}
		i++
	}
	return found
}

func tambahGrupMember(onGrup *group, unameBaru string, grupIdx int) {
	dataAkun[searchAkunIdx(dataAkun, nDataAkun, unameBaru)].joinedGroupID[dataAkun[searchAkunIdx(dataAkun, nDataAkun, unameBaru)].nGroup] = grupIdx
	dataAkun[searchAkunIdx(dataAkun, nDataAkun, unameBaru)].nGroup++
	onGrup.memberUname[onGrup.nMember] = unameBaru
	onGrup.nMember++
}

func hapusMemberGrup(grupIdx, memberIdx int) {
	for i := memberIdx; i < dataGrup[grupIdx].nMember-1; i++ {
		dataGrup[grupIdx].memberUname[i] = dataGrup[grupIdx].memberUname[i+1]
	}
	dataGrup[grupIdx].nMember--
}

func hapusJoinedGrup(acc *account, grupIdx int) {
	var found bool = false
	for i := 0; i < acc.nGroup && !found; i++ {
		if acc.joinedGroupID[i] == grupIdx {
			for j := i; j < acc.nGroup-1; j++ {
				acc.joinedGroupID[j] = acc.joinedGroupID[j+1]
			}
			found = true
		}
	}
	acc.nGroup--
}

func keluarGrup(acc *account, grupIdx int) {
	var finish bool = false
	var i int

	for i < dataGrup[grupIdx].nMember && !finish {
		if dataGrup[grupIdx].memberUname[i] == acc.uname {
			hapusMemberGrup(grupIdx, i)
			hapusJoinedGrup(acc, grupIdx)
			finish = true
		}
		i++
	}
	// Perlu tambain kondisi kalo anggota grup habis brarti hapus grup
	if dataGrup[grupIdx].nMember == 0 {
		hapusGrup(grupIdx)
	}
}

func hapusGrup(grupIdx int) {
	var idx int = grupIdx
	for i := idx; i < nDataGrup-1; i++ {
		dataGrup[i] = dataGrup[i+1]
	}
	nDataGrup--
}

func searchPesanGrupID(onGrup group, inputPesanID int) int {

	var found int = -1
	var i int
	for i < onGrup.nGroupText && found == -1 {
		if inputPesanID == onGrup.groupText[i].textID {
			found = i
		}
		i++
	}
	return found
}

func hapusPesanGrup(grupIdx, pesanGrupID int) {
	var tempIdxPesanLama text
	for i := pesanGrupID; i < dataGrup[grupIdx].nGroupText-1; i++ {
		tempIdxPesanLama = dataGrup[grupIdx].groupText[i]
		dataGrup[grupIdx].groupText[i] = dataGrup[grupIdx].groupText[i+1]
		dataGrup[grupIdx].groupText[i].textID = tempIdxPesanLama.textID
	}
	dataGrup[grupIdx].nGroupText--
}

// Fungsi Insertion Sort untuk mengurutkan akun berdasarkan nama
func sortAkunByUname(data *accounts, nData int) {
	var i, j int
	var key account

	for i = 1; i < nData; i++ {
		key = data[i]
		j = i - 1
		for j >= 0 && data[j].uname > key.uname {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = key
	}
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin"{
		cmd = exec.Command("clear")
	} else {
		fmt.Println("Unsupported platform")
		return
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
