package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Second
const longDuration 	= 1 * time.Minute

func main1() {
	ctx, cancel := context.WithTimeout(context.Background(), longDuration)
	defer cancel()

	select{
	case <- time.After(2*time.Second):
		fmt.Println("finished after 1 second")
	case <- ctx.Done():
		fmt.Println(ctx.Err())
	}
}
// short duration -> context deadline exceeded
// long duration  -> finished after 1 second