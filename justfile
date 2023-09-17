build:
    templ generate
    bun x tailwindcss -i ./assets/_index.css -o ./assets/tailwind.css
    go build
    ./gogame

watch:
    find -path '*.go' -and -not -path '*_templ.go' -or -path '*.templ' | entr -rc just build

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

build-browser-sync:
    templ generate
    bun x tailwindcss -i ./assets/_index.css -o ./assets/tailwind.css
    go build
    ./gogame &
    browser-sync start --no-open --proxy "localhost:8080"

browser-sync:
    ls **.templ **.go | entr -rc just build-browser-sync

