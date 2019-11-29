package core

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codecommit"
	"os"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func ListRepos(awsProfile, awsRegion string) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(awsRegion)},
		Profile: awsProfile,
	})
	if err != nil {
		exitErrorf("Unable to establish Session, %v", err)
	}
	_, sessErr := sess.Config.Credentials.Get()
	if sessErr != nil {
		exitErrorf("Unable to utilize Session, %v", sessErr)
	}
	svc := codecommit.New(sess)
	result, err := svc.ListRepositories(nil)
	if err != nil {
		exitErrorf("Unable to list buckets, %v", err)
	}

	fmt.Println("Repos:")

	for _, b := range result.Repositories {
		fmt.Printf("* %s id %s\n",
			aws.StringValue(b.RepositoryName), aws.StringValue(b.RepositoryId))
	}
}
