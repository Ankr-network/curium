import {
    decodeData,
    DEFAULT_TIMEOUT,
    defaultLease,
    encodeData,
    getSdk,
    newSdkClient
} from "../../helpers/client-helpers/sdk-helpers";
import {useChaiAsPromised} from "testing/lib/globalHelpers";
import {BluzelleSdk, DbSdk} from "../../../src/bz-sdk/bz-sdk";
import {expect} from "chai";
import {Lease} from "../../../src/codec/crud/lease";
import {createKeys} from "../../helpers/client-helpers/sdk-helpers";
import {defaultRegistryTypes} from "@cosmjs/stargate";

describe('tx.Delete()', function () {
    this.timeout(DEFAULT_TIMEOUT);

    let sdk: BluzelleSdk;
    let uuid: string;
    beforeEach(async () => {
        useChaiAsPromised();
        sdk = await getSdk();
        uuid = Date.now().toString()
    });

    // it('should resolve to chain information', async () => {
    //     await bz.create('myKey', 'someValue', defaultGasParams());
    //     const result = await bz.delete('myKey', defaultGasParams())
    //     expect(result.txhash).to.be.a('string');
    //     expect(result.height).to.be.a('number');
    // });

    it('should delete a key in the database', async () => {
        await sdk.db.tx.Create({
            creator: sdk.db.address,
            uuid,
            key: 'myKeys',
            value: new TextEncoder().encode('myValue'),
            lease: defaultLease,
            metadata: new Uint8Array()
        });
        expect(await sdk.db.tx.Read({
            creator: sdk.db.address,
           uuid,
            key: 'myKeys'
        }).then(resp => resp.value).then(decodeData)).to.equal('myValue');
        await sdk.db.tx.Delete({
            creator: sdk.db.address,
            uuid,
            key: 'myKeys'
        });
        await expect(sdk.db.tx.Delete({
            creator: sdk.db.address,
            uuid,
            key: 'myKeys'
        })).to.be.rejectedWith(/key not found/);
    });

    it('should be able to delete an empty value', async () => {
        await sdk.db.tx.Create({
            creator: sdk.db.address,
           uuid,
            key: 'emptyValue',
            value: encodeData(''),
            lease: defaultLease,
            metadata: new Uint8Array()
        });

        await sdk.db.tx.Delete({
            creator: sdk.db.address,
            uuid,
            key: 'emptyValue'
        });

        await expect(sdk.db.tx.Read({
            creator: sdk.db.address,
            uuid,
            key: 'emptyValue'
        })).to.be.rejectedWith(/key not found/);

    })

    it('should throw an error if a key does not exist', async () => {
        expect(
            sdk.db.tx.Delete({
                creator: sdk.db.address,
                uuid,
                key: 'voided'
            })
        ).to.be.rejectedWith(/Key does not exist/);
    });


    it('should handle parallel deletes', async () => {
        const {keys} = await createKeys(sdk.db, 5, uuid);
        await Promise.all(keys.map(key => sdk.db.tx.Delete({
            creator: sdk.db.address,
            uuid,
            key
        })));
        expect(await sdk.db.tx.KeyValues({
            creator: sdk.db.address,
            uuid,
        }).then(val => val.keyValues)).to.have.length(0);
    });

    it("should only be able to delete someone else's key-value", async () => {

        const otherSdk = await newSdkClient(sdk);

        await sdk.db.tx.Create({
            creator: sdk.db.address,
            uuid,
            key: 'myKey',
            value: encodeData('myValue'),
            lease: defaultLease,
            metadata: new Uint8Array()
        });

        await expect(otherSdk.db.tx.Delete({
            creator: otherSdk.db.address,
            uuid,
            key: 'myKey'
        })).to.be.rejectedWith(/incorrect owner/)

        expect(await sdk.db.tx.Read({
            creator: sdk.db.address,
            uuid,
            key: 'myKey'
        }).then(resp => decodeData(resp.value))).to.equal('myValue')

    })
});