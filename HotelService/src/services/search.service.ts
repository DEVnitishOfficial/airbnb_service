// src/services/search.service.ts

import { ElasticsearchRepository } from "../repositories/elasticsearch.repository";

export class SearchService {
  constructor(private esRepo = new ElasticsearchRepository()) {}

  async searchHotels(params : any) {
    const { q, location, minPrice, maxPrice, checkIn, checkOut, page = 0, size = 20 } = params;
    const must = [];
    const filter = [];

    if (q) {
      must.push({
        multi_match: {
          query: q,
          fields: ["name^3", "name.autocomplete^4", "address", "categories"],
          fuzziness: "AUTO"
        }
      });
    } else {
      must.push({ match_all: {} });
    }

    if (location) filter.push({ term: { location } });

    if (minPrice || maxPrice) {
      const range: any = {};
      if (minPrice) range.gte = Number(minPrice);
      if (maxPrice) range.lte = Number(maxPrice);
      filter.push({ range: { min_price: range } });
    }

    if (checkIn && checkOut) {
      filter.push({
        nested: {
          path: "rooms",
          query: {
            bool: {
              must: [
                { range: { "rooms.date_of_availability": { gte: checkIn, lte: checkOut } } },
                { term: { "rooms.booked": false } }
              ]
            }
          }
        }
      });
    }

    const body = {
      query: { bool: { must, filter } },
      sort: [{ _score: { order: "desc" } }, { min_price: { order: "asc" } }],
      from: page * size,
      size
    };

    const res = await this.esRepo.search(body);
    return res;
  }
}
