# ConcorenciaGolang

## Descrição

Este projeto é uma simulação de um sistema de cozinha concorrente implementado em Go. Ele demonstra o uso de **goroutines**, **canais bufferizados**.

## Como funciona

- **Goroutines**: Cada cozinheiro funciona como uma goroutine. Eles processam pedidos de forma independente e simultânea.
- **Canal Bufferizado**: Um canal é usado como semáforo para limitar o número máximo de goroutines que podem operar simultaneamente.
- **Canal de Conclusão**: Um canal adicional é usado para sinalizar quando cada goroutine termina seu trabalho, garantindo que todos os pedidos sejam concluídos antes que o programa termine.

## Principais Variáveis

- `numCozinheiros`: Número de cozinheiros disponíveis (representa o limite de concorrência).
- `numPedidos`: Número total de pedidos a serem processados.

## Como executar o código

1. Certifique-se de ter o Go instalado no seu sistema.
2. Clone este repositório ou copie os arquivos para o seu ambiente local.
3. No terminal, navegue até o diretório do projeto e execute:

   ```bash
   go run cozinha.go
   ```

## Saída Esperada

O programa imprimirá mensagens no terminal indicando:

- Qual cozinheiro está processando qual pedido.
- Quando cada pedido foi concluído.
- O tempo total gasto para processar todos os pedidos.
