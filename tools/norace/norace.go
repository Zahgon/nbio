package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	root, _ = filepath.Abs("./")

	skipPaths = []string{
		"_test.go",
		"norace.go",
	}

	chTask = make(chan func(), 32)
)

func main() {
	defer close(chTask)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for f := range chTask {
				f()
			}
		}()
	}

	wg := &sync.WaitGroup{}

	walk(wg, root)

	// wg.Done()
	fmt.Println("wait")
	wg.Wait()

	fmt.Println("exit")
}

func run(f func()) { _ = "STUB: not implemented"; return }

func walk(wg *sync.WaitGroup, currRoot string) { _ = "STUB: not implemented"; return }

func shouldSkip(path string) bool { _ = "STUB: not implemented"; return false }

func addNorace(path string, info os.FileInfo) { _ = "STUB: not implemented"; return }
