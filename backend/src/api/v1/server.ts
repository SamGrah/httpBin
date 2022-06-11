import app from './app';
const port = process.env.PORT;

app
  .listen(port, (): void => {
    console.log(`Connected succesfully on port ${port}`);
  })
  .on("error", (error) => {
    console.error(`Error occured: ${error.message}`);
  });