package model

import (
	"context"
	"encoding/json"
	"fmt"
	twitter "github.com/g8rswimmer/go-twitter/v2"
	"log"
	"net/http"
	"os"
)

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func validateTwitter(names []string) {
	fmt.Println("first", names[0])
	client := &twitter.Client{
		Authorizer: authorize{
			Token: os.Getenv("TwitterToken"),
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}
	opts := twitter.UserLookupOpts{
		Expansions: []twitter.Expansion{twitter.ExpansionPinnedTweetID},
	}

	fmt.Println("Callout to user lookup callout")

	userResponse, err := client.UserNameLookup(context.Background(), names, opts)
	print(names)
	if err != nil {
		log.Panicf("user lookup error: %v", err)
	}

	dictionaries := userResponse.Raw.UserDictionaries()

	enc, err := json.MarshalIndent(dictionaries, "", "    ")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(enc))
}

func ValidateTwitter(data []*EventData) {
	var usernames []string
	fmt.Println("Usernames starts from here")
	for _, d := range data {
		id, status := d.ValidateSnsTwitter()
		if status != NG && len(id) > 0 {
			fmt.Println(id)
			usernames = append(usernames, id)
		}
	}
	validateTwitter(usernames)
}
