package main

import (
	"context"
	"fmt"
	pb "github.com/AnthonyLaiuppa/redsense/redsense"
	"github.com/turnage/graw/reddit"
	"google.golang.org/grpc"
	"log"
	"strings"
	"time"
)

const address = "localhost:50051"

func shipIt(out pb.MediaPage)(*pb.MediaAnalysis,error){
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewRedsenseServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	r, err := c.Generate(ctx, &out)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Media from %s\n Has sentiment %s\n", r.Sentiment, r.Url)

	return r,nil
}


func readdit(){

	bot, err := reddit.NewBotFromAgentFile("../config/example.agent", 0)
	if err != nil {
		fmt.Println("Failed to create bot handle: ", err)
		return
	}

	harvest, err := bot.Listing("/r/investing", "")
	if err != nil {
		fmt.Println("Failed to fetch /r/investing: ", err)
		return
	}

	for _, post := range harvest.Posts[:20] {

		if containsSubs(post.SelfText){
			out := pb.MediaPage{
				Source: "reddit.com/r/investing",
				Url:post.URL,
				Title:post.Title,
				Body:post.SelfText,
				Datetime:post.CreatedUTC,
			}
			shipIt(out)
		}
	}

}

func containsSubs(text string) bool{
	subs := []string{"MSFT", "SPY", "TSLA"}
	for _, sub := range subs {
		if strings.Contains(text, sub) {
			return true
		}
	}
	return false
}

func main() {
	readdit()
}