package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	WorldWidth    = 30
	WorldHeight   = 15
	UpdatesPerSec = 10
	MaxLives      = 5
	Port          = 8080
	MaxPlayers    = 4 // Suporte para até 4 jogadores
)

type Entity struct {
	X, Y int
	Ch   rune
}

type Bullet struct {
	Entity
	DirectionX int
	DirectionY int
	Active     bool
}

type Player struct {
	Entity
	Lives    int
	ID       int
	Active   bool
	LastSeen time.Time
}

type GameState struct {
	Players map[int]*Player
	Bullets []Bullet
	Tick    int
	Started bool
}

type NetworkMessage struct {
	Type      string    `json:"type"` // "state", "input", "join", "ping", "pong"
	PlayerID  int       `json:"player_id"`
	GameState GameState `json:"game_state,omitempty"`
	Input     string    `json:"input,omitempty"`
	Timestamp int64     `json:"timestamp"`
}

type Connection struct {
	conn     net.Conn
	playerID int
	encoder  *json.Encoder
	decoder  *json.Decoder
	active   bool
	mutex    sync.Mutex
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func render(gameState GameState, playerID int) {
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

	// posiciona jogadores ativos
	for _, player := range gameState.Players {
		if player.Active {
			grid[player.Y][player.X] = player.Ch
		}
	}

	// posiciona projéteis
	for _, b := range gameState.Bullets {
		if b.Active && b.Y > 0 && b.Y < WorldHeight-1 && b.X > 0 && b.X < WorldWidth-1 {
			grid[b.Y][b.X] = b.Ch
		}
	}

	// limpa e desenha
	clearScreen()
	fmt.Printf("=== BULLET HELL MULTIPLAYER (Player %d) ===\n", playerID)
	fmt.Printf("Jogadores ativos: %d\n", len(gameState.Players))

	for _, row := range grid {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}

	// Desenha barras de vida
	for _, player := range gameState.Players {
		if player.Active {
			fmt.Printf("\nPlayer %d Vidas: ", player.ID)
			for i := 0; i < player.Lives; i++ {
				fmt.Print("♥ ")
			}
			for i := player.Lives; i < MaxLives; i++ {
				fmt.Print("♡ ")
			}
		}
	}
	fmt.Printf("\nTick: %d\n", gameState.Tick)
}

func checkCollision(bullet Bullet, player Entity) bool {
	return bullet.X == player.X && bullet.Y == player.Y
}

func updateBullets(gameState *GameState) {
	// Atualiza posição dos projéteis ativos
	for i := range gameState.Bullets {
		if gameState.Bullets[i].Active {
			gameState.Bullets[i].X += gameState.Bullets[i].DirectionX
			gameState.Bullets[i].Y += gameState.Bullets[i].DirectionY

			// Verifica colisão com todos os jogadores ativos
			for _, player := range gameState.Players {
				if player.Active && checkCollision(gameState.Bullets[i], player.Entity) {
					gameState.Bullets[i].Active = false
					if player.Lives > 0 {
						player.Lives--
					}
					break
				}
			}

			// Desativa projéteis que atingiram as bordas
			if gameState.Bullets[i].X <= 0 || gameState.Bullets[i].X >= WorldWidth-1 ||
				gameState.Bullets[i].Y <= 0 || gameState.Bullets[i].Y >= WorldHeight-1 {
				gameState.Bullets[i].Active = false
			}
		}
	}
}

func spawnBullet() Bullet {
	side := rand.Intn(4)

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

func checkGameOver(gameState GameState) (bool, string) {
	activePlayers := 0
	alivePlayers := 0
	var lastAlivePlayer int

	for _, player := range gameState.Players {
		if player.Active {
			activePlayers++
			if player.Lives > 0 {
				alivePlayers++
				lastAlivePlayer = player.ID
			}
		}
	}

	// Se o jogo não começou ainda, não termina
	if !gameState.Started {
		return false, ""
	}

	// Se não há jogadores ativos, não termina (aguarda conexões)
	if activePlayers == 0 {
		return false, ""
	}

	// Se há apenas 1 jogador ativo, não termina (aguarda mais jogadores)
	if activePlayers == 1 {
		return false, ""
	}

	// Só verifica game over se há pelo menos 2 jogadores
	if alivePlayers == 0 {
		return true, "EMPATE! Todos os jogadores morreram!"
	} else if alivePlayers == 1 {
		return true, fmt.Sprintf("PLAYER %d VENCEU!", lastAlivePlayer)
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

	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			break
		}
	}
}

func showMenu() int {
	clearScreen()
	fmt.Println("╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                    BULLET HELL MULTIPLAYER                   ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════╣")
	fmt.Println("║  1. Iniciar servidor (Host)                                 ║")
	fmt.Println("║  2. Conectar como cliente                                   ║")
	fmt.Println("║  3. Sobre                                                   ║")
	fmt.Println("║  4. Sair                                                    ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════╝")
	fmt.Print("Escolha uma opção: ")

	var choice int
	fmt.Scanf("%d", &choice)
	return choice
}

func showAbout() {
	clearScreen()
	fmt.Println("╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                           SOBRE                              ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════╣")
	fmt.Println("║  Bullet Hell Multiplayer - Sistema Distribuído              ║")
	fmt.Println("║                                                              ║")
	fmt.Println("║  Características:                                           ║")
	fmt.Println("║  • Suporte a até 4 jogadores simultâneos                    ║")
	fmt.Println("║  • Compartilhamento de estado em tempo real                 ║")
	fmt.Println("║  • Comunicação via TCP                                      ║")
	fmt.Println("║  • Reconexão automática                                     ║")
	fmt.Println("║                                                              ║")
	fmt.Println("║  Controles:                                                 ║")
	fmt.Println("║  • Setas direcionais: Mover                                 ║")
	fmt.Println("║  • ESC: Sair                                                ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════╝")
	fmt.Print("Pressione qualquer tecla para voltar...")

	var input string
	fmt.Scanf("%s", &input)
}

func runServer() {
	fmt.Println("Iniciando servidor...")

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Estado inicial do jogo
	gameState := GameState{
		Players: make(map[int]*Player),
		Bullets: make([]Bullet, 0),
		Tick:    0,
		Started: false,
	}

	// Adiciona o host como Player 1
	gameState.Players[1] = &Player{
		Entity:   Entity{X: 2, Y: 2, Ch: '1'},
		Lives:    MaxLives,
		ID:       1,
		Active:   true,
		LastSeen: time.Now(),
	}

	// Inicia servidor
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(Port))
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Printf("Servidor iniciado na porta %d\n", Port)
	fmt.Println("Aguardando conexões...")

	// Canais para comunicação
	inputChan := make(chan string)
	done := make(chan bool)
	connections := make(map[int]*Connection)
	var connMutex sync.Mutex

	// Goroutine para captura de input do host
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				ev := termbox.PollEvent()
				if ev.Type == termbox.EventKey {
					switch ev.Key {
					case termbox.KeyArrowUp:
						if gameState.Players[1].Y > 1 {
							gameState.Players[1].Y--
							inputChan <- "up"
						}
					case termbox.KeyArrowDown:
						if gameState.Players[1].Y < WorldHeight-2 {
							gameState.Players[1].Y++
							inputChan <- "down"
						}
					case termbox.KeyArrowLeft:
						if gameState.Players[1].X > 1 {
							gameState.Players[1].X--
							inputChan <- "left"
						}
					case termbox.KeyArrowRight:
						if gameState.Players[1].X < WorldWidth-2 {
							gameState.Players[1].X++
							inputChan <- "right"
						}
					case termbox.KeyEsc:
						done <- true
						return
					}
				}
			}
		}
	}()

	// Goroutine para aceitar conexões
	go func() {
		playerID := 2
		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}

			if playerID > MaxPlayers {
				conn.Close()
				continue
			}

			connMutex.Lock()
			connections[playerID] = &Connection{
				conn:     conn,
				playerID: playerID,
				encoder:  json.NewEncoder(conn),
				decoder:  json.NewDecoder(conn),
				active:   true,
			}

			// Adiciona novo jogador
			gameState.Players[playerID] = &Player{
				Entity:   Entity{X: WorldWidth - 3, Y: WorldHeight - 3, Ch: rune('0' + playerID)},
				Lives:    MaxLives,
				ID:       playerID,
				Active:   true,
				LastSeen: time.Now(),
			}

			fmt.Printf("Player %d conectado!\n", playerID)
			playerID++
			connMutex.Unlock()

			// Goroutine para receber input do jogador
			go func(conn *Connection) {
				for {
					var msg NetworkMessage
					if err := conn.decoder.Decode(&msg); err != nil {
						connMutex.Lock()
						conn.active = false
						if gameState.Players[conn.playerID] != nil {
							gameState.Players[conn.playerID].Active = false
						}
						connMutex.Unlock()
						return
					}

					if msg.Type == "input" && gameState.Players[conn.playerID] != nil {
						player := gameState.Players[conn.playerID]
						switch msg.Input {
						case "up":
							if player.Y > 1 {
								player.Y--
							}
						case "down":
							if player.Y < WorldHeight-2 {
								player.Y++
							}
						case "left":
							if player.X > 1 {
								player.X--
							}
						case "right":
							if player.X < WorldWidth-2 {
								player.X++
							}
						}
						player.LastSeen = time.Now()
					}
				}
			}(connections[playerID-1])
		}
	}()

	// Loop principal do jogo
	ticker := time.NewTicker(time.Second / UpdatesPerSec)
	defer ticker.Stop()

	spawnTicker := time.NewTicker(time.Second)
	defer spawnTicker.Stop()

	for {
		select {
		case <-done:
			return
		case input := <-inputChan:
			// Envia input para todos os clientes
			connMutex.Lock()
			for _, conn := range connections {
				if conn.active {
					msg := NetworkMessage{
						Type:      "input",
						PlayerID:  1,
						Input:     input,
						Timestamp: time.Now().UnixNano(),
					}
					conn.encoder.Encode(msg)
				}
			}
			connMutex.Unlock()
		case <-spawnTicker.C:
			if gameState.Started {
				gameState.Bullets = append(gameState.Bullets, spawnBullet())
			}
		case <-ticker.C:
			gameState.Tick++

			// Verifica se deve iniciar o jogo (pelo menos 2 jogadores)
			connMutex.Lock()
			activeCount := 0
			for _, conn := range connections {
				if conn.active {
					activeCount++
				}
			}
			connMutex.Unlock()

			if !gameState.Started && activeCount >= 1 { // Host + pelo menos 1 cliente
				gameState.Started = true
				fmt.Println("Jogo iniciado! Há pelo menos 2 jogadores conectados.")
			}

			if gameState.Started {
				updateBullets(&gameState)
			}

			// Verifica game over
			if gameOver, message := checkGameOver(gameState); gameOver {
				render(gameState, 1)
				showGameOver(message)
				return
			}

			// Envia estado atualizado para todos os clientes
			connMutex.Lock()
			for _, conn := range connections {
				if conn.active {
					msg := NetworkMessage{
						Type:      "state",
						PlayerID:  1,
						GameState: gameState,
						Timestamp: time.Now().UnixNano(),
					}
					conn.encoder.Encode(msg)
				}
			}
			connMutex.Unlock()

			render(gameState, 1)
		}
	}
}

func runClient(serverAddr string) {
	fmt.Println("Conectando ao servidor...")

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Conecta ao servidor
	conn, err := net.Dial("tcp", serverAddr+":"+strconv.Itoa(Port))
	if err != nil {
		fmt.Printf("Erro ao conectar: %v\n", err)
		fmt.Println("Pressione qualquer tecla para sair...")
		var input string
		fmt.Scanf("%s", &input)
		return
	}
	defer conn.Close()

	fmt.Println("Conectado ao servidor!")

	// Estado local do jogo
	var gameState GameState
	var playerID int

	// Canais para comunicação
	inputChan := make(chan string)
	done := make(chan bool)

	// Goroutine para captura de input
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				ev := termbox.PollEvent()
				if ev.Type == termbox.EventKey {
					switch ev.Key {
					case termbox.KeyArrowUp:
						inputChan <- "up"
					case termbox.KeyArrowDown:
						inputChan <- "down"
					case termbox.KeyArrowLeft:
						inputChan <- "left"
					case termbox.KeyArrowRight:
						inputChan <- "right"
					case termbox.KeyEsc:
						done <- true
						return
					}
				}
			}
		}
	}()

	// Goroutine para enviar input ao servidor
	go func() {
		encoder := json.NewEncoder(conn)
		for {
			select {
			case <-done:
				return
			case input := <-inputChan:
				msg := NetworkMessage{
					Type:      "input",
					PlayerID:  playerID,
					Input:     input,
					Timestamp: time.Now().UnixNano(),
				}
				encoder.Encode(msg)
			}
		}
	}()

	// Goroutine para receber estado do servidor
	go func() {
		decoder := json.NewDecoder(conn)
		for {
			var msg NetworkMessage
			if err := decoder.Decode(&msg); err != nil {
				fmt.Println("Erro ao receber estado:", err)
				done <- true
				return
			}

			if msg.Type == "state" {
				gameState = msg.GameState
			}
		}
	}()

	// Loop principal
	ticker := time.NewTicker(time.Second / UpdatesPerSec)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			// Verifica game over
			if gameOver, message := checkGameOver(gameState); gameOver {
				render(gameState, playerID)
				showGameOver(message)
				return
			}

			render(gameState, playerID)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		choice := showMenu()

		switch choice {
		case 1:
			runServer()
		case 2:
			var serverAddr string
			fmt.Print("Digite o endereço do servidor (ex: localhost): ")
			fmt.Scanf("%s", &serverAddr)
			runClient(serverAddr)
		case 3:
			showAbout()
		case 4:
			fmt.Println("Obrigado por jogar!")
			return
		default:
			fmt.Println("Opção inválida!")
			time.Sleep(2 * time.Second)
		}
	}
}
