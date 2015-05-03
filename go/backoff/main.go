package main

import (
	"errors"
	"fmt"

	"github.com/cenkalti/backoff"
	"github.com/coreos/go-log/log"
)

var state int

func opFail2() error {
	if state != 2 {
		state++
		println(fmt.Sprintf("failed(%d)", state))
		return errors.New("error")
	}
	println("passed")
	state = 0
	return nil
}

func opFailAll() error {
	println("failed")
	return errors.New("error")
}

func main() {
	var err error

	println("fail2")
	err = backoff.Retry(opFail2, backoff.NewExponentialBackOff())
	if err != nil {
		log.Error(err)
	}

	println("fail and retry controll")
	b := backoff.NewExponentialBackOff()
	ticker := backoff.NewTicker(b)
	retry := 0
	for range ticker.C {
		if err = opFailAll(); err != nil {
			retry++
			if retry == 3 {
				break
			}
			continue
		}
		break
	}
	ticker.Stop()

	println("always fail...(Ctrl-C stop)")
	err = backoff.Retry(opFailAll, backoff.NewExponentialBackOff())
	if err != nil {
		log.Error(err)
	}
}
