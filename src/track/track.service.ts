import { Injectable } from '@nestjs/common';
import { TopTracksDto } from './dto/top-track.dto';
import { TrackDto } from './dto/track.dto';

interface SpotifyArtist {
  name: string;
}
interface SpotifyImage {
  url: string;
}
interface SpotifyTrackItem {
  id: string;
  name: string;
  popularity: number;
  artists: SpotifyArtist[];
  album: { images: SpotifyImage[] };
}
interface SpotifyResponse {
  items: SpotifyTrackItem[];
}

@Injectable()
export class TrackService {
  private async fetchSpotifyTopTracks(
    token: string,
    range: string,
    limit: number,
  ): Promise<TrackDto[]> {
    const response = await fetch(
      `https://api.spotify.com/v1/me/top/tracks?time_range=${range}&limit=${limit}`,
      { headers: { Authorization: `Bearer ${token}` } },
    );

    if (!response.ok) throw new Error('Erro ao buscar dados do Spotify');

    const data = (await response.json()) as SpotifyResponse;

    return data.items.map((item) => ({
      id: item.id,
      title: item.name,
      artist: item.artists.map((a) => a.name).join(', '),
      plays: item.popularity.toString(),
      cover: item.album.images[0]?.url || '',
    }));
  }

  async getTopTracks(token: string, limit: number): Promise<TopTracksDto> {
    const [week, month, year] = await Promise.all([
      this.fetchSpotifyTopTracks(token, 'short_term', limit),
      this.fetchSpotifyTopTracks(token, 'medium_term', limit),
      this.fetchSpotifyTopTracks(token, 'long_term', limit),
    ]);

    return { week, month, year };
  }
}
