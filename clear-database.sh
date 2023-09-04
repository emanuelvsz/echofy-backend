cd config/docker
docker compose rm -sf && docker compose stop
cd ../..
source run-project.sh
sudo pkill docker