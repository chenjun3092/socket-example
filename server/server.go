package main

import (
  "fmt"
  "net"
  "os"
  "time"
)

func main() {
  service := ":8000"
  tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
  checkErr(err)
  listner, err := net.ListenTCP("tcp", tcpAddr)
  checkErr(err)
  fmt.Fprintf(os.Stdout, "Listning on %s", tcpAddr.String())

  for {
    conn, err := listner.Accept()
    if err != nil {
      continue
    }
    go handleClient(conn)
  }
}

func handleClient(conn net.Conn) {
  defer conn.Close()
  daytime := time.Now().String()
  conn.Write([]byte(daytime))
}

func checkErr(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "Fatal Error: %s", err.Error())
    os.Exit(1)
  }

}
