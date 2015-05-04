package main

import (
	"bufio"
	"flag"
	"io"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/Sirupsen/logrus"
)

var ln net.Listener
var restarting bool
var logger *logrus.Entry

func handleSignal(c chan os.Signal) func() {
	return func() {
		for {
			<-c
			logger.Info("sighup received")
			fl, err := ln.(*net.TCPListener).File()
			if err != nil {
				logger.Error(err)
				return
			}
			cmd := exec.Command(os.Args[0], "-restart")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.ExtraFiles = []*os.File{fl}
			cmd.Start()
			logger.Info("FINISHED")
			os.Exit(0)
		}
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	_, err := io.Copy(w, r)
	if err != nil {
		logger.Error(err)
	}
}

func main() {
	logger = logrus.WithFields(logrus.Fields{"PID": os.Getpid()})
	logger.Info("START")

	restarting = false
	restart := flag.Bool("restart", false, "")

	flag.Parse()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)
	go handleSignal(c)()

	var err error
	if *restart {
		f := os.NewFile(3, "")
		ln, err = net.FileListener(f)
	} else {
		ln, err = net.Listen("tcp", ":12345")
	}
	if err != nil {
		logger.Fatal(err)
	}
	for !restarting {
		conn, err := ln.Accept()
		if err != nil {
			logger.Error(err)
			continue
		}
		go handleConnection(conn)
	}
}
