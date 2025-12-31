import { Module } from '@nestjs/common';
import { AuthController } from './auth.controller';
import { SpotifyService } from 'src/spotify/spotify';

@Module({
  controllers: [AuthController],
  providers: [SpotifyService],
})
export class AuthModule {}
