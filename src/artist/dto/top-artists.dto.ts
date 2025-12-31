import { ArtistDto } from './artist.dto';

export class TopArtistsDto {
  week: ArtistDto[];
  month: ArtistDto[];
  year: ArtistDto[];
}
