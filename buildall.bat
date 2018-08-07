@echo off 
goto comment
    Build the command lines and tests in Windows.
    Must install gcc tool before building.
:comment

echo on

go build -ldflags "-s -w" -o ./build/gin-quick-api.exe ./cmd/api
@echo "Done gin-quick-api building release"

pause
