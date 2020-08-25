package main

import (
	pb "github.com/kiddy_line_processor_test.git/cmd/grpcserver/proto/github.com/protocolbuffers/protobuf/examples/go/tutorialpb"
	"github.com/kiddy_line_processor_test.git/database"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"math"
	"net"
	"sort"
	"time"
)

type server struct{}
var mas_req []string
var key int8

var koeff_data = map[string]float32{}
var koeff_data_req = map[string]float32{}

func (s server) SubscribeOnSportsLines(srv pb.Sport_SubscribeOnSportsLinesServer) error {

	err := database.Connect("postgres", "1234", "web")
	if err != nil {
		log.WithFields(log.Fields{
			"database" : "error",
			"error" : err,
		}).Fatal("database connection error")
	}
	log.WithFields(log.Fields{
		"gRPC_server" : "start",
	}).Info("start new server")

	ctx := srv.Context()
	for {
		// exit if context is done
		// or continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// receive data from stream
		req, err := srv.Recv()
		if err == io.EOF {
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}
		time.Sleep(time.Duration(req.Y) *time.Second)
		req_data := req.X
		sort.Strings(req.X)
		sort.Strings(mas_req)

		key = 0
		if len(req.X) == len(mas_req) {
			for i, value := range req.X {
				if value != mas_req[i]{
					key = 1
					break
				}
			}
		} else {
			key = 1
		}

		if key == 1 {
			for _, value := range req.X {
				koeff_data[value] = database.Select_sport(value)
				koeff_data_req[value] = koeff_data[value]
			}
		} else {
			for _, value := range req.X {
				koeff_data[value] = database.Select_sport(value)
				delta := float64(koeff_data[value] - koeff_data_req[value])
				koeff_data_req[value] = koeff_data[value]
				koeff_data[value] = float32(math.Round(delta*1000)/1000)
			}
		}

		resp := pb.Response{Result: koeff_data, Xresult: req_data }
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
		mas_req = req.X
	}
}

func main() {
	// create listiner
	lis, err := net.Listen("tcp", ":55555")
	if err != nil {
		log.WithFields(log.Fields{
			"Create_listener" : "error",
			"error" : err,
		}).Fatal("failed to listen")
	}

	// create grpc server
	s := grpc.NewServer()
	pb.RegisterSportServer(s, server{})

	// and start...
	if err := s.Serve(lis); err != nil {
		log.WithFields(log.Fields{
			"Start_server" : "error",
			"error" : err,
		}).Fatal("Start error")
	}
}
