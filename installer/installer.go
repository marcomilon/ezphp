package installer

import (
    "archive/zip"
    "path/filepath"
    "fmt"
    "io"
    "net/http"
    "os"
)

const (
    PhpDir = "php"
    PhpDownloadUrl = "https://windows.php.net/downloads/releases/php-7.2.3-nts-Win32-VC15-x64.zip"
    PhpZipFile = "php/php-7.2.3-nts-Win32-VC15-x64.zip"
    PhpExecutable = "php.exe";
)

func Install() (string, error){
    fmt.Printf("Installing PHP\n")
    fmt.Printf("Creating PHP Directory\n")
    
    err := createDirIfNotExist(PhpDir)
    if err != nil {
        return "", err
    }
    
    fmt.Printf("Downloading PHP\n")
    err = downloadFile(PhpZipFile, PhpDownloadUrl)
    if err != nil {
        return "", err
    }
    
    fmt.Printf("Unziping PHP\n")
    unZipErr := unzip(PhpZipFile, PhpDir);
    if unZipErr != nil {
        return "", err
    }
    
    path, err := filepath.Abs(filepath.Dir(PhpDir))
    if err != nil {
        return "", err
    }

    completePath := path + string(os.PathSeparator) + PhpDir + string(os.PathSeparator) + PhpExecutable
    
    if _, err := os.Stat(completePath); os.IsNotExist(err) {
        if err != nil {
            return "", err
        }
    }
    
    return completePath, nil
    
}

func createDirIfNotExist(dir string) error {
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        err = os.MkdirAll(dir, 0755)
        if err != nil {
            return err
        }
    }
    
    return nil
}

func downloadFile(filepath string, url string) error {
    
    if _, err := os.Stat(filepath); err == nil {
        return nil
    }
    
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

func unzip(src, dest string) error {
    r, err := zip.OpenReader(src)
    if err != nil {
        return err
    }
    defer func() {
        if err := r.Close(); err != nil {
            panic(err)
        }
        }()
        
        os.MkdirAll(dest, 0755)
        
        // Closure to address file descriptors issue with all the deferred .Close() methods
        extractAndWriteFile := func(f *zip.File) error {
            rc, err := f.Open()
            if err != nil {
                return err
            }
            defer func() {
                if err := rc.Close(); err != nil {
                    panic(err)
                }
                }()
                
                path := filepath.Join(dest, f.Name)
                
                if f.FileInfo().IsDir() {
                    os.MkdirAll(path, f.Mode())
                    } else {
                        os.MkdirAll(filepath.Dir(path), f.Mode())
                        f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
                        if err != nil {
                            return err
                        }
                        defer func() {
                            if err := f.Close(); err != nil {
                                panic(err)
                            }
                            }()
                            
                            _, err = io.Copy(f, rc)
                            if err != nil {
                                return err
                            }
                        }
                        return nil
                    }
                    
                    for _, f := range r.File {
                        err := extractAndWriteFile(f)
                        if err != nil {
                            return err
                        }
                    }
                    
                    return nil
                }