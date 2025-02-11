package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Mahasiswa struct {
    ID    int   `json:"id"`
    Nama  string  `json:"nama"`
    Prodi string `json:"prodi"`
    Fakultas string `json:"fakultas"`
}

var mahasiswa = []Mahasiswa{
    {ID: 1, Nama: "Budi", Prodi: "Teknik Informatika", Fakultas: "Teknik"},
    {ID: 2, Nama: "Andi", Prodi: "Sistem Informasi", Fakultas: "Ilmu Komputer"},
    {ID: 3, Nama: "Caca", Prodi: "Teknik Industri", Fakultas: "Teknik"},
    {ID: 4, Nama: "Dedi", Prodi: "Pendidikan Sosiologi", Fakultas: "Ilmu Sosial"},
}


func getMahasiswa(c *gin.Context) {
    c.JSON(http.StatusOK, mahasiswa)
}

func main() {
    router := gin.Default()
    router.GET("/mahasiswa", getMahasiswa)

    router.Run(":8080")
}