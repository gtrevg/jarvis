// Root handler/dispatcher for all "commands" from slack.
// Commands are anything of the form
//    jarvis {ACTION} {arguments}
// Examples
//    jarvis PING www.google.com
//    jarvis LOVE mikehock
//    jarvis GIF cats
//    jarvis GOOGLE weird fetish porn
//    jarvis YOUTUBE Rick Astley

package handlers

import (
  "github.com/jbrukh/bayesian"
  "github.com/mhoc/jarvis/commands"
  "github.com/mhoc/jarvis/log"
  "github.com/mhoc/jarvis/util"
  "github.com/mhoc/jarvis/ws"
  "strings"
)

var commandManifest = []util.Command{
  commands.Help{},
  commands.Status{},
}

var cmdCh = make(chan util.IncomingSlackMessage)
var Classifier *bayesian.Classifier

func InitCommands() {
  log.Info("Initing command listener")
  ws.SubscribeToMessages(cmdCh)
  TrainCommandClassifier()
  go BeginCommandLoop()
}

func TrainCommandClassifier() {
  classes := []bayesian.Class{}
  for _, command := range commandManifest {
    classes = append(classes, command.Class())
  }
  Classifier = bayesian.NewClassifier(classes...)
  for _, command := range commandManifest{
    Classifier.Learn(command.TrainingStrings(), command.Class())
  }
}

func BeginCommandLoop() {
  for {
    msg := <-cmdCh
    if !IsCommand(msg.Text) {
      continue
    }
    cmd := MatchCommand(msg.Text)
    go cmd.Execute(msg)
  }
}

func IsCommand(text string) bool {
  if !strings.HasPrefix(text, "jarvis") && !strings.HasPrefix(text, "Jarvis") {
    return false
  }
  return true
}

func MatchCommand(text string) util.Command {
  _, likely, _ := Classifier.LogScores(strings.Split(text, " "))
  return commandManifest[likely]
}
