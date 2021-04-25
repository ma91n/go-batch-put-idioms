package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ma91n/batchidioms"
)

func main() {
	ctx := context.Background()
	loadForums := batchidioms.LoadForums()

	batch := make([]batchidioms.Forum, 0, 25)
	for _, v := range loadForums {

		batch = append(batch, v) // 1行枚にスライスに追加

		if len(batch) >= 25 {
			if err := batchidioms.BatchWrite(ctx, batch); err != nil {
				log.Fatal(err)
			}
			batch = batch[:0] // スライスをクリア
		}
	}

	if len(batch) > 0 { // 25の剰余が1~24の場合の救済
		if err := batchidioms.BatchWrite(context.Background(), batch); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("finished")
}
