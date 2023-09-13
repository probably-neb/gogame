build:
    templ generate
    go build
    ./gogame

watch:
    ls **.templ **.go | entr -rc just build

test:
    bun test

test-open:
    OPEN_BROWSER=1 pnpx mocha --exit

test-watch:
    pnpx mocha --watch

fmt:
    #!/bin/bash
    dirs=$(find . -type f -name "*.go" -exec dirname {} \; | sort -u)
    go fmt $dirs
