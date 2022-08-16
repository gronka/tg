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

func TimeNowMilli() int64 {
	return time.Now().UnixNano() / 1e6
}

func CheckAndPanic(err error, msg string) {
	if err != nil {
		panic(errors.Wrap(err, msg))
	}
}

func CheckAndWrap(err error, msg string) error {
	if err != nil {
		return errors.Wrap(err, msg)
	}
	return nil
}

func CheckQuery(err error) error {
	if err != nil {
		return errors.Wrap(err, "query failed")
	}
	return nil
}

func PrintBytes(bytes []byte) {
	for _, n := range bytes {
		fmt.Printf("%8b", n) // prints 1111111111111101
	}
	fmt.Printf("\n")
}

func Panic(i interface{}) { // 5
	log.Panic().Msg(fmt.Sprintf("%v", i))
}

func Fatal(i interface{}) { // 4
	log.Fatal().Msg(fmt.Sprintf("%v", i))
}

func Error(i interface{}) { // 3
	log.Error().Msg(fmt.Sprintf("%v", i))
}

func Warn(i interface{}) { // 2
	log.Warn().Msg(fmt.Sprintf("%v", i))
}

func Info(i interface{}) { // 1
	log.Info().Msg(fmt.Sprintf("%v", i))
}

func Debug(i interface{}) { // 0
	log.Debug().Msg(fmt.Sprintf("%v", i))
}

func Trace(i interface{}) { // -1
	log.Trace().Msg(fmt.Sprintf("%v", i))
}
