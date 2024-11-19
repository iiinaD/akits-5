package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	proto "main/grpc"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type AuctionNode struct {
	proto.UnimplementedAuctionServiceServer
	port        string
	logicalTime uint64
	clients     map[string]proto.AuctionServiceClient
	highestBid  uint64
	timeLeft    uint64
}

var mutuallyExclusiveUniversalLock sync.Mutex

func main() {
	clientPort := os.Args[1] //The port of this node

	node := &AuctionNode{
		port:        clientPort,
		clients:     make(map[string]proto.AuctionServiceClient),
		logicalTime: 0,
	}
	go node.startServer()

	if len(os.Args) > 2 {
		joinOnPort := os.Args[2] //The port for this node to join the network on

		node.startClient(joinOnPort)

		message := proto.JoinMessage{
			Port: clientPort,
		}

		response, err := node.clients[joinOnPort].Join(context.Background(), &message) //Trying to join using the first node

		if err != nil {
			panic(err)
		}

		if !response.Success {
			panic("Failed to join cluster")
		}

		// get this nodes ID and start clients up for each of the other nodes in the network
		for _, port := range response.Ports {
			node.startClient(port)
		}
	}
	go node.TakeInputs()
	prev := len(node.clients)
	fmt.Printf("Connected nodes: %d\n", prev+1)
	for {
		// print only when changed
		length := len(node.clients)
		if length != prev {
			fmt.Printf("Connected nodes: %d\n", length+1)
			prev = length
		}

	}
}

func (s *AuctionNode) TakeInputs() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		splitText := strings.Split(text, " ")
		if splitText[0]+" "+splitText[1] == "Start auction" {
			mutuallyExclusiveUniversalLock.Lock()
			textTime, err := strconv.Atoi(splitText[2])
			if err != nil {
				panic(err)
			}
			s.timeLeft = uint64(textTime)
			mutuallyExclusiveUniversalLock.Unlock()
			go s.time()
		}
	}
}

func (s *AuctionNode) startClient(port string) { // start up a new client for the node to send information through the given port
	conn, err := grpc.NewClient("localhost:"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	s.clients[port] = proto.NewAuctionServiceClient(conn)
}

func (s *AuctionNode) startServer() { // start up a new server and listen on the nodes port
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":"+s.port)
	fmt.Println("Created listener")

	if err != nil {
		panic(err)
	}

	proto.RegisterAuctionServiceServer(grpcServer, s)
	fmt.Printf("Now listening on port %s\n", s.port)
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}

// Join Called from node, to make a request to join the network
func (s *AuctionNode) Join(context context.Context, message *proto.JoinMessage) (*proto.JoinResponse, error) {
	//Sends a.
	var ports []string
	for _, client := range s.clients {
		res, err := client.AddNode(context, message)
		if err != nil {
			return nil, err
		}
		ports = append(ports, res.Port)
	}
	res, err := s.AddNode(context, message)
	if err != nil {
		return nil, err
	}
	ports = append(ports, res.Port)
	reply := proto.JoinResponse{
		Ports:   ports,
		Success: true,
	}
	return &reply, nil
}

// AddNode This function adds a client to the node so it can send information to the newly joined node
func (s *AuctionNode) AddNode(context context.Context, message *proto.JoinMessage) (*proto.JoinMessage, error) {
	port := message.Port
	s.startClient(port)
	return &proto.JoinMessage{
		Port: s.port,
	}, nil
}

func (s *AuctionNode) Bid(context context.Context, message *proto.BidMessage) (*proto.Reply, error) {
	reply := proto.Reply{}
	mutuallyExclusiveUniversalLock.Lock()
	if s.highestBid < message.Amount {
		s.highestBid = message.Amount
		s.logicalTime++
		for _, client := range s.clients {
			message := proto.BidUpdateMessage{
				Amount:      s.highestBid,
				TimeLeft:    s.timeLeft,
				LogicalTime: s.logicalTime,
			}
			_, err := client.UpdateBid(context, &message)
			if err != nil {
				return nil, err
			}
		}
		reply.Acknowledgement = fmt.Sprintf("You have the new highest bid at: %d", s.highestBid)
	} else {
		reply.Acknowledgement = "Your bid is too low"
	}
	mutuallyExclusiveUniversalLock.Unlock()
	return &reply, nil
}

func (s *AuctionNode) Result(context context.Context, message *proto.Empty) (*proto.ResultResponse, error) {
	response := &proto.ResultResponse{}
	mutuallyExclusiveUniversalLock.Lock()
	if s.timeLeft <= 0 {
		response.Outcome = fmt.Sprintf("Auction sold for: %d", s.highestBid)
	} else {
		response.Outcome = fmt.Sprintf("Highest bid is: %d", s.highestBid)
	}
	mutuallyExclusiveUniversalLock.Unlock()
	return response, nil
}

func (s *AuctionNode) UpdateBid(context context.Context, message *proto.BidUpdateMessage) (*proto.Reply, error) {
	mutuallyExclusiveUniversalLock.Lock()
	s.highestBid = message.Amount
	s.timeLeft = message.TimeLeft
	s.logicalTime = message.LogicalTime
	mutuallyExclusiveUniversalLock.Unlock()

	reply := proto.Reply{Acknowledgement: "OK"}
	return &reply, nil
}

func (s *AuctionNode) time() {
	for s.timeLeft > 0 {
		mutuallyExclusiveUniversalLock.Lock()
		s.timeLeft -= 1
		fmt.Printf("Time left: %d\n", s.timeLeft)
		mutuallyExclusiveUniversalLock.Unlock()
		time.Sleep(1000 * time.Millisecond)
	}
}
