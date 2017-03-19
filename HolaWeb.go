package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/user"
)

// Usuario de la pagina.
type Usuario struct {
	Nombre     string
	Contraseña string
}

func mostrarUsuario(w http.ResponseWriter, r *http.Request) {

	usuario1 := Usuario{"user", "000"}
	//json.NewDecoder(r)
	json.NewEncoder(w).Encode(usuario1)
}

func main() {
	u, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}
	usuarioAndres := Usuario{u.Name, u.Username}
	fmt.Println("Usuario actual:")
	fmt.Println(usuarioAndres.Nombre)

	http.HandleFunc("/usuario", mostrarUsuario)

	http.HandleFunc("/", func(w http.ResponseWriter, arg2 *http.Request) {
		//usuarioAndres := Usuario{"andrés", "contraseña123"}
		json.NewEncoder(w).Encode(usuarioAndres)
	})
	http.ListenAndServe(":8002", nil)

}
