// src/lib/elasticsearch.client.ts
import { Client } from "@elastic/elasticsearch";
import esConfig from "../config/es.config";

const auth = esConfig.username && esConfig.password ? { username: esConfig.username, password: esConfig.password } : undefined;

export const esClient = new Client({
  node: esConfig.node,
   auth,
  tls: {
    rejectUnauthorized: false, // allow self-signed cert
  },
  maxRetries: 3,
  requestTimeout: 30000
});
