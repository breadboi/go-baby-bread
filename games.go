package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

// GetVoiceMembers Returns a list of member IDs as strings from the channel
func GetVoiceMembers(ctx *dgc.Ctx, channelID string) []string {
	var members []string
	guild, err := ctx.Session.State.Guild(ctx.Event.GuildID)

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
func RandomizeTeams(ctx *dgc.Ctx) {
	if ctx.Arguments.Amount() == 3 {
		guild, guildError := ctx.Session.Guild(ctx.Event.GuildID)

		if guildError != nil {
			print(guildError)
		} else {
			guildRoleMap := make(map[string]*discordgo.Role) // [role_id]role
			for _, role := range guild.Roles {
				guildRoleMap[role.ID] = role
			}

			hasPermission := false
			for _, role := range ctx.Event.Member.Roles {
				if guildRoleMap[role].Name == "Matchmaker" {
					hasPermission = true
				}
			}

			if hasPermission {
				// Get source and destination arguments
				sourceInput := ctx.Arguments.Get(0).Raw()
				destOneInput := ctx.Arguments.Get(1).Raw()
				destTwoInput := ctx.Arguments.Get(2).Raw()

				channelMap := make(map[string]*discordgo.Channel) // [channel_name]channel

				guildState, guildError := ctx.Session.State.Guild(guild.ID)

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
				members := GetVoiceMembers(ctx, channelMap["source"].ID)

				// Randomize the members
				rand.Seed(time.Now().UnixNano())
				rand.Shuffle(len(members), func(i, j int) { members[i], members[j] = members[j], members[i] })

				teamSize := len(members) / 2

				ctx.RespondText("Team size: " + strconv.Itoa(teamSize))

				// Ensure we don't panic
				if teamSize >= 1 {
					teamOne := members[:teamSize]
					teamTwo := members[teamSize:]

					for _, player := range teamOne {
						ctx.Session.GuildMemberMove(guild.ID, player, &channelMap["team1"].ID)
					}

					for _, player := range teamTwo {
						ctx.Session.GuildMemberMove(guild.ID, player, &channelMap["team2"].ID)
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
						currentMember, _ := ctx.Session.GuildMember(guild.ID, member)
						embedFields[0].Value += currentMember.Mention()
					}

					for _, member := range teamTwo {
						currentMember, _ := ctx.Session.GuildMember(guild.ID, member)
						embedFields[1].Value += currentMember.Mention()
					}

					embedMessage := &discordgo.MessageEmbed{
						Color:       0x00ff00, // Green
						Description: "Please stay in your assigned voice channels.",
						Title:       "Baby Bread Team Randomizer",
						Fields:      embedFields,
					}

					ctx.RespondEmbed(embedMessage)

				} else {
					ctx.RespondText("You need more than one player to create a match...")
				}

			} else {
				ctx.RespondText("You don't have the correct role to do this...")
			}
		}

	} else {
		ctx.RespondText("Too few arguments were provided...")
	}
}
