package main

import (
	"fmt"
	"time"
)

const (
	WorldWidth    = 30 // largura do seu “mapa”
	WorldHeight   = 15 // altura do seu “mapa”
	UpdatesPerSec = 30 // quantos frames por segundo
)

type Entity struct {
	X, Y int  // coordenadas dentro do mapa (0..Width-1 / 0..Height-1)
	Ch   rune // caractere a desenhar
}

func clearScreen() {
	// ANSI escape: home cursor & clear screen
	fmt.Print("\033[H\033[2J")
}

func render(entities []Entity, tick int) {
	// prepara grid vazio
	grid := make([][]rune, WorldHeight)
	for y := range grid {
		grid[y] = make([]rune, WorldWidth)
		for x := range grid[y] {
			grid[y][x] = ' '
		}
	}

	// desenha borda
	for x := 0; x < WorldWidth; x++ {
		grid[0][x] = '#'
		grid[WorldHeight-1][x] = '#'
	}
	for y := 0; y < WorldHeight; y++ {
		grid[y][0] = '#'
		grid[y][WorldWidth-1] = '#'
	}

	// posiciona entidades (players, bullets…)
	for _, e := range entities {
		if e.Y > 0 && e.Y < WorldHeight-1 && e.X > 0 && e.X < WorldWidth-1 {
			grid[e.Y][e.X] = e.Ch
		}
	}

	// limpa e desenha
	clearScreen()
	for _, row := range grid {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}
	fmt.Printf("\nTick: %d\n", tick)
}

func main() {
	// estado inicial
	bullet := Entity{X: WorldWidth / 2, Y: WorldHeight / 2, Ch: '*'}
	player1 := Entity{X: 2, Y: 2, Ch: '1'}
	player2 := Entity{X: WorldWidth - 3, Y: WorldHeight - 3, Ch: '2'}

	ticker := time.NewTicker(time.Second / UpdatesPerSec)
	defer ticker.Stop()

	tick := 0
	for range ticker.C {
		tick++
		// aqui pode atualizar posições de players/balas antes de renderizar
		render([]Entity{bullet, player1, player2}, tick)
	}
}
