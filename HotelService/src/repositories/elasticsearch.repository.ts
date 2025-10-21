import esConfig from "../config/es.config";
import { esClient } from "../lib/elasticsearch.client";

const INDEX = esConfig.index;

  console.log('request came for indexing');

export class ElasticsearchRepository {
  async indexHotel(doc: Record<string, any>) {
    console.log('Indexing hotel document>>>:', doc);
    return esClient.index({
      index: INDEX,
      id: String(doc.id),
      body: doc
    });
  }

  async updateHotel(id: string | number, partial: Record<string, any>) {
    return esClient.update({
      index: INDEX,
      id: String(id),
      body: { doc: partial, doc_as_upsert: true }

    } as any);
  }

  async deleteHotel(id: string | number) {
    return esClient.delete({
      index: INDEX,
      id: String(id)
    }).catch(err => {
      // ignore not_found
      if (err?.meta?.statusCode !== 404) throw err;
    });
  }

  async bulkIndex(docs: Record<string, any>[]) {
    if (!docs.length) return null;
    const body = docs.flatMap(d => [{ index: { _index: INDEX, _id: String(d.id) } }, d]);
    return esClient.bulk({ refresh: false, body });
  }

  async search(body: any) {
    console.log('Searching hotels with body:', body);
    return esClient.search({ index: INDEX, body });
  }
}
