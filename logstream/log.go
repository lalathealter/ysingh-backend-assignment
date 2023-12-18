package logstream

import (
	"fmt"
	"log"
	"os"
	"time"
)

const logsDir = "./logs/"
const logTimeLayout = "2006-01-02_15-04-05"

func ProduceLogStreamer() *LogStreamer {
	timeNow := time.Now().Format(logTimeLayout)
	fname := logsDir + timeNow + ".log"
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan LogPair, 2)
	return &LogStreamer{f, c, 2}
}

type LogPair [2]string

func (lp LogPair) String() string {
	res := fmt.Sprintf("INPUT: `%v`\nOUTPUT: `%v`\n", lp[0], lp[1])
	return res
}

type LogStreamer struct {
	File     *os.File
	Channel  chan LogPair
	QueueMax int
}

func (ls *LogStreamer) Send(input, output string) {
	go func() {
		lastInp := LogPair{input, output}
		if len(ls.Channel) < ls.QueueMax-1 {
			ls.Channel <- lastInp
			return
		}

		load := time.Now().Format(time.RFC3339) + ":\n"
		for len(ls.Channel) > 0 {
			curr, ok := <-ls.Channel
			if !ok {
				log.Fatal("FATAL ERROR IN LOGGER")
			}
			load += curr.String()
		}
		load += lastInp.String()

		_, err := ls.File.Write([]byte(load))
		if err != nil {
			log.Fatal(err)
		}

	}()
}
