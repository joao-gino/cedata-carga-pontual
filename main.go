package main

import (
	"cedata-carga-pontual/authorize"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {

	fmt.Println("Generating Authorize API Token...")

	authorizationToken := authorize.GenerateAuthorizationToken()

	fmt.Println("Authorization Token: " + authorizationToken)

	fmt.Println("Calling API...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://raizdaserra.cargapontual.com/sistema/Integracao/bi_grupocesari_1/servidor/v100/agendamento/consumoBI/", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("tokenCombinado", authorizationToken)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(string(bodyBytes))
}
