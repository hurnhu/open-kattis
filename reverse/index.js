const readline = require('readline');

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

let lines = [];
rl.on('line', (line) => {
    lines.push(line)
    if(lines.length - 1 == lines[0]){
        lines.slice(1).reverse().forEach(element => {
            console.log(element)
        });
        rl.close()
    }
});