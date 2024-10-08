package main

import (
  "io"
  "os"
  "fmt"
  "net"
  "path/filepath"
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

func check(err error) {
  if err != nil {
    panic(err)
  }
}

func handleRequest(conn net.Conn) {
  defer conn.Close()
  currentTime := time.Now()
  buf := make([]byte, 1024)
  reqLen, err := conn.Read(buf)
  if err != nil {
    fmt.Println("Error reading from: ", err.Error())
  } else {
    // fmt.Println("Connection from: ", conn.RemoteAddr())
  }

  request := strings.TrimSpace(string(buf[:reqLen]))
  fmt.Println(currentTime.Format(time.RFC3339), conn.RemoteAddr(), request)

  parts := strings.Split(request, " ")
  wide := false
  user := parts[0]

  if parts[0] == "/W" && len(parts) == 2 {
    wide = true
    user = parts[1]
  } else if parts[0] == "/W" && len(parts) == 1 {
    conn.Write([]byte("\r\n"))
    return
  }

  if strings.Index(user, "@") != -1 {
    conn.Write([]byte("Forwarding not supported\r\n"))
  } else {
    if wide {
      //TODO
    } else {
      pwd, err := os.Getwd()
      filePath := filepath.Join(pwd, "plans", filepath.Base(user + ".plan"))
      filePath, err = filepath.Abs(filePath)
      check(err)
      if filepath.Dir(filePath) == filepath.Join(pwd, "plans") {
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
  }
}
