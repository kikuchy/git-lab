package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/xanzy/go-gitlab"
	"log"
	"net/url"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func CheckoutDelegate(settings *GitLabSettings, c *cli.Context) {
	branchIdentity := c.Args()[0]
	startsWithIID := regexp.MustCompile(`^#[0-9]+$`)
	if startsWithIID.MatchString(branchIdentity) {
		iid, _ := strconv.Atoi(strings.TrimPrefix(branchIdentity, "#"))
		checkoutMergeRequestBranch(settings, iid)
	} else {
		checkout(branchIdentity)
	}
}

func checkoutMergeRequestBranch(settings *GitLabSettings, mergeRequestIID int) {
	gl := gitlab.NewClient(nil, settings.Token)
	e := gl.SetBaseURL(settings.EndPoint)
	if e != nil {
		log.Println(e)
		return
	}

	id := url.QueryEscape(settings.ProjectPath)
	project, _, err := gl.Projects.GetProject(id)
	if err != nil {
		log.Println(err)
		return
	}
	projId := *project.ID

	mr, _, err := gl.MergeRequests.GetMergeRequest(projId, mergeRequestIID)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("checking out ", mr.TargetBranch)
	checkout(mr.TargetBranch)
}

func checkout(refspec string) {
	exec.Command("git", "checkout", refspec)
}
