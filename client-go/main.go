package main

import "net"
import "log"
import "encoding/binary"
import "io/ioutil"
import "fmt"

func main() {
  template, err := ioutil.ReadFile("template.mjml")

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("write", len(template))

  conn, err := net.Dial("tcp", "localhost:8686")

  defer conn.Close()

  if err != nil {
    log.Fatal(err)
  }

  var sizeSend [4]byte
  binary.BigEndian.PutUint32(sizeSend[:], uint32(len(template)))

  conn.Write(sizeSend[:])
  conn.Write(template[:])

  data, err := ioutil.ReadAll(conn)

  if err != nil {
    log.Fatal(err)
  }

  sizeReceive := binary.BigEndian.Uint32(data[:4])

  fmt.Println("read", sizeReceive)

  if len(data[4:]) == int(sizeReceive) {
    err = ioutil.WriteFile("output.html", data[4:], 0755)

    if err != nil {
      log.Fatal(err)
    }
  }
}
