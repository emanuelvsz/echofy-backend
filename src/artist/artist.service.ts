import { Injectable } from '@nestjs/common';
import { ArtistDto } from './dto/artist.dto';
import { TopArtistsDto } from './dto/top-artists.dto';
interface SpotifyArtistItem {
  id: string;
  name: string;
  popularity: number;
  images: { url: string }[];
}

interface SpotifyArtistResponse {
  items: SpotifyArtistItem[];
}

@Injectable()
export class ArtistService {
  private async fetchSpotifyTopArtists(
    token: string,
    range: string,
    limit: number,
  ): Promise<ArtistDto[]> {
    const response = await fetch(
      `https://api.spotify.com/v1/me/top/artists?time_range=${range}&limit=${limit}`,
      { headers: { Authorization: `Bearer ${token}` } },
    );

    if (!response.ok) throw new Error('Erro ao buscar artistas do Spotify');

    const data = (await response.json()) as SpotifyArtistResponse;

    return data.items.map(
      (item) =>
        ({
          id: item.id,
          name: item.name,
          popularityPercentage: `${item.popularity}%`,
          photo: item.images[0]?.url || '',
        }) as ArtistDto,
    );
  }

  async getTopArtists(token: string, limit: number): Promise<TopArtistsDto> {
    const [week, month, year] = await Promise.all([
      this.fetchSpotifyTopArtists(token, 'short_term', limit),
      this.fetchSpotifyTopArtists(token, 'medium_term', limit),
      this.fetchSpotifyTopArtists(token, 'long_term', limit),
    ]);

    return { week, month, year };
  }
}
