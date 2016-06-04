# go-async

Library helper to produce parallel execution
- with well defined set of tasks
- without explicitly using channel

https://godoc.org/github.com/mh-cbon/go-async

# Install

```sh
glide get github.com/mh-cbon/go-async
```

# Usage

#### async.Parallel

```go
package main

import (
  "fmt"
  "time"
  "errors"
  "github.com/mh-cbon/go-async"
)


func main () {

  fmt.Println("start")

  results := async.Parallel([]async.Task{
    func (done async.Done) {
      time.Sleep(time.Second * 1)
      fmt.Println("work one")
      done(nil)
    },
    func (done async.Done) {
      time.Sleep(time.Second * 1)
      fmt.Println("work two")
      done(errors.New("can't work with 42"))
    },
  })

  if async.HasErrors(results) {
    for _, res := range results {
      if res.Err!=nil {
        fmt.Printf("Task #%d: %s\n", res.Index, res.Err)
      }
    }
  }

  fmt.Println("end")

}

```


#### async.ParallelLimit

```go
package main

import (
  "fmt"
  "time"
  "errors"
  "github.com/mh-cbon/go-async"
)


func main () {

  fmt.Println("start")

  results, _ := async.ParallelLimit(2, []async.Task{
    func (done async.Done) {
      time.Sleep(time.Second * 1)
      fmt.Println("work one")
      done(nil)
    },
    func (done async.Done) {
      time.Sleep(time.Second * 1)
      fmt.Println("work two")
      done(errors.New("can't work with 42"))
    },
  })

  if async.HasErrors(results) {
    for _, res := range results {
      if res.Err!=nil {
        fmt.Printf("Task #%d: %s\n", res.Index, res.Err)
      }
    }
  }

  fmt.Println("end")

}

```
