# gitscore

Gitscore aggregates all contributions to a github organisations repositories.

It then displays it in a barchart or a list.

Example barchart for the golang organisation(on github):

![golang organisation example](/example-golang.png)

## setup

`go get -v github.com/oem/gitscore/...`

`go install github.com/oem/gitscore/cmds/gitscore`

`go install github.com/oem/gitscore/cmds/gitscore-dashboard`

## usage

You will need a github token to use gitscore. This token also limits what gitscore can create stats for.

github provides a simple guide on how to create your token:

[github: create a personal token](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/)

### via docker

You can easily run gitscore in a docker container. There is an image ready for usage:

`docker pull oembot/gitscore`

And then simply run it:

#### dashboard

`docker run --rm -ti oembot/gitscore dashboard --token <your github token> --orga <the github organisation you are interested in>`
#### simple list

`docker run --rm -ti oembot/gitscore list --token <your github token> --orga <the github organisation you are interested in>`

### If you have a go development environment set up

There is two binaries you can use: gitscore and gitscore-dashboard.

`gitscore` returns a highscore list, `gitscore-dashboard` a dashboard with charts helping to visualize the contributions.

Both commands take the "token" parameter, which is your github token. Keep in mind that this token determines what repos can be accessed.

You also neeed to provide the "orga" parameter, which is the github organisation you want the aggregated stats for.

`gitscore --token <github token> --orga <github organisation>`

`gitscore-dashboard --token <github token> --orga <github organisation>`
