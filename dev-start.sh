source dev_variables
docker rm -f mongo-server
docker run -d --name mongo-server -p 27017:27017 mongo
go build
./motivationapp
