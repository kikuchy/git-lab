# git-lab

Yet another gitlab command

## Requires

* `git` command is on your `PATH`
* `$GOPATH/bin` is in your `PATH`


## Install

```
go get github.com/kikuchy/git-lab
```


## Usage

"git-lab" adds `git lab` subcommand for your git.


### Configure

Configure below using `git config (--global)`.

* `gitlab.url` ... Gitlab's base URL.
* `gitlab.token` ... Gitlab private token.
* `gitlab.project` ... Gitlab project ID. `NAMESPACE/PROJECT_NAME` format or numeric id for API calling. It's necessary to setting every project.

For example,

```
cd path/to/existing/project
git config --global gitlab.url https://gitlab.mycompany.com/
git configure --global gitlab.token THIS_IS_MY_TOKEN
git config gitlab.project mygroup/myproject
```


### `merge-request`

Show opend Merge Requests.

```
git lab merge-request
```

* Switches
	* `-b` ... Shows branch name of the Merge Request.


... and WIP :runner: