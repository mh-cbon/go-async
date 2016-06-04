package main

import (
  "fmt"
  // "os"
  "time"
  "errors"

  "../"
)


func main () {
  fmt.Println("start")
  results := async.Parallel([]async.Task{
    func (done async.Done) {
      fmt.Println("work one")
      done(errors.New("42"))
    },
    func (done async.Done) {
      time.Sleep(time.Second * 1)
      fmt.Println("work two")
      done(errors.New("can't work with 42"))
    },
    func (done async.Done) {
      time.Sleep(time.Second * 1)
      fmt.Println("work three")
      done(nil)
    },
    func (done async.Done) {
      time.Sleep(time.Second * 1)
      fmt.Println("work four")
      done(nil)
    },
    func (done async.Done) {
      time.Sleep(time.Second * 1)
      fmt.Println("work five")
      done(nil)
    },
  })
  if async.HasErrors(results) {
    for _, res := range results {
      if res.Err!=nil {
        fmt.Printf("Task #%d: %s\n", res.Index, res.Err)
      }
    }
    // os.Exit(1)
  } else {
    fmt.Println("all done")
  }
  fmt.Println("end")
}
