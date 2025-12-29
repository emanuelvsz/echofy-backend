import { Test, TestingModule } from '@nestjs/testing';
import { ArtistController } from './artist.controller';
import { ArtistService } from './artist.service';

describe('ArtistController', () => {
  let controller: ArtistController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [ArtistController],
      providers: [ArtistService],
    }).compile();

    controller = module.get<ArtistController>(ArtistController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
