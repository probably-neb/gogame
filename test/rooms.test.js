const { expect, should: init_should, assert } = require("chai");
init_should();
const {
    URL,
    E,
    openPage,
    createRoom,
    joinRoom,
    textContent,
    waitUntil,
    createAndJoin,
    playGame,
} = require("./lib");
const puppeteer = require("puppeteer");

let browser;
beforeAll(async function () {
    let headless = "new";
    if (process.env.OPEN_BROWSER == 1) {
        headless = false;
    }
    browser = await puppeteer.launch({ headless });
});
afterAll(async function () {
    await browser.close();
});

afterEach(async function () {
    for (const page of await browser.pages()) {
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
    let entry = await joinPage.$("#" + name);
    expect(entry).to.exist;
});

test("join", async function () {
    let hostName = "host",
        guestName = "guest",
        roomName = "foo";
    let hostPage = await createRoom(browser, roomName, hostName);
    let guestPage = await joinRoom(browser, roomName, guestName);
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
            await textContent(page, E.room.host),
            `host name not present in ${pageType} page`,
        ).to.equal(hostName);
        expect(
            await textContent(page, "#" + guestName),
            `guest name not present in ${pageType} page`,
        ).to.equal(guestName);
    }
    hostPage.close();
    guestPage.close();
});
test("play Tic Tac Toe", async function () {
    let hostName = "host",
        guestName = "guest",
        roomName = "foo";
    let pages = await createAndJoin(browser, roomName, hostName, guestName);
    let [hostPage, guestPage] = pages;
    await chooseGame(hostPage, E.room.play.tic_tac_toe);
});
