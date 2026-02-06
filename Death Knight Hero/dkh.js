const readline = require('readline');

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

let lines = [];
let numLines = NaN;
rl.on('line', (line) => {
    if (isNaN(numLines)){
        numLines = line
    } else {
        lines.push(line)
    }
    if(lines.length == numLines){
        console.log(lines.filter(b => !b.includes("CD")).length)
        rl.close()
    }
});