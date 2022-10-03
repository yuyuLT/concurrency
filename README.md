# search-wordpress

The goal is to collect WordPress sites by traversing URLs in a parallel process in the Go language.

## Original development environment

https://github.com/platetech/docker-golang-boilerplate

## How to develop (quote) 
1. Write codes in {ProjectRoot}/app dir.

    - Set main script name to main.go.

    - Command go mod get or go mod tidy is unnecessary.

      (Only write import section in codes)

1. Copy .env.template to .env and configure it.

1. Execute ./init.sh shell script.

1. Execute ./run.sh shell script.

1. Modify codes in {ProjectRoot}/app dir.

1. Return to step 4.
