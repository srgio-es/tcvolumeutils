@echo off

rem *********************************************
rem VARIABLES
rem *********************************************

rem Specify tc_menu script (usually inside $TC_ROOT/tc_menu/foo.bat)
set TC_MENU=CHANGE_ME

rem Specify volumeutils executables folder path
set VOLUMEUTILS_PATH=C:\volumeutils

rem User with DBA role permissions which will invoke review_volumnes command
set USER=infodba
rem Specify TC Password file location for the user invoking the review_volumes command
set PWD_FILE=CHANGE_ME

rem Set output folder to store review_volumes result files. A new folder with execution date will be created
rem inside this one. ie: C:\foo\bar\review_volumes\20210121_00_10_10\
set RESULT_BASE_DIR=C:\foo\bar\review_volumes

rem *********************************************
rem **** LOGGING
rem *********************************************
for /f "tokens=2 delims==" %%a in ('wmic OS Get localdatetime /value') do set "dt=%%a"
set "YY=%dt:~2,2%" & set "YYYY=%dt:~0,4%" & set "MM=%dt:~4,2%" & set "DD=%dt:~6,2%"
set "HH=%dt:~8,2%" & set "Min=%dt:~10,2%" & set "Sec=%dt:~12,2%"

goto set_variables

:set_variables
set OUTPUT_FOLDER=%RESULT_BASE_DIR%\%YYYY%%MM%%DD%_%HH%_%Min%_%Sec%

set LOG_FILE=%OUTPUT_FOLDER%\%~n0_%YYYY%%MM%%DD%.log
set REPORT_FILE=%OUTPUT_FOLDER%\%~n0_%YYYY%%MM%%DD%.xlsx

IF NOT EXIST %OUTPUT_FOLDER% mkdir %OUTPUT_FOLDER%
IF NOT EXIST %OUTPUT_FOLDER%\log mkdir %OUTPUT_FOLDER%\log
IF NOT EXIST %OUTPUT_FOLDER%\before mkdir %OUTPUT_FOLDER%\before
IF NOT EXIST %OUTPUT_FOLDER%\after mkdir %OUTPUT_FOLDER%\after

echo Review Volumes - %~n0%~x0 - %YYYY%%MM%%DD%_%HH%_%Min%_%Sec% >> %LOG_FILE%
echo. >> %LOG_FILE%

call %TC_MENU%

echo ----------- Review Volumes begins ------------ >> %LOG_FILE%
echo. >> %LOG_FILE%
review_volumes -u=%USER% -pf=%PWD_FILE% -g=dba -parallel=5 -rfolder=%OUTPUT_FOLDER%\before >> %LOG_FILE%

echo review_volumes finished with exit code=%ERRORLEVEL% >> %LOG_FILE%
echo. >> %LOG_FILE%

for %%f in (%OUTPUT_FOLDER%\before\*.txt) do call :clean_volume %%~nf

echo ----------- Review Volumes After clean begins ------------ >> %LOG_FILE%
echo. >> %LOG_FILE%
review_volumes -u=%USER% -pf=%PWD_FILE% -g=dba -parallel=5 -rfolder=%OUTPUT_FOLDER%\after >> %LOG_FILE%

echo ----------- Generate Missing Files Report XLSX ------------ >> %LOG_FILE%
echo. >> %LOG_FILE%

set PATH=%PATH%;%VOLUMEUTILS_PATH%
tcvolumeutils reportmissing -f %OUTPUT_FOLDER%\after -r %REPORT_FILE% -v >> %LOG_FILE%

echo Script finished >> %LOG_FILE%

goto :EOF

:clean_volume
if NOT [%1]==[] (
    set VOLUME=%1
    echo Cleaning volume %VOLUME%.txt >> %LOG_FILE%
    echo. >> %LOG_FILE%

    review_volumes -u=%USER% -pf=%PWD_FILE% -g=dba -v=%VOLUME% -if=%OUTPUT_FOLDER%\before\%VOLUME%.txt -of=%OUTPUT_FOLDER%\log\%VOLUME%.log >> %LOG_FILE%

    echo review_volumes for volume %VOLUME% finished with exit code=%ERRORLEVEL% >> %LOG_FILE%
    echo. >> %LOG_FILE%
)