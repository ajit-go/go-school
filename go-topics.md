---
# Testing
## Running test Testing
    Go test -v
    Go test -v -cover
## Docs in tests with Output: keyword 
   myGame_test.go
     ExampleMyFucn() { 
         MyFunc("call 1")
         MyFunc("call 2")

        -> following is outout kernel which generates docs + matches actual output of above calls to what is documented chronologically below
         //output: 
     }
   Generating docs: godoc -http ":8082"
Benchmarking
    myGame_test.go
        BenchmarkMyFucn(b *tester.B) { for i< b.N .... etc } Go automatically determines value of N to run iterations
    Go test -bench=.
## t.Run("testName), func(){....}
## t.Helper()
    Helper functions so that test fails returns actual line of failure
Running go program ?
packaging / artifact ? 
---
# Type conversion vs type assertion 
   type conversion for basic types
   assertion is when an empty interface is defined (object of java), to get the underlaying type e.g. Emp e = i.type(i)??

---
## Go routines, channels and synching
    calls to channels are bloking calls can be used to synch routines
     always run tests with go test -v -race . will output if there is a race issue
    sharing data between go routines, one way is global object
    synch/atomic if atomic.addUint64  method is used to modify shared data and atomic.loadUint64
### Channles
    use make to create channel
    channel could be in/out or listen only 
    a function blocks when it listens for a chennels values
    fucntion 
### select 
    like switch on channels
    wait on channels indefinitely - like poll git changes for leadup project and push data to db
    if there is data on channel do that case else move to next case
    selct { case }
    #### time.Afte: 
    case <- time.After(time.millisecond * 500)
    #### time.NewTicker 
    https://gobyexample.com/tickers
    send time to ticker channel every 500 milliseconds
### range on channel

### synch 
    https://gobyexample.com/channel-synchronization
    https://gobyexample.com/waitgroups
    ### synch.mutex lock and unloack

### links
    https://wiki.hybris.com/display/prodandtech/Cloud+Lab%3A+Concurrent+programming+made+easy+with+GoLang
    https://gobyexample.com/channels
    https://golang.org/doc/effective_go.html#channels
    https://www.golangprograms.com/goroutines.html

    synch: https://stackoverflow.com/questions/46906647/wait-result-of-multiple-goroutines
    https://stackoverflow.com/questions/16590778/wait-for-the-termination-of-n-goroutines?rq=1
---- finish routines ---

## custom errors

## web server in go 
### http package to create a simple server
### go get github.com/gorrilla/mux
### Rest/json https://tutorialedge.net/golang/creating-restful-api-with-golang/
### rest test 
#### REST Testing with Mux:
    https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql 

---
# Concurrency in details
## buffered channel 
Concepts

- buffered channels 
- through-put control 


Introduction material:
- tour.golang.org -> Concurrency  > buffered channels
- tour.golang.org -> Concurrency  > Default Selection
- https://gobyexample.com/channel-buffering
- https://gobyexample.com/non-blocking-channel-operations
- https://gobyexample.com/worker-pools
- https://gobyexample.com/timeouts
- https://gobyexample.com/waitgroups
- https://gobyexample.com/rate-limiting
- https://gobyexample.com/timers
- https://gobyexample.com/tickers

Task:
- Take and rewrite some of those above as Example Tests

Further reading (for all channel steps)
- Space Invaders with GoRoutines https://github.com/kenlomaxhybris/spaceInvadersWithGoRoutines

---
# defer
    like finally of java, executed after the parent function
    exception must be handled in this
    the parameters are evealuated at call time even though 

---
# panic and recover

---
# method vs functions
   func is a stand alone function
   method is adding behaviour to a custom type ( a struct)  or alias to inbuilt types e.g. type myFloat float64
   e.g.
   typ s struct { lat int}
   func (mys s) F(){mys.xxxxx}

   s1 := 

---

---
# pointers 
  & vs * in method parameters ???


---
# Databases
## mongo: https://github.com/mongodb/mongo-go-driver

---
---
Links:
1. All topics: https://www.golangprograms.com/goroutines.html
2. gobyexample: https://gobyexample.com/defer
   links to go play-ground
3. All topics with executable example . www.tour.golang.org on the right is list of all topics
