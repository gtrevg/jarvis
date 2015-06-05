# A Better Slackbot

Or, a worse hubot, depending on how you see it.

# Installing

You need to provie an authentication token in the envvar SLACK_AUTH_TOKEN before starting. You can get this by going to the slack admin console, setting up a new Bot integration, giving it all the parameters you want (like naming it jarvis), and the key should be on that page. 

Other than that, `npm install` once then `npm start` to start up the server. You should see log output immediately. 

# Writing Commands

Every js file in the commands directory contains a command. Wait, that's obvious. Each command is formatted like so:

```
module.exports = {
  description: "Imagine jarvis himself is saying this. That's how this should sound.",
  match: [ /regex/, /array/ ],
  run: function(msg, respond) {}
}
```

Or

```
module.exports = [
  { see above },
  { oh here's another }
]
```

So the base element can be an object or an array of the objects, depending on how you want to organize the js files you create.

### Match

`match` is an array of regex matches. You can look at examples in the provided commands, but the gist is that if a given message matches a regex you provide, the result of that match is provided in the slack message object under the `_matchResult` key. This means you can do grouping very easily, which is pretty useful.

If two commands provide matches which could both match a given message, consider the behavior undefined. It will definitely match and execute one of them but not both and it won't complain while doing so.

### Talking Back

The `respond` parameter on run is a function through which you provide a slackMsg object. This object should at minimum have two fields:

```
{ "text": "The body of the message to send back.", "channel": "the channel id" }
```

It is adequate enough to simply modify the text element of the `msg` passed in then reuse the rest. If you want to post to a different channel then you'll need something