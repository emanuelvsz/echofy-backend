package spotifyrepo

import (
	"context"
	"echofy_backend/src/core/domain/album"
	"echofy_backend/src/core/domain/artist"
	"echofy_backend/src/core/domain/playlist"
	"echofy_backend/src/core/domain/song"
	"echofy_backend/src/core/domain/user"
	"echofy_backend/src/core/errors"
	"echofy_backend/src/core/interfaces/repository"
	"echofy_backend/src/core/messages"
	"time"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
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

func (u UserSpotifyRepository) FindPlaylistByID(playlistID string) (*playlist.Playlist, errors.Error) {
	ctx := context.Background()
	token := getConnection(ctx)

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)

	playlistRow, fetchErr := client.GetPlaylist(ctx, spotify.ID(playlistID))
	if fetchErr != nil {
		return nil, errors.NewUnexpectedError(messages.UnexpectedErrorMessage, fetchErr)
	}

	playlistBuilder := playlist.NewBuilder()
	playlistBuilder.WithID(string(playlistRow.ID)).WithName(playlistRow.Name)
	playlistBuilder.WithSongAmount(playlistRow.Tracks.Total).WithDescription(playlistRow.Description)
	playlistBuilder.WithFollowersAmount(int(playlistRow.Followers.Count))
	playlist, createdEr := playlistBuilder.Build()
	if createdEr != nil {
		return nil, errors.NewUnexpectedError(messages.UnexpectedErrorMessage, createdEr)
	}

	return playlist, nil
}

func (u UserSpotifyRepository) FindSongsByAlbumID(albumID string) ([]song.Song, errors.Error) {
	ctx := context.Background()
	token := getConnection(ctx)

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)
	album, err := client.GetAlbum(ctx, spotify.ID(albumID))
	if err != nil {
		return nil, errors.NewUnexpectedError(messages.UnexpectedErrorMessage, err)
	}

	songs := make([]song.Song, 0)
	for _, track := range album.Tracks.Tracks {
		trackID := string(track.ID)
		trackName := track.Name
		trackNumber := track.TrackNumber
		trackDuration := track.Duration
		trackURL := track.PreviewURL
		trackRelease := album.ReleaseDateTime()

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

func (instance UserSpotifyRepository) FindUserBasicInfo() (*user.User, errors.Error) {
	ctx := context.Background()

	// ! PASTE A VALID ACCESS TOKEN HERE
	token := oauth2.Token{
		AccessToken: "",
	}

	httpClient := spotifyauth.New().Client(ctx, &token)
	client := spotify.New(httpClient)
	userInfo, err := client.CurrentUser(ctx)
	if err != nil {
		return nil, errors.NewUnexpectedError(messages.UnexpectedErrorMessage, err)
	}

	userBuilder := user.NewBuilder().WithID(&userInfo.ID).WithName(userInfo.DisplayName).
		WithEmail(userInfo.Email).WithPictureURL(userInfo.Images[0].URL).WithURI(string(userInfo.ExternalURLs["spotify"]))

	user, err := userBuilder.Build()
	if err != nil {
		return nil, errors.NewUnexpectedError(messages.UnexpectedErrorMessage, err)
	}

	return user, nil
}

func (u UserSpotifyRepository) FindArtistAlbumsByID(artistID string) ([]album.Album, errors.Error) {
	ctx := context.Background()
	token := getConnection(ctx)

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)
	artist, err := client.GetArtistAlbums(ctx, spotify.ID(artistID), nil)
	if err != nil {
		return nil, errors.NewUnexpectedError(messages.UnexpectedErrorMessage, err)
	}

	albums := make([]album.Album, 0)
	for _, single_album := range artist.Albums {
		albumID := single_album.ID
		albumName := single_album.Name

		albumRealeaseDate, err := time.Parse("2006-01-02", single_album.ReleaseDate)
		if err != nil {
			return nil, errors.NewUnexpectedError(messages.UnexpectedErrorMessage, err)
		}

		albumBuilder := album.NewBuilder()
		albumBuilder.WithArtistID(artistID).WithID(albumID.String()).WithName(albumName).WithReleaseDate(albumRealeaseDate)
		albumBuilded, createdError := albumBuilder.Build()
		if createdError != nil {
			return nil, errors.NewUnexpectedError(messages.UnexpectedErrorMessage, createdError)
		}

		albums = append(albums, *albumBuilded)

	}

	return albums, nil

}

func NewUserSpotifyRepository() *UserSpotifyRepository {
	return &UserSpotifyRepository{}
}
