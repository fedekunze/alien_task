# Aliens!

Golang coding task for Cosmos

### Assumptions

- File provided can only have __.txt__ format
- A fight happens at the moment that an alien moves to another city and encounters to another alien, no matter how many aliens are on the city
- When a city is destroyed, it:
  - Destroys all the roads from it to other cities, as well as the roads from any other city to it. In particular, that means it all the status of the roads to `destoyed = true`.
  - Kills all the aliens in the destroyed city (sets their status to `alive = false`)

## Instalation

To run the Alien app you'll have to previously download [Go](https://golang.org/) and set your `$GOPATH`.
After you download Golang, install the app by running:

```
go get -u github.com/fedekunze/alien_task
```

## Usage

Once installed you can run the program directly from the command line:

```
cd $GOPATH/src/github.com/fedekunze/alien_task
go build
```

### Run Aliens

You can run the program by running the following command on your terminal:

```
alien_task --file=<path_to_map.txt> -N=<total_number_of_aliens>
```

For simplicity, the file privided with the map *MUST* have a `.txt` format.
You can provide a full path to the file (__e.g__ `/Users/<usename>/Desktop/map.txt`) or a relative path to the file on the same folder that you're running the program (__e.g__ `map.txt`)

## Test App

Run tests for existing types and logic of the program by typing:

```
cd $GOPATH/src/github.com/fedekunze/alien_task/cosmos
go test cosmos_test.go
go test types_test.go
```
