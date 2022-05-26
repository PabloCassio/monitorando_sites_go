package main // indica que esse é o pacote principal da execução

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

const monitoramentos = 3
const delay = 5

// para usar a função print ln precisa importar o fmt. Toda função externa começa fom a letra maiuscula.

func main() {

	exibeIntroducao()
	for {
		exibeMenu()

		comando := leComando()

		switch comando { // só executa cada caso, em ordem. Não precisa colocar break
		case 1:

			iniciarMonitoramento()
		case 2:
			fmt.Println("Logs...")
			imprimelogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)

		default:
			fmt.Println("Não conheço este comando.")
			os.Exit(-1)

		}

	}
	/*
		if comando == 1 { // sempre necessita ser expressões, nunca teste de variável sem condicional. Retorno sempre necessita ser bool.
			fmt.Println("Monitorando...")

		} else if comando == 2 {
			fmt.Println("Logs...")

		} else if comando == 0 {
			fmt.Println("Saindo do programa...")
		} else {
			fmt.Println("Não conheço este comando.")

		}
	*/
}

func exibeIntroducao() {
	// começa com valor 0 (int), vazio para string, todas as variaveis tem que ser utilizadas, não precio declarar sempre o tipode variável.
	var nome string = "Pablo"
	versao := 1.1 // float é recomendável tipar. Coloca a maior quando não especificada. := relacionado com o var.
	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Esse programa está na versão", versao) // go build (linguagem compilada) go run (compila e executa)

}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido) // & endereço da variavel, %d para receber inteiros.

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir os logs")
	fmt.Println("0 - Sair do programa")

}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// slices são estruturas de tamanho dinâmico.
	/* var sites [4]string // array tem tamanho fixo
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br"
	*/
	// sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}
	sites := leSitesDoArquivo()
	for i := 0; i < monitoramentos; i++ {

		for i, site := range sites {
			fmt.Println("Testando site", i+1, ":", site)

			testaSite(site)

		}
		time.Sleep(delay * time.Second)
		fmt.Println("")

	}

	fmt.Println("")
	// se não me interesso em um dos retornos, uso o _
	//fmt.Println(resp)

}
func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)

	} else {
		fmt.Println("Site:", site, "está com problema. Status Code: ", resp.StatusCode)
		registraLog(site, false)
	}

}

// funções com mais de um retorno precisam ser especificados entre parenteses ()
//arquivo, err := ioutil.ReadFile("sites.txt") //array de bites (ler tudo)

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n') //informa um bite (aspas simples)
		linha = strings.TrimSpace(linha)      //retira espaços do fim da linha
		sites = append(sites, linha)

		if err == io.EOF {
			break

		}
	}
	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) //flags:
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()

}

func imprimelogs() {
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	fmt.Println(string(arquivo))

}
