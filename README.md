# KingGizzBot

**"Is a King Gizzard song being played on @UniversityRadioYork, and if so send a message to people saying we should VoNC @tovels ?"**

Runs as a cronjob every now and again.

### Command Line Flags

`./kingGizzBot [flags]` - No flags runs as default, checking the endpoint, and posting the standard message to Slack if a King Gizz song is being played

-   `-m` - Manual Run. Use the standard message even if a King Gizz song isn't playing
-   `-c (string)` - Custom Message. Send a custom message, regardless of the currently playing song.
-   `-e (w | y)` - Emoji. Choose to send the message in Slack's white letter emoji or yellow letter emoji. Excluding this flag will send the message as plain text.
-   `-d` - Development Mode. If included, the message won't be sent to the Slack endpoint, instead outputted to the console.

###### Michael Grace 2020, 2021
