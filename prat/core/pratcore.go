package core

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	. "github.com/aws/aws-sdk-go/service/codecommit"
	"os"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func GetSession(awsProfile, awsRegion string) *CodeCommit {
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
	return New(sess)
}
func ListPullRequest(svc *CodeCommit, repositoryName string) {
	input := ListPullRequestsInput{
		//AuthorArn:         nil,
		//MaxResults:        nil,
		//NextToken:         nil,
		//PullRequestStatus: nil,
		RepositoryName:    aws.String(repositoryName),
	}
	result, err := svc.ListPullRequests(&input)
	if err != nil {
		exitErrorf("Unable to list PRs, %v", err)
	}

	fmt.Printf("Repos: %v\n", result.PullRequestIds)

	for _, b := range result.PullRequestIds {
		fmt.Printf("* id %v\n", *b)
	}
}

func GetPullRequest(svc *CodeCommit, pullRequestId string) {
	input := GetPullRequestInput{
		// id's are unique to account/region so we don't need to specify repo
		PullRequestId: aws.String(pullRequestId),
	}
	result, err := svc.GetPullRequest(&input)
	if err != nil {
		exitErrorf("Unable get PR, %v", err)
	}

	fmt.Printf("Pull Request: %v\n", result.PullRequest)

}

func ListRepos(svc *CodeCommit) {
	result, err := svc.ListRepositories(nil)
	if err != nil {
		exitErrorf("Unable to list repos, %v", err)
	}

	fmt.Println("Repos:")

	for _, b := range result.Repositories {
		fmt.Printf("* %s id %s\n",
			aws.StringValue(b.RepositoryName), aws.StringValue(b.RepositoryId))
	}
}

func CreatePullRequest(session *CodeCommit, title, desc, sourceBranch, destBranch, repository string) {
	input := CreatePullRequestInput{
		ClientRequestToken: nil,
		Description:        aws.String(desc),
		Targets: []*Target{
			&Target{
				DestinationReference: aws.String(destBranch),
				RepositoryName:       aws.String(repository),
				SourceReference:      aws.String(sourceBranch),
			},
		},
		Title: aws.String(title),
	}
	result, err := session.CreatePullRequest(&input)
	if err != nil {
		exitErrorf("Error response from CreatePullRequest, %v", err)
	}
	//{
	//PullRequest: {
	//ApprovalRules: [],
	//	AuthorArn: "arn:aws:iam::193391773412:user/samkeen",
	//		ClientRequestToken: "4CD53B5C-CDE7-4745-9A2E-254FD1983AC3",
	//		CreationDate: 2019-12-17 04:31:09 +0000 UTC,
	//		Description: "made some changes",
	//		LastActivityDate: 2019-12-17 04:31:09 +0000 UTC,
	//		PullRequestId: "3",
	//		PullRequestStatus: "OPEN",
	//		PullRequestTargets: [{
	//		DestinationCommit: "a8b3fd2df078453751b3a12e736ef887f673618a",
	//		DestinationReference: "refs/heads/master",
	//		MergeBase: "a8b3fd2df078453751b3a12e736ef887f673618a",
	//		MergeMetadata: {
	//			IsMerged: false
	//		},
	//		RepositoryName: "CodePipelineTest",
	//		SourceCommit: "b64e8b8bec2bbad9981dc93512f41cef0feb8376",
	//		SourceReference: "refs/heads/test-branch"
	//	}],
	//RevisionId: "0a74a03d5f9a7fd6e31f1be85448832d9979a8850d4c53be0f284775df3f1796",
	//Title: "testing PR"
	//}
	//}
	fmt.Printf("Create PR result: %v\n", result)
}
