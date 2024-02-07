package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s" validate:"gt=0"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 333333333}

	res, err := json.Marshal(conta)

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
	fmt.Println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	jsonPuro := []byte(`{"n":1,"s":30000000}`)
	var contaX Conta

	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		panic(err)
	}

	fmt.Println(contaX)
}
