import express, { Application } from "express";
import postRouter from "./modules/post/post.router.js";

const app: Application = express();

app.use(express.json());

app.use("/posts", postRouter);

app.get("/", (req, res) => {
  res.send("Hello from express!");
});

export default app;
