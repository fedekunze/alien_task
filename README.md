# Aliens!

Golang coding task for Cosmos

### Assumptions

- Total number of aliens provided (`N`) is less than the total amount of cities
- File provided can only have __.txt__ format
- A fight happens at the moment that an alien moves to another city and encounters to another alien
- When a city is destroyed, it:
  - Destroys all the roads from it to other cities, as well as the roads from any other city to it. In particular, that means it deletes pointers to any `Road`
  - Kills all the aliens in the destroyed city (sets their status to `alive = false`)


## Instalation

To run the Alien app you'll have to previously download Go and set your `$GOPATH`.
After you download Golang, install the app by running:

```
go get -u github.com/fedekunze/alien_task
```

## Usage

Once installed you can run the program directly from the command line:

### Run Aliens

You can run the program by running the following command on your terminal:

```
aliens run --file=<path_to_map.txt> -N=<total_number_of_aliens>
```

For simplicity, the file privided with the map *MUST* have a `.txt` format.
You can provide a full path to the file (__e.g__ `/Users/federico/Desktop/map.txt`) or a relative path to the file on the same folder that you're running the program (__e.g__ `map.txt`)

## Test App

Run tests for existing types and logic of the program by typing:

```
aliens test
```
