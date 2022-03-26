import express, { Application, Request, Response } from "express";

const app: Application = express();
const port = 3000;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get("/", async (req: Request, res: Response): Promise<Response> => {
  return res.status(200).json({
    message: "initial body",
  });
});

app
  .listen(port, (): void => {
    console.log(`Connected succesfully on port ${port}`);
  })
  .on("error", (error) => {
    console.error(`Error occured: ${error.message}`);
  });
