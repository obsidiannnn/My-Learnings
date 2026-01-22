// Ques - create a random number from 1-9 if it is below 5 resolve else reject.

var ans = new Promise((res,rej)=>{
    let num = Math.floor(Math.random()*10)

    if(num<=5){
        return res(num)
    } else {
        return rej(num)
    }
})
ans
.then((num)=>{
    console.log(`${num},is less than or equal to 5`)
})
.catch((num)=>{
    console.log(`${num}, is greater than 5`)
})

// create a simple cab booking interface using js callbacks function

function cabbooking(callback){
    console.log("Cab booking started....")
    setTimeout(()=>{
        console.log("Cab Confirmed")
        console.log("Ride started with Rahul")
    },3000)

}
cabbooking(()=>{
    setTimeout(()=>{
        console.log("Ride Completed")
        setTimeout(()=>{
            console.log("Payment Successfull")
        },2000)
    },3000)
})