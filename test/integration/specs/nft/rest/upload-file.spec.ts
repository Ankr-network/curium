import {expect} from 'chai'
import {API, bluzelle} from "bluzelle";
import {promises} from "fs"


import {getBzClient} from "../../../helpers/bluzelle-client";
import * as fs from "fs";
import * as path from "path";

describe('file upload', function()  {
    this.timeout(10000);
    let bz: API

    beforeEach(() => bz = getBzClient());


    it('should create a directory nft-upload in .blzd', () => {
        return fetch(`http://localhost:1317/nft/upload/someHash/1`, {
            method: 'POST',
            body: "1"
        })
            .then(() => fs.existsSync(path.resolve(__dirname, `${process.env.HOME}/.blzd/nft-upload`)))
            .then(resp => expect(resp).to.be.true)
    });

    it('should create a file containing the chunk number', () => {
        let hash = Date.now().toString();
        return fetch(`http://localhost:1317/nft/upload/${hash}/1`, {
            method: 'POST',
            body: "1"
        })
            .then(() => fs.existsSync(path.resolve(__dirname, `${process.env.HOME}/.blzd/nft-upload/${hash}-0001`)))
            .then(resp => expect(resp).to.be.true)
            .then(() => fs.readFileSync(path.resolve(__dirname, `${process.env.HOME}/.blzd/nft-upload/${hash}-0001`)))
            .then(resp => expect(new TextDecoder().decode(resp)).to.equal("1"))
    });

});
