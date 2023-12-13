const readline = require('readline');

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

let lines = [];
rl.on('line', (line) => {
    lines.push(line)
    if(lines.length === 2){
        console.log(lines.reduce((accu, item) =>{
            return parseInt(accu) + parseInt(item)
        }, 0))
        rl.close()
    }
});