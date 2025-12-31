import {
  Controller,
  Get,
  Query,
  Res,
  Req,
  HttpStatus,
  UnauthorizedException,
} from '@nestjs/common';
import type { Request, Response } from 'express';
import { SpotifyService } from 'src/spotify/spotify';

interface SpotifyUser {
  id: string;
  display_name: string;
  email: string;
  images: { url: string; height: number; width: number }[];
  external_urls: { spotify: string };
}

@Controller('api')
export class AuthController {
  constructor(private readonly spotifyService: SpotifyService) {}

  @Get('callback')
  async callback(
    @Query('code') code: string,
    @Query('error') error: string,
    @Res() res: Response,
  ) {
    const FRONTEND_URL = 'http://localhost:3000';

    if (error) return res.redirect(`${FRONTEND_URL}/?error=access_denied`);

    if (!code)
      return res
        .status(HttpStatus.BAD_REQUEST)
        .json({ message: 'Code not provided' });

    try {
      const tokenData = await this.spotifyService.requestAccessToken(code);

      res.cookie('spotify_access_token', tokenData.access_token, {
        httpOnly: true,
        secure: true,
        sameSite: 'none',
        maxAge: tokenData.expires_in * 1000,
        path: '/',
      });

      return res.redirect(FRONTEND_URL);
    } catch (err) {
      console.error('Erro no callback do Spotify:', err);
      return res.redirect(`${FRONTEND_URL}/?error=auth_failed`);
    }
  }

  @Get('me')
  async getMe(@Req() req: Request): Promise<SpotifyUser> {
    const cookies = req.cookies as Record<string, string | undefined>;
    const token = cookies?.['spotify_access_token'];

    if (!token) {
      throw new UnauthorizedException('Nenhum token encontrado');
    }

    const response = await fetch('https://api.spotify.com/v1/me', {
      headers: { Authorization: `Bearer ${token}` },
    });

    if (!response.ok) {
      throw new UnauthorizedException('Token inválido ou expirado');
    }

    return (await response.json()) as SpotifyUser;
  }
}
