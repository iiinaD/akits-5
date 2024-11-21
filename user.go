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
	"time"
)

var client proto.AuctionServiceClient

func main() {

	//port := os.Args[1]
	startUser(&client, "5050")

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
	fmt.Println("Connected to auction.")
	fmt.Println("Type \"bid <amount>\" to bid or \"result\" to see the current highest bid or the result of the auction.")
}

func TakeUserInput() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		splitText := strings.Split(text, " ")
		if strings.ToLower(splitText[0]) == "bid" {
			bidAmount, err := strconv.Atoi(splitText[1])
			if err != nil {
				panic(err)
			}
		createbid:
			response, err := CreateBid(uint64(bidAmount))
			if err != nil {
				time.Sleep(250 * time.Millisecond)
				goto createbid
			}
			fmt.Println(response.Acknowledgement)
		} else if strings.ToLower(splitText[0]) == "result" {
			response, err := GetResult()
			if err != nil {
				panic(err)
			}
			fmt.Println(response.Outcome)
		} else {
			fmt.Println("Command not recognized. Try again!")
		}
	}
}

func GetResult() (*proto.ResultResponse, error) {
	return client.Result(context.Background(), &proto.Empty{})
}

func CreateBid(bidamount uint64) (*proto.Reply, error) {
	message := &proto.BidMessage{
		Amount: bidamount,
	}
	return client.Bid(context.Background(), message)
}
