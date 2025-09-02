// por Fernando Dotti - fldotti.github.io - PUCRS - Escola Politécnica
// EXERCÍCIO:  dado o programa abaixo
//    1) quantos processos concorrentes são gerados ? são criadas 41 goroutines, umas vez que uma go routine é executada dentro de um for, as quais ficam
//		rodando infinitamente

//    2) execute e observe: que se pode supor sobre a velocidade relativa dos mesmos ?
//		teoricamente iguais 

// OBSERVACAO:o sleep no método main serve para este nao acabar, o que acabaria todos processos em execucao.
//     mais adiante veremos outras formas de sincronizar isto

package main

import (
	"fmt"// print
	//"time"
)

var N int = 40 

func funcaoA(id int, s string) { 
	for {
		fmt.Println(s, id) 
	}
}

func geraNespacos(n int) string {
	s := "  "
	for j := 0; j < n; j++ {
		s = s + "   "
	}
	return s
}

func main() {
	for i := 0; i < N; i++ {
		go funcaoA(i, geraNespacos(i))
	}
	for true {
		
	}
	
}
