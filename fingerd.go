package main

import (
  "io"
  "os"
  "fmt"
  "net"
  "path"
  "time"
  "strings"
)

const (
  CONN_HOST = "0.0.0.0"
  CONN_PORT = "79"
  CONN_TYPE = "tcp"
)

func main () {
  l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
  if err != nil {
    fmt.Println("Error opening port: ", err.Error())
    os.Exit(1)
  }

  defer l.Close()
  for {
    conn, err := l.Accept()
    if err != nil {
      fmt.Println("Error accepting connection: ", err.Error())
      continue
    }
    go handleRequest(conn)
  }
}

func handleRequest(conn net.Conn) {
  currentTime := time.Now()
  buf := make([]byte, 1024)
  reqLen, err := conn.Read(buf)
  fmt.Println(currentTime.Format(time.RFC3339))
  if err != nil {
    fmt.Println("Error reading from: ", err.Error())
  } else {
    fmt.Println("Connection from: ", conn.RemoteAddr())
  }

  request := strings.TrimSpace(string(buf[:reqLen]))
  fmt.Println(request)

  parts := strings.Split(request, " ")
  wide := false
  user := parts[0]

  if parts[0] == "/W" {
    wide = true
    user = parts[1]
  }

  if strings.Index(user, "@") != -1 {
    fmt.Println("remote")
    conn.Write([]byte("Forwarding not supported\r\n"))
  } else {
    if wide {
      //TODO
    } else {
      pwd, err := os.Getwd()
      filePath := path.Join(pwd, "plans", path.Base(user + ".plan"))
      filePath = path.Clean(filePath)
      fmt.Println(filePath)
      file, err := os.Open(filePath)
      if err != nil {
        //not found
        // io.Write([]byte("Not Found\r\n"))
      } else {
        defer file.Close()
        io.Copy(conn,file)
        conn.Write([]byte("\r\n"))
      }
    }
  }


  conn.Close()
}
