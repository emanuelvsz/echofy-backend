import { Injectable } from '@nestjs/common';
import { TopTracksDto } from './dto/top-track.dto';

@Injectable()
export class TrackService {
  getTopTracks(): TopTracksDto {
    return {
      week: [
        {
          id: 1,
          title: 'Not Like Us',
          artist: 'Kendrick Lamar',
          plays: '142',
          cover:
            'https://i.scdn.co/image/ab67616d0000b2731ea0c62b2339cbf493a999ad',
        },
        {
          id: 2,
          title: 'Espresso',
          artist: 'Sabrina Carpenter',
          plays: '98',
          cover:
            'https://i.scdn.co/image/ab67616d0000b273659cd4673230913b3918e0d5',
        },
      ],
      month: [],
      year: [],
    };
  }
}
