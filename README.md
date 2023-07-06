# Golang Slack Bot
A slack bot written in Go which can upload files on slack channels and calculate age

##Setup 

Initially, change the Slack Bot token and App token along with the channel ID according to your workspace
```bash
SLACK_BOT_TOKEN = xoxb-xxxxxxxxxxxxx-xxxxxxxxxxxx-xxxxxxxxxxxxxxxxxxxxxxxx
xoxb-xxxxxxxxxxxxx-xxxxxxxxxxxx-nW6L4pjkicLXYL9569PeNlar
SLACK_APP_TOKEN = xapp-1-xxxxxxxxxxx-xxxxxxxxxxxxx-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

CHANNEL_ID = XXXXXXXXXXX
```

To run from the source code:
```bash
go mod tidy
go build && go run main.go
```


###FileUploader
![alt text](./.img/1.png?raw=true)

To upload a file from your local machine to slack through a command
```bash
@go-bot upload <path>
```
where path is the path relative to root where the file resides

###Age Calculator
![alt text](./.img/2.png?raw=true)

To calculate your age given your birth yera with slack through a command
```bash
@go-bot i was born in <year>
```
where year is your birth year