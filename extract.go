package main 

import (
"io"
"net/http"
"os"
"os/exec"
"log"
"fmt"
)

func main() {
	fmt.Println("--> Downloading 7-zip...")
	fileName := "7z.zip"
	URL := "https://www.dropbox.com/s/1g9vchwrpyzoxjx/7z.zip?dl=1" //Downloads 7-zip
	err := downloadFile(URL, fileName)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("--> Extracting the download...")
	cmd := exec.Command("powershell", "-Command", "Expand-Archive", "-Path", "7z.zip", "-DestinationPath", "7z") //Extract the download
	cmd.Stdout = os.Stdout
	cmd.Run()
	
	dir, err := os.Getwd() //gets the currect directory
	if err != nil {
		panic(err)
	}
	cmd = exec.Command("cmd", "/c", "robocopy", "7z", dir, "/mov", ">nul", "&", "rd", "7z") //copy it to the current dir
	cmd.Stdout = os.Stdout
	cmd.Run()
	
	fmt.Println("--> Extracting Discord Modules...")
	cmd = exec.Command("cmd", "/c", "cd", "modules", "&", "7z", "x", "*.zip", "-o*", "-y", ">nul") //extract the modules
	cmd.Stdout = os.Stdout
	cmd.Run()
	
	fmt.Println("--> Cleaning...")
	cmd = exec.Command("cmd", "/c", "del", "/s", "7z.dll", "7z.exe", "7z.zip", ">nul", "&", "cd", "modules", "&", "del", "/s", "*.zip", ">nul") //delete 7zip and the zipped modules 
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("==> Done.")
}

func downloadFile(URL, fileName string) error { //Function to download 7-zip
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		panic("[ ! ] Download of 7 Zip failed, check your connection and try again or download 7 Zip in https://7-zip.org")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
	
}