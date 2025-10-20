// scripts/reindex-hotels.ts
// import { Sequelize } from "sequelize";
// import dbConfig from "../src/config/sequelize.config"; // adjust path
// import { SequelizeHotelRepository } from "../src/repositories/hotel.repository";
// import { ElasticsearchRepository } from "../src/repositories/elasticsearch.repository";
// import { transformHotelToESDoc } from "../src/utils/transformers/hotel.transformer";
// import esConfig from "../src/config/es.config";

import esConfig from "../config/es.config";
import { ElasticsearchRepository } from "../repositories/elasticsearch.repository";
import { transformHotelToESDoc } from "../utils/transformers/hotel.transformer";

async function run() {
  // initialize sequelize as your app does. Skip if apps bootstrap handles it.
  const hotelRepo = new SequelizeHotelRepository();
  const esRepo = new ElasticsearchRepository();

  const pageSize = esConfig.bulkChunkSize || 500;
  let page = 0;
  while (true) {
    const hotels = await hotelRepo.getHotelsWithRoomsPaginated(page, pageSize);
    if (!hotels || hotels.length === 0) break;

    const docs = hotels.map(h => transformHotelToESDoc(h));
    console.log(`Indexing chunk ${page} size=${docs.length}`);
    await esRepo.bulkIndex(docs);
    page += 1;
  }
  console.log("Reindexing completed.");
  process.exit(0);
}

run().catch(err => {
  console.error("Reindex failed", err);
  process.exit(1);
});
