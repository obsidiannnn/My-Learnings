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

//create a simple cab booking interface using js callbacks function

function cabbooking(callback){
    console.log("Cab booking started....")
    setTimeout(()=>{
        console.log("Cab Confirmed")
        console.log("Ride started with Rahul")
        callback()
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

// send a OTP to a user, later ask user to enter OTP then verify it's valid or not


function randomotp(max,min){
    return Math.floor(Math.random()*(max+min)-min)

}
function OTP(){
    let otp=randomotp(9999,1000) 
    console.log(otp)

    console.log("OTP Sent")

    const verify = new Promise((res,rej)=>{
        setTimeout(()=>{
            let en = parseInt(prompt("Enter OTP"))
            if (en === otp){
                return res("OTP Verified")
            } else{
                return rej("Invalid OTP")
            }
        },3000)

    })
    return verify
}

OTP()
.then((res)=>{
    console.log(res)
}).catch((rej)=>{
    console.log(rej)
})