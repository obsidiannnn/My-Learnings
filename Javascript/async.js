// async code is when all task are given and we don't know which task will complete first.

// they don't block the execution of other tasks. The code which will take time to execute will be 
// executed in the background and once it is completed the result will be returned before other tasks which are wating to be executed.
 
// if code contains these it is a async code:
//setTimeout
//setInterval
//promises
//fetch
//axios
//async/await
//XMLHttpRequest

// Sometimes our code depends on other severs. In that we don't know what time each task will take to complete.
// so we can't write sync code for that. To counter it we have async code.

// setTimeout(funtion(){}, time in ms) is an example of async code.)
// this fnc inside setTimeout is called as callback fnc.

console.log("hey")
console.log("hey2")
setTimeout(function(){
    console.log("hey3")
},0)
console.log("heyr4")

// promises are code that will produce a value in future.
// it can be in one of the three states:
//1. pending - initial state, neither fullfilled noe rejected.
//2. fullfiled - meaning that the operation completed successfully.
//3. rejected - meaning that the operation failed.

// to create a promise we use the promise constructor which takes a fucntion as an arugument.
// this function takes two parameters resolve and reject.



let aa = new Promise(function(resolve,reject){
    if(true){
        return resolve

    } else {
        return reject
    }
})
aa 
.then(function(){
    console.log("Work Done")
})
.catch(function(){
    console.log("work not done")
})


