package main

import (
	"a21hc3NpZ25tZW50/helper"
	"fmt"
	"strings"
)

var Students = []string{
	"A1234_Aditira_TI",
	"B2131_Dito_TK",
	"A3455_Afis_MI",
}

var StudentStudyPrograms = map[string]string{
	"TI": "Teknik Informatika",
	"TK": "Teknik Komputer",
	"SI": "Sistem Informasi",
	"MI": "Manajemen Informasi",
}

type studentModifier func(string, *string)
// func main(){
// 	fmt.Println(Login("A1234", "Aditira"))
// }
func Login(id string, name string) string {
	if id == "" || name ==""{
		return "ID or Name is undefined!"
	} else {
		StudentData := []string{}
		CheckName := false
		checkID := false
		for _, info:= range Students {
			StudentData = strings.Split(info, "_")
			CheckName= strings.Contains(StudentData[1], name)
			checkID = strings.Contains(StudentData[0], id)
			if CheckName && checkID{
				return fmt.Sprintf("Login berhasil: %s", name)
			} 
		}
		return "Login gagal: data mahasiswa tidak ditemukan"
	}
	return "" // TODO: replace this
}


func Register(id string, name string, major string) string {
	// Cek apakah id, name, atau major kosong
    if id == "" || name == "" || major == "" {
        return "ID, Name or Major is undefined!"
    }

    // Cek apakah id sudah digunakan sebelumnya
	StudentData := []string{}
	checkID := false
    for _, info := range Students {
		StudentData = strings.Split(info, "_")
		checkID = strings.Contains(StudentData[0], id)
        if checkID {
            return "Registrasi gagal: id sudah digunakan"
        }
    }

    // Tambahkan data mahasiswa baru ke dalam slice Students
    var student string = fmt.Sprintf("%s_%s_%s", id, name, major)
    Students = append(Students, student)

    // Berhasil mendaftarkan mahasiswa baru
    return "Registrasi berhasil: " + name + " (" + major + ")"
}
// func main() {
// 	fmt.Println(GetStudyPrograms("mi"))
// }
func GetStudyProgram(code string) string {
	checkMajor := false
	for key := range StudentStudyPrograms {
		checkMajor =strings.Contains(key, strings.ToUpper(code))
		if checkMajor {
			return StudentStudyPrograms[key]
		}
	}
	return "Kode program studi tidak ditemukan"
}

// func main() {
// 	fmt.Println(ModifyStudnt("TI", "Aditira", UpdateStudyProgram))
	// "A1234_Aditira_TI"e
	// "B1231_Ditto_TK"
	// "A3455_Afis_MI"
// }
func ModifyStudent(programStudi, nama string, fn studentModifier) string {
	for _, info:=range Students{
		StudentData := strings.Split(info, "_")
		CheckName := strings.Contains(StudentData[1], nama)
		if CheckName {
			fn(programStudi, &StudentData[2])
			DataFinal :=strings.Join(StudentData, "_")
			fmt.Println(DataFinal)
			return "Program studi mahasiswa berhasil diubah"
		}
	}
	return "Mahasiswa tidak ditemukan"
}


func UpdateStudyProgram(programStudi string, students *string) {
	// TODO: answer here
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		for i, student := range Students {
			fmt.Println(i+1, student)
		}

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Change student study program")
		fmt.Println("5. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			helper.ClearScreen()
			var id, name string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)

			fmt.Println(Login(id, name))

			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			var id, name, jurusan string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)
			fmt.Print("Masukkan jurusan: ")
			fmt.Scan(&jurusan)
			fmt.Println(Register(id, name, jurusan))

			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			var kode string
			fmt.Print("Masukkan kode: ")
			fmt.Scan(&kode)

			fmt.Println(GetStudyProgram(kode))
			helper.Delay(5)
		case "4":
			helper.ClearScreen()
			var nama, programStudi string
			fmt.Print("Masukkan nama mahasiswa: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan program studi baru: ")
			fmt.Scanln(&programStudi)

			fmt.Println(ModifyStudent(programStudi, nama, UpdateStudyProgram))
			helper.Delay(5)
		case "5":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
