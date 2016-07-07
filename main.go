package main

import(
  "fmt"
  "time"
  "math"
)

func main() {
  t := time.NewTicker(1 * time.Second)
  for {
    select {
    case <- t.C:
      var specified_time time.Time = time.Date(2020,7,5,15,14,0,0,time.Local)
      duration := -time.Since(specified_time)
      duration_all_sec := int64(math.Ceil(float64(duration) / 1000000000))
      duration_day := math.Ceil(float64(duration_all_sec) / (60 * 60 * 24))
      duration_hour := math.Ceil(float64(duration_all_sec % (60 * 60 * 24)) / (60 * 60))
      duration_min := math.Ceil(float64(duration_all_sec % (60 * 60)) / 60)
      duration_sec := duration_all_sec % 60
      fmt.Println(duration_day,"日",duration_hour,"時間",duration_min,"分",duration_sec,"秒")
    }
  }

}
