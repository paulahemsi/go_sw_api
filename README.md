# go_sw_api
go studies

*in_process*

studying go, api and databases

# Go

## Instalção

```s
$ sudo apt-get update
$ sudo apt-get upgrade
$ sudo apt-get install golang-go
```
Um ponto importante ao se trabalhar com go é que precisamos deixar a pasta dos nossos projetos dentro do diretório `src` que está dentro do diretório `go`, e configurar dentro do `.bashrc` as seguintes variáveis de ambiente:

```s
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

## Documentação

Dentro da [página oficial](golang.org) do **Go** é super simples navegar pelos pacotes e documentações oficiais 

## Tipos

* int
* float
* string
* boolean

## Estruturas_de_controle

* IF
* For (possui uma maneira de ser usado em que se torna um *while* como se usa em C)
* Switch

## Funções

Muito semelhantes ao C, possuem a seguinte sintaxe:

```go
func nomeDaFunc (parametro tipoDoParametro) retorno {
  corpoDaFunc
  return retorno
}
```

sendo possível, assim como no C, o uso de variadic functions

```go
func nomeDaFunc (parametro ...tipoDoParametro) retorno {
  corpoDaFunc
  return retorno
}
```

## Arrays

Para definir um array a sintaxe é a seguinte:

```go
var nome_do_array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
```

possuindo também opções mais sintéticas de declaração:

```go
nome_do_array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
```

ou ainda

```go
nome_do_array := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
```

em que a linguagem é capaz de deduzir sozinha as informações faltantes

## Slices

Para resolver um dos grandes problemas do array que, como em C, precisam de um tamanho fico ao serem declarados, o **Go** conta Slices e com a função Make, que faz as vezes de um malloc, separando o espaço de memória necessário

```go
var nome_do_slice []int
```

É bastante simples com **Go** incluirmos, excluirmos, copiarmos ou modificarmos o tamanho de um slice

## Maps

O map tem uma estrutura semelhante a structs ou objetos, tendo pares de valores relacionados, como por exemplo números e strings

```go
var agenda map[string]int
```

## Structs

Bastante semelhantes as structs de C, onde podemos criar nosso tipo *Struct* customizado com quaisquer tipo de dados dentro delas

```go
type NomeDaStruct struct {
    idade     int
    nome      string
    inscrito  bool
} 
```

## Ponteiros

Ponteiros estão presentes em Go e possuem exatamente a mesma sintaxe e uso do C, sendo super útil para mover grandes estruturas de dados pelo nosso programa e passar variáveis por valor ou por referência conforme for conveniente

## comandos

O go possui uma série de possíveis comandos a serem executados no terminal como o que compila os arquivos necessários recebendo apenas o arquivo em que se encontra a função main (entrada do programa, como em C)

```bash
go run main.go
```

Ou este comando muito útil que formata determinado arquivo automaticamente

```bash
go fmt <nome do arquivo>
```

Ou ainda o seu próprio manual embutido para acessar todas as funções dos pacotes 

```bash
go doc <nome do pacote>.<nome da função>
```

Tem também seu próprio *get* para baixar novos pacotes

```bash
go get golang.org/caminho_para_o_pacote
```

Dentro muitos outros

## Pacotes

A inclusão das bibliotecas, chamadas pacotes, é feita de maneira super simples como segunda informação dentro de todo arquivo .go

```go
import "fmt"
```

ou, para múltiplos pacotes

```go
import (
    "fmt"
    "net/http"
    "errors"
)
```

A única informação que vem antes de importar os pacotes é o pacote em si a que aquele arquivo pertence

```go
package main
```

O main, por exemplo, é obrigatório de estar presente, coerentemente no próprio arquivo em que teremos a função de entrada `main`.

Essa classificação de todo arquivo em um pacote torna muito simples o uso de pacotes costumizados, basta incluir o caminho para ele nos imports de um arquivo e está feito! 
Ah, importante lembrar que o caminho deve sempre ser relativo a pasta `go/src`

```go
import {
    "fmt"
    "caminho/para/meu_pacote/local"
}
```

Além de incluir o pacote local nos imports de um arquivo, para tornar as funções de um pacote visíveis para outros, é necessário que elas tenham **a primeira letra maiúscula** e **um comentário na linha imediatmanete acima**

```go
//comentario sobre a função a seguir que está sendo exportada
func NomeDaFunc (parametro tipoDoParametro) retorno {
  corpoDaFunc
  return retorno
}
```

Feito isso, para acessar a função de outros arquivos, basta chamar o nome do pacote *ponto* o nome da função

```go
nomedopacote.NomeDaFunc
```

## Testes

**Go** possui uma maneira super simples e intuitiva de realizar testes unitários, basta criar um arquivo que tenha o package "testing" e criar uma função com o nome `Test` + `nome da função que quer testar` e dentro dela definir os testes que deseja fazer, sempre indicando qual seria o valor esperado e chamando a função com os parâmetros que deveriam supostamente dar aquele resultado. A biblioteca de testes lida com o processo e quando o teste é executado o terminal nos indica `PASS` ou `FAIL` e quaisquer outras mensagens costumizadas que programamos.

## Gestão_de_Erros

O Go possui um pensamento muito interessante com relação ao gerenciamento de erros, sempre prevendo os possíveis erros durante a execução do programa e já definindo como lidar com eles. 

É bastante comum a seguinte sintaxe depois de cada instrução passível de dar errado

```go
if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("failed to request 42_cursus")
	}
```

## Materiais_de_estudo

* [introdução ao mySQL do guanabara](https://www.youtube.com/watch?v=Ofktsne-utM)
* [package sql do go](https://golang.org/pkg/database/sql/)
* [Different approaches to HTTP routing in Go](https://benhoyt.com/writings/go-routing/)
* [Go By Example json](https://gobyexample.com/json)
* [Panic](https://www.digitalocean.com/community/tutorials/handling-panics-in-go-pt)
* [Rest Api](https://tutorialedge.net/software-eng/what-is-a-rest-api/)


