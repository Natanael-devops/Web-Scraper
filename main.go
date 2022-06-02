package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"

	"github.com/gocolly/colly"
)

type Projeto struct {
	Nome      string `json:"nome"`
	Publicado string `json:"publicado"`
	Propostas string `json:"propostas"`
	Valor     string `json:"valor"`
	Descricao string `json:"descricao"`
	Categoria string `json:"categoria"`
	Autor     string `json:"autor"`
}

func main() {
	todosDados := make([]Projeto, 0)
	todosValores := []string{}
	todasPropostas := []string{}
	projeto := Projeto{}

	collector := colly.NewCollector(
		colly.AllowedDomains("workana.com", "www.workana.com"),
	)

	collector.OnHTML("div.project-item", func(elemento *colly.HTMLElement) {
		elemento.ForEach("h2.h3", func(i int, h *colly.HTMLElement) {
			projeto.Nome = h.ChildAttr("span", "title") //falta descobrir isso aqui

		})

		elemento.ForEach("div.expander", func(i int, h *colly.HTMLElement) {
			textoBruto := h.Text
			descricao := TratandoDados(`(.*)Categoria`, textoBruto)
			projeto.Descricao = descricao
			categoria := TratandoDados(`(Categoria: (.*)Subcategoria)`, textoBruto)
			projeto.Categoria = categoria

		})

		elemento.ForEach("span.values", func(i int, h *colly.HTMLElement) {
			projeto.Valor = h.Text
			todosValores = append(todosValores, h.Text)
		})
		elemento.ForEach("span.date", func(i int, h *colly.HTMLElement) {
			projeto.Publicado = h.Attr("title")

		})

		elemento.ForEach(".project-main-details", func(i int, h *colly.HTMLElement) {
			propostaBruta := h.ChildText("span.bids")
			proposta := TratandoDados(`(Propostas: (.*))`, propostaBruta)
			projeto.Propostas = proposta
			todasPropostas = append(todasPropostas, proposta)

		})

		elemento.ForEach(".project-author", func(i int, h *colly.HTMLElement) {
			projeto.Autor = h.ChildText("a")

		})

		todosDados = append(todosDados, projeto)
	})
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visitando o site: ", request.URL.String())
	})

	for i := 1; i < 101; i++ {
		collector.Visit("https://www.workana.com/jobs?language=pt&page=" + strconv.Itoa(i))
	}

	EscreveJSON(todosDados)

	fmt.Println("Calculando os valores dos projetos: ")
	valores := SelecionaNumeros(todosValores)
	CalculaMedia(valores)

	fmt.Println("Calculando quantidades de propostas dos projetos")
	projetos := SelecionaNumeros(todasPropostas)
	CalculaMedia(projetos)

}

func EscreveJSON(dados []Projeto) {
	arquivo, err := json.MarshalIndent(dados, "", " ")
	if err != nil {
		log.Println("Não foi possível criar o arquivo json")
		return
	}

	_ = ioutil.WriteFile("projetos-workana.json", arquivo, 0644)
}

func SelecionaNumeros(dados []string) []string {
	keys := []string{}

	re := regexp.MustCompile("[0-9.]+")

	for _, runa := range dados {
		numeros := re.FindAllString(runa, -1)
		keys = append(keys, numeros...)

	}

	valoresFinais := []string{}
	for _, decimal := range keys {
		ree := regexp.MustCompile("[.]+")
		inteiros := ree.ReplaceAllString(decimal, "")
		valoresFinais = append(valoresFinais, inteiros)
	}
	return valoresFinais
}

func CalculaMedia(valor []string) {
	var valores = []float64{}
	for _, i := range valor {
		j, err := strconv.ParseFloat(i, 64)
		if err != nil {
			panic(err)
		}
		valores = append(valores, j)
	}
	var soma float64 = 0
	for i := 0; i < len(valores); i++ {
		soma = valores[i] + float64(soma)
	}
	var divide = float64(soma)
	fmt.Println("somando todos os dados, o resultado é: ", divide)
	fmt.Println("O tamanho do array é de: ", len(valor))
	var denominador = float64(len(valor))
	divide = divide / denominador
	fmt.Println("a média total dos valores dos dados da categoria é de: ", divide)
}

func TratandoDados(regex string, variavel string) string {
	re := regexp.MustCompile(regex)
	var retorno string
	var categoria = re.FindStringSubmatch(variavel)

	var minhaString = categoria[len(categoria)-1]
	if len(minhaString) >= 1 {
		retorno = minhaString
	}

	return retorno
}
