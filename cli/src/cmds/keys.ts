import {Arguments, Argv} from "yargs";
import {join} from "path";


export const command = 'keys <method>'
export const desc = 'generate keys'

export const builder = (yargs: Argv) => {
    return yargs
        .commandDir(join(__dirname,`keys`))
        .option('recover', {
            describe: 'recover account by providing mnemonic',
            default: false,
            demandOption: true
        })
        .help()
        .demandCommand()
}
export const handler = (argv: Arguments) => {

}
