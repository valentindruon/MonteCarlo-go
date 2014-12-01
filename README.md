MonteCarlo-go
=============

A Go implementation of my project Montecarlo

Example of executable to run project
```go
package main

import "fmt"
import "github.com/vdruon/montecarlo"
import "sync"
import "time"

func main() {

  // Scores
  var bank_score, player_score montecarlo.Score
  // WaitGroup
  var wg sync.WaitGroup

  // Time for benchmarking
  start := time.Now()

  for i := 0; i < 1000; i++ {
    wg.Add(1)
    go montecarlo.Launch(&wg, &bank_score, &player_score);
  }
  wg.Wait()

  // Time since start
  elapsed := time.Since(start)

  fmt.Println("Bank score : ", bank_score.Value())
  fmt.Println("Player score : ", player_score.Value())
  rate := float64(bank_score.Value()) / float64(player_score.Value())
  fmt.Printf("Bank win rate : %.2f\n", rate)
  fmt.Printf("Benchmark : %s\n", elapsed)
}
```
