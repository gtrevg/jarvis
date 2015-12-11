package main

import (
  "jarvis/commands"
  "jarvis/config"
  "jarvis/data"
  "jarvis/handlers"
  "jarvis/log"
  "jarvis/ws"
  "runtime"
)

func main() {
  log.Info("Starting Jarvis")
  runtime.GOMAXPROCS(runtime.NumCPU())
  config.LoadYaml()
  data.CheckRedisConn()
  handlers.Init()
  commands.StartReminderLoop()
  ws.Init()
  log.Info("Jarvis is live and receiving messages")
  select {}
}
