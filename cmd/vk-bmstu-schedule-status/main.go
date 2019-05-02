package main

import (
	"flag"
	"fmt"
	"os"

	vkstatus "github.com/trubitsyn/vk-bmstu-schedule-status"
)

var (
	login    = flag.String("login", "", "VK login")
	password = flag.String("password", "", "VK password")

	chatId = flag.Uint("chat", 0, "Chat ID")
	prefix = flag.String("prefix", "", "Chat title prefix")
)

func main() {
	flag.Parse()
	checkError(*prefix == "", "Invalid chat title prefix")
	checkError(*chatId == 0 || *chatId > vkstatus.MaxChatId, "Invalid chat ID")
	checkError(*login == "" || *password == "", "Invalid credentials")

	var client vkstatus.DefaultVkApiClient
	if err := client.Connect(*login, *password); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	weekNum := vkstatus.CurrentWeek()
	weekName := vkstatus.WeekTypeFromNumber(weekNum).ShortWeekName()

	title := vkstatus.Title{
		Prefix:   *prefix,
		WeekAbbr: "нед.",
		WeekNum:  weekNum,
		WeekName: weekName,
	}.Format("p (s) d a")

	if err := client.EditChat(*chatId, title); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Set chat %d title to '%s'\n", *chatId, title)
}

func checkError(condition bool, message string) {
	if condition {
		fmt.Println("Error:", message)
		flag.Usage()
		os.Exit(1)
	}
}
