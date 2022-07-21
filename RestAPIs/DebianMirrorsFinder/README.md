# Fastest Debian Mirrors Finder API
## Problem
Build a REST service that returns the information of the fastest mirror to download a given OS from a huge list of 
mirrors. List: https://www.debian.org/mirror/list

## Design:
HTTP Verb: GET
PATH: /fastest-mirror
ACTION: fetch
RESOURCE: URL: string

## Implementation
1. create project directory, add `mirrorFinder` directory
2. create `main.go` file inside `mirrorFinder` directory
3. create `mirrors` directory for packaging the list of mirrors
4. add `data.go` file in mirrors directory