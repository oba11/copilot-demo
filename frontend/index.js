const express = require("express");
const axios = require('axios');
const app = express();
const path = require('path');

app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'pug');

app.get("/", async (req, res) => {
    let giphy = [];
    try {
        const query = await axios.get(`http://api.${process.env.COPILOT_SERVICE_DISCOVERY_ENDPOINT}/giphy`);
        // const query = await axios.get(`http://127.0.0.1/giphy`);
        giphy = query.data
    } catch (error) {
        console.error("Error retrieving giphy images")
    }

    res.render("index", { giphy: giphy });
})

app.get('/health', (req, res) => {
    res.send('OK')
})

app.listen(3000, () => {
    console.log("server started on 3000")
})
