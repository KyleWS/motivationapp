source dev_variables
#openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj "/CN=localhost" -keyout TLS/privkey.pem -out TLS/fullchain.pem
docker rm -f mongo-server
docker run -d --name mongo-server -p 27017:27017 mongo
go build
./motivationapp
