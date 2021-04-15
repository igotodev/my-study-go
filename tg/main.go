// Feed Telegram Bot 
// https://habr.com/ru/rss/best/daily/?fl=ru              - лучшее за день по всем темам
// https://www.opennet.ru/opennews/opennews_all_utf.rss   - rss от opennet

package main

import (
	"encoding/xml"
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

func main() {
	bot, err := tb.NewBot(tb.Settings{
		Token:  "YOUR_TOKEN", // telegram bot token 
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatalln(err)
	}

	//b.Handle("Hello!", func(m *tb.Message) {
	//	b.Send(m.Sender, "Hi!")
	//})
	bot.Handle("/getHabr", func(m *tb.Message) {
		fmt.Println("Loadind...")
		fp := gofeed.NewParser()
		feed, _ := fp.ParseURL("https://habr.com/ru/rss/best/daily/?fl=ru")
		for i, v := range feed.Items {
			if i < 10 {
				str := ""
				strOneNotFullFix, strOneFullFix := "", ""
				strTwoNotFullFix, strTwoFullFix := "", ""

				bTitle, _ := xml.Marshal(v.Title)
				strOneNotFullFix += strings.ReplaceAll(string(bTitle), "<string>", "")
				strOneFullFix += strings.ReplaceAll(strOneNotFullFix, "</string>", "")
				str += strOneNotFullFix + "\t\t\t\n"

				bLink, _ := xml.Marshal(v.Link)
				strTwoNotFullFix += strings.ReplaceAll(string(bLink), "<string>", "")
				strTwoFullFix += strings.ReplaceAll(strTwoNotFullFix, "</string>", "")
				str += strTwoFullFix + "\t\t\t\n"
				bot.Send(m.Sender, str)
			}
		}

		bot.Send(m.Sender, "oho")

		fmt.Println("Success!")
	})
	bot.Handle("/getOpenNet", func(m *tb.Message) {
		fmt.Println("Loadind...")
		str := make([]string, 20)
		fp := gofeed.NewParser()
		feed, err := fp.ParseURL("https://www.opennet.ru/opennews/opennews_all_utf.rss")
		if err != nil {
			log.Fatalln(err)
		}
		for i, v := range feed.Items {
			if i < 20 {
				str[i] += fmt.Sprintln(v.Title)
				str[i] += fmt.Sprintln(v.Link)
				str[i] += "\t\t\t\n"
				str[i] += fmt.Sprintln(v.Description)
				str[i] += "\t\t\t\n\n\n"
				bot.Send(m.Sender, str[i])
			}
		}
		bot.Send(m.Sender, "oho")
		fmt.Println("Success!")
	})
	bot.Handle(tb.OnText, func(m *tb.Message) {
		bot.Send(m.Sender, "If you would like to receive the news, \nplease type \"/getOpennet\" or \"/getHabr\"")
		bot.Send(m.Sender, "Thanks!")

	})
	bot.Handle("Hi!", func(m *tb.Message) {
		bot.Send(m.Sender, "Hi!")
		bot.Send(m.Sender, "If you would like to receive the news, \nplease type \"/getOpennet\" or \"/getHabr\"")
		bot.Send(m.Sender, "Thanks!")
	})
	bot.Handle("What is your name?", func(m *tb.Message) {
		bot.Send(m.Sender, "T1000")
	})
	bot.Handle("stopFeedBot", func(m *tb.Message) {
		bot.Send(m.Sender, "I am out...")
		bot.Stop()
	})
	bot.Start()

}
