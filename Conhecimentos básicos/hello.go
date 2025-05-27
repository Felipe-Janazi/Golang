// Para executarmos esse programa precisamos compilar e ai executar, mas através do comando "go run hello.go" ele já compila e executa nosso código

// Todo inicio de código informamos que este é o package Main
package main

// Fazemos a instalação dos pacotes que desejamos utilizar
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

func main() {
	exibeIntroducao()

	// Essa é a criação de um for infinito, ele só vai sair do for quando ele cair na opção 0
	for {
		exibeMenu()

		// Uma declaração mais limpa de variavel, pois através do = a variável vai analisar e saber qual seu tipo e através do : ela sabe que ela vai ser uma variável.
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLogs()
		case 0:
			fmt.Printf("Saindo do programa...")
			// Encerrando o for saindo dele e enviando para o sistema o número 0 que significa que saiu do programa
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
		}
	}
}

func exibeIntroducao() {
	nome := "Douglas"
	versao := 1.1

	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir os logs")
	fmt.Println("0- Sair do programa")

}

// Informamos se ela devolve algo depois do nome dela
func leComando() int {

	var comandoLido int

	// Criação de um scaner em go
	fmt.Scan(&comandoLido)

	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {

	fmt.Println("Monitorando...")

	// Criação de uma lista, informando o nome, usando o := e qual o tipo da lista com []string
	// sites := []string{"https://www.alura.com.br", "https://random-status-code.herokuapp.com/", "https://www.caelum.com.br"}

	sites := lerSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {

		// Falamos que com o I vamos pegar a posição e com o sites pegamos o conteúdo da posição, falando que o range vai ser até o final do slice sites
		for i, site := range sites {

			fmt.Println("Estou passando na posição", i)
			fmt.Println("O site que estou é o", site)

			testaSite(site)
			fmt.Println("")

		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

}

// Informamos o nome da variável que vai receber e o seu tipo
func testaSite(site string) {

	// A funcção http.get possui duas respostas, o resp significa o código que foi devolvido, se foi 200 o site está rodando corretamente, o segundo parâmetro é de erro, onde quando não precisamos usar colocamos apenas um _ que significa que não queremos aquela resposta.

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas, Status Code:", resp)
		registraLog(site, false)
	}
}

func lerSitesDoArquivo() []string {

	// Criação de uma lista
	var sites []string

	// Abrindo um arquivo dentro da mesma pasta deste, mas não consegue exibir o conteúdo
	arquivo, err := os.Open("sites.txt")

	// Abre o arquivo, lê o mesmo em binário mas se colocarmos string(NomeVariavél) conseguimos visualizar todo o conteúdo dentro
	// arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	// Usamos uma nova biblioteca, ela vai ler e jogar tudo dentro de um arquivo e depois, dentro deste arquivo leitor nós informamos até onde ele vai ler, neste caso até a quebra de linha, assim ele lê linha por linha.
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')

		// Usamos a biblioteca strings para tirar o quebra linha e espaços do final da linha
		linha = strings.TrimSpace(linha)

		// Colocando a linha dentro do slice
		sites = append(sites, linha)

		// Informamos que se ele chegar no fim do arquivo ele vai quebrar o FOR
		if err == io.EOF {
			break
		}

	}

	// Após usamos o arquivo que abrimos nós fechamos o mesmo, assim sempre deixando ele fechado após a utilização
	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

	//  Aqui nos abrimos um arquivo com OpenFile e damos a permissão de RDWR, caso o arquivo não exista ele é criado e para que seja escrito e não sobrescrito damos a permissão de APPEND e falamos que é com o permissionamento de 666
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	// Aqui estamos escrevendo no arquivo, usando o formato de horas e data que está na documentação do GO, concatenando tudo e como o status era um boolean usamos a biblioteca que converte valores pare strings
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online:" + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	// Agora usamos a função para exibir todo o conteúdo que tem dentro do nosso arquivo de logs 
	arquivo, err := ioutil.ReadFile("log.txt")
	
	if err != nil {
		fmt.Println(err)
	}

	//  Aqui transformamos o arquvio que vamos ver em uma string 
	fmt.Println (string(arquivo))
}
