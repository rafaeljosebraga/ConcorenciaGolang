package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	// "sync"
)

var (
	numCozinheiros = 8 // Representa o tamanho do buffer do canal (capacidade do semáforo)
	numPedidos     = 100
	pedidos        = make([]int, 0) // Simula uma lista de pedidos
	mu             sync.Mutex       // Mutex para proteger o acesso ao slice "pedidos"
)

func doThing(id int, pedido int, completion chan bool) {
	fmt.Printf("Cozinheiro %d preparando pedido %d\n", id, pedido)
	mu.Lock()
	pedidos = append(pedidos, pedido)                       // Simula o processamento do pedido
	time.Sleep(time.Duration(rand.Float64()) * time.Second) // Simula o tempo de trabalho
	// pedidos = pedidos[1:]                                   // Remove o pedido da lista após o processamento
	mu.Unlock()
	fmt.Printf("Cozinheiro %d terminou pedido %d\n", id, pedido)
	completion <- true // Sinaliza que o pedido foi processado
}

func main() {
	sem := make(chan bool, numCozinheiros)
	completion := make(chan bool, numPedidos)

	fmt.Println("Iniciando a cozinha (usando canal bufferizado como semáforo)...")
	startTime := time.Now()
	for i := 0; i < numPedidos; i++ {
		sem <- bool(true) // Envia um valor para o canal semáforo, bloqueando se o canal estiver cheio
		pedido := i       // Captura a variável do loop para a goroutine
		go func(cozinheiroID int, p int, completion chan bool) {
			doThing(cozinheiroID, p, completion) // Executa o trabalho de cozinhar
			<-sem                                // Libera o "permit" de volta para o semáforo
		}(i%numCozinheiros, pedido, completion) // Atribui um "ID de cozinheiro" para fins de impressão
	}

	// Aguarda todas as goroutines sinalizarem conclusão
	for i := 0; i < numPedidos; i++ {
		<-completion
	}

	// Testes e Resultados.
	elapsedTime := time.Since(startTime)
	fmt.Printf("\nTempo total: %s para %d pedidos.\n", elapsedTime, numPedidos)
	fmt.Printf("Com %d cozinheiros (limite de concorrência usando canal bufferizado).\n", numCozinheiros)

	// Teste importante (todos os pedidos foram processados?)
	for i := 0; i < numPedidos; i++ {
		j := 0
		for ; j < len(pedidos); j++ {
			if pedidos[j] == i {
				break // Verifica se o pedido foi processado
				// Verifica se todos os pedidos foram processados
			}
		}
		if pedidos[j] != i {
			fmt.Printf("Pedido %d não processado.\n", i)
		}
	}
}
