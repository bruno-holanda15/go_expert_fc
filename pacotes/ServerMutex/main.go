package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
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
	mux := http.NewServeMux()
	mux.HandleFunc("/", BuscaCepHandler)


	muxOther := http.NewServeMux()
	muxOther.HandleFunc("/check", CheckOtherServer)
	
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		http.ListenAndServe(":8080", mux)
	}()

	go func() {
		defer wg.Done()
		http.ListenAndServe(":8089", muxOther)
	}()

	wg.Wait()

	fmt.Println("service has shut down")
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	cep, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// result, err := json.Marshal(cep)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(result)

	json.NewEncoder(w).Encode(cep)
	// w.Write([]byte("Olá amigo!"))
}

func BuscaCep(cep string) (*ViaCep, error) {
	res, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var c ViaCep

	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func CheckOtherServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Top", "Brunin")
	w.WriteHeader(http.StatusNoContent)
}