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

func mostrarUsuario(w http.ResponseWriter, r *http.Request) {
	usuario1 := Usuario{"usuario", "usuario1"}
	json.NewDecoder(r)
	json.NewEncoder(w).Encode(usuario1)
}

func main() {
	http.HandleFunc("/usuario", mostrarUsuario)

	http.HandleFunc("/", func(w http.ResponseWriter, arg2 *http.Request) {
		usuarioAndres := Usuario{"andrés", "contraseña123"}
		json.NewEncoder(w).Encode(usuarioAndres)
	})
	http.ListenAndServe(":8002", nil)
}
