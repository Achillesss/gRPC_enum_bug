package daysTest

import (
	pb "../daysproto"
	srv "../server"
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"testing"
)

type test struct {
	t          *testing.T
	testserver *srv.A
	srv        *grpc.Server
	cc         *grpc.ClientConn
}

var (
	serverAddr = flag.String("addr", "127.0.0.1:20010", "The server address in the format of host:port")
	port       = flag.Int("port", 20010, "The server port")
)

func (t *test) startServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	var opts = []grpc.ServerOption{}
	t.srv = grpc.NewServer(opts...)
	t.testserver = srv.NewServer()
	pb.RegisterGetWeekDayServiceServer(t.srv, t.testserver)
	go t.srv.Serve(lis)
}

func (t *test) clientConn() *grpc.ClientConn {
	flag.Parse()
	if t.cc != nil {
		return t.cc
	}
	var opts = []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	return conn
}
func getDays(client pb.GetWeekDayServiceClient, day *pb.Day) *pb.DayResponse {
	status, err := client.GetWeekDay(context.Background(), day)
	if err != nil {
		fmt.Printf("%v getDays error:%v", client, err)
	}
	return status
}
func TestStartServer(te *testing.T) {
	flag.Parse()
	var t = test{}
	te.Logf("Test start\n")
	t.startServer()
	client := pb.NewGetWeekDayServiceClient(t.clientConn())
	var day = &pb.Day{Day: pb.Day_Monday}
	te.Logf("Monday0: %v\tIsWeekDay: %v\n", *day, *getDays(client, day))
	day = &pb.Day{Day: pb.Day_WeekDay(5)}
	te.Logf("Monday1: %v\tIsWeekDay: %v\n", *day, *getDays(client, day))
	day = &pb.Day{Day: pb.Day_WeekDay(10000)}
	te.Logf("Expecting an error when WeekDay=10000: %v", *getDays(client, day))
}
