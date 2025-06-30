# SOLUÇÕES PARA MULTIPLAYER BULLET HELL

## 🎯 PROBLEMA IDENTIFICADO
Os jogos estão dessincronizados porque cada terminal roda uma instância independente.

## 🚀 SOLUÇÕES DISPONÍVEIS

### 1. 🎮 **SOLUÇÃO SIMPLES (2 Jogadores no Mesmo Terminal)**
**RECOMENDADA!**

```bash
.\compilar_2players.bat
```

**Características:**
- ✅ 2 jogadores no mesmo terminal
- ✅ Player 1: Setas direcionais
- ✅ Player 2: WASD
- ✅ Estado compartilhado
- ✅ Não precisa de rede

**Como jogar:**
1. Execute: `.\compilar_2players.bat`
2. Se o Go não estiver instalado, o script te guiará
3. Ambos jogam no mesmo terminal

### 2. 🌐 **SOLUÇÃO AVANÇADA (Multiplayer via Rede)**
Para jogar em computadores diferentes:

```bash
.\compilar_multiplayer.bat
```

**Características:**
- ✅ Jogadores em terminais/computadores diferentes
- ✅ Sincronização via rede
- ✅ Estado compartilhado em tempo real

### 3. 🎲 **SOLUÇÃO BÁSICA (Jogos Independentes)**
Para testar rapidamente:

```bash
# Terminal 1
.\bulletHell.exe

# Terminal 2  
.\bulletHell.exe
```

**Características:**
- ❌ Jogos independentes
- ❌ Sem sincronização
- ✅ Funciona imediatamente

## 📋 INSTRUÇÕES DETALHADAS

### Para a SOLUÇÃO SIMPLES (Recomendada):

1. **Abra um terminal**
2. **Navegue para a pasta:**
   ```bash
   cd t2-jogo\bulletHell
   ```
3. **Execute o compilador:**
   ```bash
   .\compilar_2players.bat
   ```
4. **Se o Go não estiver instalado:**
   - O script te dará opções
   - Ou instale o Go: https://golang.org/dl/
5. **Jogue:**
   - Player 1: Setas direcionais
   - Player 2: WASD
   - ESC para sair

### Para a SOLUÇÃO AVANÇADA:

1. **Instale o Go** (se não tiver)
2. **Compile:**
   ```bash
   .\compilar_multiplayer.bat
   ```
3. **Execute:**
   - Terminal 1: `.\bulletHell_multiplayer.exe` → Opção 1
   - Terminal 2: `.\bulletHell_multiplayer.exe` → Opção 2 → `localhost`

## 🎯 CONTROLES

### Solução Simples (2 jogadores no mesmo terminal):
- **Player 1**: Setas direcionais (↑↓←→)
- **Player 2**: WASD (W=↑, S=↓, A=←, D=→)
- **Ambos**: ESC para sair

### Solução Avançada (via rede):
- **Player 1**: Setas direcionais
- **Player 2**: WASD
- **Ambos**: ESC para sair

## 🔧 RESOLVENDO PROBLEMAS

### "Go não é reconhecido":
- Instale o Go: https://golang.org/dl/
- Ou use a solução básica

### "Erro de compilação":
- Execute: `go mod tidy`
- Verifique se as dependências estão instaladas

### "Arquivo não encontrado":
- Use `.\` antes do nome do arquivo no PowerShell
- Exemplo: `.\compilar_2players.bat`

## 🏆 RECOMENDAÇÃO FINAL

**Use a SOLUÇÃO SIMPLES** (`compilar_2players.bat`) porque:
- ✅ Funciona com 2 jogadores no mesmo terminal
- ✅ Estado compartilhado
- ✅ Mais fácil de configurar
- ✅ Não precisa de rede

**Teste agora e me diga se funcionou!** 🎮 