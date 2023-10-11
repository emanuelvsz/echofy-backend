echo "\
+----------------------+
| Starting database... |
+----------------------+\
"

cd config/docker
docker compose rm -sf && docker compose up --build -d
cd ../..

echo "\
+-------------------------------------+
| Downloading project dependencies... |
+-------------------------------------+\
"

go mod tidy

echo "\
+---------------------------------+
| Generating API documentation... |
+---------------------------------+\
"

bash -c "cd src/app/api && swag init -g ../../main.go --output ./docs --dir ./endpoints/handlers"

echo "\
+-----------------+
| Starting API... |
+-----------------+\
"

go run ./src/app/api/main.go
