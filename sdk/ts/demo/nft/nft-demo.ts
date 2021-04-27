import {bluzelle} from "../../src/bz-sdk/bz-sdk";
import {passThroughAwait} from "promise-passthrough";
import {readFile} from "fs/promises";

bluzelle({
    mnemonic:  "foam card blouse leaf convince scrub marble pencil camp hover food install waste aunt minute alarm gauge cabin garbage over kitten jealous draft miracle",
    url: "http://localhost:26657",
    gasPrice: 0.002,
    maxGas: 100000000
})
    .then(bz =>
        readFile("./test.tiff")
            .then(data => bz.helpers.nft.uploadNft({
                mime: 'image/tiff',
                meta: ''
            }, data, (chunk, size) => console.log(chunk, size)))
            .then(({id}) => console.log(id))
    )


