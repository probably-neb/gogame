const { expect, should: init_should, assert } = require("chai");
init_should();
const { URL, E, createRoom, joinRoom, textContent, waitUntil } = require("./lib");
const puppeteer = require("puppeteer");

let browser;
beforeAll(async function () {
    let headless = "new";
    if (process.env.OPEN_BROWSER == 1) {
        headless = false;
    }
    browser = await puppeteer.launch({ headless });
});

afterEach(async function () {
    for (const page of await browser.pages()) {
        await page.close();
    }
});

test("create", async function () {
    let page = await createRoom(browser, "foo", "host");
});
test("join", async function () {
    let hostName = "host",
        guestName = "guest",
        roomName = "foo";
    let hostPage = await createRoom(browser, roomName, hostName);
    let guestPage = await joinRoom(browser, roomName, guestName);
    for (const page of [hostPage, guestPage]) {
        let pageType = page.url().includes("host=true") ? "host" : "guest";
        console.log(page.url(), { pageType });
        expect(
            await textContent(page, E.room.host),
            `host name not present in ${pageType} page`,
        ).to.equal(hostName);
        await page.waitForTimeout(10000);
        expect(
            await textContent(page, "#" + guestName),
            `guest name not present in ${pageType} page`,
        ).to.equal(guestname);
    }
    hostPage.close();
    guestPage.close();
});
test("start", async function () {
    let hostName = "host",
        guestName = "guest",
        roomName = "foo";
    let page = await createRoom(browser, roomName, hostName);
    let joinPage = await joinRoom(browser, roomName, guestName);
});
