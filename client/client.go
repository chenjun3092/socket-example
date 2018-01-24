package main

import (
  "fmt"
  "io/ioutil"
  "net"
  "os"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Fprintf(os.Stderr, "Usage %s host:port", os.Args[0])
  }
  service := os.Args[1]
  tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
  checkErr(err)
  fmt.Fprintf(os.Stdout, "Requestingt to %s\n", tcpAddr.String())
  conn, err := net.DialTCP("tcp", nil, tcpAddr)
  checkErr(err)

  // Like HTTP Requesting
  _, err = conn.Write([]byte("HEAD / HTTP1.0\r\n\r\n"))
  checkErr(err)

  result, err := ioutil.ReadAll(conn)
  checkErr(err)
  fmt.Println(string(result))
  os.Exit(0)
}

func checkErr(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "Fatal Error: %s", err.Error())
    os.Exit(1)
  }
}
