package main

import (
	"http_test/config"
	"http_test/internal/api"
	"log"
	"sync"
	"time"
)

func run() error {
	// err := godotenv.Load()
	// if err != nil {
	// 	//log.Printf("Error loading .env file")
	// }
	_ = config.Get()
	testClient := api.Get()

	var wg sync.WaitGroup

	cnt := 0
	start := 0
	number := 1
	for i := start; i < (start + number); i++ {
		wg.Add(1)
		if i%20 == 0 {
			// log.Printf("pausa: %d", i)
			time.Sleep(500 * time.Millisecond)
		}
		go func(i int) {
			defer wg.Done()
			err := testClient.UpdFlat(i)
			if err != nil {
				log.Printf("err create i=%d : %s", i, err.Error())
				return
			}
			cnt++
			//log.Printf("upd: %d", i)
		}(i)

	}
	wg.Wait()
	log.Printf("upd flats: %d", cnt)

	// token, err := b3_api.GetToken(0, false)
	// if err != nil {
	// 	log.Printf("i=%d create: %s", 0, err.Error())
	// 	return err
	// }
	// log.Printf("token = %s", token)
	//b3_api.Test()
	return nil
}
func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
