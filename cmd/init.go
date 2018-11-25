package cmd

import (
	"encoding/json"
	"sync"

	"github.com/marhaupe/json-to-struct/internal"
)

func Start() {
	c := make(chan json.Token)
	var wg sync.WaitGroup
	wg.Add(2)
	go internal.Lex(`{ 
		"Hallo": "Hey", 
		"DasisteinTest": { 
			"Schoen": true
			 } 
		}`, c, &wg)
	go internal.Parse(c, &wg)
	wg.Wait()
}
