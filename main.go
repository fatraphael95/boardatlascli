package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/fatraphael95/boardatlascli/api"
)

func main() {
	// bga --name "ticket to ride" --clientID GsLThh5vKR --skip 10 --limit 5
	// Define the command line arguments
	query := flag.String("query", "", "Boardgame name to search")
	clientId := flag.String("clientId", "", "Boardgame Atlas client_id")
	limit := flag.Uint("limit", 10, "Limits the number of results returned")
	skip := flag.Uint("skip", 0, "Skips the number of results provided")
	timeout := flag.Uint("timeout", 10, "Timeout")

	// Parse the command line
	flag.Parse()

	// fmt.Printf("query=%s, clientId=%s, limit=%d, skip=%d\n", *query, *clientId, *limit, *skip)
	if isNull(*query) {
		log.Fatalln("Please use --query to set the boardgame name to search")
	}

	if isNull(*clientId) {
		log.Fatalln("Please use --clientId to set the Boardgame Atlas client id")
	}

	fmt.Printf("query%s, clientId=%s, limit=%d, skip=%d\n", *query, *clientId, *limit, *skip)

	// Instatiate a BoardgameAtlas struct
	bga := api.New(*clientId)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout*uint(time.Second)))
	defer cancel() // defer will allow any functions to be called

	// time.Sleep(5 * time.Second) // uncomment this and set --timeout=1 in CLI

	// Make the invocation
	result, err := bga.Search(ctx, *query, *limit, *skip)
	if nil != err {
		log.Fatalf("Cannot search for boardgame: %v", err)
	}

	// Colors
	boldGreen := color.New(color.Bold).Add(color.FgHiGreen).SprintFunc()

	for i, g := range result.Games {
		fmt.Printf("Index No. %d\n", i)
		fmt.Printf("%s %s\n", boldGreen("Name:"), g.Name)
		fmt.Printf("%s %s\n", boldGreen("Description:"), g.Description)
		fmt.Printf("%s %s\n", boldGreen("Url:"), g.Url)
	}
}

func isNull(s string) bool {
	return len(strings.TrimSpace(s)) <= 0
}
