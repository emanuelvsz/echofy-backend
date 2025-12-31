import {
  Controller,
  Get,
  Req,
  Query,
  UnauthorizedException,
  ParseIntPipe,
  DefaultValuePipe,
} from '@nestjs/common';
import type { Request } from 'express';
import { ArtistService } from './artist.service';

@Controller('artist')
export class ArtistController {
  constructor(private readonly artistService: ArtistService) {}

  @Get('top-artists')
  async getTopArtists(
    @Req() req: Request,
    @Query('limit', new DefaultValuePipe(10), ParseIntPipe) limit: number,
  ) {
    const cookies = req.cookies as Record<string, string | undefined>;
    const token = cookies['spotify_access_token'];

    if (!token) {
      throw new UnauthorizedException('Sessão expirada. Faça login novamente.');
    }

    return this.artistService.getTopArtists(token, limit);
  }
}
