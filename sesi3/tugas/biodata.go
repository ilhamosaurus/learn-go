package main

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	Nama              string
	Alamat            string
	Pekerjaan         string
	AlasanPilihGolang string
}

var dataTeman = []Teman{
	{Nama: "Budi", Alamat: "Jalan Sudirman No. 1", Pekerjaan: "Mahasiswa", AlasanPilihGolang: "Ingin menjadi programmer"},
	{Nama: "Ani", Alamat: "Jalan Merdeka No. 2", Pekerjaan: "Freelancer", AlasanPilihGolang: "Belajar bahasa pemrograman baru"},
	{Nama: "Caca", Alamat: "Jalan Mawar No. 3", Pekerjaan: "Web Developer", AlasanPilihGolang: "Menyukai sintaks Golang yang mudah dipelajari"},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run biodata.go <nomor absen>")
		return
	}

	nomorAbsenString := os.Args[1]
	nomorAbsen, err := strconv.Atoi(nomorAbsenString)
	if err != nil {
		fmt.Println("Nomor Absen tidak valid")
		return
	}

	indexTeman := nomorAbsen - 1
	if indexTeman < 0 || indexTeman >= len(dataTeman) {
		fmt.Println("Data teman dengan nomor absen", nomorAbsen, "tidak ditemukan")
		return
	}

	teman := dataTeman[indexTeman]

	fmt.Println("## Data teman ##")
	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan Pilih Golang:", teman.AlasanPilihGolang)
}
