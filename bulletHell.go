package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	WorldWidth    = 30 // largura do seu "mapa"
	WorldHeight   = 15 // altura do seu "mapa"
	UpdatesPerSec = 10 // quantos frames por segundo
	MaxLives      = 5  // número máximo de vidas
)

type Entity struct {
	X, Y int  // coordenadas dentro do mapa (0..Width-1 / 0..Height-1)
	Ch   rune // caractere a desenhar
}

type Bullet struct {
	Entity
	DirectionX int  // direção X do projétil (-1: esquerda, 0: nenhuma, 1: direita)
	DirectionY int  // direção Y do projétil (-1: cima, 0: nenhuma, 1: baixo)
	Active     bool // se o projétil está ativo no jogo
}

type Player struct {
	Entity
	Lives int // número de vidas restantes
}

func clearScreen() {
	// ANSI escape: home cursor & clear screen
	fmt.Print("\033[H\033[2J")
}

func render(entities []Entity, bullets []Bullet, player1 Player, player2 Player, tick int) {
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

	// posiciona entidades (players)
	grid[player1.Y][player1.X] = player1.Ch
	grid[player2.Y][player2.X] = player2.Ch
	for _, e := range entities {
		if e.Y > 0 && e.Y < WorldHeight-1 && e.X > 0 && e.X < WorldWidth-1 {
			grid[e.Y][e.X] = e.Ch
		}
	}

	// posiciona projéteis
	for _, b := range bullets {
		if b.Active && b.Y > 0 && b.Y < WorldHeight-1 && b.X > 0 && b.X < WorldWidth-1 {
			grid[b.Y][b.X] = b.Ch
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

	// Desenha a barra de vida do Player 1
	fmt.Print("\nPlayer 1 Vidas: ")
	for i := 0; i < player1.Lives; i++ {
		fmt.Print("♥ ")
	}
	for i := player1.Lives; i < MaxLives; i++ {
		fmt.Print("♡ ")
	}

	// Desenha a barra de vida do Player 2
	fmt.Print("\nPlayer 2 Vidas: ")
	for i := 0; i < player2.Lives; i++ {
		fmt.Print("♥ ")
	}
	for i := player2.Lives; i < MaxLives; i++ {
		fmt.Print("♡ ")
	}
	fmt.Printf("\nTick: %d\n", tick)
}

func handleInput(player1 *Player, done chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			ev := termbox.PollEvent()
			if ev.Type == termbox.EventKey {
				switch ev.Key {
				case termbox.KeyArrowUp:
					if player1.Y > 1 {
						player1.Y--
					}
				case termbox.KeyArrowDown:
					if player1.Y < WorldHeight-2 {
						player1.Y++
					}
				case termbox.KeyArrowLeft:
					if player1.X > 1 {
						player1.X--
					}
				case termbox.KeyArrowRight:
					if player1.X < WorldWidth-2 {
						player1.X++
					}
				case termbox.KeyEsc:
					done <- true
					return
				}
			}
		}
	}
}

func checkCollision(bullet Bullet, player Entity) bool {
	return bullet.X == player.X && bullet.Y == player.Y
}

func updateBullets(bullets []Bullet, player1 *Player, player2 *Player) []Bullet {
	// Atualiza posição dos projéteis ativos
	for i := range bullets {
		if bullets[i].Active {
			bullets[i].X += bullets[i].DirectionX
			bullets[i].Y += bullets[i].DirectionY

			// Verifica colisão com o player1
			if checkCollision(bullets[i], player1.Entity) {
				bullets[i].Active = false
				if player1.Lives > 0 {
					player1.Lives--
				}
			}

			// Verifica colisão com o player2
			if checkCollision(bullets[i], player2.Entity) {
				bullets[i].Active = false
				if player2.Lives > 0 {
					player2.Lives--
				}
			}

			// Desativa projéteis que atingiram as bordas
			if bullets[i].X <= 0 || bullets[i].X >= WorldWidth-1 ||
				bullets[i].Y <= 0 || bullets[i].Y >= WorldHeight-1 {
				bullets[i].Active = false
			}
		}
	}
	return bullets
}

func spawnBullet() Bullet {
	// Escolhe uma borda aleatória para spawnar o projétil
	side := rand.Intn(4) // 0: topo, 1: direita, 2: baixo, 3: esquerda

	var bullet Bullet
	bullet.Ch = '*'
	bullet.Active = true

	switch side {
	case 0: // topo
		bullet.X = rand.Intn(WorldWidth-2) + 1
		bullet.Y = 1
		bullet.DirectionX = 0
		bullet.DirectionY = 1
	case 1: // direita
		bullet.X = WorldWidth - 2
		bullet.Y = rand.Intn(WorldHeight-2) + 1
		bullet.DirectionX = -1
		bullet.DirectionY = 0
	case 2: // baixo
		bullet.X = rand.Intn(WorldWidth-2) + 1
		bullet.Y = WorldHeight - 2
		bullet.DirectionX = 0
		bullet.DirectionY = -1
	case 3: // esquerda
		bullet.X = 1
		bullet.Y = rand.Intn(WorldHeight-2) + 1
		bullet.DirectionX = 1
		bullet.DirectionY = 0
	}

	return bullet
}

func checkGameOver(player1 Player, player2 Player) (bool, string) {
	if player1.Lives <= 0 && player2.Lives <= 0 {
		return true, "EMPATE! Ambos os jogadores morreram!"
	} else if player1.Lives <= 0 {
		return true, "PLAYER 2 VENCEU! Player 1 foi eliminado!"
	} else if player2.Lives <= 0 {
		return true, "PLAYER 1 VENCEU! Player 2 foi eliminado!"
	}
	return false, ""
}

func showGameOver(message string) {
	clearScreen()
	fmt.Println("╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                        FIM DE JOGO                           ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════╣")
	fmt.Printf("║  %-52s  ║\n", message)
	fmt.Println("╠══════════════════════════════════════════════════════════════╣")
	fmt.Println("║  Pressione qualquer tecla para sair...                       ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════╝")

	// Aguarda uma tecla ser pressionada
	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			break
		}
	}
}

func main() {
	// Inicializa o gerador de números aleatórios
	rand.Seed(time.Now().UnixNano())

	// Inicializa termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// estado inicial
	player1 := Player{
		Entity: Entity{X: 2, Y: 2, Ch: '1'},
		Lives:  MaxLives,
	}
	player2 := Player{
		Entity: Entity{X: WorldWidth - 3, Y: WorldHeight - 3, Ch: '2'},
		Lives:  MaxLives,
	}

	// Lista de projéteis
	bullets := make([]Bullet, 0)

	// Canal para sinalizar quando o jogo deve terminar
	done := make(chan bool)

	// Inicia a goroutine para captura de teclado
	go handleInput(&player1, done)

	ticker := time.NewTicker(time.Second / UpdatesPerSec)
	defer ticker.Stop()

	tick := 0
	spawnTicker := time.NewTicker(time.Second) // Spawna um projétil a cada segundo
	defer spawnTicker.Stop()

	for {
		select {
		case <-done:
			return
		case <-spawnTicker.C:
			// Adiciona um novo projétil
			bullets = append(bullets, spawnBullet())
		case <-ticker.C:
			tick++
			// Atualiza posição dos projéteis
			bullets = updateBullets(bullets, &player1, &player2)

			// Verifica se o jogo acabou
			if gameOver, message := checkGameOver(player1, player2); gameOver {
				// Renderiza uma última vez para mostrar o estado final
				render([]Entity{}, bullets, player1, player2, tick)
				// Mostra a tela de game over
				showGameOver(message)
				return
			}

			// Renderiza o jogo
			render([]Entity{}, bullets, player1, player2, tick)
		}
	}
}
