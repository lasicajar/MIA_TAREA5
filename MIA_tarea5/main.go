package main

import (
	"fmt"
	"net/http"
	"time"
)

func apiTarea5(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("[%s] Petición recibida desde AWS EC2 %s a %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.URL.Path)

	// Responder con JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"Nombre": "Lusvin Alexander Sicaja Ramirez - 201602630", "timestamp": "%s"}`, time.Now().Format(time.RFC3339))
}

func main() {
	http.HandleFunc("/", apiTarea5)
	fmt.Println("🚀 API escuchando en http://localhost:8080")

	// Iniciar servidor
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %v\n", err)
	}
}
