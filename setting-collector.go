package main

import (
	"errors"
	"os/exec"
	"strings"
)

type GitLabSettings struct {
	EndPoint    string
	Token       string
	ProjectPath string
}

func CollectGitLabSettings() (*GitLabSettings, error) {
	configs, err := getGitConfig()
	if err != nil {
		return nil, err
	}
	ep, err := getEndPoint(configs)
	if err != nil {
		return nil, err
	}
	token, err := getToken(configs)
	if err != nil {
		return nil, err
	}
	path, err := getProjectPath(configs)
	if err != nil {
		return nil, err
	}
	return &GitLabSettings{
		EndPoint:    ep,
		Token:       token,
		ProjectPath: path,
	}, nil
}

func getEndPoint(configs map[string]string) (string, error) {
	gitlabUrl, found := configs["gitlab.url"]
	if !found {
		return "", errors.New("gitlab URL is not found")
	}
	gitlabUrl = strings.TrimRight(gitlabUrl, "/")
	return gitlabUrl + "/api/v3/", nil
}

func getToken(configs map[string]string) (string, error) {
	gitlabToken, found := configs["gitlab.token"]
	if !found {
		return "", errors.New("gitlab private token is not found.")
	}
	return gitlabToken, nil
}

func getProjectPath(configs map[string]string) (string, error) {
	gitlabProjectFullPath, found := configs["gitlab.project"]
	if !found {
		return "", errors.New("gitlab project path (NAMESPACE/PROJECT_NAME) is not found.")
	}
	return gitlabProjectFullPath, nil
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
