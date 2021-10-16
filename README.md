thist - a go package for calculating online histograms with plotting to the terminal and images
===============================================================================================
[![Documentation](https://godoc.org/github.com/botond-sipos/thist?status.svg)](http://godoc.org/github.com/botond-sipos/thist)
[![Go Report Card](https://goreportcard.com/badge/github.com/botond-sipos/thist)](https://goreportcard.com/report/github.com/botond-sipos/thist)

This is a maintained fork of [bsipos/thist](https://github.com/bsipos/thist).

Check out the `watch` subcommands from the [csvtk](https://github.com/shenwei356/csvtk) and [seqkit](https://github.com/shenwei356/seqkit) tools to see the code in action.

Example
-------

```go
package main

import (
        "fmt"
        "github.com/bsipos/thist"
        "math/rand"
        "time"
)

// randStream return a channel filled with endless normal random values
func randStream() chan float64 {
        c := make(chan float64)
        go func() {
                for {
                        c <- rand.NormFloat64()
                }
        }()
        return c
}

func main() {
        // create new histogram
        h := thist.NewHist(nil, "Example histogram", "auto", -1, true)
        c := randStream()

        i := 0
        for {
                // add data point to hsitogram
                h.Update(<-c)
                if i%50 == 0 {
                        // draw histogram
                        fmt.Println(h.Draw())
                        time.Sleep(time.Second)
                }
                i++
        }
}

```

[![demo video](http://img.youtube.com/vi/7mrs1QGDyys/0.jpg)](http://www.youtube.com/watch?v=7mrs1QGDyys)

TODO
----

- Add more details on online histogram generation.
- Add separate object for online estimation of moments.
- Maybe add tcell as a back-end?
