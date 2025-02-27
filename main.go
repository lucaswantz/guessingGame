package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da adivinhação")
	fmt.Println("Um número aleatório será sorteado. Tente acertar. O número é um inteiro entre 0 e 100")

	x := rand.Int64N(101)

	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		fmt.Println("Qual é o seu chute?")

		scanner.Scan()

		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("O chute deve ser um número inteiro")
			return
		}

		switch {
		case chuteInt < x:
			fmt.Println("Você errou. O número sorteado é maior que", chuteInt)
		case chuteInt > x:
			fmt.Println("Você errou. O número sorteado é menor que", chuteInt)
		case chuteInt == x:
			fmt.Printf(
				"Parabéns! Você acertou e o número era: %d\n"+
					"Você acretou em %d tentativas.\n"+
					"Esses foram os números escolhidos: %v\n", x, i+1, chutes[:i])

			return
		}

		chutes[i] = chuteInt
	}

	fmt.Printf(
		"Infelizmente você não acertou o número, o número sorteado é %d\n"+
			"Você teve 10 tentativas.\n"+
			"Esses foram os números escolhidos: %v\n", x, chutes)
}
