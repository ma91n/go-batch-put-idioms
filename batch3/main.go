package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ma91n/batchidioms"
)

func main() {
	ctx := context.Background()
	forums := batchidioms.LoadForums()

	for i := 0; i < len(forums); i += 25 {
		end := i + 25
		if end > len(forums) {
			end = len(forums)
		}

		if err := batchidioms.BatchWrite(ctx, forums[i:end]); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("finished")
}
