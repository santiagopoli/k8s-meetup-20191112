const express = require('express');

const app = express()

process.on('SIGTERM', () => process.exit(1));
process.on('SIGINT', () => process.exit());

const findUser = (id) => ({
    id,
    name: `User ${id}`,
    age: 28
})

app.get("/users/:id", (req, res) => {
    res.json(findUser(req.params.id))
})

console.log("Starting users service")
app.listen(8080, () => {
    console.log("Server listening on :8080")
})