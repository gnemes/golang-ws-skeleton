#/bin/sh

ROOTPATH=/go/src/github.com/GoIntegro/go5-reco-api/
PROGRAM_PATH=/go/src/github.com/GoIntegro/go5-reco-api/Infrastructure/Commander/main

docker-compose exec api go build -o $PROGRAM_PATH $ROOTPATH/Infrastructure/Commander/main.go
docker-compose exec api chmod +x $PROGRAM_PATH $ROOTPATH

docker-compose exec api $PROGRAM_PATH "indexUsers"
