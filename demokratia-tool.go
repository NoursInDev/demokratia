package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func main() {

	BotToken  := "MTE2MjQzMzQ4ODE0MjczMzM3Mw.GomDvH.4rUW3hAI6cumcn-FaLYgRzhV4nRB6GQvGLhfLA"
	//AdminRole := "1162432736091455600"
	OwnerID   := "675016635869954075"
	//GuildID	  := "1159147284613824587"
	//Intents := discordgo.Intents(3276799)


	dg, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		fmt.Println("erreur : type 1")
	}
	defer dg.Close()

	dg.Identify.Intents = discordgo.IntentsAll

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// Vérifiez si le message commence par "+admin add"
		if m.Content == "+admin add" && m.Author.ID == OwnerID{
			s.ChannelMessageSend(m.ChannelID, "Fonctionnel.")
		}
	})

	err = dg.Open()
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la session DiscordGo:", err)
		return
	}
	fmt.Println("bot started")

	// Bloquez le programme pour qu'il s'exécute indéfiniment (vous pouvez gérer proprement la fermeture de la session lorsque vous le souhaitez).
	select {}


}