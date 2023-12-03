package main

import (
	"flag"
	"log"
)

func main() {
	t = mustToken()
	// tgClient = telegram.New(token)
	// fetcher = fetcher.New(tgClient)
	// processor = processor.New(tgClient)

	// consumer.start(fetcher, processor)
}

func mustToken() string {
	token := flag.String("token", "", "token for access tg api")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
}
