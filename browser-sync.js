#!/usr/bin/env node

var bs = require("browser-sync").create();
var { spawn } = require("child_process");

let server_proc = null;

function just(cmd) {
    let proc = spawn("just", [cmd]);
    function strlog(logtype) {
        return function (data) {
            console[logtype](data.toString());
        };
    }
    proc.stdout.on("data", strlog("log"));
    proc.stderr.on("data", strlog("error"));
    return proc;
}

function reload_server() {
    if (server_proc) {
        server_proc.on("exit", () => {
            server_proc = just("build")
        })
        console.log("killing server");
        server_proc.kill();
    } else {
        server_proc = just("build");
    }
}

bs.watch(["**/*.html", "**/*.templ", "**/*.go"], function (event, file) {
    if (file.endsWith("_templ.go"))
        return
    if (file.includes("node_modules"))
        return
    if (event == "add")
        return
    console.log(`[${event}]`, file, "...", "rebuilding");
    reload_server();
    bs.reload();
});

bs.init({
    proxy: {
        target: "localhost:8080",
        ws: false, },
    open: false,
});
reload_server();
