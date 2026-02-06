let letters = ["H","i","p"]

/*
this is a horrible way to do this, but i didnt want to take the stright path :)
*/
for (let index = 0; index < 20; index++) {
    let toLog = ""
    let fun = [...letters, ...letters]
    fun.forEach((element, index) => {
        toLog += element;
        if(index == letters.length-1 || index==fun.length-1){
            toLog += element + " ";
        }
    });
    console.log(toLog + "hurra!")
}