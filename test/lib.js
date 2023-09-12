const { expect, should: init_should, assert } = require("chai");
init_should();

export const URL = "http://localhost:8080/";

export const E = {
    home: {
        name: "#name",
        create_room: "#create-room",
        join_room: "#join-room",
    },
    session: "#session",
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
        play: {
            tic_tac_toe: "#start-tic-tac-toe",
        }
    },
};
export async function openPage(browser) {
    let page = await browser.newPage();
    await page.goto(URL);
    await page.waitForSelector(E.create_room.open);
    return page;
}
export async function createRoom(browser, name, hostName) {
    let page = await openPage(browser);
    let open = await page.$(E.create_room.open);
    expect(open).exists.and.to.have.property("click");
    await open.click();
    await page.waitForSelector(E.create_room.modal);
    let modal = await page.$(E.create_room.modal);
    if (!name) return page;
    await page.type(E.create_room.room_name, name);
    await page.type(E.create_room.display_name, hostName);
    let create = await page.$(E.create_room.create);
    expect(create).exists.and.to.have.property("click");
    await create.click()
    await waitUntil(async () => {
        return page.url().includes(name);
    }, 100, 10)
    // await page.waitForSelector(E.room.ws);
    // expect(page.$(E.room.name)).exists;
    // expect(page.url()).to.include(name);
    return page;
}

export async function joinRoom(browser, name, displayName) {
    let page = await openPage(browser);
    await page.click(E.join_room.join);
    await page.waitForSelector(E.join_room.rooms);
    let entries = await page.$$(E.join_room.entries);
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

export async function createAndJoin(browser, roomName = "room", hostName = "host", guestName = "guest") {
    let hostPage = await createRoom(browser, roomName, hostName);
    let guestPage = await joinRoom(browser, roomName, guestName);
    return [hostPage, guestPage];
}

export async function chooseGame(page, game) {
    let startGame = await page.$(game);
    // if not probably not passed host page
    let errMsg = "cannot find button to start game: " + game.replace("start-", "") + " Are you sure this player is the host?";
    expect(startGame,errMsg).to.exist;
    await startGame.click();
}

export async function textContent(p, e) {
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
export async function waitUntil(cb, timeout = 1000, interval = 10) {
    if (await cb()) return true;
    let time = 0;
    return new Promise((res, rej) => {
        const intervalId = setInterval(async () => {
            if (!!(await cb())) {
                clearInterval(intervalId);
                res(true)
            }
            time += interval;
            if (time >= timeout) {
                clearInterval(intervalId);
                res(false);
            }
        }, interval);
    });
}
