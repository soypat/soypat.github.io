package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//go:generate go run .

func main() {
	genWasm()
	genJSBootloader()
	log.Println("generation finished. enjoy your day!")
}

func genJSBootloader() {
	out, err := exec.Command("go", "env", "GOROOT").Output()
	if err != nil {
		log.Fatal(err)
	}
	f := filepath.Join(strings.TrimSpace(string(out)), "misc", "wasm", "wasm_exec.js")
	fp, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	jsb, err := os.Create("wasm_exec.js")
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(jsb, fp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("js bootloader generated.")
}

func genWasm() {
	cmd := exec.Command("go", "build", "-o=main.wasm", ".")
	cmd.Dir = "./app"
	cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(string(output), err)
	}
	err = os.Rename("app/main.wasm", "main.wasm")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("generate WASM finished.")
}
