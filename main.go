package main

import (
	"context"
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	ctx := context.Background()

	// Header Section
	headerSection := slack.NewSectionBlock(
		slack.NewTextBlockObject("mrkdwn", "You have a new request:\n*<fakeLink.toEmployeeProfile.com|Fred Enriquez - New device request>*", false, false),
		nil,
		nil,
	)

	// Fields
	fieldsSection := slack.NewSectionBlock(
		nil,
		[]*slack.TextBlockObject{
			slack.NewTextBlockObject("mrkdwn", "*Type:*\nComputer (laptop)", false, false),
			slack.NewTextBlockObject("mrkdwn", "*When:*\nSubmitted Aut 10", false, false),
			slack.NewTextBlockObject("mrkdwn", "*Last Update:*\nMar 10, 2015 (3 years, 5 months)", false, false),
			slack.NewTextBlockObject("mrkdwn", "*Reason:*\nAll vowel keys aren't working.", false, false),
			slack.NewTextBlockObject("mrkdwn", "*Specs:*\n\"Cheetah Pro 15\" - Fast, really fast\"", false, false),
		},
		nil,
	)

	slackToken := os.Getenv("SLACK_API_TOKEN")
	slackChannel := os.Getenv("SLACK_CHANNEL")
	slackBotName := os.Getenv("SLACK_BOT_NAME")
	slackClient := slack.New(slackToken)

	_, _, err := slackClient.PostMessageContext(ctx, slackChannel,
		slack.MsgOptionUsername(slackBotName),
		slack.MsgOptionBlocks(headerSection, fieldsSection),
		slack.MsgOptionAttachments(
			slack.Attachment{
				Color: "#FFC0CB",
				//Color: "good",		// good/warning/danger do not work ??
				Blocks: slack.Blocks {
					BlockSet: []slack.Block{
						fieldsSection,
					},
				},
			},
		),
	)
	if err != nil {
		fmt.Println(err)
	}
}
