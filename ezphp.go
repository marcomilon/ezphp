package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

func main() {
    fmt.Printf("EzPHP\n")
    createDirIfNotExist("php")
    
    fileUrl := "https://windows.php.net/downloads/releases/php-7.2.3-nts-Win32-VC15-x64.zip"
    filePath := "php/php-7.2.3-nts-Win32-VC15-x64.zip"
    
    err := DownloadFile(filePath, fileUrl)
    if err != nil {
        panic(err)
    }
}

func createDirIfNotExist(dir string) {
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        err = os.MkdirAll(dir, 0755)
        if err != nil {
            panic(err)
        }
    }
}

func DownloadFile(filepath string, url string) error {

    // Create the file
    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    if err != nil {
        return err
    }

    return nil
}