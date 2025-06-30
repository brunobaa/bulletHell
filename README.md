# Bullet Hell Multiplayer

## 🎮 Visão Geral

**Bullet Hell Multiplayer** é um jogo competitivo 2D em Go que roda no terminal, permitindo que até 4 jogadores participem simultaneamente através de conexão de rede. O jogo apresenta:

* Arena retangular desenhada em ASCII com bordas (`#`)
* Suporte a até 4 jogadores simultâneos (`1`, `2`, `3`, `4`)
* Projéteis (`*`) que aparecem das bordas e se movem pelo mapa
* Sistema de vidas para cada jogador (♥ ♡)
* Comunicação em tempo real via TCP
* Interface de menu intuitiva

## 🚀 Como Rodar

### Pré-requisitos

* [Go](https://go.dev/dl/) (versão 1.18+ recomendada)
* Terminal compatível com ANSI escape codes
* Conexão de rede local (Wi-Fi/LAN) para multiplayer

### Instalação e Compilação

1. **Clone o repositório**
   ```bash
   git clone https://github.com/brunobaa/bulletHell-.git
   cd bulletHell
   ```

2. **Compile o jogo**
   ```bash
   # Windows (usando o script)
   compilar_network.bat
   
   # Ou manualmente
   go build -o bulletHell_network.exe bulletHell_network.go
   ```

### Como Jogar

#### **Opção 1: Jogar no Mesmo Computador**
1. Execute `bulletHell_network.exe`
2. Escolha opção 1 (Iniciar servidor)
3. Em outro terminal, execute novamente e escolha opção 2
4. Digite `localhost` como endereço do servidor

#### **Opção 2: Jogar em Computadores Diferentes**

**No computador servidor (Host):**
1. Execute `bulletHell_network.exe`
2. Escolha opção 1 (Iniciar servidor)
3. Anote o IP que aparecer (ex: `192.168.1.100`)
4. Aguarde os jogadores se conectarem

**No computador cliente:**
1. Execute `bulletHell_network.exe`
2. Escolha opção 2 (Conectar como cliente)
3. Digite o IP do servidor (ex: `192.168.1.100`)
4. Conecte e jogue!

### Controles

* **Setas direcionais:** Mover o jogador
* **ESC:** Sair do jogo

### Como Descobrir o IP do Servidor

**Windows:**
```cmd
ipconfig
```
Procure por "IPv4 Address" (ex: `192.168.1.100`)

**Linux/Mac:**
```bash
ifconfig
# ou
ip addr
```
Procure por "inet" seguido do IP

## 🎯 Regras do Jogo

1. **Objetivo:** Ser o último jogador vivo
2. **Vidas:** Cada jogador começa com 5 vidas (♥)
3. **Projéteis:** Aparecem das bordas e se movem pelo mapa
4. **Colisão:** Tocar um projétil perde 1 vida
5. **Game Over:** Quando apenas 1 jogador permanece vivo
6. **Início:** O jogo só inicia quando há pelo menos 2 jogadores conectados

## 🔧 Configurações

Você pode ajustar as constantes no início de `bulletHell_network.go`:

```go
const (
    WorldWidth    = 30  // largura do mapa (colunas)
    WorldHeight   = 15  // altura do mapa (linhas)
    UpdatesPerSec = 10  // frames por segundo
    MaxLives      = 5   // vidas por jogador
    Port          = 8080 // porta de conexão
    MaxPlayers    = 4   // máximo de jogadores
)
```

## 🌐 Configuração de Rede

### Firewall
Certifique-se de que a porta 8080 está liberada no firewall do computador servidor.

### Rede Local
Ambos os computadores devem estar na mesma rede Wi-Fi/LAN.

### Exemplos de IPs
* `localhost` - mesmo computador
* `192.168.1.100` - outro computador na rede Wi-Fi
* `10.0.0.50` - rede local

## 🏗️ Arquitetura Técnica

### Sistema Distribuído
- **Compartilhamento de Estado**: Todos os jogadores veem o mesmo estado do jogo
- **Sincronização em Tempo Real**: Atualizações instantâneas via TCP
- **Suporte a Múltiplos Jogadores**: Até 4 jogadores simultâneos
- **Reconexão Automática**: Sistema robusto de tratamento de desconexões

### Estruturas de Dados
```go
type GameState struct {
    Players map[int]*Player  // Jogadores ativos
    Bullets []Bullet         // Projéteis no jogo
    Tick    int              // Contador de frames
    Started bool             // Se o jogo iniciou
}

type NetworkMessage struct {
    Type      string     // Tipo da mensagem
    PlayerID  int        // ID do jogador
    GameState GameState  // Estado do jogo
    Input     string     // Input do jogador
    Timestamp int64      // Timestamp da mensagem
}
```

### Fluxo de Comunicação
1. **Host inicia servidor** na porta 8080
2. **Clientes se conectam** via TCP
3. **Host processa lógica** do jogo
4. **Estados são enviados** para todos os clientes
5. **Inputs são recebidos** e aplicados
6. **Ciclo se repete** a cada frame

## 🛠️ Solução de Problemas

### Erro de Compilação
```bash
# Verificar se o Go está instalado
go version

# Atualizar dependências
go mod tidy

# Compilar novamente
go build bulletHell_network.go
```

### Erro de Conexão
1. Verifique se ambos os computadores estão na mesma rede
2. Confirme se o firewall permite conexões na porta 8080
3. Teste com `localhost` primeiro

### Jogo Não Inicia
1. Certifique-se de que há pelo menos 2 jogadores conectados
2. Verifique se todos os jogadores estão ativos

### Performance
1. Reduza `UpdatesPerSec` se houver lag
2. Feche outros programas que usem muita rede

### "Go não é reconhecido"
- Instale o Go: https://golang.org/dl/
- Ou use os executáveis pré-compilados

### "Arquivo não encontrado"
- Use `.\` antes do nome do arquivo no PowerShell
- Exemplo: `.\compilar_network.bat`

## 📁 Arquivos do Projeto

* `bulletHell_network.go` - Código fonte principal
* `bulletHell_network.exe` - Executável compilado
* `compilar_network.bat` - Script de compilação (Windows)
* `iniciar_network.bat` - Script de inicialização
* `instalar_go.bat` - Script de instalação do Go
* `go.mod` - Dependências do Go
* `README.md` - Este arquivo

## 🎮 Características Técnicas

* **Linguagem:** Go
* **Interface:** Terminal ASCII
* **Rede:** TCP/IP
* **Sincronização:** Tempo real
* **Máximo de jogadores:** 4
* **Plataformas:** Windows, Linux, macOS

## 🔮 Melhorias Implementadas

### Comparado ao Sistema Anterior
- ✅ **Suporte a 4 jogadores** (vs 2 anteriormente)
- ✅ **Interface melhorada** com menus
- ✅ **Tratamento de erros** mais robusto
- ✅ **Reconexão automática** de clientes
- ✅ **Sincronização de timestamp** para evitar desync
- ✅ **Sistema de conexões** mais organizado

### Funcionalidades Avançadas
- **Detecção de jogadores ativos**
- **Sistema de vidas individual**
- **Colisão com múltiplos jogadores**
- **Interface de menu interativa**
- **Logs de conexão e desconexão**

## 📊 Performance

### Requisitos de Rede
- **Largura de banda mínima**: 1 KB/s por jogador
- **Latência recomendada**: < 100ms
- **Protocolo**: TCP para confiabilidade

### Otimizações
- **Compressão de mensagens** via JSON
- **Sincronização eficiente** de estados
- **Gerenciamento de memória** otimizado

## 📞 Contato

Para dúvidas, sugestões ou contribuições, entre em contato:

**Email:** [brunoandradeprof7@gmail.com](mailto:brunoandradeprof7@gmail.com)

## 📝 Licença

Este projeto é parte do trabalho de FPPD (Fundamentos de Programação Paralela e Distribuída).

---

**Divirta-se jogando Bullet Hell Multiplayer!** 🎮 