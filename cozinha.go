package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	numCozinheiros = 5 // Representa o tamanho do buffer do canal (capacidade do semáforo)
	numPedidos     = 100
)

func doThing(id int, pedido int) {
	fmt.Printf("Cozinheiro %d preparando pedido %d\n", id, pedido)
	time.Sleep(time.Duration(rand.Float64()) * time.Second) // Simula o tempo de trabalho
	fmt.Printf("Cozinheiro %d terminou pedido %d\n", id, pedido)
}

func main() {
	sem := make(chan bool, numCozinheiros)
	fmt.Println("Iniciando a cozinha (usando canal bufferizado como semáforo)...")
	startTime := time.Now()
	for i := 0; i < numPedidos; i++ {
		sem <- bool(true) // Envia um valor para o canal semáforo, bloqueando se o canal estiver cheio
		pedido := i       // Captura a variável do loop para a goroutine
		go func(cozinheiroID int, p int) {
			doThing(cozinheiroID, p) // Executa o trabalho de cozinhar
			<-sem                    // Libera o "permit" de volta para o semáforo
		}(i%numCozinheiros, pedido) // Atribui um "ID de cozinheiro" para fins de impressão
	}
	// Wait for all goroutines to finish by draining the sem channel
	for i := 0; i < numCozinheiros; i++ {
		sem <- true
	}
	elapsedTime := time.Since(startTime)
	fmt.Printf("\nTempo total: %s para %d pedidos.\n", elapsedTime, numPedidos)
	fmt.Printf("Com %d cozinheiros (limite de concorrência usando canal bufferizado).\n", numCozinheiros)
}
