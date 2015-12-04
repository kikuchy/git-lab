package internal

import (
	//"fmt"
	"log"
	"os/exec"
	"strings"
)

type GitLabSettings struct {
	EndPoint    string
	Token       string
	ProjectPath string
}

func CollectGitLabSettings() *GitLabSettings {
	configs, err := getGitConfig()
	if err != nil {
		log.Fatal(err)
	}
	return &GitLabSettings{
		EndPoint:    getEndPoint(configs),
		Token:       getToken(configs),
		ProjectPath: getProjectPath(configs),
	}
}

func getEndPoint(configs map[string]string) string {
	gitlabUrl, found := configs["gitlab.url"]
	if !found {
		log.Fatal("gitlab URL is not found")
	}
	gitlabUrl = strings.TrimRight(gitlabUrl, "/")
	return gitlabUrl + "/api/v3/"
}

func getToken(configs map[string]string) string {
	gitlabToken, found := configs["gitlab.token"]
	if !found {
		log.Fatal("gitlab private token is not found.")
	}
	return gitlabToken
}

func getProjectPath(configs map[string]string) string {
	gitlabProjectFullPath, found := configs["gitlab.project"]
	if !found {
		log.Fatal("gitlab project path (NAMESPACE/PROJECT_NAME) is not found.")
	}
	return gitlabProjectFullPath
}

func getGitConfig() (map[string]string, error) {
	rawConfigs, err := exec.Command("git", "config", "-l").Output()
	configs := strings.Split(strings.TrimRight(string(rawConfigs), "\n"), "\n")
	if len(configs) < 1 {
		return nil, err
	}
	result := make(map[string]string)
	for _, line := range configs {
		c := strings.Split(string(line), "=")
		result[c[0]] = c[1]
	}
	return result, err
}
