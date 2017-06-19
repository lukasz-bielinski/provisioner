package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

var flWait = flag.Int("wait", envInt("LOOP_WAIT", 5), "number of seconds to wait before next sync")

func envInt(key string, def int) int {
	if env := os.Getenv(key); env != "" {
		val, err := strconv.Atoi(env)
		if err != nil {
			log.Printf("invalid value for %q: using default: %q", key, def)
			return def
		}
		return val
	}
	return def
}

func printFile(ignoreDirs []string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Print(err)
			return nil
		}
		if info.IsDir() {
			dir := filepath.Base(path)
			for _, d := range ignoreDirs {
				if d == dir {
					return filepath.SkipDir
				}
			}
		}
		fmt.Println(path)
		fi, err := os.Stat(path)
		if err != nil {
			panic(err)
		}
		if fi.IsDir() {
			fmt.Println("it's a directory")
		} else {
			fmt.Println("it's not a directory")
			h := md5.New()
			f, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			if _, err := io.Copy(h, f); err != nil {
				log.Fatal(err)
			}
			os.Stdout.WriteString(hex.EncodeToString(h.Sum(nil)))

		}
		return nil
	}
}

// func findFiles() string {
// 	searchDir := "/src"
// 	fileList := []string{}
// 	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
// 		fileList = append(fileList, path)
// 		return nil
// 	})
// 	_ = err
// 	for _, file := range fileList {
// 		fmt.Println(file)
// 		// h := md5.New()
// 		// f, err := os.Open(file)
// 		// if err != nil {
// 		// 	log.Fatal(err)
// 		// }
// 		// defer f.Close()
// 		// if _, err := io.Copy(h, f); err != nil {
// 		// 	log.Fatal(err)
// 		// }
// 		// os.Stdout.WriteString(hex.EncodeToString(h.Sum(nil)))
// 	}
// 	return ""
// }

func main() {
	flag.Parse()
	if _, err := exec.LookPath("kubectl"); err != nil {
		log.Fatalf("required kubectl executable not found: %v", err)
	}

	for {
		log.SetFlags(log.Lshortfile)
		dir := "/src/"
		ignoreDirs := []string{".hg", ".git"}
		err := filepath.Walk(dir, printFile(ignoreDirs))
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Duration(*flWait) * time.Second)
	}
}
