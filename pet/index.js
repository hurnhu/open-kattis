const readline = require('readline');

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

let lines = [];
let currMax = {
    index: 0,
    max: 0
};
rl.on('line', (line) => {
    lines.push(line)
    if(lines.length === 5){
        lines.forEach((item, index) => {
            let reduced = item.split(' ').reduce((accu, curr) => {
                return parseInt(accu) + parseInt(curr)
            }, 0)
            if(currMax.max < reduced){
                currMax.max = reduced
                currMax.index = index
            }
            return reduced
        })
        console.log(currMax.index + 1, currMax.max)
        rl.close()
    }
});