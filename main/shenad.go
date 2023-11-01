//go:generate $GOPATH/bin/statik -src=../dist -dest . -f
//go:generate go fmt statik/statik.go
package main

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/rakyll/statik/fs"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"shenad/main/opts"
	_ "shenad/main/statik"
	"strings"
)

var (
	Stderr io.Writer = os.Stderr
	Stdout io.Writer = os.Stdout
)

type Config struct {
	ServerHost string
	ConfigPath string
}

var config *Config

func main() {

	parseFlags()

	statikFS, err := fs.New()

	if err != nil {
		log.Fatalf(err.Error())
	}

	http.FileServer(statikFS)

	var mux = http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	opts.WakeOpts(opts.InitOptions{
		Mux: mux,
	})

	go func(serverAddr string, m *http.ServeMux) {
		if err = http.ListenAndServe(serverAddr, m); err != nil {
			fmt.Println("Can not start server:", err)
			os.Exit(-1)
		}
	}(config.ServerHost, mux)

	// for test do not open the browser automatic
	// trigger it on conf.toml
	if !opts.LoadConfig(config.ConfigPath).App.Test {
		openPage()
	}

	opts.InitDB()

	handleSignals()

}

func parseFlags() {
	config = &Config{}
	flag.StringVar(&config.ServerHost, "p", "0.0.0.0:20000", "local server address")
	flag.StringVar(&config.ConfigPath, "c", "./conf.toml", "config file path")
	flag.Usage = func() {
		fmt.Fprintf(Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
}

func handleSignals() {
	go handleFileSysCall()
	// 确保最后监听系统调用事件
	handleSysCall()
}

func handleSysCall() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}

func handleFileSysCall() {
	var confPath = config.ConfigPath
	fmt.Println("config path:", confPath)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					opts.TriggerReload(confPath)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(confPath)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func openPage() {
	url := fmt.Sprintf("http://%v", config.ServerHost)
	fmt.Println("shenad console open", url, "in browser")
	var err error
	switch runtime.GOOS {
	case "linux":
		err = runCmd("xdg-open", url)
	case "darwin":
		err = runCmd("open", url)
	case "windows":
		r := strings.NewReplacer("&", "^&")
		err = runCmd("cmd", "/c", "start", r.Replace(url))
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		fmt.Println(err)
	}
}

func runCmd(prog string, args ...string) error {
	cmd := exec.Command(prog, args...)
	cmd.Stdout = Stdout
	cmd.Stderr = Stderr
	return cmd.Run()
}
