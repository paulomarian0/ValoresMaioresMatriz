package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
 Leia uma matriz 4 x 4, conte e escreva quantos e quais valores maiores que 10 ela possui.

 */


func main() {

	var matriz[5][5] int

	slice:= [] int{}

	contador:= 0;


	s1:= rand.NewSource(time.Now().UnixNano())
	r1:= rand.New(s1)

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			matriz[i][j] = r1.Intn(20)

			if(matriz[i][j] > 10){
				slice = append(slice ,matriz[i][j])
				contador++
			}

		}

	}

	fmt.Printf("Possui %d numeros maiores que 10\n", contador)
	fmt.Println("Esses numeros sao:", slice[:contador])


}
