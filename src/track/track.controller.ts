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
import { TrackService } from './track.service';

@Controller('track')
export class TrackController {
  constructor(private readonly trackService: TrackService) {}

  @Get('top-tracks')
  async getTopTracks(
    @Req() req: Request,
    @Query('limit', new DefaultValuePipe(10), ParseIntPipe) limit: number,
  ) {
    const cookies = req.cookies as Record<string, string | undefined>;
    const token = cookies['spotify_access_token'];

    if (!token) {
      throw new UnauthorizedException('Sessão expirada. Faça login novamente.');
    }

    return this.trackService.getTopTracks(token, limit);
  }
}
