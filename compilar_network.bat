@echo off
title COMPILANDO BULLET HELL NETWORK
color 0A

echo.
echo ========================================
echo    COMPILANDO BULLET HELL NETWORK
echo ========================================
echo.

echo Verificando dependencias...
go mod tidy

echo.
echo Compilando bulletHell_network.go...
go build -o bulletHell_network.exe bulletHell_network.go

if %ERRORLEVEL% EQU 0 (
    echo.
    echo ========================================
    echo    COMPILACAO CONCLUIDA COM SUCESSO!
    echo ========================================
    echo.
    echo Arquivo gerado: bulletHell_network.exe
    echo.
    echo Para executar:
    echo 1. Player 1 (Host): bulletHell_network.exe
    echo 2. Player 2 (Cliente): bulletHell_network.exe
    echo.
    echo Pressione qualquer tecla para sair...
    pause > nul
) else (
    echo.
    echo ========================================
    echo    ERRO NA COMPILACAO!
    echo ========================================
    echo.
    echo Verifique se o Go esta instalado:
    echo go version
    echo.
    echo Pressione qualquer tecla para sair...
    pause > nul
) 