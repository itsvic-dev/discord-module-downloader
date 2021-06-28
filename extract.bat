@echo off
echo --^> Downloading 7-zip...
powershell -Command Invoke-WebRequest -Uri https://www.dropbox.com/s/1g9vchwrpyzoxjx/7z.zip?dl=1 -OutFile 7z.zip
echo --^> Extracting the download...
powershell -Command Expand-Archive -Path 7z.zip -DestinationPath 7z
if exist modules\ (
robocopy 7z %cd%\modules /mov
del 7z.zip
) else (
echo [ ! ] Looks like you didn't downloaded the modules, exiting now
pause
exit
rem basic check to make sure nothing will fail
)
echo --^> Extracting Discord Modules...
cd modules
7z x *.zip -o* -y >nul
echo --^> Cleaning...
del /s 7z.dll 7z.exe
del /s *.zip
cd ..
rd 7z
echo ==^> Done.
pause