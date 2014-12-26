package main

import (
    zmq "github.com/alecthomas/gozmq"
    "log"
    "github.com/howeyc/fsnotify"
    "os"
)

func watchFile(fileName string, fileChannel chan string, logChannel chan []byte) {
    log_file, err := os.Open(fileName)

    stat, _ := os.Stat(fileName)
    size := stat.Size()
    log_file.Seek(0, 2)

    watcher, err := fsnotify.NewWatcher()

    if err != nil {
        log.Fatal(err)
    }

    go func() {
        for {
            ev := <-watcher.Event

            if (ev != nil && ev.IsModify()) {
                // Create a buffer for reading the new data
                stat, err = os.Stat(fileName)

                if err != nil {
                    continue
                }

                new_size := stat.Size()
                bytes := make([]byte, new_size - size)

                if (new_size - size > 0) {
                    _, err := log_file.Read(bytes)

                    if err != nil {
                        log.Fatal(err)
                    }

                    fileChannel <- fileName
                    logChannel <- bytes

                    if err != nil {
                       log.Fatal(err)
                    }
                }
                size = new_size
            }
        }
    }()

    err = watcher.Watch(fileName)

    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    files := []string {
        "/var/log/system.log",
        "/Users/ashish/server_log",
    }

    // Create and bind socket
    context, _ := zmq.NewContext()
    socket, _ := context.NewSocket(zmq.PUB)
    defer context.Close()
    defer socket.Close()
    socket.Bind("tcp://*:5556")

    fileChannel := make(chan string)
    logChannel := make(chan []byte)

    println("Watching file for changes")

    for _, v := range files {
        watchFile(v, fileChannel, logChannel)
    }

    for {
        fName := <-fileChannel
        logData := <-logChannel
        outData := append([]byte(fName + ":"), logData...)
        _ = socket.Send(outData, 0)
    }
}
