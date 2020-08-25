package main

import (
	"context"
	"fmt"
	pb "github.com/kiddy_line_processor_test.git/cmd/grpcserver/proto/github.com/protocolbuffers/protobuf/examples/go/tutorialpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	log.WithFields(log.Fields{
		"gRPC_client" : "start",
	}).Info("Start new client")

	//Enter stream data
	sport_mass := []string{"baseball","football"}
	var time_stream int32
	time_stream = 4

	conn, err := grpc.Dial("kiddy_line_processor_testgit_server_1:55555", grpc.WithInsecure())
	if err != nil {
		log.WithFields(log.Fields{
			"connection" : "kiddy_line_processor_testgit_server_1:55555",
			"error" : err,
		}).Fatal("can not connect with server")

	}
	client := pb.NewSportClient(conn)
	resp, err := http.Get("http://kiddy_line_processor_testgit_kiddy_line_1:8088/ready")
	if err != nil {
		log.WithFields(log.Fields{
			"connection" : "kiddy_line_processor_testgit_server_1:55555",
			"error" : err,
		}).Info("can not connect with server")
		return
	}
	if resp.StatusCode !=200 {
		log.WithFields(log.Fields{
			"StatusCode" : resp.StatusCode,
		}).Fatal("StatusCode error")
	}
	// Stream
	stream, err := client.SubscribeOnSportsLines(context.Background())
	if err != nil {
		log.WithFields(log.Fields{
			"Error" : err,
		}).Fatal("Open stream Error")
	}
	koeff_data := map[string]float32{}
	ctx := stream.Context()
	done := make(chan bool)

	//Request to server
	go func() {
		for {
			req := pb.Request{X:sport_mass, Y: time_stream }
			if err := stream.Send(&req); err != nil {
				log.WithFields(log.Fields{
					"Error" : err,
				}).Fatal("Can not send stream")
			}
		}
		if err := stream.CloseSend(); err != nil {
			log.WithFields(log.Fields{
				"Error" : err,
			}).Fatal("Stream error")
		}
	}()
	//Recieve data from server
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			koeff_data = resp.Result
			log.Println("")
			for _, value := range resp.Xresult {
				fmt.Print(value,":", koeff_data[value]," ")
			}
			fmt.Println()
		}
	}()

	//Close ch if ctx done
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()
	<-done
}
