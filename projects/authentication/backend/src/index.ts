import express from "express";
import cors from "cors";
import { createConnection } from "typeorm";

const app = express();
app.use(cors());

app.post("/login", (req, res) => {
});

app.listen(3000, () => {
    console.log("Server is running on port 3000");
});