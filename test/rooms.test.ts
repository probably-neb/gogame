import { expect, assert } from "chai";
import {beforeAll, afterAll, afterEach, test } from "bun:test"
const {
    URL,
    E,
    openPage,
    createRoom,
    joinRoom,
    textContent,
    waitUntil,
    createAndJoin,
    chooseGame,
} = require("./lib");
const puppeteer = require("puppeteer");

let browser;
beforeAll(async function () {
    let headless = "new";
    if (process.env.OPEN_BROWSER === "1") {
        headless = "false";
    }
    browser = await puppeteer.launch({ headless });
});
afterAll(async function () {
    await browser.close();
});

afterEach(async function () {
    for (const page of await browser.pages()) {
        if (page.isClosed()) {
            continue;
        }
        await page.close();
    }
});

test("create", async function () {
    await createRoom(browser, "foo", "host");
    let joinPage = await openPage(browser);
    await joinPage.click(E.join_room.join);
    await joinPage.waitForSelector(E.join_room.rooms);
    let entries = await joinPage.$$(E.join_room.entries);
    // console.error(await page.content())
    expect(entries).to.have.lengthOf.at.least(1, "did not find room entries");
    let entry = await joinPage.$("#" + "foo");
    expect(entry).to.exist;
});

test("join", async function () {
    let hostName = "host",
        guestName = "guest",
        roomName = "foo";
    await createRoom(browser, roomName, hostName);
    await joinRoom(browser, roomName, guestName);
});
test("create and join", async function () {
    let hostName = "host",
        guestName = "guest",
        roomName = "foo";
    let pages = await createAndJoin(browser, roomName, hostName, guestName);
    for (let i = 0; i < pages.length; i++) {
        let page = pages[i];
        let pageType = "guest";
        if (i == 0) {
            pageType = "host";
        }
        expect(
            await textContent(await page.$(E.room.host)),
            `host name not present in ${pageType} page`,
        ).to.equal(hostName);
        expect(
            await textContent(await page.$("#" + guestName)),
            `guest name not present in ${pageType} page`,
        ).to.equal(guestName);
    }
});
test("find Tic Tac Toe", async function () {
    let hostName = "host",
        guestName = "guest",
        roomName = "foo";
    let pages = await createAndJoin(browser, roomName, hostName, guestName);
    let [hostPage, guestPage] = pages;
    let game = E.room.play.tic_tac_toe
    let startGame = await hostPage.$(game);
    if (!startGame) {
        throw new Error("could not find game: " + game)
    }
    // FIXME: this causes TargetClosed Error if it is ran
    // await startGame.click()
});
test("promote to host", async function () {
    let hostName = "host",
        guestName = "guest",
        roomName = "foo";
    let [hostPage, guestPage] = await createAndJoin(
        browser,
        roomName,
        hostName,
        guestName,
    );
    await hostPage.close();
    let guestBecameHost = await waitUntil(async () => {
        (await textContent(await guestPage.$(E.room.host))) == guestName;
    });
    expect(
        guestBecameHost,
        "guest did not become host after host disconnected",
    );
});

test("disconnect guest", async function() {
    let hostName = "host",
        guestName = "guest",
        roomName = "foo";
    let [hostPage, guestPage] = await createAndJoin(
        browser,
        roomName,
        hostName,
        guestName,
    );
    await guestPage.close();
    let guestNoLongerListed = await waitUntil(async () => {
        (await hostPage.$('#' + guestName)) == null
    });
    expect(
        guestNoLongerListed,
        "guest did not get removed from guest list",
    );
})
