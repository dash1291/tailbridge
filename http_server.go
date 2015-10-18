// http server with socket.io
package main

import (
    "net/http"

    "github.com/googollee/go-socket.io"
    "os/exec"
    "log"
    "strings"
    "bytes"
)

func Tail(machine string, file_name string, out_bytes chan string) {
    tail_cmd := "tail -f " + file_name
    cmd := exec.Command("ssh", machine, tail_cmd)
    stdout, err := cmd.StdoutPipe()

    if err != nil {
      log.Fatal(err)
    }

    cmd.Start()
    defer cmd.Wait()

    line := bytes.NewBuffer([]byte{})
    r := make([]byte, 1)

    for {
        stdout.Read(r)

        if r[0] == '\n' {
            out_bytes <- line.String()
            line.Reset()
        } else {
            line.WriteByte(r[0])
        }
    }
}

func main() {
    sio_server, err := socketio.NewServer(nil)

    if err != nil {
        log.Fatal(err)
    }

    sio_server.On("connection", func(socket socketio.Socket) {
        log.Println("Got new connection")

        socket.On("init", func(msg string) {
            msg_parts := strings.Split(msg, ",")
            if len(msg_parts) < 2 {
                println("Insufficient data received from client.")
                return
            }

            out_bytes := make(chan string)
            go Tail(msg_parts[0], msg_parts[1], out_bytes)

            for {
                socket.Emit("stream", string(<-out_bytes))
            }
        })

        socket.On("stream", func(msg string) {
            log.Println("emit:", socket.Emit("chat message", msg))
            socket.BroadcastTo("chat", "chat message", msg)
        })

        socket.On("disconnection", func() {
            log.Println("on disconnect")
        })
    })

    sio_server.On("error", func(so socketio.Socket, err error) {
        log.Println("error:", err)
    })

    http.Handle("/socket.io/", sio_server)
    http.Handle("/", http.FileServer(http.Dir("./static")))
    log.Println("Serving at localhost:5000...")
    log.Fatal(http.ListenAndServe(":5000", nil))
}
