# COMO JOGAR BULLET HELL MULTIPLAYER

## SOLUÇÃO SIMPLES (Sem Compilação)

Se você está tendo problemas com compilação, use esta solução:

### Passo 1: Player 1
1. Abra um terminal/prompt de comando
2. Navegue para a pasta: `cd t2-jogo\bulletHell`
3. Execute: `multiplayer_simples.bat`
4. Siga as instruções na tela

### Passo 2: Player 2
1. Abra **outro** terminal/prompt de comando
2. Navegue para a pasta: `cd t2-jogo\bulletHell`
3. Execute: `player2.bat`
4. Ou simplesmente execute: `bulletHell.exe`

## SOLUÇÃO AVANÇADA (Com Rede)

Se você quer o verdadeiro multiplayer via rede:

### Pré-requisitos:
- Go instalado no sistema
- Acesso à internet para baixar dependências

### Compilação:
```bash
go mod tidy
go build bulletHell_multiplayer.go
```

### Execução:
1. **Player 1 (Servidor):**
   ```bash
   ./bulletHell_multiplayer
   # Escolha opção 1
   ```

2. **Player 2 (Cliente):**
   ```bash
   ./bulletHell_multiplayer
   # Escolha opção 2
   # Digite: localhost
   ```

## CONTROLES

### Player 1:
- **Setas direcionais**: Mover
- **ESC**: Sair

### Player 2:
- **Setas direcionais**: Mover (solução simples)
- **WASD**: Mover (solução avançada)
- **ESC**: Sair

## REGRAS DO JOGO

- Evite os projéteis (*) que aparecem das bordas
- Cada jogador tem 5 vidas
- O último sobrevivente vence!
- Se ambos morrerem ao mesmo tempo, é empate

## RESOLVENDO PROBLEMAS

### Erro "go não é reconhecido":
- Use a **SOLUÇÃO SIMPLES** acima
- Ou instale o Go: https://golang.org/dl/

### Erro de compilação:
- Execute: `go mod tidy`
- Verifique se todas as dependências estão instaladas

### Erro de conexão:
- Verifique se o firewall não está bloqueando
- Use `localhost` para teste local
- Use o IP da máquina para jogar em rede

### Jogo não inicia:
- Verifique se está na pasta correta
- Execute: `bulletHell.exe` diretamente

## TESTE RÁPIDO

Para testar rapidamente:
1. Execute `multiplayer_simples.bat` em um terminal
2. Execute `player2.bat` em outro terminal
3. Ambos jogarão simultaneamente!

## DIFERENÇAS ENTRE AS SOLUÇÕES

### Solução Simples:
- ✅ Funciona imediatamente
- ✅ Não precisa de compilação
- ❌ Não sincroniza estado via rede
- ❌ Cada jogador tem seu próprio jogo

### Solução Avançada:
- ✅ Sincronização real via rede
- ✅ Estado compartilhado
- ❌ Precisa de compilação
- ❌ Mais complexa de configurar 