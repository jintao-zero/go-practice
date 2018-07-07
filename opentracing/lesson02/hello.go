package main

import (
    "os"
    "fmt"
    "go-practice/opentracing/lib/tracing"
    "github.com/opentracing/opentracing-go"
    "context"
    "github.com/opentracing/opentracing-go/log"
)

func main()  {
    if len(os.Args) != 2 {
        panic("ERROR: Expecting one argument")
    }

    tracer, closer := tracing.Init("lession02")
    defer closer.Close()
    opentracing.SetGlobalTracer(tracer)

    helloTo := os.Args[1]

    span := tracer.StartSpan("say-hello")
    span.SetTag("hello-to", helloTo)
    defer span.Finish()

    span.SetBaggageItem("hahah", "eee")
    span.Context().ForeachBaggageItem(func(k, v string) bool {
        fmt.Println(k, v)
        return true
    })
    ctx := opentracing.ContextWithSpan(context.Background(), span)
    helloStr := formatString(ctx, helloTo)
    printHello(ctx, helloStr)
}

func formatString(ctx context.Context, helloTo string) string {
    span, _ := opentracing.StartSpanFromContext(ctx, "formatString")
    defer span.Finish()

    helloStr := fmt.Sprintf("Hello, %s!", helloTo)
    span.LogFields(
        log.String("event", "string-format"),
        log.String("value", helloStr),
    )
    return helloStr
}

func printHello(ctx context.Context, helloStr string) {
    span, _ := opentracing.StartSpanFromContext(ctx, "printHello")
    defer span.Finish()

    println(helloStr)
    span.LogKV("event", "println")
}


