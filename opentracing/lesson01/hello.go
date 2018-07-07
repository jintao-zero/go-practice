package main

import (
	"os"
	"go-practice/opentracing/lib/tracing"
	"fmt"
	"github.com/opentracing/opentracing-go/log"
)

func main() {
	//
	if len(os.Args) != 2 {
		panic("ERROR: Expectiong one argument")
	}
	tracer, closer := tracing.Init("lession01")
	defer closer.Close()

	helloTo := os.Args[1]

	span := tracer.StartSpan("say-hello")
	span.SetTag("hello-to", helloTo)
	helloStr := fmt.Sprintf("Hello, %s!", helloTo)
	span.LogFields(
		log.String("event", "string-format"),
		log.String("value", helloStr),
	)
	println(helloStr)
	span.LogKV("event", "println")
	span.Finish()
}
