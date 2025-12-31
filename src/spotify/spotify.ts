import { Injectable } from '@nestjs/common';

interface SpotifyAuthResponse {
  access_token: string;
  token_type: string;
  expires_in: number;
  refresh_token: string;
  scope: string;
}

interface SpotifyErrorResponse {
  error: string;
  error_description?: string;
}

@Injectable()
export class SpotifyService {
  private readonly client_id = process.env.SPOTIFY_CLIENT_ID;
  private readonly client_secret = process.env.SPOTIFY_CLIENT_SECRET;
  private readonly redirect_uri = process.env.SPOTIFY_REDIRECT_URI;
  private readonly token_endpoint = 'https://accounts.spotify.com/api/token';

  private get basicAuth() {
    return Buffer.from(`${this.client_id}:${this.client_secret}`).toString(
      'base64',
    );
  }

  async requestAccessToken(code: string): Promise<SpotifyAuthResponse> {
    console.log('CLIENT_ID:', this.client_id);
    console.log('CLIENT_SECRET:', this.client_secret?.substring(0, 4) + '****');
    const response = await fetch(this.token_endpoint, {
      method: 'POST',
      headers: {
        Authorization: `Basic ${this.basicAuth}`,
        'Content-Type': 'application/x-www-form-urlencoded',
      },
      body: new URLSearchParams({
        grant_type: 'authorization_code',
        code,
        redirect_uri: this.redirect_uri ?? '',
      }),
    });

    if (!response.ok) {
      const errorData = (await response.json()) as SpotifyErrorResponse;

      throw new Error(
        `Spotify Auth Error: ${errorData.error_description || errorData.error}`,
      );
    }

    return response.json() as Promise<SpotifyAuthResponse>;
  }
}
