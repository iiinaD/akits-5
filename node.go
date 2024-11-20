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
	leader      bool
}

var mutuallyExclusiveUniversalLock sync.Mutex
var electionLock sync.Mutex
var electionRunning bool

func main() {
	var clientPort string
	if len(os.Args) < 2 {
		clientPort = "5050"
	} else {
		clientPort = os.Args[1] //The port of this node
	}

	node := &AuctionNode{
		port:        clientPort,
		clients:     make(map[string]proto.AuctionServiceClient),
		logicalTime: 0,
		leader:      clientPort == "5050",
	}
	go node.startServer()

	if len(os.Args) > 1 {
		node.startClient("5050")

		message := proto.JoinMessage{
			Port: clientPort,
		}

		response, err := node.clients["5050"].Join(context.Background(), &message) //Trying to join using the first node

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
	if node.port != "5050" {
		go node.WatchLeaderPulse()
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
		if strings.ToLower(splitText[0]+" "+splitText[1]) == "start auction" {
			mutuallyExclusiveUniversalLock.Lock()
			textTime, err := strconv.Atoi(splitText[2])
			if err != nil {
				panic(err)
			}
			s.timeLeft = uint64(textTime)
			go s.time()
			mutuallyExclusiveUniversalLock.Unlock()
			message := proto.TimeMessage{
				Time: uint64(textTime),
			}
			for _, client := range s.clients {
				_, err := client.StartAuction(context.Background(), &message)
				if err != nil {
					panic(err)
				}
			}
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
		s.timeLeft -= 1
		if s.port == "5050" {
			fmt.Printf("Time left: %d\n", s.timeLeft)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func (s *AuctionNode) WatchLeaderPulse() {
	for s.port != "5050" {
		response, err := s.clients["5050"].CheckPulse(context.Background(), &proto.Empty{})
		if err != nil {
			fmt.Printf("Node on port: %s detected a leader crash\n", s.port)
			s.StartElection()
			return
		}
		if response.Acknowledgement != "OK" {
			panic("got response but isn't OK.")
		} else {
			fmt.Println("Got OK from leader")
		}
		time.Sleep(2000 * time.Millisecond / 2)
	}
}

func (s *AuctionNode) StartElection() {
	fmt.Printf("Node on port: %s started an election\n", s.port)
	electionLock.Lock()
	if !electionRunning {
		_, err := s.clients["5050"].CheckPulse(context.Background(), &proto.Empty{})
		if err != nil {
			electionRunning = true
		} else {
			electionLock.Unlock()
			return
		}
		electionLock.Unlock()
	} else {
		electionLock.Unlock()
		return
	}
	minPortString := s.port
	minPort, err := strconv.Atoi(minPortString)
	if err != nil {
		panic(err)
	}
	for port, _ := range s.clients {
		portNumber, err := strconv.Atoi(port)
		if err != nil {
			panic(err)
		}
		if portNumber == 5050 {
			continue
		}
		if portNumber < minPort {
			minPort = portNumber
			minPortString = port
		}
	}
	if minPortString == s.port {
		fmt.Printf("Node on port: %s is running the election\n", s.port)
		_, err := s.RunElection(context.Background(), &proto.Empty{})
		if err != nil {
			return
		}
	} else {
		_, err := s.clients[minPortString].RunElection(context.Background(), &proto.Empty{})
		if err != nil {
			return
		}
	}
}

func (s *AuctionNode) RunElection(context context.Context, _ *proto.Empty) (*proto.Empty, error) {
	minLogicalTime := s.logicalTime
	minLogicalTimePort := s.port
	for port, client := range s.clients {
		if port == "5050" {
			continue
		}
		timeMessage, err := client.GetLogicalTime(context, &proto.Empty{})
		if err != nil {
			return nil, err
		}
		if timeMessage.Time < minLogicalTime {
			minLogicalTime = timeMessage.Time
			minLogicalTimePort = port
		}
	}
	electionResult := &proto.PortMessage{
		Port: minLogicalTimePort,
	}
	for port, client := range s.clients {
		if port == "5050" {
			continue
		}
		_, err := client.SetLeader(context, electionResult)
		if err != nil {
			return nil, err
		}
	}
	_, err := s.SetLeader(context, electionResult)
	if err != nil {
		return nil, err
	}
	return &proto.Empty{}, nil
}

func (s *AuctionNode) SetLeader(context context.Context, message *proto.PortMessage) (*proto.Empty, error) {
	port := message.Port
	if port == s.port {
		s.port = "5050"
		go s.startServer()
	} else {
		fmt.Printf("Node previously on port: %s is the new leader.\n", port)
		delete(s.clients, port)
		go s.WatchLeaderPulse()
	}
	return &proto.Empty{}, nil
}

func (s *AuctionNode) GetLogicalTime(context context.Context, message *proto.Empty) (*proto.TimeMessage, error) {
	reply := &proto.TimeMessage{
		Time: s.logicalTime,
	}
	return reply, nil
}

func (s *AuctionNode) CheckPulse(context context.Context, message *proto.Empty) (*proto.Reply, error) {
	reply := &proto.Reply{
		Acknowledgement: "OK",
	}
	return reply, nil
}

func (s *AuctionNode) StartAuction(context context.Context, message *proto.TimeMessage) (*proto.Empty, error) {
	mutuallyExclusiveUniversalLock.Lock()
	s.timeLeft = message.Time
	mutuallyExclusiveUniversalLock.Unlock()
	go s.time()
	return &proto.Empty{}, nil
}
