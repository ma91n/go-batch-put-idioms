package main

import (
	"context"
	"fmt"
	"log"

	. "github.com/ma91n/batchidioms"
)

func main() {
	ctx  := context.Background()
	forums := LoadForums()

	batch := make([]Forum, 0, 25)
	for i, v := range forums {

		batch = append(batch, v) // 1行枚にスライスに追加

		if len(batch) >= 25 || i == len(forums) -1 { // 25個になったか、最終行の場合
			if err := BatchWrite(ctx, batch); err != nil {
				log.Fatal(err)
			}
			batch = batch[:0] // スライスをクリア
		}
	}

	fmt.Println("finished")
}
