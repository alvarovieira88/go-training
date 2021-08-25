package main

import (
	"fmt"
	"log"
	"time"
)

// FIXME criar o objeto Pessoa
// irmaos
// pais
// pets

// criar o metodo criaEU

// mauricio -- irmao > tiago
// mauricio -- pet  > dudu
// mauricio -- pais > magda

func main() {
	type pessoa struct {
		nome            string
		genero          string
		peso            float64
		altura          float64
		dataNascimento  time.Time
		nacionalidade   string // brasileiro
		naturalidade    string // joinville
		numeroDocumento string
	}

	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		log.Fatalf("ERRRO NA GERACAO DO LOCATION")
	}
	alvaro := pessoa{nome: "Alvaro", genero: "masculino", peso: 70.00, altura: 1.75,
		dataNascimento: time.Date(1991, time.February, 22, 01, 01, 01, 00, location),
		nacionalidade: "brasileiro", naturalidade: "Juiz de Fora - MG", numeroDocumento: "999.999.999-99",
	}
	fmt.Println("Nome =", alvaro.nome)
	fmt.Println("Gênero =", alvaro.genero)
	fmt.Println("Peso =", alvaro.peso, "Kg")
	fmt.Println("Altura =", alvaro.altura, "cm")
	fmt.Println(alvaro.dataNascimento)
	fmt.Println("Nacionalidade =", alvaro.nacionalidade)
	fmt.Println("Naturalidade =", alvaro.naturalidade)
	fmt.Println("Documento =", alvaro.numeroDocumento)
}
