package handlers

import (
	"fmt"
	"godis/internal/storage"
	"net"
)

func parseBytesToSlice(b []byte) []string {
	command := make([]string, 0)
	comm := ""
	for idx, char := range string(b) {
		if string(char) == " " {
			command = append(command, comm)
			comm = ""
		} else if idx == len(b)-1 {
			comm += string(char)
			command = append(command, comm)
		} else {
			comm += string(char)
		}
	}
	return command
}

func HandleConnection(conn net.Conn, s *storage.Storage) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024*4)
		n, err := conn.Read(buf)
		if n == 0 || err != nil {
			break
		}
		if string(buf[0:3]) == "set" {
			command := parseBytesToSlice(buf)
			if len(command) != 3 {
				return
			} else {
				err = s.Set(command[1], command[2])
				if err != nil {
					return
				}
			}
		} else if string(buf[0:3]) == "get" {
			command := parseBytesToSlice(buf)
			if len(command) != 2 {
				return
			} else {
				fmt.Println(command[1])
				data := s.Get(command[1])
				// conn.Write([]byte(data))
				fmt.Println(data)
			}
		}
	}
}
