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
	sem := make(chan struct{}, numCozinheiros) // struct{} é usado para economizar memória, pois não precisamos de um valor real, apenas do sinal
	pedidosProntos := make(chan int, numPedidos)
	concluido := make(chan struct{}, numPedidos) // Canal para sinalizar conclusão de cada goroutine
	fmt.Println("Iniciando a cozinha (usando canal bufferizado como semáforo)...")
	startTime := time.Now()

	for i := 0; i < numPedidos; i++ {
		sem <- struct{}{}
		pedido := i // Captura a variável do loop para a goroutine
		go func(cozinheiroID int, p int) {
			defer func() { concluido <- struct{}{} }()
			defer func() { <-sem }() // Note o uso de uma função anônima para garantir que o defer seja executado corretamente

			doThing(cozinheiroID, p) // Executa o trabalho de cozinhar
			pedidosProntos <- p      // Envia o pedido concluído para o canal de prontos
		}(i%numCozinheiros, pedido) // Atribui um "ID de cozinheiro" para fins de impressão
	}
	go func() {
		for i := 0; i < numPedidos; i++ {
			<-concluido
		}
		close(pedidosProntos) // Fecha o canal de pedidos prontos
	}()
	fmt.Println("\n--- Pedidos Prontos ---")
	completedOrders := 0
	for pedido := range pedidosProntos {
		fmt.Printf("Pedido %d pronto!\n", pedido)
		completedOrders++
	}
	elapsedTime := time.Since(startTime)
	fmt.Printf("\nTempo total: %s para %d pedidos.\n", elapsedTime, completedOrders)
	fmt.Printf("Com %d cozinheiros (limite de concorrência usando canal bufferizado).\n", numCozinheiros)
}
