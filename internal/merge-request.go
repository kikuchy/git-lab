package internal

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/xanzy/go-gitlab"
	"log"
	"net/url"
	"os/exec"
)

func MergeRequestDelegate(settings *GitLabSettings, c *cli.Context) {
	gl := gitlab.NewClient(nil, settings.Token)
	e := gl.SetBaseURL(settings.EndPoint)
	if e != nil {
		log.Fatal(e)
	}

	id := url.QueryEscape(settings.ProjectPath)
	project, _, err := gl.Projects.GetProject(id)
	if err != nil {
		log.Fatal(err)
	}
	projId := *project.ID
	showMergeRequests(gl.MergeRequests, projId, c.Bool("branch"))
}

func showMergeRequests(mrService *gitlab.MergeRequestsService, projectID interface{}, beShowBranch bool) error {
	mrs, _, err := mrService.ListMergeRequests(projectID, &gitlab.ListMergeRequestsOptions{State: "opened"})
	if err != nil {
		log.Fatal(err)
	}
	for i := range mrs {
		mr := mrs[i]
		if beShowBranch {
			fmt.Printf("#%d\t%s\t%s\n", mr.IID, mr.SourceBranch, mr.Title)
		} else {
			fmt.Printf("#%d\t%s\n", mr.IID, mr.Title)
		}
	}
	return nil
}

func checkoutMergeRequestBranch(mr *gitlab.MergeRequest) {
	fmt.Println("checking out ", mr.TargetBranch)
	exec.Command("git", "checkout", mr.TargetBranch)
}
