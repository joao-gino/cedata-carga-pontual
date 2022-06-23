package authorize

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Retorno Retorno `json:"retorno"`
}

type Retorno struct {
	Status          int      `json:"status"`
	StatusDescricao string   `json:"statusDescricao"`
	Mensagem        string   `json:"mensagem"`
	Conteudo        Conteudo `json:"conteudo"`
}

type Conteudo struct {
	Token string `json:"token"`
}

func GenerateAuthorizationToken() string {
	jsonData := map[string]string{"tokenCombinado": os.Getenv("tokenCombinado"), "identificadorEmpresa": os.Getenv("identificadorEmpresa"), "loginUsuario": os.Getenv("loginUsuario"), "loginUsuarioSenha": os.Getenv("loginUsuarioSenha")}
	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post("https://raizdaserra.cargapontual.com/sistema/Integracao/bi_grupocesari_1/servidor/v100/autenticacao/token/", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return "The HTTP request failed with error"
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	post := Response{}
	err = json.Unmarshal(data, &post)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return ""
	}

	return string(post.Retorno.Conteudo.Token)
}
