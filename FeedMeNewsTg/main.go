// FeedMeNews Telegram Bot

package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
	tb "gopkg.in/tucnak/telebot.v2"
)

//type Sendable interface {
//	Send(*Bot, Recipient, *SendOptions) (*Message, error)
//}
const stopWord = "stopFeedBot" // stop word for the bot

func main() {
	bot, err := tb.NewBot(tb.Settings{
		Token:  "YOUR_TOKEN", // bot token 
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatalln(err)
	}
	//b.Handle("Hello!", func(m *tb.Message) {
	//	b.Send(m.Sender, "Hi!")
	//})
	sayHi := "Hi! " + "\nIf you would like to receive the news,\n" +
		"please type \"/getSlashdot\", \"/getMediumProg\" or\n\"/getStackOverFlow\"." +
		" And if you would like to read\n a blog about Go development,\nplease type \"/getEli\"." + "\nThanks!\n"
	// news sites
	eli := "https://eli.thegreenplace.net/feeds/all.atom.xml"
	slashdot := "http://rss.slashdot.org/Slashdot/slashdotMain"
	medium := "https://medium.com/feed/topic/programming"
	stackoverflow := "https://stackoverflow.blog/feed/"

	readRSS(slashdot, "/getSlashdot", bot)
	readRSS(medium, "/getMediumProg", bot)
	readRSS(stackoverflow, "/getStackOverFlow", bot)
	readXML(eli, "/getEli", bot)

	massage("What is your name?", "T1000", bot)
	massage(tb.OnText, sayHi, bot)

	bot.Handle(stopWord, func(m *tb.Message) {
		bot.Send(m.Sender, "I am out...")
		bot.Stop()
	})
	bot.Start()
}
// xml reading and delivery function 
func readXML(address string, command string, bot *tb.Bot) {
	bot.Handle(command, func(m *tb.Message) {
		fmt.Println("Loadind...")
		fp := gofeed.NewParser()

		feed, _ := fp.ParseURL(address)
		for i, v := range feed.Items {
			if i < 10 {
				str := ""
				strOneNotFullFix, strOneFullFix := "", ""
				strTwoNotFullFix, strTwoFullFix := "", ""

				strOneNotFullFix += strings.ReplaceAll(v.Title, "<string>", "")
				strOneFullFix += strings.ReplaceAll(strOneNotFullFix, "</string>", "")
				str += strOneNotFullFix + "\t\t\t\n"

				str += v.Published + "\t\t\t\n"

				strTwoNotFullFix += strings.ReplaceAll(v.Link, "<string>", "")
				strTwoFullFix += strings.ReplaceAll(strTwoNotFullFix, "</string>", "")
				str += strTwoFullFix + "\t\t\t\n"

				bot.Send(m.Sender, str)
			}
		}
		bot.Send(m.Sender, "oho")

		fmt.Println("Success!")
	})
}
// rss reading and delivery function 
func readRSS(address string, command string, bot *tb.Bot) {
	bot.Handle(command, func(m *tb.Message) {
		fmt.Println("Loadind...")
		str := make([]string, 10)
		fp := gofeed.NewParser()
		feed, err := fp.ParseURL(address)
		if err != nil {
			log.Fatalln(err)
		}
		for i, v := range feed.Items {
			if i < 10 {
				str[i] += fmt.Sprintln(v.Title)
				str[i] += fmt.Sprintln(v.Link)
				bot.Send(m.Sender, str[i])
			}
		}
		bot.Send(m.Sender, "oho")
		fmt.Println("Success!")
	})
}
// for message delivery
func massage(userTxt string, botTxt string, bot *tb.Bot) {
	bot.Handle(userTxt, func(m *tb.Message) {
		bot.Send(m.Sender, botTxt)
	})
}
