const { expect, should: init_should, assert } = require("chai");
init_should();
const puppeteer = require("puppeteer");

var { spawn } = require("child_process");

const URL = "http://localhost:8080/";
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
        console.log("killing server");
        server_proc.kill();
    } else {
        server_proc = just("build");
        server_proc.on("close", () => {
            server_proc = just("build");
        });
    }
}
const E = {
    home: {
        name: "#display-name",
        create_room: "#create-room",
        join_room: "#join-room",
    },
    create_room: {
        open: "#create-room",
        modal: "#create-room-modal",
        room_name: "#room-name",
        display_name: "#display-name",
        create: "#create", // TODO: add this id
        cancel: "#cancel", // TODO: add this id
    },
    join_room: {
        join: "#join-room",
        rooms: "ul",
        entries: "li",
        entry: {
            name: "h3",
            join: "#join-room",
        },
        modal: {
            id: "#join-room-modal",
            display_name: "#display-name",
            join: "#join",
        },
    },
    room: {
        ws: "#ws-connection",
        name: "h3",
        players: "#players",
        host: "#host",
    },
};

async function openPage(browser) {
    let page = await browser.newPage();
    await page.goto(URL);
    await page.waitForSelector(E.create_room.open);
    return page;
}
async function createRoom(browser, name, hostName) {
    let page = await openPage(browser);
    let open = await page.$(E.create_room.open);
    expect(open).exists.and.to.have.property("click");
    await open.click();
    await page.waitForSelector(E.create_room.modal);
    let modal = await page.$(E.create_room.modal);
    if (!name) return page;
    await page.type(E.create_room.room_name, name);
    await page.type(E.create_room.display_name, hostName);
    await page.click(E.create_room.create);
    await page.waitForSelector(E.room.ws);
    expect(page.$(E.room.name)).exists;
    expect(page.url()).to.include("rooms");
    return page;
}

async function joinRoom(browser, name, displayName) {
    let page = await openPage(browser);
    await page.click(E.join_room.join);
    await page.waitForSelector(E.join_room.rooms);
    entries = await page.$$(E.join_room.entries);
    // console.error(await page.content())
    expect(entries).to.have.lengthOf.at.least(1, "did not find room entries");
    let entry = await page.$("#" + name);
    expect(entry).to.exist;
    let join = await entry.$(E.join_room.entry.join);
    await join.click();
    await page.waitForSelector(E.join_room.modal.id);
    expect(await page.$(E.join_room.modal.id)).to.exist;
    await page.type(E.join_room.modal.display_name, displayName);
    expect(await page.$(E.join_room.modal.join)).to.exist;
    await page.click(E.join_room.modal.join);
    expect(page.$(E.room.name)).exists;
    let modalDissapeared = await waitUntil(async () => {
        // model doesn't exist
        let modal = await page.$(E.join_room.modal.id);
        return !modal || (await textContent(modal)) == "";
    });
}
async function textContent(p, e) {
    if (!!e && typeof e == "string") {
        let s = e;
        e = await p.$(e);
        expect(e, "could not find element with selector: " + s).to.exist;
    } else {
        e = p;
    }
    let hndl = await e.getProperty("textContent");
    return await hndl.jsonValue();
}
async function waitUntil(cb, timeout, interval = 10) {
    if (await cb()) return true;
    let time = 0;
    const intervalId = setInterval(async () => {
        if (!!(await cb())) {
            clearInterval(intervalId);
            return true;
        }
        time += interval;
        if (time >= timeout) {
            return false;
        }
    }, interval);
}

describe("ui", async function () {
    let browser;
    before(async function () {
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

    describe.skip("room", async function () {
        it("create", async function () {
            let page = await createRoom(browser, "foo", "host");
        });
        it("join", async function () {
            let hostName = "host",
                guestName = "guest",
                roomName = "foo";
            let hostPage = await createRoom(browser, roomName, hostName);
            let guestPage = await joinRoom(browser, roomName, guestName);
            for (const page of [hostPage, guestPage]) {
                let pageType = page.url().includes("host=true")
                    ? "host"
                    : "guest";
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
        it("start", async function () {
            let hostName = "host",
                guestName = "guest",
                roomName = "foo";
            let page = await createRoom(browser, roomName, hostName);
            let joinPage = await joinRoom(browser, roomName, guestName);
        });
    });
});
