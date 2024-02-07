package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		res, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error ao fazer chamada http %s", err)
		}

		if res.StatusCode != 200 {
			fmt.Printf("Status retorno - %d\n", res.StatusCode)
			return
		}

		defer res.Body.Close()

		resParsed, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error ao parsear response %s", err)
			return
		}

		var vc ViaCep

		err = json.Unmarshal(resParsed, &vc)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error ao parsear resParsed para Struct %s", err)
			return
		}

		file, err := os.Create("cidades_buscadas.txt")
		defer file.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error ao criar arquivo %s", err)
			return
		}

		_, err = file.WriteString(vc.Localidade)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error ao adicionar cidade ao arquivo %s", err)
			return
		}

		fmt.Println(vc.Localidade)
	}
}
