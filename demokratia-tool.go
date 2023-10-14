package main

import (
	"fmt"
	"strings"
	"strconv"
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
				if strings.HasPrefix(m.Content, "+admin add <@") && strings.HasSuffix(m.Content, ">") {
					if len(m.Mentions) == 0 {
						s.ChannelMessageSend(m.ChannelID, "Aucune mention d'utilisateur trouvée dans le message.")
						return
					}
					// Extrait l'ID de l'utilisateur mentionné dans le message
					id_user := m.Mentions[0].ID
					// Ajoute le rôle d'administrateur à l'utilisateur
					err := s.GuildMemberRoleAdd(m.GuildID, id_user, AdminID)
					if err != nil {
						fmt.Println("Erreur lors de l'ajout du rôle :", err)
						s.ChannelMessageSend(m.ChannelID, "Erreur lors de l'ajout du rôle.")
					} else {
						s.ChannelMessageSend(m.ChannelID, "Rôle d'administrateur ajouté à l'utilisateur.")
					}
				} else {
					s.ChannelMessageSend(m.ChannelID, "Le format valide de l'utilisateur est <@id>")
				}
			}
		}
	})

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == OwnerID {
			// Vérifiez si le message commence par "+admin add"
			if strings.HasPrefix(m.Content, "+admin rm ") || strings.HasPrefix(m.Content, "+admin remove ") {
				if (strings.HasPrefix(m.Content, "+admin rm <@") || strings.HasPrefix(m.Content, "+admin remove <@")) && strings.HasSuffix(m.Content, ">") {
					if len(m.Mentions) == 0 {
                    s.ChannelMessageSend(m.ChannelID, "Aucune mention d'utilisateur trouvée dans le message.")
                    return
                }
					// Extrait l'ID de l'utilisateur mentionné dans le message
					id_user := m.Mentions[0].ID
					// Retire le rôle d'administrateur à l'utilisateur
					err := s.GuildMemberRoleRemove(m.GuildID, id_user, AdminID)
					if err != nil {
						fmt.Println("Erreur lors de la suppression du rôle :", err)
						s.ChannelMessageSend(m.ChannelID, "Erreur lors de la suppression du rôle.")
					} else {
						s.ChannelMessageSend(m.ChannelID, "Rôle d'administrateur retiré à l'utilisateur.")
					}
				} else {
					s.ChannelMessageSend(m.ChannelID, "Le format valide de l'utilisateur est <@id>")
				}
			}
		}
	})

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate){
		if m.Author.ID == OwnerID {
			// Vérifie si le message commence par "+clear"
			if strings.HasPrefix(m.Content, "+clear") {
				// Divise le message en mots
				parts := strings.Fields(m.Content)
	
				// Vérifie qu'il y a au moins deux parties dans le message
				if len(parts) >= 2 {
					// Essaie de convertir la deuxième partie en un nombre
					numMessages, err := strconv.Atoi(parts[1])
					if err == nil && numMessages > 0 {
						// Récupère les messages dans le canal
						messages, getErr := s.ChannelMessages(m.ChannelID, numMessages, "", "", "")
						if getErr != nil {
							fmt.Printf("Erreur lors de la récupération des messages: %v\n", getErr)
							return
						}
	
						// Supprime les messages
						for _, message := range messages {
							delErr := s.ChannelMessageDelete(m.ChannelID, message.ID)
							if delErr != nil {
								fmt.Printf("Erreur lors de la suppression du message %s: %v\n", message.ID, delErr)
							}
						}
	
						fmt.Printf("Suppression de %d messages dans le canal %s\n", numMessages, m.ChannelID)
					} else {
						s.ChannelMessageSend(m.ChannelID, "Veuillez spécifier un nombre valide de messages à supprimer.")
					}
				} else {
					s.ChannelMessageSend(m.ChannelID, "Utilisation: +clear x (où x est le nombre de messages à supprimer)")
				}
			}
		}
	})

	err = dg.Open()
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la session DiscordGo:", err)
		return
	}
	fmt.Println("Bot démarré, Ctrl-C pour quitter.")

	select {}
}
