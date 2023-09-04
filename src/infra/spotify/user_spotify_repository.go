package spotifyrepo

import (
	"context"
	"echofy_backend/src/core/domain/artist"
	"echofy_backend/src/core/domain/song"
	"echofy_backend/src/core/errors"
	"echofy_backend/src/core/interfaces/repository"
	"echofy_backend/src/core/messages"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

var _ repository.UserLoader = &UserSpotifyRepository{}

type UserSpotifyRepository struct{}

func (u UserSpotifyRepository) FindSongsByPlaylistID(playlistID string) ([]song.Song, errors.Error) {
	ctx := context.Background()
	token := getConnection(ctx)

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)
	playlistItems, err := client.GetPlaylistItems(ctx, spotify.ID(playlistID))
	if err != nil {
		return nil, errors.NewUnexpectedError(messages.UnexpectedErrorMessage, err)
	}

	songs := make([]song.Song, 0)
	for _, item := range playlistItems.Items {
		track := item.Track.Track
		trackID := string(track.ID)
		trackName := track.Name
		trackNumber := track.TrackNumber
		trackDuration := track.Duration
		trackURL := track.PreviewURL
		trackRelease := track.Album.ReleaseDateTime()

		artists := make([]artist.Artist, 0)
		for _, each := range track.Artists {
			artistID := each.ID
			artistName := each.Name
			artistURI := each.URI

			artistBuilder := artist.NewBuilder()
			artistBuilder.WithID(string(artistID)).WithName(artistName).WithSpotifyURL((*string)(&artistURI))
			artistBuilded, createdError := artistBuilder.Build()
			if createdError != nil {
				return nil, errors.NewUnexpectedError(messages.UnexpectedErrorMessage, createdError)
			}

			artists = append(artists, *artistBuilded)
		}

		songBuilder := song.NewBuilder()
		songBuilder.WithID(trackID).WithName(trackName).WithTrackNumber(trackNumber)
		songBuilder.WithDuration(trackDuration).WithSpotifyURL(trackURL).WithReleaseDate(trackRelease)
		songBuilder.WithArtists(artists)
		songBuilded, createdError := songBuilder.Build()
		if createdError != nil {
			return nil, errors.NewUnexpectedError(messages.UnexpectedErrorMessage, createdError)
		}

		songs = append(songs, *songBuilded)
	}

	return songs, nil
}

func (u UserSpotifyRepository) FindArtistByID(artistID string) (*artist.Artist, errors.Error) {
	ctx := context.Background()
	token := getConnection(ctx)

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)

	artistRow, fetchError := client.GetArtist(ctx, spotify.ID(artistID))
	if fetchError != nil {
		return nil, errors.NewUnexpectedError(messages.UnexpectedErrorMessage, fetchError)
	}

	artistBuilder := artist.NewBuilder()
	artistBuilder.WithID(string(artistRow.ID)).WithName(artistRow.Name).WithImageURL(&artistRow.Images[0].URL)
	artistBuilder.WithSpotifyURL((*string)(&artistRow.URI))
	artist, createError := artistBuilder.Build()
	if createError != nil {
		return nil, errors.NewUnexpectedError(messages.UnexpectedErrorMessage, createError)
	}

	return artist, nil
}

func NewUserSpotifyRepository() *UserSpotifyRepository {
	return &UserSpotifyRepository{}
}
