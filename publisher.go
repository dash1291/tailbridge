package main

import (
    zmq "github.com/alecthomas/gozmq"
    "log"
    "github.com/howeyc/fsnotify"
    "os"
)

func main() {
    // Create and bind socket
    context, _ := zmq.NewContext()
    socket, _ := context.NewSocket(zmq.PUB)
    defer context.Close()
    defer socket.Close()
    socket.Bind("tcp://*:5556")
    
    // Open file and seek to end
    FILE := "/var/log/system.log"
    log_file, err := os.Open(FILE)
    
    if err != nil {
        println(err)
    }

    stat, _ := os.Stat(FILE)
    size := stat.Size()

    log_file.Seek(0, 2)

    if err != nil {
        log.Fatal(err)
    }

    watcher, err := fsnotify.NewWatcher()

    done := make(chan bool)

    println("Watching file for changes...")
    // Process events
    go func() {
        for {
            ev := <-watcher.Event
            if (ev.IsModify()) {

                // Create a buffer for reading the new data
                stat, _ = os.Stat(FILE)
                new_size := stat.Size()
                bytes := make([]byte, new_size - size)
                
                if (new_size - size > 0) {
                    _, err := log_file.Read(bytes)

                    if err != nil {
                        log.Fatal(err)
                    }

                    err = socket.Send(bytes, 0)

                    if err != nil {
                       log.Fatal(err)
                    }
                }
                size = new_size
            }
        }
    }()

    err = watcher.Watch(FILE)
    if err != nil {
        log.Fatal(err)
    }

    <-done

    watcher.Close()
}
