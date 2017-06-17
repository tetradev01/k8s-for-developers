package main

import (
    "github.com/gin-gonic/gin"
    "os"
    "sync"
    "time"
    "log"
    "math"
)

var healthy = true

func index (c *gin.Context){
    hostname,err := os.Hostname()
    checkErr(err)
    c.String(200,"v3 "+ hostname)
}

func healthz (c *gin.Context){
    if healthy==true {
     c.String(200,"OK")
    }
}

func cancer (c *gin.Context){
     healthy = false
     c.String(500,"NOT_OK")
}


func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

/*******************  MAIN Function **************/
func main(){
  app := gin.Default()
  app.GET("/", index)
  app.GET("/healthz", healthz)
  app.GET("/cancer", cancer)
  app.GET("/thrash",thrash)
  app.Run(":8000")
}
/******************* End MAIN Function **************/




func thrash(c *gin.Context){
    var wg sync.WaitGroup
    start := time.Now()
    for i := 0; i <= 10; i++ {
        wg.Add(1)
        go sqrt(&wg)
    }
    wg.Wait()
    elapsed := time.Since(start)
    log.Printf("Runtime took %s", elapsed)
    c.String(200,"OK")
}

func sqrt(wg *sync.WaitGroup) {
        defer wg.Done()
        x := 0.0001
        for i := 0; i <= 4000000000; i++ {
                x += math.Sqrt(x)
        }
}
