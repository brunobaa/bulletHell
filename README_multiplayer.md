# Bullet Hell Multiplayer

Este é um jogo multiplayer distribuído onde dois jogadores podem jogar simultaneamente em terminais diferentes, compartilhando o mesmo estado do jogo através de comunicação de rede.

## Como Jogar

### Configuração

1. **Compile o jogo:**
   ```bash
   go build bulletHell_multiplayer.go
   ```

2. **Execute o jogo:**
   ```bash
   ./bulletHell_multiplayer
   ```

### Iniciando uma Partida

#### Passo 1: Iniciar o Servidor (Player 1)
1. Execute o programa
2. Escolha a opção `1` (Iniciar servidor)
3. O servidor aguardará a conexão do Player 2

#### Passo 2: Conectar o Cliente (Player 2)
1. Em outro terminal, execute o programa
2. Escolha a opção `2` (Conectar como cliente)
3. Digite o endereço do servidor (ex: `localhost` se estiver na mesma máquina)
4. O cliente se conectará ao servidor

### Controles

#### Player 1 (Servidor):
- **Setas direcionais**: Mover o jogador
- **ESC**: Sair do jogo

#### Player 2 (Cliente):
- **W**: Mover para cima
- **S**: Mover para baixo
- **A**: Mover para esquerda
- **D**: Mover para direita
- **ESC**: Sair do jogo

### Objetivo

- Evite os projéteis (*) que aparecem das bordas
- Cada jogador tem 5 vidas
- O último jogador sobrevivente vence
- Se ambos morrerem ao mesmo tempo, é empate

## Arquitetura do Sistema

### Componentes Principais

1. **Servidor (Player 1)**:
   - Gerencia o estado principal do jogo
   - Processa inputs do Player 1
   - Recebe inputs do Player 2 via rede
   - Envia estado atualizado para o Player 2
   - Controla spawn de projéteis

2. **Cliente (Player 2)**:
   - Recebe estado do jogo do servidor
   - Processa inputs do Player 2
   - Envia inputs para o servidor
   - Renderiza o jogo localmente

### Comunicação de Rede

- **Protocolo**: TCP
- **Porta**: 8080
- **Formato**: JSON
- **Tipos de mensagem**:
  - `state`: Estado completo do jogo
  - `input`: Comando de movimento

### Estruturas de Dados

```go
type GameState struct {
    Player1 Player
    Player2 Player
    Bullets []Bullet
    Tick    int
}

type NetworkMessage struct {
    Type      string     // "state" ou "input"
    PlayerID  int        // 1 ou 2
    GameState GameState  // Estado do jogo
    Input     string     // Comando de movimento
}
```

## Características do Sistema Distribuído

### Compartilhamento de Estado
- O servidor mantém o estado autoritativo
- O cliente recebe atualizações em tempo real
- Ambos os jogadores veem o mesmo estado do jogo

### Sincronização
- O servidor processa a lógica do jogo
- O cliente apenas renderiza e envia inputs
- Comunicação bidirecional para inputs

### Concorrência
- Múltiplas goroutines para:
  - Captura de input
  - Comunicação de rede
  - Loop principal do jogo
  - Renderização

## Testando em Dois Computadores

1. **No computador do Player 1:**
   - Execute o servidor
   - Anote o IP do computador

2. **No computador do Player 2:**
   - Execute o cliente
   - Digite o IP do computador do Player 1

## Dependências

- `github.com/nsf/termbox-go`: Para interface de terminal
- Bibliotecas padrão do Go: `net`, `encoding/json`, `time`, `math/rand`

## Limitações Atuais

- Apenas 2 jogadores por partida
- Sem reconexão automática
- Latência pode afetar a jogabilidade
- Sem criptografia na comunicação 