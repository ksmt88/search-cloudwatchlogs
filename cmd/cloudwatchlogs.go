package cmd

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
    "github.com/manifoldco/promptui"
)

type Input struct {
    Profile string
}

func Search(input *Input) {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
        Profile: input.Profile,
    }))

    svc := cloudwatchlogs.New(sess)

    logGroupsInput := cloudwatchlogs.DescribeLogGroupsInput{
        Limit:              nil,
        LogGroupNamePrefix: nil,
    }
    logGroupsOutput, err := svc.DescribeLogGroups(&logGroupsInput)
    if err != nil {
        fmt.Println(err)
    }

    var items []string
    for _, group := range logGroupsOutput.LogGroups {
        items = append(items, *group.LogGroupName)
    }
    promptSelect := promptui.Select{
        Label: "Select LogGroup",
        Items: items,
    }

    _, logGroup, err := promptSelect.Run()

    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
        return
    }

    promptInput := promptui.Prompt{
        Label: "Input search text",
    }

    searchText, err := promptInput.Run()

    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
        return
    }

    filter := searchText
    filterLogEventsInput := cloudwatchlogs.FilterLogEventsInput{
        // EndTime:    nil,
        FilterPattern: &filter,
        LogGroupName:  &logGroup,
        // StartTime:  nil,
    }
    filterLogEventsOutput, err := svc.FilterLogEvents(&filterLogEventsInput)
    if err != nil {
        fmt.Println(err)
    }

    for _, event := range filterLogEventsOutput.Events {
        fmt.Println(*event.Message, *event.Timestamp)
    }
}
