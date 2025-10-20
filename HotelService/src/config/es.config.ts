const ES_NODE = process.env.ES_NODE || "https://localhost:9200";
const ES_USERNAME = process.env.ES_USERNAME || "";
const ES_PASSWORD = process.env.ES_PASSWORD || "";

export default {
  node: ES_NODE,
  username: ES_USERNAME,
  password: ES_PASSWORD,
  index: process.env.ES_HOTEL_INDEX || "hotels",
  bulkChunkSize: Number(process.env.ES_BULK_CHUNK_SIZE || 500)
};
