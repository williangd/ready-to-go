package main

import "fmt"

func main() {
	simples()
	argumento("Isso é um exemplo de argumento.")
	argumento("Isso é outro exemplo de argumento.")
	fmt.Println("Uma função com retorno que entrega o \"resultado\" da sua funcionalidade", retorno())
	fmt.Println(múltiplos(8, 7))
	fmt.Println(variádica(5, 4, 7, 8, 9, 10, -75, 654))
	fmt.Println("Ao invés de passar vários argumentos individualmente, tambem podemos jogar uma slice numa função variádica...")
	números := []int{12, 23, 34, 45, 56, 67, 78, 89, 90}
	fmt.Println(variádica(números...))
	func(s string) {
		fmt.Println("Tambem podemos ter funções", s, "(descartáveis).")
	}("anônimas")
	variável := func() {
		fmt.Println("E salvar funções como variáveis.")
	}
	variável()
	função := retornaFunção()
	função()
	euReceboFunçãoComoArgumento(euFaloArgumento)

}

func simples() {
	fmt.Println("Uma função simples executa uma funcionalidade declarada em outro local.")
}

func argumento(s string) {
	fmt.Println("Uma função com argumentos recebe um valor e trabalha em cima do mesmo.", s)
}

func retorno() string {
	return "Isso é um exemplo de retorno."
}

func múltiplos(a, b int) (string, int) {
	return "Funções podem ter argumentos e retornos na mesma função. E não só isso, mas podem ter vários de cada um. E multiplicando aqueles números temos:", a * b
}

func variádica(args ...int) (string, int) {
	soma := 0
	for _, x := range args {
		soma += x
	}
	return "Tambem temos funções que recebem argumentos \"infinitos\", e usamos aqueles três pontinhos ali. Um argumento variádico deve ser sempre o último da lista, caso tenhamos múltiplos argumentos. E a soma de todos esses valores é:", soma
}

func retornaFunção() func() {
	return func() {
		fmt.Println("E olha só, podemos ter funções como valor de retorno.")
	}
}

func euReceboFunçãoComoArgumento(f func() string) {
	fmt.Println("E a essa altura você já deve esperar, corretamente, que podemos passar funções como", f())
}

func euFaloArgumento() string {
	return "argumento."
}

// - Função simples
// - Função que aceita um argumento
// - Função com retorno
// - Função com múltiplos argumentos e múltiplos retornos
// - Função variádica
// - Passando um slice como argumento para uma função variádica
// - Função anônima
// - Função como expressão: f := func(p params){ code }
// - Retornando uma função
// 	- func f() func() int { [...]; return func() int{ return [int] } }
// - Callbacks: passando funções como argumentos
// - Closure
