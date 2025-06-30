@echo off
title BULLET HELL MULTIPLAYER - GUIA AUTOMATICO
color 0E

echo.
echo ========================================
echo    BULLET HELL MULTIPLAYER
echo    GUIA AUTOMATICO
echo ========================================
echo.
echo Este script vai te ajudar a jogar multiplayer!
echo.
echo ESCOLHA UMA OPCAO:
echo.
echo 1. Iniciar como Player 1 (este terminal)
echo 2. Iniciar como Player 2 (este terminal)
echo 3. Ver instrucoes completas
echo 4. Sair
echo.
set /p escolha="Digite sua escolha (1-4): "

if "%escolha%"=="1" goto player1
if "%escolha%"=="2" goto player2
if "%escolha%"=="3" goto instrucoes
if "%escolha%"=="4" goto sair
goto erro

:player1
cls
echo.
echo ========================================
echo    INICIANDO PLAYER 1
echo ========================================
echo.
echo INSTRUCOES:
echo.
echo 1. Este terminal sera o Player 1
echo 2. Abra OUTRO terminal (Ctrl+Shift+T ou nova janela)
echo 3. No segundo terminal, navegue para esta pasta:
echo    cd t2-jogo\bulletHell
echo 4. Execute: player2.bat
echo 5. Ou simplesmente: bulletHell.exe
echo.
echo CONTROLES:
echo - Setas direcionais para mover
echo - ESC para sair
echo.
echo Pressione qualquer tecla para iniciar...
pause > nul

cls
echo Iniciando Player 1...
echo.
timeout /t 2 /nobreak > nul

bulletHell.exe
goto fim

:player2
cls
echo.
echo ========================================
echo    INICIANDO PLAYER 2
echo ========================================
echo.
echo INSTRUCOES:
echo.
echo 1. Este terminal sera o Player 2
echo 2. Certifique-se de que o Player 1 ja esta rodando
echo    em outro terminal antes de continuar
echo 3. Se o Player 1 nao estiver rodando, volte e
echo    escolha a opcao 1 primeiro
echo.
echo CONTROLES:
echo - Setas direcionais para mover
echo - ESC para sair
echo.
echo Pressione qualquer tecla para iniciar...
pause > nul

cls
echo Iniciando Player 2...
echo.
timeout /t 2 /nobreak > nul

bulletHell.exe
goto fim

:instrucoes
cls
echo.
echo ========================================
echo    INSTRUCOES COMPLETAS
echo ========================================
echo.
echo PARA JOGAR MULTIPLAYER:
echo.
echo 1. Abra DOIS terminais diferentes
echo 2. Em ambos, navegue para: t2-jogo\bulletHell
echo 3. No primeiro terminal, execute: multiplayer_simples.bat
echo 4. No segundo terminal, execute: player2.bat
echo.
echo OU:
echo.
echo 1. Abra DOIS terminais diferentes
echo 2. Em ambos, navegue para: t2-jogo\bulletHell
echo 3. No primeiro terminal, execute: bulletHell.exe
echo 4. No segundo terminal, execute: bulletHell.exe
echo.
echo CONTROLES:
echo - Ambos usam setas direcionais
echo - ESC para sair
echo.
echo REGRAS:
echo - Evite os projeteis (*) das bordas
echo - Cada jogador tem 5 vidas
echo - O ultimo sobrevivente vence!
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
echo Obrigado por jogar!
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