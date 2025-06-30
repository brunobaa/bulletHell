@echo off
title BULLET HELL NETWORK - SISTEMA MULTIPLAYER
color 0E

echo.
echo ========================================
echo    BULLET HELL NETWORK MULTIPLAYER
echo ========================================
echo.
echo Este e o sistema multiplayer melhorado!
echo.
echo ESCOLHA UMA OPCAO:
echo.
echo 1. Iniciar como Host (Player 1)
echo 2. Conectar como Cliente (Player 2+)
echo 3. Compilar o jogo
echo 4. Ver instrucoes
echo 5. Sair
echo.
set /p escolha="Digite sua escolha (1-5): "

if "%escolha%"=="1" goto host
if "%escolha%"=="2" goto client
if "%escolha%"=="3" goto compile
if "%escolha%"=="4" goto instructions
if "%escolha%"=="5" goto sair
goto erro

:host
cls
echo.
echo ========================================
echo    INICIANDO COMO HOST (Player 1)
echo ========================================
echo.
echo INSTRUCOES:
echo.
echo 1. Este terminal sera o Host (Player 1)
echo 2. O jogo iniciara automaticamente
echo 3. Outros jogadores podem se conectar
echo    usando o endereco IP desta maquina
echo.
echo CONTROLES:
echo - Setas direcionais para mover
echo - ESC para sair
echo.
echo Pressione qualquer tecla para iniciar...
pause > nul

cls
echo Iniciando Host...
echo.
timeout /t 2 /nobreak > nul

if exist bulletHell_network.exe (
    bulletHell_network.exe
) else (
    echo ERRO: bulletHell_network.exe nao encontrado!
    echo Execute a opcao 3 para compilar primeiro.
    echo.
    pause
)
goto fim

:client
cls
echo.
echo ========================================
echo    CONECTANDO COMO CLIENTE
echo ========================================
echo.
echo INSTRUCOES:
echo.
echo 1. Este terminal sera um Cliente
echo 2. Digite o endereco IP do Host
echo 3. Use 'localhost' se estiver na mesma maquina
echo.
echo CONTROLES:
echo - WASD para mover
echo - ESC para sair
echo.
echo Pressione qualquer tecla para continuar...
pause > nul

cls
echo Conectando como cliente...
echo.
timeout /t 2 /nobreak > nul

if exist bulletHell_network.exe (
    bulletHell_network.exe
) else (
    echo ERRO: bulletHell_network.exe nao encontrado!
    echo Execute a opcao 3 para compilar primeiro.
    echo.
    pause
)
goto fim

:compile
cls
echo.
echo ========================================
echo    COMPILANDO O JOGO
echo ========================================
echo.
echo Compilando bulletHell_network.go...
echo.

go build -o bulletHell_network.exe bulletHell_network.go

if %ERRORLEVEL% EQU 0 (
    echo.
    echo ========================================
    echo    COMPILACAO CONCLUIDA!
    echo ========================================
    echo.
    echo O jogo foi compilado com sucesso!
    echo Agora voce pode executar as opcoes 1 ou 2.
    echo.
) else (
    echo.
    echo ========================================
    echo    ERRO NA COMPILACAO!
    echo ========================================
    echo.
    echo Verifique se o Go esta instalado:
    echo go version
    echo.
)
pause
goto inicio

:instructions
cls
echo.
echo ========================================
echo    INSTRUCOES COMPLETAS
echo ========================================
echo.
echo SISTEMA MULTIPLAYER VIA REDE:
echo.
echo 1. COMPILACAO:
echo    - Execute a opcao 3 para compilar
echo    - Ou use: go build bulletHell_network.go
echo.
echo 2. HOST (Player 1):
echo    - Execute a opcao 1
echo    - O servidor iniciara na porta 8080
echo    - Anote o IP da maquina
echo.
echo 3. CLIENTE (Player 2+):
echo    - Execute a opcao 2
echo    - Digite o IP do Host
echo    - Use 'localhost' para teste local
echo.
echo 4. CONTROLES:
echo    - Host: Setas direcionais
echo    - Cliente: WASD
echo    - Ambos: ESC para sair
echo.
echo 5. CARACTERISTICAS:
echo    - Suporte a ate 4 jogadores
echo    - Estado compartilhado em tempo real
echo    - Reconexao automatica
echo    - Interface melhorada
echo.
echo Pressione qualquer tecla para voltar...
pause > nul
goto inicio

:erro
echo.
echo Opcao invalida! Tente novamente.
echo.
timeout /t 2 /nobreak > nul
goto inicio

:sair
echo.
echo Obrigado por usar o Bullet Hell Network!
echo.
timeout /t 2 /nobreak > nul
exit

:fim
echo.
echo Jogo finalizado!
echo.
timeout /t 3 /nobreak > nul

:inicio
cls
goto :eof 