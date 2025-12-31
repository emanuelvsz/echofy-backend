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

interface SpotifyAudioFeature {
  id: string;
  energy: number;
  valence: number;
}

interface SpotifyAudioFeaturesResponse {
  audio_features: (SpotifyAudioFeature | null)[];
}

interface SpotifyRecentlyPlayedItem {
  track: SpotifyTrackItem;
  played_at: string;
}

interface SpotifyRecentlyPlayedResponse {
  items: SpotifyRecentlyPlayedItem[];
}

export interface HistoryResponseDto {
  daily: TrackDto[];
  weekly: TrackDto[];
  monthly: TrackDto[];
  yearly: TrackDto[];
  all: TrackDto[];
}

@Injectable()
export class TrackService {
  private getVibeLabel(energy: number, valence: number): string {
    if (energy > 0.7 && valence > 0.7) return 'Eufórica';
    if (energy > 0.7 && valence < 0.4) return 'Agressiva';
    if (energy < 0.4 && valence > 0.7) return 'Chill / Relax';
    if (energy < 0.4 && valence < 0.4) return 'Melancólica';
    return 'Equilibrada';
  }

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

  private async fetchAudioFeatures(
    token: string,
    trackIds: string[],
  ): Promise<SpotifyAudioFeaturesResponse> {
    const ids = trackIds.join(',');
    const response = await fetch(
      `https://api.spotify.com/v1/audio-features?ids=${ids}`,
      { headers: { Authorization: `Bearer ${token}` } },
    );
    return (await response.json()) as SpotifyAudioFeaturesResponse;
  }

  async getTopTracks(token: string, limit: number): Promise<TopTracksDto> {
    const [week, month, year] = await Promise.all([
      this.fetchSpotifyTopTracks(token, 'short_term', limit),
      this.fetchSpotifyTopTracks(token, 'medium_term', limit),
      this.fetchSpotifyTopTracks(token, 'long_term', limit),
    ]);

    return { week, month, year };
  }

  async getOnLoopTrack(token: string): Promise<TrackDto> {
    const tracks = await this.fetchSpotifyTopTracks(token, 'short_term', 1);
    if (!tracks || tracks.length === 0) {
      throw new Error('Nenhuma música encontrada no período recente.');
    }
    return tracks[0];
  }

  async getHistory(token: string): Promise<HistoryResponseDto> {
    const [dailyRaw, weekly, monthly, yearly] = await Promise.all([
      fetch(`https://api.spotify.com/v1/me/player/recently-played?limit=20`, {
        headers: { Authorization: `Bearer ${token}` },
      }).then((res) => res.json() as Promise<SpotifyRecentlyPlayedResponse>),
      this.fetchSpotifyTopTracks(token, 'short_term', 20),
      this.fetchSpotifyTopTracks(token, 'medium_term', 20),
      this.fetchSpotifyTopTracks(token, 'long_term', 20),
    ]);

    const daily: TrackDto[] = dailyRaw.items.map((item) => ({
      id: item.track.id,
      title: item.track.name,
      artist: item.track.artists.map((a) => a.name).join(', '),
      plays: 'Recente',
      cover: item.track.album.images[0]?.url || '',
    }));

    const allTracks = [...daily, ...weekly, ...monthly, ...yearly];
    const uniqueIds = Array.from(new Set(allTracks.map((t) => t.id)));

    const featuresResponse = await this.fetchAudioFeatures(token, uniqueIds);

    const featuresMap = new Map<string, SpotifyAudioFeature>(
      (featuresResponse.audio_features || [])
        .filter((f): f is SpotifyAudioFeature => f !== null)
        .map((f) => [f.id, f]),
    );

    const enrich = (tracks: TrackDto[]): TrackDto[] =>
      tracks.map((track) => {
        const feat = featuresMap.get(track.id);
        return {
          ...track,
          energy: feat ? Math.round(feat.energy * 100) : 50,
          vibe: feat
            ? this.getVibeLabel(feat.energy, feat.valence)
            : 'Desconhecida',
          loyalty: yearly.find((t) => t.id === track.id)
            ? 'Alta'
            : 'Nova Obsessão',
        };
      });

    return {
      daily: enrich(daily),
      weekly: enrich(weekly),
      monthly: enrich(monthly),
      yearly: enrich(yearly),
      all: enrich(yearly),
    };
  }
}
