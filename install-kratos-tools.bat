@echo off
:: ==============================================
::  Kratos v2 Protobuf 插件一键安装脚本 (Windows)
::  功能：安装 protoc-gen-go-http, go-kratos, go-errors
:: ==============================================

echo 🚀 开始安装 Kratos v2 Protobuf 代码生成插件...
echo.

:: 检查是否安装了 go
where go >nul 2>&1
if %errorlevel% neq 0 (
    echo ❌ 错误：未找到 'go' 命令，请先安装 Go (https://golang.org/dl)
    echo    并确保已加入系统 PATH。
    pause
    exit /b 1
)

echo ✅ 检测到 Go 已安装

:: 获取 GOPATH
for /f "delims=" %%i in ('go env GOPATH') do set GOPATH=%%i
set GOBIN=%GOPATH%\bin

echo.
echo 📦 GOPATH: %GOPATH%
echo    GOBIN: %GOBIN%
echo.

:: 创建 bin 目录（如果不存在）
if not exist "%GOBIN%" (
    mkdir "%GOBIN%"
    echo ✅ 已创建 %GOBIN%
)

:: 开始安装插件
echo.
echo ⚙️ 正在安装 Kratos 插件...

call :install_tool github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
call :install_tool github.com/go-kratos/kratos/cmd/protoc-gen-go-kratos/v2@latest
call :install_tool github.com/go-kratos/kratos/cmd/protoc-gen-go-errors@latest

echo.
echo ✅ 所有插件安装完成！

:: 检查 %GOBIN% 是否在 PATH 中
echo.
echo 🔍 正在检查 %GOBIN% 是否在系统 PATH 中...
setx PATH "%PATH%;%GOBIN%" >nul 2>&1
echo 🛠️ 已尝试将 %GOBIN% 添加到用户 PATH。
echo    请关闭并重新打开终端，使环境变量生效。

echo.
echo 💡 使用方法：
echo    protoc --go-http_out=. --go-kratos_out=. your.proto
echo.
echo 🎉 安装完成！重启终端后即可使用 protoc 插件。
pause
exit /b 0

:: ----------------------------
:: 函数：安装单个工具
:: ----------------------------
:install_tool
set PACKAGE=%~1
echo.
echo   📥 安装: %PACKAGE%
go install %PACKAGE%
if %errorlevel% equ 0 (
    echo   ✅ 成功: %PACKAGE%
) else (
    echo   ❌ 失败: %PACKAGE%
    echo   请检查网络或 GOPROXY 设置。
    echo   推荐设置 GOPROXY:
    echo   go env -w GOPROXY=https://goproxy.cn,direct
)
goto :eof