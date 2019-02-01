package main

import(
	"os"
	"strings"
	"os/signal"
	"syscall"
	"math/rand"
	"fmt"
	"regexp"
	"strconv"
	"github.com/bwmarrin/discordgo"
)

var(
	//Botのトークン
	Token string = "BotToken"
)

func main(){
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		panic(err)
	}

	go dg.AddHandler(messageCreate)

	//開始
	err = dg.Open()
	if err != nil {
		panic(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}

//メッセージハンドラ
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){

	//メッセージの分割
	commands := strings.Split(m.Content, " ")
	command:= commands[0]

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping"{
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	//ダイスコマンドの判定
	if command == "/d"{
		var times int
		var diceType int

		diceIs := commands[1]
		diceD := strings.Split(diceIs, "D")

		//diceDの含まれている文字の判定
		if check(`[0-9]`, diceD[0]) == true{
			if check(`[0-9]`, diceD[1]) == true {
				times, _= strconv.Atoi(diceD[0])
				diceType, _ = strconv.Atoi(diceD[1])

				fmt.Println()

				resultInt := roll(diceType, times)

				resultSting := strconv.Itoa(resultInt)

				s.ChannelMessageSend(m.ChannelID, "合計 " + resultSting)
			}else{
				s.ChannelMessageSend(m.ChannelID, "数字を入力してください")
			}
		}else {
			s.ChannelMessageSend(m.ChannelID, "数字を入力してください")
		}
	}
}
//ダイス関数
func roll(diceType int, times int) int {

	var reslut int

	for times > 0 {
		times--
		reslut += rand.Intn(diceType)
	}

	return reslut + 1
}

//含まれている文字の判定
func check(reg, str string) bool {
	if regexp.MustCompile(reg).Match([]byte(str)) == true{
		return true

	} else{

		return false
	}
}