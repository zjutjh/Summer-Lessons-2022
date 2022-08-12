package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now) // 2022-08-10 19:58:24.190927 +0800 CST m=+0.000226710
	t := time.Date(2022, 8, 10, 19, 58, 24, 0, time.UTC)
	t2 := time.Date(2022, 9, 11, 21, 5, 24, 0, time.UTC)
	fmt.Println(t)                                                  // 2022-08-10 19:58:24 +0000 UTC
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()) // 2022 August 10 19 58
	fmt.Println(t.Format("2006-01-02 15:04:05"))                    // 2022-08-10 19:58:24
	diff := t2.Sub(t)
	fmt.Println(diff)                           // 1h7m0s
	fmt.Println(diff.Minutes(), diff.Seconds()) // 67 4020
	t3, err := time.Parse("2006-01-02 15:04:05", "2022-08-10 19:58:23")
	if err != nil {
		panic(err)
	}
	fmt.Println(t3 == t)    // true
	fmt.Println(now.Unix()) // 1660132932
}
