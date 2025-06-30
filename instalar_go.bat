@echo off
title VERIFICANDO E INSTALANDO GO
color 0B

echo.
echo ========================================
echo    VERIFICANDO INSTALACAO DO GO
echo ========================================
echo.

echo Verificando se o Go esta instalado...
go version >nul 2>&1
if %errorlevel% equ 0 (
    echo.
    echo ✅ GO JA ESTA INSTALADO!
    echo.
    go version
    echo.
    echo Pressione qualquer tecla para continuar...
    pause > nul
    goto compilar
) else (
    echo.
    echo ❌ GO NAO ESTA INSTALADO!
    echo.
    echo Para instalar o Go:
    echo.
    echo 1. Acesse: https://golang.org/dl/
    echo 2. Baixe a versao para Windows
    echo 3. Execute o instalador
    echo 4. Reinicie o terminal
    echo.
    echo Ou use o instalador automatico:
    echo.
    set /p instalar="Deseja tentar instalar automaticamente? (s/n): "
    if /i "%instalar%"=="s" goto instalar_auto
    goto manual
)

:instalar_auto
echo.
echo Tentando instalar o Go automaticamente...
echo.
echo Baixando Go...
powershell -Command "Invoke-WebRequest -Uri 'https://go.dev/dl/go1.21.5.windows-amd64.msi' -OutFile 'go_installer.msi'"
if %errorlevel% equ 0 (
    echo.
    echo Instalando Go...
    msiexec /i go_installer.msi /quiet
    echo.
    echo Go instalado! Reinicie o terminal e tente novamente.
    echo.
    pause
    exit /b 0
) else (
    echo.
    echo Erro no download automatico.
    goto manual
)

:manual
echo.
echo ========================================
echo    INSTALACAO MANUAL NECESSARIA
echo ========================================
echo.
echo 1. Abra o navegador
echo 2. Vá para: https://golang.org/dl/
echo 3. Clique em "Download" para Windows
echo 4. Execute o arquivo baixado
echo 5. Siga as instrucoes do instalador
echo 6. Reinicie este terminal
echo 7. Execute novamente: .\instalar_go.bat
echo.
echo Pressione qualquer tecla para abrir o site...
pause > nul
start https://golang.org/dl/
exit /b 1

:compilar
echo.
echo ========================================
echo    COMPILANDO BULLET HELL MULTIPLAYER
echo ========================================
echo.

echo Baixando dependencias...
go mod tidy

echo.
echo Compilando versao de 2 jogadores...
go build bulletHell_2players.go

if %errorlevel% equ 0 (
    echo ✅ Compilacao concluida!
    echo.
    echo Compilando versao multiplayer via rede...
    go build bulletHell_multiplayer.go
    
    if %errorlevel% equ 0 (
        echo ✅ Compilacao multiplayer concluida!
        echo.
        echo ========================================
        echo    TODAS AS VERSOES COMPILADAS!
        echo ========================================
        echo.
        echo Arquivos criados:
        echo - bulletHell_2players.exe (2 jogadores no mesmo terminal)
        echo - bulletHell_multiplayer.exe (multiplayer via rede)
        echo.
        echo Para jogar:
        echo.
        echo 1. 2 jogadores no mesmo terminal:
        echo    .\bulletHell_2players.exe
        echo.
        echo 2. Multiplayer via rede:
        echo    Terminal 1: .\bulletHell_multiplayer.exe (opcao 1)
        echo    Terminal 2: .\bulletHell_multiplayer.exe (opcao 2)
        echo.
        echo Pressione qualquer tecla para testar a versao de 2 jogadores...
        pause > nul
        
        echo.
        echo Executando versao de 2 jogadores...
        echo.
        bulletHell_2players.exe
    ) else (
        echo ❌ Erro na compilacao multiplayer
        echo.
        echo Mas a versao de 2 jogadores foi compilada com sucesso!
        echo Execute: .\bulletHell_2players.exe
        pause
    )
) else (
    echo ❌ Erro na compilacao
    echo.
    echo Verifique se todas as dependencias estao instaladas
    echo Execute: go mod tidy
    pause
    exit /b 1
) 