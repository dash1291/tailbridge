package main

import zmq "github.com/alecthomas/gozmq"

func main() {
  context, _ := zmq.NewContext()
  socket, _ := context.NewSocket(zmq.SUB)
  defer context.Close()
  defer socket.Close()

  socket.SetSubscribe("")
  socket.Connect("tcp://localhost:5556")
  println("connected")

  for {
    msg, err := socket.Recv(0)
    
    if err != nil {
        println(err)
    }
    println(string(msg))
  }
}
