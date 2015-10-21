package main

import "./tailbridge"

func main() {
  config := tailbridge.ReadConfig("./config.yml")
  tailbridge.InitServer(config.Listen)
}
