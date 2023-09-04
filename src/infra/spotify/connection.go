package spotifyrepo

import (
	"context"
	"log"

	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func getConnection(ctx context.Context) *oauth2.Token {
	config := &clientcredentials.Config{
		ClientID:     "3348f4b437614e8b9c742c305eb9865b",
		ClientSecret: "cd3d51d52a724d18aac6d3910534420c",
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	return token
}
