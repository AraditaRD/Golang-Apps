package main

import (
	
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Struktur data untuk tugas (To-Do)
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// Data sementara (database tiruan)
var todos = []Todo{
	{ID: 1, Title: "Belajar Golang", Status: "pending"},
	{ID: 2, Title: "Bikin API dengan Gin", Status: "done"},
}

// GET: Ambil semua tugas
func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

// POST: Tambah tugas baru
func addTodo(c *gin.Context) {
	var newTodo Todo

	// Bind JSON ke struct Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
		return
	}

	// Set ID baru (auto-increment sederhana)
	newTodo.ID = len(todos) + 1
	todos = append(todos, newTodo)

	c.JSON(http.StatusCreated, newTodo)
}

// PUT: Perbarui tugas berdasarkan ID dari path parameter
func updateTodo(c *gin.Context) {
	idParam := c.Param("id") // Ambil ID dari URL
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	var updatedTodo Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
		return
	}

	// Cari dan update tugas
	for i, t := range todos {
		if t.ID == id {
			todos[i].Title = updatedTodo.Title
			todos[i].Status = updatedTodo.Status
			c.JSON(http.StatusOK, todos[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Tugas tidak ditemukan"})
}

// DELETE: Hapus tugas berdasarkan ID dari path parameter
func deleteTodo(c *gin.Context) {
	idParam := c.Param("id") // Ambil ID dari URL
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Tugas berhasil dihapus"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Tugas tidak ditemukan"})
}

// Main function untuk menjalankan server
func main() {
	r := gin.Default()
	r.Use(gin.Logger()) // Logging middleware untuk debugging

	// Routing
	r.GET("/todos", getTodos)        // Lihat semua tugas
	r.POST("/todos", addTodo)        // Tambah tugas baru
	r.PUT("/todos/:id", updateTodo)  // Perbarui tugas berdasarkan ID
	r.DELETE("/todos/:id", deleteTodo) // Hapus tugas berdasarkan ID

	// Jalankan server di port 8080
	r.Run(":8080")
}
