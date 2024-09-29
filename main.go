package main

import (
	"log"
	"main/src/engine"
	"net/http"
	_"net/http/pprof"
	"fmt"
    "os"
    
)

func main() {
	var e engine.Engine

	e.Init()
	e.Load()
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	e.Run()
	e.Unload()
	e.Close()
}
func e() {
    engine := engine.NewEngine()
   
    if len(os.Args) > 1 {
        command := os.Args[1]
        switch command {
        case "help", "git-h":
            engine.StateMenu = engine.HELP  
        case "start":
            engine.StateMenu = engine.HOME 
        case "quit":
            os.Exit(0)  
        default:
            fmt.Println("Commande inconnue :", command)
            os.Exit(1)
        }
    }

    engine.Run() 
}
