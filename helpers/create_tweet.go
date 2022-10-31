package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Aberry-byte/ThatsOuttaPocket/auth"
	"github.com/g8rswimmer/go-twitter/v2"
)

func Post(text string) {
	authorizer := auth.Bearer_auth{Token: os.Getenv("Bearer_Token")}
	if authorizer.Token == "" {
		fmt.Printf("Invalid\nBearer_Token == %s\n", authorizer.Token)
		os.Exit(1)
	}

	client := twitter.Client{
		Authorizer: authorizer,
		Client:     http.DefaultClient,
		Host:       "https://api.twitter.com",
	}

	twitter_req := twitter.CreateTweetRequest{Text: text}

	tweetResponse, err := client.CreateTweet(context.Background(), twitter_req)
	if err != nil {
		log.Panicf("create tweet error: %v", err)
	}

	enc, err := json.MarshalIndent(tweetResponse, "", "    ")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(enc))
}
