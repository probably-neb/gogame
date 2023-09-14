build:
    templ generate
    go build
    ./gogame

watch:
    ls **.templ **.go | entr -rc just build

test:
    bun test

test-open:
    OPEN_BROWSER=1 bun test

test-watch:
    bun test --watch

fmt:
    #!/bin/bash
    dirs=$(find . -type f -name "*.go" -exec dirname {} \; | sort -u)
    go fmt $dirs
