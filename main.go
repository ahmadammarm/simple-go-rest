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
    { NIM: "220535601234", Nama: "Budi", Prodi: "Teknik Informatika", Fakultas: "Teknik" },
    { NIM: "221085802134", Nama: "Andi", Prodi: "Sistem Informasi", Fakultas: "Ilmu Komputer" },
    { NIM: "220535610023", Nama: "Caca", Prodi: "Teknik Industri", Fakultas: "Teknik" },
    { NIM: "220751601425", Nama: "Dedi", Prodi: "Pendidikan Sosiologi", Fakultas: "Ilmu Sosial" },
}


// Handler untuk mengambil data mahasiswa
func getMahasiswa(c *gin.Context) {
    c.JSON(http.StatusOK, mahasiswa)
}


// Handler untuk membuat data mahasiswa
func postMahasiswa(c *gin.Context) {
    var input Mahasiswa

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    mahasiswa = append(mahasiswa, input)
    c.JSON(http.StatusOK, input)
}

// Handler untuk melihat data setiap mahasiswa berdasarkan NIM
func getMahasiswaByNIM(c *gin.Context) {
    NIM := c.Param("nim")

    for _, item := range mahasiswa {
        if item.NIM == NIM {
            c.JSON(http.StatusOK, item)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"message": "Mahasiswa tidak ditemukan"})
}

// Handler untuk mengedit data setiap mahasiswa berdasarkan NIM

func editMahasiswaByNIM(c *gin.Context) {
    NIM := c.Param("nim")
    var input Mahasiswa

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for i, item := range mahasiswa {
        if item.NIM == NIM {
            mahasiswa[i] = input
            c.JSON(http.StatusOK, input)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"message": "Mahasiswa tidak ditemukan"})
}

// Handler untuk menghapus data setiap mahasiswa berdasarkan NIM

func deleteMahasiswaByNIM(c *gin.Context) {
    NIM := c.Param("nim")

    for i, item := range mahasiswa {
        if item.NIM == NIM {
            mahasiswa = append(mahasiswa[:i], mahasiswa[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "Mahasiswa berhasil dihapus"})
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"message": "Mahasiswa tidak ditemukan"})
}


func main() {
    // Membuat router baru menggunakan gin
    router := gin.Default()

    // Menambahkan route untuk membuat data mahasiswa dengan methodnya
    router.GET("/mahasiswa", getMahasiswa)
    router.POST("/mahasiswa", postMahasiswa)
    router.GET("/mahasiswa/:nim", getMahasiswaByNIM)
    router.PUT("/mahasiswa/:nim", editMahasiswaByNIM)
    router.DELETE("/mahasiswa/:nim", deleteMahasiswaByNIM)



    router.Run(":8080")
}