const express = require("express");
const axios = require('axios');
const path = require('path');
const AWSXRay = require('aws-xray-sdk');
AWSXRay.captureHTTPsGlobal(require('http'));
AWSXRay.captureHTTPsGlobal(require('https'));

const http = require('http');
const https = require('https');
const app = express();

app.use(AWSXRay.express.openSegment('frontend'));

app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'pug');

app.get("/", async (req, res) => {
    AWSXRay.capturePromise();
    let giphy = [];
    try {
        const instance = axios.create({
            httpAgent: new http.Agent(),
            httpsAgent: new https.Agent(),
        });
        const query = await instance.get(`http://api.${process.env.COPILOT_SERVICE_DISCOVERY_ENDPOINT}/giphy`);
        // const query = await instance.get(`http://127.0.0.1/giphy`);
        giphy = query.data
    } catch (error) {
        console.error("Error retrieving giphy images")
    }

    res.render("index", { giphy: giphy });
})

app.get('/health', (req, res) => {
    res.send('OK')
})

app.use(AWSXRay.express.closeSegment());

app.listen(3000, () => {
    console.log("server started on 3000")
})
