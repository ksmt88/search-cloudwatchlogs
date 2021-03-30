package main

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func main() {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    svc := cloudwatchlogs.New(sess)

    logGroupsInput := cloudwatchlogs.DescribeLogGroupsInput{
        Limit:              nil,
        LogGroupNamePrefix: nil,
        NextToken:          nil,
    }
    logGroupsOutput, err := svc.DescribeLogGroups(&logGroupsInput)
    if err != nil {
        fmt.Println(err)
    }

    for _, group := range logGroupsOutput.LogGroups {
        fmt.Println(*group.LogGroupName)
    }

    filter := ""
    filterLogEventsInput := cloudwatchlogs.FilterLogEventsInput{
        // EndTime:    nil,
        FilterPattern: &filter,
        LogGroupName:  logGroupsOutput.LogGroups[2].LogGroupName,
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
