# gitscore

Gitscore aggregates all contributions to a github organisations repositories.

It then displays it in a barchart or a list.

## setup

`go get github.com/oem/gitscore`

## usage

There is two binaries you can use: gitscore and gitscore-dashboard.

`gitscore` returns a highscore list, `gitscore-dashboard` a dashboard with charts helping to visualize the contributions.

Both commands take the "token" parameter, which is your github token. Keep in mind that this token determines what repos can be accessed.

You also neeed to provide the "orga" parameter, which is the github organisation you want the aggregated stats for.

`gitscore --token <github token> --orga <github organisation>`

`gitscore-dashboard --token <github token> --orga <github organisation>`
