package main

import (
	"fmt"
	p "github.com/mrtuuro/cli-tool/pattern"
	"os"
	"time"
)

func main() {
	commandArgs := os.Args
	commandArgsParsed := commandArgs[1:]
	startTime := time.Now()
	switch commandArgsParsed[0] {
	case "--help":
		fmt.Printf(`
Usage:    go run main.go [OPTIONS] COMMAND

Options: // Not implemented yet.
		--
		--
		--
		--
		--
Commands:
	generator	Run "generator concurrency pattern"
	queuing		Run "queuing concurrency pattern"
	workerpool	Run "workerpool concurrency pattern"
	pipeline	Run "pipeline concurrency pattern"
	fanout		Run "fanout concurrency pattern"
	fanin		Run "fanin concurrency pattern"

`)
	case "generator":
		p.GeneratorMain()
		fmt.Println("Process time: ", time.Since(startTime))
	case "fanin":
		p.FaninMain()
		fmt.Println("Process time: ", time.Since(startTime))

	case "fanout":
		p.FanoutMain()
		fmt.Println("Process time: ", time.Since(startTime))
	case "pipeline":
		p.PipelineMain()
		fmt.Println("Process time: ", time.Since(startTime))
	case "workerpool":
		p.WorkerpoolMain()
		fmt.Println("Process time: ", time.Since(startTime))
	case "queueing":
		p.QueueingMain()
		fmt.Println("Process time: ", time.Since(startTime))
	}
}
