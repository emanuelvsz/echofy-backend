package response

import "echofy_backend/src/core/domain/playlist"

type PlaylistDTO struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	SongAmount      int    `json:"song_amount"`
	FollowersAmount int    `json:"followers_amount"`
}

func NewPlaylistDTO(playlist playlist.Playlist) *PlaylistDTO {
	return &PlaylistDTO{
		ID:              playlist.ID(),
		Name:            playlist.Name(),
		Description:     playlist.Description(),
		SongAmount:      playlist.SongAmount(),
		FollowersAmount: playlist.FollowersAmount(),
	}
}
