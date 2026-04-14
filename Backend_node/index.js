const express = require('express') // used to use express from dependencies


const app = express() // using express inside app

const port = 3000

app.get('/',(req,res)=>{
    res.send("Hello")
})
// '/' means it is listening on home route

app.get('/insta', (req,res)=>{
    res.send("insta@name")
    // now server is listening on insta route
})


app.listen(port, ()=>{
    console.log(`App listening on port ${port}`);
})
// server will listen on the route we hit