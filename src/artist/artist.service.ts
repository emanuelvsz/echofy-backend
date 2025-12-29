import { Injectable } from '@nestjs/common';
import { TopArtistDto } from './dto/top-artist.dto';

@Injectable()
export class ArtistService {
  getTopArtist(): TopArtistDto {
    return {
      name: 'Kendrick Lamar',
      plays: 142,
    };
  }
}
