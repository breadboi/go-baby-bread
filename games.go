package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

// GetVoiceMembers Returns a list of member IDs as strings from the channel
func GetVoiceMembers(dgoSession *discordgo.Session, i *discordgo.InteractionCreate, channelID string) []string {
	var members []string

	guild, err := dgoSession.Guild(i.GuildID)

	if err != nil {
		return members
	}

	for _, voiceState := range guild.VoiceStates {
		if voiceState.ChannelID == channelID {
			members = append(members, voiceState.UserID)
		}
	}

	return members
}

// RemoveMember Removes value from i in an array
func RemoveMember(s []string, i int) []string {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}

// RandomizeTeams Logic for the randomizeteams command
func RandomizeTeams(dgoSession *discordgo.Session, i *discordgo.InteractionCreate) {
	if len(i.ApplicationCommandData().Options) == 3 {
		guild, guildError := dgoSession.Guild(i.GuildID)

		if guildError != nil {
			print(guildError)
		} else {
			guildRoleMap := make(map[string]*discordgo.Role) // [role_id]role
			for _, role := range guild.Roles {
				guildRoleMap[role.ID] = role
			}

			hasPermission := false
			for _, role := range i.Member.Roles {
				if guildRoleMap[role].Name == "Matchmaker" {
					hasPermission = true
				}
			}

			if hasPermission {
				// Get source and destination arguments
				sourceInput := i.ApplicationCommandData().Options[0].StringValue()
				destOneInput := i.ApplicationCommandData().Options[1].StringValue()
				destTwoInput := i.ApplicationCommandData().Options[2].StringValue()

				channelMap := make(map[string]*discordgo.Channel) // [channel_name]channel

				guildState, guildError := dgoSession.Guild(guild.ID)

				if guildError != nil {
					print(guildError)
				}

				// Map the desired channels
				for _, channel := range guildState.Channels {
					if channel.Name == sourceInput {
						channelMap["source"] = channel
					} else if channel.Name == destOneInput {
						channelMap["team1"] = channel
					} else if channel.Name == destTwoInput {
						channelMap["team2"] = channel
					}
				}

				// Get everyone in the source voice channel
				members := GetVoiceMembers(dgoSession, i, channelMap["source"].ID)

				// Randomize the members
				rand.Seed(time.Now().UnixNano())
				rand.Shuffle(len(members), func(i, j int) { members[i], members[j] = members[j], members[i] })

				teamSize := len(members) / 2

				dgoSession.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Team size: " + strconv.Itoa(teamSize),
					},
				})

				// Ensure we don't panic
				if teamSize >= 1 {
					teamOne := members[:teamSize]
					teamTwo := members[teamSize:]

					for _, player := range teamOne {
						dgoSession.GuildMemberMove(guild.ID, player, &channelMap["team1"].ID)
					}

					for _, player := range teamTwo {
						dgoSession.GuildMemberMove(guild.ID, player, &channelMap["team2"].ID)
					}

					// Build our embed
					embedFields := make([]*discordgo.MessageEmbedField, len(members))

					embedFields[0] = &discordgo.MessageEmbedField{
						Name:  "Team 1",
						Value: "",
					}

					embedFields[1] = &discordgo.MessageEmbedField{
						Name:  "Team 2",
						Value: "",
					}

					for _, member := range teamOne {
						currentMember, _ := dgoSession.GuildMember(guild.ID, member)
						embedFields[0].Value += currentMember.Mention()
					}

					for _, member := range teamTwo {
						currentMember, _ := dgoSession.GuildMember(guild.ID, member)
						embedFields[1].Value += currentMember.Mention()
					}

					embedMessage := &discordgo.MessageEmbed{
						Color:       0x00ff00, // Green
						Description: "Please stay in your assigned voice channels.",
						Title:       "Baby Bread Team Randomizer",
						Fields:      embedFields,
					}

					dgoSession.ChannelMessageSendEmbed(i.ChannelID, embedMessage)

				} else {
					dgoSession.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "You need more than one player to create a match...",
						},
					})
				}

			} else {
				dgoSession.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "You don't have the correct role to do this...",
					},
				})
			}
		}

	}
}
