@echo off
goto comment
    Build the command lines and tests in Windows.
    Must install gcc tool before building.
:comment

set para=%*
if not defined para (
    set act=all
)else (
    set act=%1
)

call :versionSet
set PROJECT_DIR=gin-quick
set BASEDIR=%~dp0
REM build with verison infos
set versionDir=github.com/wxcsdb88/gin-quick/version
set ldflagsRelase="-s -w -X %versionDir%.gitBranch=%gitBranch% -X %versionDir%.gitTag=%gitTag% -X %versionDir%.buildDate=%buildDate% -X %versionDir%.gitCommit=%gitCommit% -X %versionDir%.gitTreeState=%gitTreeState%"
set ldflagsOrigin="-X %versionDir%.gitBranch=%gitBranch% -X %versionDir%.gitTag=%gitTag% -X %versionDir%.buildDate=%buildDate% -X %versionDir%.gitCommit=%gitCommit% -X %versionDir%.gitTreeState=%gitTreeState%"


REM echo "choose cmd [%act%]"

if "%act%"=="all" (
    call :all 
) else ( 
    if "%act%"=="api" (
        call :api 
    ) else ( 
        if "%act%" == "release" (
            call :release
        ) else (
            if "%act%" == "clean" (
                call :clean
            )
        )
    )
)
exit

:all 
call :release 
goto:eof

:api
echo "call api"
echo on
go build -v -ldflags %ldflagsOrigin% -o ./build/bin/gin-quick-api.exe ./cmd/api
@echo "Done gin-quick-api building debug"
@echo off
goto:eof

:release
echo "call release"
echo on
go build -v -ldflags %ldflagsRelase% -o ./build/bin/gin-quick-api.exe ./cmd/api
@echo "Done gin-quick-api building release"
@echo off
goto:eof

:clean
echo "call clean"
echo "clean the build/bin dir"
del build\bin\* /q /f /s
@echo off
goto:eof


:versionSet
for /F %%i in ('"git symbolic-ref --short -q HEAD"') do ( set gitBranch=%%i)

if "%gitBranch%" == "" (
   for /F %%i in ('"git describe --tags --abbrev=0"') do ( set gitTag=%%i)
) 

REM fixed the hour error fill with space when hour is less than 10
set hour=%time:~,2%
if "%time:~,1%"==" " set hour=0%time:~1,1%

set buildDate=%date:~0,4%-%date:~5,2%-%date:~8,2%T%hour%:%time:~3,2%:%time:~6,2%  
for /F %%i in ('"git rev-parse HEAD"') do ( set gitCommit=%%i)

for /F %%i in ('"git status|findstr 'clean'"') do ( set gitTreeState=%%i)

git status|findstr "clean" && set gitTreeState=clean || set gitTreeState=dirty

@echo off
goto:eof