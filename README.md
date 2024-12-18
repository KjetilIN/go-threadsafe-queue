# Go Thread-Safe Queue

Group members: 
- Kjetil K. Indrehus, kjetiki@uio.no
- Thomas N. Kristiansen, thomank@uio.no

A thread safe queue implemented in Go. Concurrency in Go happens through message passing with channels:

More specifically, the task is to implement a simplified
map-reduce system that computes the sum of all squares up to a certain integer, using a thread-safe linked queue and thread pools. 


### Usage

Run the source code by changing the directory to the `scr` folder:

```terminal
cd src/
```

Then run the go code with: 

```terminal
go run .
```

### Performance 

Max number to square and add: `1_000_000`

Execution times:
- 1.670913 seconds
- 1.609310 seconds
- 1.681716 seconds
- 1.705196 seconds
- 1.615626 seconds
- 1.602665 seconds

