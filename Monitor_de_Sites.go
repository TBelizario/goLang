package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
			mostrarLogs()
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
	siteTxt := "sites.txt"
	sites := leArquivos(siteTxt)

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
		log(site, true)
	} else {
		fmt.Println("O site Esta com problemas")
		log(site, false)
	}
}

func leArquivos(site string) []string {
	var sites []string

	arquivo, err := os.Open(site)
	if err != nil {
		fmt.Println("Erro detectado: ", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func log(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("Falha ao ler ou para criar o arquivo")
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - Online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func mostrarLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo")
	}
	println(string(arquivo))
}
