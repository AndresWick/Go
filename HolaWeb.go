package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/user"
)

// Usuario de la pagina.
type Usuario struct {
	Nombre     string
	Contraseña string
}

func mostrarUsuarios(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(BloqueIndices)
}

// BloqueIndices a usuarios
var BloqueIndices = make(map[string]Usuario)

func main() {
	BloqueIndices = make(map[string]Usuario)
	u, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}
	usuarioAndres := Usuario{u.Name, u.Username}
	usuarioPedro := Usuario{"Pedro", "698"}
	usuarioRamon := Usuario{"Ramon", "456"}
	usuarioJulia := Usuario{"Julia", "123456"}

	fmt.Println("Usuario actual:")
	fmt.Println(usuarioAndres.Nombre)

	BloqueIndices["001"] = (usuarioAndres)
	BloqueIndices["002"] = (usuarioPedro)
	BloqueIndices["003"] = (usuarioRamon)
	BloqueIndices["004"] = (usuarioJulia)

	http.HandleFunc("/usuarios", mostrarUsuarios)

	http.HandleFunc("/", func(w http.ResponseWriter, arg2 *http.Request) {
		fmt.Fprintf(w, "<h1> PAGINA PRINCIPAL </h1>")
	})

	http.ListenAndServe(":8002", nil)

	b, err := ioutil.ReadFile("matriz.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", b)
	fmt.Printf("%v\n", BloqueIndices["001"])
}
