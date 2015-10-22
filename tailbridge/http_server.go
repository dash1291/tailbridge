// http server with socket.io
package tailbridge

import (
    "net/http"
    "strconv"
    "github.com/googollee/go-socket.io"
    "os/exec"
    "log"
    "strings"
    "bytes"
)

func Tail(machine string, user string, port int, file_name string, out_bytes chan string) {
    tail_cmd := "tail -f " + file_name
    cmd := exec.Command(
        "ssh", "-p", strconv.Itoa(port),
        user + "@" + machine,
        tail_cmd)

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

func InitServer(port int) {
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

            user, port, ok := GetMachineParams(msg_parts[0])

            if !ok {
                println("Some problem with the IP provided.")
            }

            out_bytes := make(chan string)
            go Tail(msg_parts[0], user, port, msg_parts[1], out_bytes)

            for {
                socket.Emit("stream", string(<-out_bytes))
            }
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

    port_str := strconv.Itoa(port)
    log.Println("Listening on 0.0.0.0:" + port_str)
    log.Fatal(http.ListenAndServe("0.0.0.0:" + port_str, nil))
}
