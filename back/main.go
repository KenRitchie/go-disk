package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"go-disk/back/common/tools"
	"go-disk/back/router"
	"time"
)

func main() {

	router.Routes()
	log.Println("handle")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("listen")
	ctx,cancel := context.WithDeadline(context.Background(),time.Now().Add(5 * time.Second))
	go doTimeOutStuff(ctx)
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Println("listenAndServer fail")
		cancel()
	}

}

func doTimeOutStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		//返回绑定当前context的任务被取消的截止时间；
		if deadline, ok := ctx.Deadline(); ok { //设置了deadLine
			fmt.Println("deadline set")
			if time.Now().After(deadline) {
				fmt.Println(ctx.Err().Error())
				//return
			}

		}

		select {
		case <-ctx.Done():
			fmt.Println("finish done")
			err := tools.Open(`http://127.0.0.1:8082/`)
			if err != nil {
				fmt.Println("fail")
			}
			return
		default:
			fmt.Println("work")
		}
	}
}
