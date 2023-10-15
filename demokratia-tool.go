package main

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/bwmarrin/discordgo"
)

func main() {
	BotToken := "---" // Token Bot
	OwnerID := "675016635869954075"   // Owner ID
	AdminID := "1162432736091455600" // Admin role ID
	//MemberID := "1163118801915760720"

	dg, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		fmt.Println("Erreur lors de la création de la session DiscordGo:", err)
		return
	}
	defer dg.Close()

	dg.Identify.Intents = discordgo.IntentsAll

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == OwnerID {
			// +admin add check
			if strings.HasPrefix(m.Content, "+admin add ") {
				if strings.HasPrefix(m.Content, "+admin add <@") && strings.HasSuffix(m.Content, ">") {
					if len(m.Mentions) == 0 {
						s.ChannelMessageSend(m.ChannelID, "Aucune mention d'utilisateur trouvée dans le message.")
						return
					}
					// ID extraction
					id_user := m.Mentions[0].ID
					// Role adding
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
			// +admin rm/remove check
			if strings.HasPrefix(m.Content, "+admin rm ") || strings.HasPrefix(m.Content, "+admin remove ") {
				if (strings.HasPrefix(m.Content, "+admin rm <@") || strings.HasPrefix(m.Content, "+admin remove <@")) && strings.HasSuffix(m.Content, ">") {
					if len(m.Mentions) == 0 {
                    s.ChannelMessageSend(m.ChannelID, "Aucune mention d'utilisateur trouvée dans le message.")
                    return
                }
					// ID extraction
					id_user := m.Mentions[0].ID
					// Role removing
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
			// +clear check
			if strings.HasPrefix(m.Content, "+clear") {
				// Message division
				parts := strings.Fields(m.Content)
	
				// Message Conformity check - 2 sides message
				if len(parts) >= 2 {
					// 2nd part int convertion
					numMessages, err := strconv.Atoi(parts[1])
					if err == nil && numMessages > 0 {
						// Tacking back messages
						messages, getErr := s.ChannelMessages(m.ChannelID, numMessages, "", "", "")
						if getErr != nil {
							fmt.Printf("Erreur lors de la récupération des messages: %v\n", getErr)
							return
						}
	
						// messages deleting
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