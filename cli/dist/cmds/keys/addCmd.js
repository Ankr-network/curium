"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.promptToOverrideUser = exports.promptForMnemonic = exports.handler = exports.builder = exports.desc = exports.command = void 0;
const sdk_helpers_1 = require("../../helpers/sdk-helpers");
const sdk_js_1 = require("@bluzelle/sdk-js");
exports.command = 'add <user>';
exports.desc = 'Add key to local system and generate mnemonic';
const builder = (yargs) => {
    return yargs
        .option('recover', {
        describe: 'recover account by providing mnemonic',
        type: 'boolean',
        default: false
    })
        .positional('user', {
        describe: 'name of user account to create',
        type: 'string'
    })
        .help();
};
exports.builder = builder;
const handler = (argv) => {
    let yourMnemonic;
    return sdk_helpers_1.makeCliDir()
        .then(() => exports.promptForMnemonic(argv.recover))
        .then(mnemonic => sdk_helpers_1.createUserFile(argv.user, mnemonic, exports.promptToOverrideUser))
        .then(() => sdk_helpers_1.readUserMnemonic(argv.user))
        .then(mnemonic => yourMnemonic = mnemonic)
        .then(sdk_helpers_1.getAccountInfoFromMnemonic)
        .then(info => ({ ...info, mnemonic: yourMnemonic }))
        .then(console.log)
        .then(() => process.exit());
};
exports.handler = handler;
const readline = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout
});
const promptForMnemonic = (recover) => recover ? new Promise((resolve) => readline.question("Please provide BIP39 mnemonic\n", (mnemonic) => {
    readline.pause();
    return resolve(mnemonic);
})) : Promise.resolve(sdk_js_1.newMnemonic());
exports.promptForMnemonic = promptForMnemonic;
const promptToOverrideUser = () => new Promise((resolve) => readline.question("User already exists, would you like to override? [y/N]\n", (ans) => {
    readline.pause();
    return resolve(ans.trim().toLowerCase() === 'y');
}));
exports.promptToOverrideUser = promptToOverrideUser;
//# sourceMappingURL=addCmd.js.map