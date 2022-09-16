package model

import (
	"context"
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

type TwitterInfo struct {
	Name     string
	Username string
}

func verifyTwitter(names []string) map[string]TwitterInfo {
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
	var resp map[string]TwitterInfo
	for _, d := range dictionaries {
		var newInfo TwitterInfo
		newInfo.Name = d.User.Name
		newInfo.Username = d.User.UserName
		resp[newInfo.Username] = newInfo
	}
	//return list of verified accounts
	return resp
}

func ValidateTwitter(data []*EventData) {
	var entries []string
	var targets []*EventData
	//load entries
	for _, d := range data {
		id, status := d.ValidateSnsTwitter()
		if status != NG && len(id) > 0 {
			entries = append(entries, id)
			targets = append(targets, d)
		}
	}
	//get list of verified accounts in entries
	//accounts := verifyTwitter(entries)
	//
	//for _, d := range targets {
	//
	//
	//}
}
