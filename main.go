package main

// Mengimpor package net/http untuk membuat server HTTP
// Mengimpor package github.com/gin-gonic/gin sebagai web framework
import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// Menginisialisasi tipe data Mahasiswa yang memiliki 4 properti
type Mahasiswa struct {
    NIM   string `json:"nim"`
    Nama  string  `json:"nama"`
    Prodi string `json:"prodi"`
    Fakultas string `json:"fakultas"`
}

// Menginisialisasi data awal mahasiswa dalam bentuk array dari tipe data Mahasiswa
var mahasiswa = []Mahasiswa{
    {NIM: "220535601234", Nama: "Budi", Prodi: "Teknik Informatika", Fakultas: "Teknik"},
    {NIM: "221085802134", Nama: "Andi", Prodi: "Sistem Informasi", Fakultas: "Ilmu Komputer"},
    {NIM: "220535610023", Nama: "Caca", Prodi: "Teknik Industri", Fakultas: "Teknik"},
    {NIM: "220751601425", Nama: "Dedi", Prodi: "Pendidikan Sosiologi", Fakultas: "Ilmu Sosial"},
}


// Handler untuk mengambil data mahasiswa dalam format JSON
func getMahasiswa(c *gin.Context) {
    c.JSON(http.StatusOK, mahasiswa)
}


// Handler untuk membuat data mahasiswa dalam format JSON
func postMahasiswa(c *gin.Context) {
    var input Mahasiswa

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    mahasiswa = append(mahasiswa, input)
    c.JSON(http.StatusOK, input)
}

func main() {
    // Membuat router baru menggunakan gin
    router := gin.Default()

    // Menambahkan route untuk membuat data mahasiswa dengan methodnya
    router.GET("/mahasiswa", getMahasiswa)
    router.POST("/mahasiswa", postMahasiswa)

    router.Run(":8080")
}