build:
    templ generate
    go run .

watch:
    ls **.templ **.go | entr -rc just build

test:
    bun test

test-open:
    OPEN_BROWSER=1 pnpx mocha --exit

test-watch:
    pnpx mocha --watch
