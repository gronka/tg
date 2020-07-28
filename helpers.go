package tg

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var EmptyUuidString = "00000000-0000-0000-0000-000000000000"
var EmptyBytes = [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var OneDayMilli = 86400000
var OneDaySecs = 86400

func timeNowMilli() int64 {
	return time.Now().UnixNano() / 1e6
}

func check(err error, msg string) {
	if err != nil {
		panic(errors.Wrap(err, msg))
	}
}

func printBytes(bytes []byte) {
	for _, n := range bytes {
		fmt.Printf("%8b", n) // prints 1111111111111101
	}
	fmt.Printf("\n")
}

func logPanicf(i interface{}) { // 5
	log.Panic().Msg(fmt.Sprintf("%v", i))
}

func logFatalf(i interface{}) { // 4
	log.Fatal().Msg(fmt.Sprintf("%v", i))
}

func logErrorf(i interface{}) { // 3
	log.Error().Msg(fmt.Sprintf("%v", i))
}

func logWarnf(i interface{}) { // 2
	log.Warn().Msg(fmt.Sprintf("%v", i))
}

func logInfof(i interface{}) { // 1
	log.Info().Msg(fmt.Sprintf("%v", i))
}

func logDebugf(i interface{}) { // 0
	log.Debug().Msg(fmt.Sprintf("%v", i))
}

func logTracef(i interface{}) { // -1
	log.Trace().Msg(fmt.Sprintf("%v", i))
}
