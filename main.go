package main

import (
	"fmt"
	"sync"
)

func main() {
	const limit = 10

	// Criando os canais
	pingChan := make(chan bool)
	pongChan := make(chan bool)

	// WaitGroup para esperar o término das goroutines
	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine para exibir "Ping"
	go func() {
		defer wg.Done() // Marca como concluída
		for i := 0; i < limit; i++ {
			<-pingChan       // Espera o sinal para exibir "Ping"
			fmt.Println("Ping") // Exibe "Ping"
			pongChan <- true // Envia o sinal para "Pong"
		}
	}()

	// Goroutine para exibir "Pong"
	go func() {
		defer wg.Done() // Marca como concluída
		for i := 0; i < limit; i++ {
			<-pongChan       // Espera o sinal para exibir "Pong"
			fmt.Println("Pong") // Exibe "Pong"
			pingChan <- true // Envia o sinal para "Ping"
		}
	}()

	// Inicia a troca enviando o primeiro sinal para o canal "Ping"
	pingChan <- true

	// Aguarda as goroutines terminarem
	wg.Wait()
}
