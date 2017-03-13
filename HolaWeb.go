package main

import (
	"encoding/json"
	"net/http"
)

// Usuario de la pagina.
type Usuario struct {
	Nombre     string
	Contraseña string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, arg2 *http.Request) {
		usuarioAndres := Usuario{"andrés", "contraseña123"}
		json.NewEncoder(w).Encode(usuarioAndres)
	})
	http.ListenAndServe(":8002", nil)
}
