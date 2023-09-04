package playlist

type Playlist struct {
	id              string
	name            string
	description     string
	songAmount      int
	followersAmount int
}

func (instance Playlist) ID() string {
	return instance.id
}

func (instance Playlist) Name() string {
	return instance.name
}

func (instance Playlist) Description() string {
	return instance.description
}

func (instance Playlist) SongAmount() int {
	return instance.songAmount
}

func (instance Playlist) FollowersAmount() int {
	return instance.followersAmount
}