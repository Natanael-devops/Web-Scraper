# Web-Scraper
Aplicação criada fazer scraping do site workana.com

<h3> Descrição </h3>
<p> A aplicação recebe o link para scraping e em seguida visualiza as cem primeiras páginas do site em questão. Utilizo o número cem pois o site da Workana
  fica fora do ar à partir de "page=101"<br>

Depois de visitar o site, a aplicação calcula a média de valor por projeto do site e a densidade de propostas por projeto, printando tudo no terminal.<br>
  A aplicação possui apenas quatro funções em seu código:<br>
  <ul>
    <li>func Main: responsável por fazer o scraping do site e chamar as outras funções;</li>
    <li>func EscreveJSON: responsável em escrever um novo arquivo com todos os dados;</li>
    <li>func SelecionaNumeros: responsável em encontrar os números contidos em strings para transformar em float64 e assim melhorar os dados.</li>
    <li>func CalculaMedia: responsável por retornar a média numérica de qualquer slice dado como parâmetro.</li>
    <li>func TratandoDados: função que ajuda a encontrar as expressões regulares (regex).</li>
</p>

  <h3>Stacks</h3>
  <p>Desenvolvido utilizando:
<ul>
  <li>GO</li>
  <li>Colly</li>
  <li>POO</li>
  </ul></p>
  
  <h3>Abrir e rodar o projeto</h3>
  <p>É necessário possuir Go(golang) instalado no computador.<br>
  Para instalar o Go, acesse o site oficial: https://go.dev/doc/install <br>
  Eu utilizei o Postman para testar a aplicação, recomendo sua uilização.<br>
  Após baixar o projeto, para abrir e rodar o projeto abra a pasta do arquivo no terminal e execute:<br>
  <code>go run main.go</code><br>
  
