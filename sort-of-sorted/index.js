const readline = require('readline');

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

let lines = [];
let tempArr = [];
let innerArrLen = 0
rl.on('line', (line) => {
    let int = parseInt(line)
    if(line == 0){
        //sort
        if(tempArr.length > 0){
            lines.push(tempArr)
            tempArr = []
        }
        lines.forEach((line) => {
            line.sort((a,b) => {
                let aFirstTwo = a.slice(0,2)
                let bFirstTwo = b.slice(0,2)
                if (aFirstTwo < bFirstTwo) {
                    return -1;
                  }
                  if (aFirstTwo > bFirstTwo) {
                    return 1;
                  }
                
                  // names must be equal
                  return 0;
            }).forEach((item) => {
                console.log(item);
            })
            console.log();
        })
        rl.close()
    } else if(!isNaN(int)){
        innerArrLen = int;
        if(tempArr.length > 0){
            lines.push(tempArr)
            tempArr = []
        }
    } else {
        tempArr.push(line)
    }
});