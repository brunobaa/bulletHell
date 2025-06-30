# Bullet Hell Network - Sistema Multiplayer Distribuído

Este é um sistema multiplayer avançado para o jogo Bullet Hell, onde múltiplos jogadores podem se conectar via rede e compartilhar o mesmo estado do jogo em tempo real.

## 🚀 Características Principais

### Sistema Distribuído
- **Compartilhamento de Estado**: Todos os jogadores veem o mesmo estado do jogo
- **Sincronização em Tempo Real**: Atualizações instantâneas via TCP
- **Suporte a Múltiplos Jogadores**: Até 4 jogadores simultâneos
- **Reconexão Automática**: Sistema robusto de tratamento de desconexões

### Arquitetura
- **Servidor Autoritativo**: O host controla a lógica do jogo
- **Clientes Sincronizados**: Clientes recebem e enviam inputs
- **Comunicação Bidirecional**: Inputs e estados são trocados continuamente

## 📋 Pré-requisitos

- Go 1.24+ instalado
- Acesso à rede (para jogar em computadores diferentes)
- Terminal compatível com termbox-go

## 🛠️ Instalação e Compilação

### Método 1: Script Automático
```bash
# Execute o script de inicialização
iniciar_network.bat
# Escolha a opção 3 para compilar
```

### Método 2: Compilação Manual
```bash
# Verificar dependências
go mod tidy

# Compilar o jogo
go build -o bulletHell_network.exe bulletHell_network.go
```

## 🎮 Como Jogar

### 1. Iniciar o Host (Player 1)
```bash
# Execute o programa
./bulletHell_network.exe
# Escolha opção 1 (Iniciar servidor)
```

O host será automaticamente o Player 1 e aguardará conexões de outros jogadores.

### 2. Conectar como Cliente (Player 2+)
```bash
# Execute o programa
./bulletHell_network.exe
# Escolha opção 2 (Conectar como cliente)
# Digite o endereço do servidor (ex: localhost ou IP)
```

### 3. Controles
- **Host (Player 1)**: Setas direcionais para mover
- **Clientes (Player 2+)**: WASD para mover
- **Todos**: ESC para sair

## 🌐 Configuração de Rede

### Jogando na Mesma Máquina
- Use `localhost` como endereço do servidor

### Jogando em Computadores Diferentes
1. **No computador do Host**:
   - Execute como servidor
   - Anote o IP da máquina (`ipconfig` no Windows)

2. **Nos computadores dos Clientes**:
   - Execute como cliente
   - Digite o IP do computador do Host

### Porta e Firewall
- **Porta padrão**: 8080
- **Protocolo**: TCP
- Certifique-se de que o firewall permite conexões na porta 8080

## 🏗️ Arquitetura Técnica

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

### Tratamento de Erros
- **Desconexões**: Jogadores são marcados como inativos
- **Reconexão**: Novos clientes podem se conectar
- **Timeout**: Conexões perdidas são detectadas automaticamente

## 🔧 Melhorias Implementadas

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

## 🐛 Solução de Problemas

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
- Verifique se o firewall não está bloqueando a porta 8080
- Confirme se o IP do host está correto
- Teste com `localhost` primeiro

### Jogadores Não Aparecem
- Verifique se o cliente se conectou com sucesso
- Confirme se o host está rodando
- Verifique os logs de conexão

## 📊 Performance

### Requisitos de Rede
- **Largura de banda mínima**: 1 KB/s por jogador
- **Latência recomendada**: < 100ms
- **Protocolo**: TCP para confiabilidade

### Otimizações
- **Compressão de mensagens** via JSON
- **Sincronização eficiente** de estados
- **Gerenciamento de memória** otimizado

## 🔮 Próximas Melhorias

### Planejadas
- [ ] Suporte a salas de jogo
- [ ] Chat entre jogadores
- [ ] Modos de jogo diferentes
- [ ] Sistema de pontuação
- [ ] Persistência de dados

### Possíveis
- [ ] Interface gráfica
- [ ] Sons e música
- [ ] Power-ups
- [ ] Diferentes tipos de projéteis

## 📝 Licença

Este projeto é parte do trabalho de FPPD (Fundamentos de Programação Paralela e Distribuída).

## 🤝 Contribuição

Para contribuir com melhorias:
1. Teste o sistema atual
2. Identifique problemas ou melhorias
3. Implemente as mudanças
4. Teste novamente
5. Documente as alterações

---

**Desenvolvido para demonstrar conceitos de sistemas distribuídos e programação concorrente em Go.** 