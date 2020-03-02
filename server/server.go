package main

import (
	"context"
	pb "github.com/AnthonyLaiuppa/redsense/redsense"
	"github.com/pkg/errors"
	grpc "google.golang.org/grpc"
	language "cloud.google.com/go/language/apiv1"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
	"log"
	"net"
)

const port = ":50051"

type redsenseService struct{}

func (s *redsenseService) Generate(ctx context.Context, in *pb.MediaPage) (*pb.MediaAnalysis, error) {
	log.Printf("Received MediaPage from %v with title %v", in.Source, in.Title)
	sentiment := analyzeSentiment(in.Body)

	return &pb.MediaAnalysis{
		Source: in.Source,
		Url: in.Url,
		Title: in.Title,
		Body: in.Body,
		Datetime: in.Datetime,
		Sentiment: sentiment,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterRedsenseServiceServer(server, &redsenseService{})
	if err := server.Serve(lis); err != nil {
		log.Fatal(errors.Wrap(err, "Failed to start server!"))
	}
}

func analyzeSentiment(text string) string{
	ctx := context.Background()

	// Creates a client.
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the text to analyze.
	t := text

	// Detects the sentiment of the text.
	sentiment, err := client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: t,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		log.Fatalf("Failed to analyze text: %v", err)
	}


	if sentiment.DocumentSentiment.Score >= 0 {
		log.Println("Sentiment: positive")
		return "Positive"
	} else {
		log.Println("Sentiment: negative")
		return "Negative"
	}
}