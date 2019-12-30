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

const monitoramento = 3
const delay = 5

func main() {
	introduction()

	for {
		showOptions()
		comando := readCommand()

		switch comando {
		case 1:
			initMonitoration()
		case 2:
			showLogs()
		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Não Conheço este comando")
			os.Exit(-1)
		}
	}
}

func introduction() {
	nome := "Gabriel"
	versao := 1.1

	fmt.Println("Olá", nome)
	fmt.Println("Este Programa está na versão", versao)
}

func showOptions() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func readCommand() int {

	var comando int
	fmt.Scan(&comando)

	return comando
}

func initMonitoration() {
	fmt.Println("Monitorando...")
	sites := readSitesOfFile()

	for i := 0; i < monitoramento; i++ {

		for _, site := range sites {
			testSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")

	}

	fmt.Println("")

}

func testSite(site string) {
	res, err := http.Get(site)

	if err != nil {
		fmt.Println("Site:", site, " esta com problemas, Status Code:", res.StatusCode)
	}

	if res.StatusCode == 200 {
		fmt.Println("Site: ", site, " foi carregado com sucesso")
		registerLog(site, true)
	} else {
		fmt.Println("Site:", site, " esta com problemas, Status Code:", res.StatusCode)
		registerLog(site, true)
	}
}

func readSitesOfFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	reader := bufio.NewReader(file)

	for {
		row, err := reader.ReadString('\n')
		row = strings.TrimSpace(row)
		sites = append(sites, row)
		if err == io.EOF {
			break
		}
	}

	file.Close()
	return sites
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func showLogs() {
	fmt.Println("Exibindo Logs...")

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(file))
}
