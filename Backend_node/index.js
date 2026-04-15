// require('dotenv').config()
//const express = require('express') // used to use express from dependencies

import dotenv from 'dotenv'
import express from 'express' // using express from dependencies with ES6 syntax
import {jokes} from './data.js'
dotenv.config()

const app = express() // using express inside app


const port = process.env.PORT || 5000


app.get('/',(req,res)=>{
    res.send("Server is ready")
})
// '/' means it is listening on home route

app.get('/insta', (req,res)=>{
    res.send("insta@name")
    // now server is listening on insta route
})

app.get('/jokes',(req,res)=>{
    res.json(jokes)
})


app.listen(port, ()=>{ // using port via env file
    console.log(`Server listening on port ${port}`);
})
// server will listen on the route we hit