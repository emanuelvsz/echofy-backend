// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package bridge

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID         uuid.UUID
	RoleID     uuid.UUID
	Name       string
	Email      string
	Hash       string
	Password   string
	LastLogin  time.Time
	PictureUrl sql.NullString
}

type Album struct {
	ID          uuid.UUID
	Name        string
	ArtistID    uuid.UUID
	ReleaseDate time.Time
	Description sql.NullString
	ImageUrl    sql.NullString
}

type AlbumGenre struct {
	ID      uuid.UUID
	AlbumID uuid.UUID
	GenreID uuid.UUID
}

type Artist struct {
	ID              uuid.UUID
	SuperArtistID   uuid.NullUUID
	Name            string
	Description     sql.NullString
	FoundedAt       time.Time
	TerminatedAt    sql.NullTime
	ImageUrl        sql.NullString
	RecordCompanyID uuid.NullUUID
	CountryID       uuid.NullUUID
	SpotifyUrl      sql.NullString
}

type ArtistGenre struct {
	ID       uuid.UUID
	ArtistID uuid.UUID
	GenreID  uuid.UUID
}

type ArtistGroup struct {
	ID            uuid.UUID
	ArtistID      uuid.UUID
	SuperArtistID uuid.UUID
	JoinedAt      time.Time
	LeftAt        sql.NullTime
	IsActive      bool
}

type ArtistNews struct {
	NewsID   uuid.UUID
	ArtistID uuid.UUID
}

type Category struct {
	ID          uuid.UUID
	Name        string
	Description sql.NullString
}

type City struct {
	ID      uuid.UUID
	Name    string
	StateID uuid.NullUUID
}

type CompanyCountry struct {
	CountryID       uuid.UUID
	RecordCompanyID uuid.UUID
}

type Country struct {
	ID   uuid.UUID
	Name string
}

type FollowedArtist struct {
	ID           uuid.UUID
	ArtistID     uuid.UUID
	AccountID    uuid.UUID
	FollowedAt   time.Time
	UnfollowedAt sql.NullTime
}

type Genre struct {
	ID          uuid.UUID
	Name        string
	Description sql.NullString
	CreatedAt   sql.NullTime
}

type LikedSong struct {
	SongID    uuid.UUID
	AccountID uuid.UUID
	LikedAt   time.Time
	UnlikedAt sql.NullTime
}

type News struct {
	ID          uuid.UUID
	Title       string
	ReleasedAt  time.Time
	Description sql.NullString
	UpdatedAt   sql.NullTime
	CreatedAt   time.Time
}

type NewsCategory struct {
	NewsID     uuid.UUID
	CategoryID uuid.UUID
}

type RecordCompany struct {
	ID         uuid.UUID
	Name       string
	FoundedAt  time.Time
	WebsiteUrl sql.NullString
	CountryID  uuid.NullUUID
}

type ResetPassword struct {
	AccountID uuid.UUID
	Email     string
	Code      string
}

type Role struct {
	ID   uuid.UUID
	Name string
	Code string
}

type Song struct {
	ID          uuid.UUID
	Name        string
	AlbumID     uuid.NullUUID
	ReleaseDate time.Time
	Duration    string
	Lyrics      sql.NullString
	TrackNumber sql.NullInt32
	SpotifyUrl  sql.NullString
}

type SongArtist struct {
	ArtistID uuid.UUID
	SongID   uuid.UUID
}

type SongGenre struct {
	ID      uuid.UUID
	SongID  uuid.UUID
	GenreID uuid.UUID
}

type State struct {
	ID        uuid.UUID
	Name      string
	CountryID uuid.NullUUID
}
