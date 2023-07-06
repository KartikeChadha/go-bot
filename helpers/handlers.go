package helpers

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
)

func AgeHandler(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
	year := r.Param("year")
	birthYear, err := strconv.Atoi(year)
	t := time.Now()
	curYear := t.Year()
	if err != nil {
		fmt.Println(err)
	}
	age := curYear - birthYear
	var output string
	if age < 0 || age > 115 {
		output = fmt.Sprintf("hmmm seems like you might be lying about your age")
	} else {
		output = fmt.Sprintf("your age is: %d", age)
	}

	w.Reply(output)
}

func FileHandler(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
	path := r.Param("path")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := path

	params := slack.FileUploadParameters{
		Channels: channelArr,
		File:     fileArr,
	}
	file, err := api.UploadFile(params)
	if err != nil {
		fmt.Printf("ERROR IS SDALSDNLAK ASLKDNALSDNALSKNDLASNDLA%s\n", err)
		// log.Print(err)
	} else {
		fmt.Printf("Name: %s, URL: %s", file.Name, file.URL)
	}

}
