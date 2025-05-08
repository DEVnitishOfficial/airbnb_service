import express from 'express'
import { serverConfig } from './config';
import v1Router from './routers/v1/index.router';
import v2Router from './routers/v2/index.router';
import { genericErrorHandler } from './middlewares/error.middleware';
import logger from './config/logger.config';
import { attachCorrelationIdMiddleware } from './middlewares/correlation.middleware';
import sequelize from './db/models/sequelize';
import Hotel from './db/models/hotel';



const app = express();
app.use(express.json());


app.use(attachCorrelationIdMiddleware)

app.use('/api/v1', v1Router);
app.use('/api/v2', v2Router);

app.use(genericErrorHandler);

app.listen(serverConfig.PORT, async () => {
    console.log(`server is listening at  http://localhost:${serverConfig.PORT}`)
    logger.info("press Ctrl + C to stop the server", { "name": "dev-server" })

    try {
        await sequelize.authenticate(); // this will check the connection to the database and if the connection is successful then it will return a promise otherwise it will throw an error.
        logger.info("Database connected successfully")

        // const hotels = await Hotel.create({ // this will create a new hotel in the database and return the hotel object.
        //     name: "citadel hotel",
        //     address: "231 main street",
        //     location: "wisconsin",
        //     rating: 3.1,
        //     ratingCount: 100,
        // })
        // logger.info("Hotel created successfully", hotels.toJSON());
        const hotels = await Hotel.findAll(); // this will find all the hotels in the database and return the hotel object.
        logger.info("Hotels fetched successfully", hotels);

    } catch (err) {
        logger.error("something went wrong while connecting to the database", { "name": "dev-server" }, err);
    }
});