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
	// Cria um canal bufferizado que atuará como nosso semáforo.
	// O tamanho do buffer (numCozinheiros) limita o número de goroutines
	// que podem "passar" por este ponto ao mesmo tempo.
	sem := make(chan bool, numCozinheiros)

	// Canal concluido para rastrear a conclusão de goroutines
	// concluido := make(chan struct{}, numPedidos) // Canal para sinalizar conclusão de cada goroutine

	fmt.Println("Iniciando a cozinha (usando canal bufferizado como semáforo)...")
	startTime := time.Now()

	// Simula a criação de pedidos e o despacho para cozinheiros
	for i := 0; i < numPedidos; i++ {
		// Adquire um "permit" do semáforo.
		// Isso enviará um valor para o canal 'sem'. Se o canal estiver cheio
		// (ou seja, 'numCozinheiros' goroutines já estão ativas),
		// este envio irá bloquear até que uma goroutine libere um "slot".
		sem <- bool(true) // Envia um valor para o canal semáforo, bloqueando se o canal estiver cheio

		// Incrementa o contador lógico para rastrear goroutines

		pedido := i // Captura a variável do loop para a goroutine
		go func(cozinheiroID int, p int) {
			doThing(cozinheiroID, p) // Executa o trabalho de cozinhar
			// concluido <- struct{}{}  // Sinaliza a conclusão da goroutine
			<-sem // Libera o "permit" de volta para o semáforo
		}(i%numCozinheiros, pedido) // Atribui um "ID de cozinheiro" para fins de impressão
	}

	go func() {
		// Espera até que todas as goroutines sinalizem conclusão no canal "concluido"
		for i := 0; i < numPedidos; i++ {
			// <-concluido
		}
	}()
	completedOrders := 0
	elapsedTime := time.Since(startTime)
	fmt.Printf("\nTempo total: %s para %d pedidos.\n", elapsedTime, completedOrders)
	fmt.Printf("Com %d cozinheiros (limite de concorrência usando canal bufferizado).\n", numCozinheiros)
}
