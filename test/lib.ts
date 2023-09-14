import { expect, assert } from "chai";
import * as ppt from "puppeteer"

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
export async function openPage(browser: ppt.Browser) {
    let page = await browser.newPage();
    await page.goto(URL);
    await page.waitForSelector(E.create_room.open);
    return page;
}
export async function createRoom(browser: ppt.Browser, name: string, hostName: string): Promise<ppt.Page> {
    let page = await openPage(browser);
    let open = await page.$(E.create_room.open);
    expect(open).to.exist.and.to.have.property("click");
    await open!.click();
    await page.waitForSelector(E.create_room.modal);
    let modal = await page.$(E.create_room.modal);
    expect(modal).to.exist
    await page.type(E.create_room.room_name, name);
    await page.type(E.create_room.display_name, hostName);
    let create = await page.$(E.create_room.create);
    expect(create).to.exist.and.to.have.property("click");
    await create!.click()
    let createResponse = await page.waitForResponse((_res) => true, {timeout: 500})
    assert(createResponse.ok(), `create response has body: ${await createResponse.text()}`)
    let urlIncludesName = await waitUntil(async () => {
        return page.url().includes(name);
    }, 1000, 10)
    assert(urlIncludesName, "url does not include room name")
    // TODO: hostName present
    return page;
}

export async function joinRoom(browser: ppt.Browser, name: string, displayName: string): Promise<ppt.Page> {
    let page = await openPage(browser);
    await page.click(E.join_room.join);
    await page.waitForSelector(E.join_room.rooms);
    let entries = await page.$$(E.join_room.entries);
    // console.error(await page.content())
    expect(entries).to.have.lengthOf.at.least(1, "did not find room entries");
    let entry = await page.$("#" + name);
    expect(entry).to.exist;
    let join = await entry!.$(E.join_room.entry.join);
    if (!join) {
        throw new Error(`join button [${E.join_room.entry.join}] not found`)
    }
    await join.click();
    await page.waitForNavigation({waitUntil: "load"});
    expect(await page.$(E.join_room.modal.id)).to.exist;
    await page.type(E.join_room.modal.display_name, displayName);
    expect(await page.$(E.join_room.modal.join)).to.exist;
    await page.click(E.join_room.modal.join);
    expect(page.$(E.room.name)).to.exist;
    let modalDissapeared = await waitUntil(async () => {
        // model doesn't exist
        let modal = await page.$(E.join_room.modal.id);
        return !modal || (await textContent(modal)) == "";
    });
    expect(modalDissapeared, "modal did not go away")
    return page
}

export async function createAndJoin(browser: ppt.Browser, roomName = "room", hostName = "host", guestName = "guest") {
    let hostPage = await createRoom(browser, roomName, hostName);
    let guestPage = await joinRoom(browser, roomName, guestName);
    expect(hostPage, "hostPage does not exist").to.exist
    expect(guestPage, "guestPage does not exist").to.exist
    return [hostPage, guestPage];
}

export async function chooseGame(page: ppt.Page, game: string) {
    let startGame = await page.$(game);
    if (!startGame) {
        throw new Error("could not find game: " + game)
    }
    // if not probably not passed host page
    let errMsg = "cannot find button to start game: " + game.replace("start-", "") + " Are you sure this player is the host?";
    expect(startGame,errMsg).to.exist;
    await startGame.click();
}

export async function textContent(e: ppt.ElementHandle) {
    if (!e) {
        throw new Error("attempt to get textContent of nonexistant element")
    }
    let hndl = await e.getProperty("textContent");
    return await hndl.jsonValue();
}
export async function waitUntil(cb: () => boolean | Promise<boolean>, timeout = 1000, interval = 10) {
    if (await cb()) return true;
    let time = 0;
    return new Promise((res) => {
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
