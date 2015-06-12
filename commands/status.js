
var exec = require('child_process').exec
var log = require('tablog')

module.exports = {
  description: "provides a cool status printout.",

  match: [
    /jarvis status/,
    /jarvis ping/,
    /jarvis you there/
  ],

  run: function(msg, respond) {
    log.trace("Running status command")
    msg.text = "Jarvis, at your service.\n"

    // Exec git status
    exec('git rev-parse HEAD', function(err, stdout, stderr) {
      if (err) {
        log.warn('Error getting latest git commit hash')
        respond(msg)
        return
      }
      msg.text += "I'm currently running jarvis version " + stdout.substring(0, 6) + " on " + require('../config').machine_name + "."
      respond(msg)
    })

  }
}
