Write-Output "--> Downloading 7-zip..."
Invoke-WebRequest -Uri https://www.dropbox.com/s/1g9vchwrpyzoxjx/7z.zip?dl=1 -OutFile 7z.zip
Expand-Archive -Path 7z.zip -DestinationPath 7z
Remove-Item 7z -Recurse -Force
Remove-Item 7z.zip -Force
if (Test-Path -Path $pwd/modules) {
	robocopy 7z $pwd\modules /mov
} else {
	"[ ! ] Looks like you didn't downloaded the modules, exiting now"
	Remove-Item 7z.zip -Force
	Remove-Item 7z -Recurse
	Exit
}
"--> Extracting Discord Modules..."
cd .\"modules"
7z x *.zip -o* -y
"--> Cleaning..."
Get-ChildItem | Where-Object Name -Like '7z*' | ForEach-Object { Remove-Item -Recurse -LiteralPath $_.Name }
Get-ChildItem * -Include *.zip -Recurse | Remove-Item
"==> Done."
Exit