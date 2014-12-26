package main

import (
  zmq "github.com/alecthomas/gozmq"
  "os"
)


func main() {
  context, _ := zmq.NewContext()
  socket, _ := context.NewSocket(zmq.SUB)
  defer context.Close()
  defer socket.Close()

  if len(os.Args) < 3 {
    println("More arguments needed. \nUsage: receiver <host> <filename>")
    return
  }

  host := os.Args[1]
  fileName := os.Args[2]

  socket.SetSubscribe(fileName)
  socket.Connect("tcp://" + host + ":5556")
  println("Connected.")

  for {
    msg, err := socket.Recv(0)
    
    if err != nil {
        println(err)
    }
    if len(msg) > (len(fileName) + 1) {
      println(string(msg)[len(fileName) + 1:])
    }
  }
}
