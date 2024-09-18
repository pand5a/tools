package util

import (
	"context"
	"fmt"
	"log"
	"time"
)

func F(ctx context.Context, mt string, a ...any) func(...any) {
	start := time.Now()	
	v := ctx.Value("code9")
	if v == nil {
		return func(...any){}
	}

	return func(b ...any) {
		elapsed := time.Since(start)
		s := fmt.Sprintf("elapsed %10s, %30s, %s", elapsed, fmt.Sprintf(mt, a...), b)
		log.Output(2, s)
	}
}
