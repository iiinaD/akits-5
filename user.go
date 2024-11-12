package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	proto "main/grpc"
	"os"
	"strconv"
	"strings"
)

var client proto.AuctionServiceClient

func main() {

	port := os.Args[1]
	startUser(&client, port)

	go TakeUserInput()
	for {

	}
}

func startUser(client *proto.AuctionServiceClient, port string) { // start up a new client for the node to send information through the given port
	conn, err := grpc.NewClient("localhost:"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	*client = proto.NewAuctionServiceClient(conn)
}

func TakeUserInput() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		splitText := strings.Split(text, " ")
		if splitText[0] != "Bid" {
			continue
		}
		bidAmount, err := strconv.Atoi(splitText[1])
		if err != nil {
			panic(err)
		}
		response, err := CreateBid(uint64(bidAmount))
		if err != nil {
			panic(err)
		}
		fmt.Println(response.Acknowledgement)
	}
}

func CreateBid(bidamount uint64) (*proto.Reply, error) {
	message := &proto.BidMessage{
		Amount: bidamount,
	}
	return client.Bid(context.Background(), message)
}
