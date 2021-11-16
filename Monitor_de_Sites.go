package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const qtdVezes = 2
const tempo = 5

func main() {

	introducao()

	for {
		menuInicial()

		comando := solicitaMenu()

		switch comando {
		case 0:
			fmt.Printf("Saindo do programa")
			os.Exit(0)
		case 1:
			iniciaMonitoramento()
		case 2:
			fmt.Println("Gerando Logs...")
		default:
			os.Exit(-1)
		}
	}
}

func introducao() {
	versao := 1.0
	fmt.Println("Ola usuario! Esse programa está na versao ", versao)
	fmt.Println("-------------------//----------------------")
}

func menuInicial() {
	fmt.Println("1- Iniciar monitoriamento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair")
}

func solicitaMenu() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://random-status-code.herokuapp.com/", "https://www.youtube.com/", "https://www.google.com.br/", "https://golang.org/"}

	for i := 0; i < qtdVezes; i++ {
		for i, site := range sites {
			fmt.Println("pos ", i, ":", site)
			statusSites(site)
			fmt.Println("")
		}
		time.Sleep(tempo * time.Second)
	}
}

func statusSites(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("O site Esta funcionando perfeitamente")
	} else {
		fmt.Println("O site Esta com problemas")
	}
}
