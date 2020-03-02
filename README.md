# RedSense

Simple example of harvesting reddit posts and 
conducting sentiment analysis on them using 
GClouds NLP API.

Its split into a client-server model using GRPC because I originally
had bigger ambitions to connect multiple clients.

### Usage

Get API key from reddit and drop info into the agent file.

>cd redsense && protoc --go_out=plugins=grpc:. *.proto

>go run server/server.go

>go run client/client.go