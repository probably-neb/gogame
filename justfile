build:
    templ generate
    go run server.go

watch:
    ./browser-sync.js

test:
    pnpx mocha --exit

test-open:
    OPEN_BROWSER=1 pnpx mocha --exit

test-watch:
    pnpx mocha --watch
