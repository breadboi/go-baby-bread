package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

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

// Removes value from i in array
func RemoveMember(s []string, i int) []string {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}

// Logic for the randomizeteams command
func RandomizeTeams(ctx *dgc.Ctx) {
	if ctx.Arguments.Amount() == 3 {
		guild, guild_error := ctx.Session.Guild(ctx.Event.GuildID)

		if guild_error != nil {
			print(guild_error)
		} else {
			guild_role_map := make(map[string]*discordgo.Role) // [role_id]role
			for _, role := range guild.Roles {
				guild_role_map[role.ID] = role
			}

			has_permission := false
			for _, role := range ctx.Event.Member.Roles {
				if guild_role_map[role].Name == "Matchmaker" {
					has_permission = true
				}
			}

			if has_permission {
				// Get source and destination arguments
				source_input := ctx.Arguments.Get(0).Raw()
				dest_one_input := ctx.Arguments.Get(1).Raw()
				dest_two_input := ctx.Arguments.Get(2).Raw()

				channel_map := make(map[string]*discordgo.Channel) // [channel_name]channel

				guild_state, guild_error := ctx.Session.State.Guild(guild.ID)

				if guild_error != nil {
					print(guild_error)
				}

				// Map the desired channels
				for _, channel := range guild_state.Channels {
					if channel.Name == source_input {
						channel_map["source"] = channel
					} else if channel.Name == dest_one_input {
						channel_map["team1"] = channel
					} else if channel.Name == dest_two_input {
						channel_map["team2"] = channel
					}
				}

				// Get everyone in the source voice channel
				members := GetVoiceMembers(ctx, channel_map["source"].ID)

				// Randomize the members
				rand.Seed(time.Now().UnixNano())
				rand.Shuffle(len(members), func(i, j int) { members[i], members[j] = members[j], members[i] })

				team_size := len(members) / 2

				ctx.RespondText("Team size: " + strconv.Itoa(team_size))

				// Ensure we don't panic
				if team_size >= 1 {
					team_one := members[:team_size]
					team_two := members[team_size:]

					for _, player := range team_one {
						ctx.Session.GuildMemberMove(guild.ID, player, &channel_map["team1"].ID)
					}

					for _, player := range team_two {
						ctx.Session.GuildMemberMove(guild.ID, player, &channel_map["team2"].ID)
					}

					// Build our embed
					embed_fields := make([]*discordgo.MessageEmbedField, len(members))

					embed_fields[0] = &discordgo.MessageEmbedField{
						Name:  "Team 1",
						Value: "",
					}

					embed_fields[1] = &discordgo.MessageEmbedField{
						Name:  "Team 2",
						Value: "",
					}

					for _, member := range team_one {
						current_member, _ := ctx.Session.GuildMember(guild.ID, member)
						embed_fields[0].Value += current_member.Mention()
					}

					for _, member := range team_two {
						current_member, _ := ctx.Session.GuildMember(guild.ID, member)
						embed_fields[1].Value += current_member.Mention()
					}

					embed_message := &discordgo.MessageEmbed{
						Color:       0x00ff00, // Green
						Description: "Please stay in your assigned voice channels.",
						Title:       "Baby Bread Team Randomizer",
						Fields:      embed_fields,
					}

					ctx.RespondEmbed(embed_message)

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
