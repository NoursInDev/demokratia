package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var (
	idOwner  = "675016635869954075"
	idAdmin  = "1162432736091455600"
	token    = "---"
)

func main() {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Erreur lors de la création de la session Discord:", err)
		return
	}

	dg.AddHandler(messageCreate) // Modifiez cette ligne

	err = dg.Open()
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la session Discord:", err)
		return
	}

	fmt.Println("Bot actif. Appuyez sur CTRL+C pour quitter.")
	select {}
}


func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == idOwner {
		// Vérifiez si le message commence par "+admin add" ou "+admin rm"
		if len(m.Content) > 11 && m.Content[:10] == "+admin add" {
			userID := m.Content[11:]
			addAdminRole(s, m.GuildID, userID)
		} else if len(m.Content) > 10 && m.Content[:9] == "+admin rm" {
			userID := m.Content[10:]
			removeAdminRole(s, m.GuildID, userID)
		}
	}
}

func addAdminRole(s *discordgo.Session, guildID, userID string) {
	// Obtenez le rôle d'admin par ID
	role, err := s.State.Role(guildID, idAdmin)
	if err != nil {
		fmt.Println("Erreur lors de la recherche du rôle d'admin:", err)
		return
	}

	// Ajoutez le rôle à l'utilisateur
	err = s.GuildMemberRoleAdd(guildID, userID, role.ID)
	if err != nil {
		fmt.Println("Erreur lors de l'ajout du rôle à l'utilisateur:", err)
		return
	}

	fmt.Println("Le rôle", role.Name, "a été ajouté à l'utilisateur", userID)
}

func removeAdminRole(s *discordgo.Session, guildID, userID string) {
	// Obtenez le rôle d'admin par ID
	role, err := s.State.Role(guildID, idAdmin)
	if err != nil {
		fmt.Println("Erreur lors de la recherche du rôle d'admin:", err)
		return
	}

	// Retirez le rôle de l'utilisateur
	err = s.GuildMemberRoleRemove(guildID, userID, role.ID)
	if err != nil {
		fmt.Println("Erreur lors de la suppression du rôle de l'utilisateur:", err)
		return
	}

	fmt.Println("Le rôle", role.Name, "a été enlevé de l'utilisateur", userID)
}
