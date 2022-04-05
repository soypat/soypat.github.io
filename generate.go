package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//go:generate go run .

func main() {
	var (
		compiler = "go"
	)
	flag.StringVar(&compiler, "c", compiler, "compiler. Either `tinygo` or `go`")
	flag.Parse()
	genWasm(compiler)
	genJSBootloader(compiler)
	genIndexHTML(compiler)
	log.Println("generation finished. enjoy your day!")
}

func genJSBootloader(compiler string) {
	root := "GOROOT"
	if compiler == "tinygo" {
		root = "TINYGOROOT"
	}
	rootEnv, err := exec.Command(compiler, "env", root).Output()
	if err != nil {
		log.Fatal(err)
	}
	var f string
	switch compiler {
	case "tinygo":
		f = filepath.Join(strings.TrimSuffix(string(rootEnv), "\n"), "targets", "wasm_exec.js")
	case "go":
		f = filepath.Join(strings.TrimSpace(string(rootEnv)), "misc", "wasm", "wasm_exec.js")
	default:
		panic("unknown compiler " + compiler)
	}

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

func genWasm(compiler string) {
	args := []string{"build", "-o=main.wasm"}
	if compiler == "tinygo" {
		args = append(args, "-target=wasm")
	}
	cmd := exec.Command(compiler, append(args, ".")...)
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
	info, err := os.Stat("main.wasm")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("generate WASM finished (main.wasm %s)\n", formatSize(info.Size()))
}

func genIndexHTML(compiler string) {
	fp, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	switch compiler {
	case "go":
		_, err = fp.WriteString(goIndexHTML)
	case "tinygo":
		_, err = fp.WriteString(tinygoIndexHTML)
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Println("generated index.html")
}

const goIndexHTML = `<!DOCTYPE html>
<!-- Polyfill for the old Edge browser -->
<script src="https://cdn.jsdelivr.net/npm/text-encoding@0.7.0/lib/encoding.min.js"></script>
<script src="wasm_exec.js"></script>
<script>
(async () => {
  const resp = await fetch('main.wasm');
  if (!resp.ok) {
    const pre = document.createElement('pre');
    pre.innerText = await resp.text();
    document.body.appendChild(pre);
  } else {
    const src = await resp.arrayBuffer();
    const go = new Go();
    const result = await WebAssembly.instantiate(src, go.importObject);
    go.argv = [];
    go.run(result.instance);
  }
  const reload = await fetch('_wait');
  // The server sends a response for '_wait' when a request is sent to '_notify'.
  if (reload.ok) {
    // location.reload();
  }
})();
</script>`

const tinygoIndexHTML = `<!DOCTYPE html>
<script src="wasm_exec.js"></script><script>
(async () => {
  const resp = await fetch('main.wasm');
  if (!resp.ok) {
    const pre = document.createElement('pre');
    pre.innerText = await resp.text();
    document.body.appendChild(pre);
    return;
  }
  const src = await resp.arrayBuffer();
  const go = new Go();
  const result = await WebAssembly.instantiate(src, go.importObject);
  go.run(result.instance);
})();
</script>`

func formatSize(size int64) (format string) {
	const (
		Byte     = 1
		Kilobyte = 1000 * Byte
		MegaByte = 1000 * Kilobyte
		GigaByte = 1000 * MegaByte
	)
	switch {
	case size > 10*GigaByte:
		format = fmt.Sprintf("%dGB", size/GigaByte)
	case size > 10*MegaByte:
		format = fmt.Sprintf("%dMB", size/MegaByte)
	case size > 10*Kilobyte:
		format = fmt.Sprintf("%dkB", size/Kilobyte)
	default:
		format = fmt.Sprintf("%dB", size)
	}
	return format
}
