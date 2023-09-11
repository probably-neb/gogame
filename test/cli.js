#!/usr/bin/env bun
import {createRoom, joinRoom, textContent } from "./lib";
import puppeteer from "puppeteer";

async function startBrowser() {
    return await puppeteer.launch({ headless: false });
}
let browser = await startBrowser();


let argv = process.argv.slice(2)
let [cmd] = argv
if (cmd === undefined) {
    console.error("no command given")
    exit(1)
}
switch (cmd) {
    case "lobby":
        try {
        let page = await createRoom(browser, "foo", "host");
        let joinPage = await joinRoom(browser, "foo", "guest");
        } catch (e) {
            console.error(e)
        }
}
