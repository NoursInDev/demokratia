package main

import (
	"fmt"
	"strings"
	"github.com/bwmarrin/discordgo"
)

func main() {
	BotToken := "---" // Remplacez par le jeton de votre bot
	OwnerID := "675016635869954075"   // Remplacez par l'ID de l'administrateur/bot owner
	AdminID := "1162432736091455600" // Remplacez par l'ID du rôle d'administrateur

	dg, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		fmt.Println("Erreur lors de la création de la session DiscordGo:", err)
		return
	}
	defer dg.Close()

	dg.Identify.Intents = discordgo.IntentsAll

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == OwnerID {
			// Vérifiez si le message commence par "+admin add"
			if strings.HasPrefix(m.Content, "+admin add ") {
				// Extrait l'ID de l'utilisateur mentionné dans le message
				mention := m.Mentions[0]
				id_user := mention.ID

				// Ajoute le rôle d'administrateur à l'utilisateur
				err := s.GuildMemberRoleAdd(m.GuildID, id_user, AdminID)
				if err != nil {
					fmt.Println("Erreur lors de l'ajout du rôle :", err)
					s.ChannelMessageSend(m.ChannelID, "Erreur lors de l'ajout du rôle.")
				} else {
					s.ChannelMessageSend(m.ChannelID, "Rôle d'administrateur ajouté à l'utilisateur.")
				}
			}
		}
	})

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == OwnerID {
			// Vérifiez si le message commence par "+admin add"
			if strings.HasPrefix(m.Content, "+admin rm ") {
				// Extrait l'ID de l'utilisateur mentionné dans le message
				mention := m.Mentions[0]
				id_user := mention.ID

				// Ajoute le rôle d'administrateur à l'utilisateur
				err := s.GuildMemberRoleRemove(m.GuildID, id_user, AdminID)
				if err != nil {
					fmt.Println("Erreur lors de l'ajout du rôle :", err)
					s.ChannelMessageSend(m.ChannelID, "Erreur lors de l'ajout du rôle.")
				} else {
					s.ChannelMessageSend(m.ChannelID, "Rôle d'administrateur retiré à l'utilisateur.")
				}
			}
		}
	})

	err = dg.Open()
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la session DiscordGo:", err)
		return
	}
	fmt.Println("Bot démarré.")

	select {}
}
