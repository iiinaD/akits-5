# akits-5

## How to run the program
All of the following commands should be run from the root directory.
To run the program, start off by starting your leader, that is done by running the following command:
```sh
go run node.go
```
To start more nodes run the following command
```sh
go run node.go <port>
```
replace _\<port\>_ with an unused port. (We use 5050 for the leader).
To start a client run
```sh
go run user.go
```

To start the auction run the following in the leader terminal
```sh
start auction <number>
```
replace _\<number>_ with the amount of seconds the auction should run for
It is then possible to bid from the clients or get result by the two following commands:
```sh
bid <number>
```
```sh
result
```