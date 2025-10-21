// src/services/search.service.ts

import { ElasticsearchRepository } from "../repositories/elasticsearch.repository";

export class SearchService {
  constructor(private esRepo = new ElasticsearchRepository()) { }

  async searchHotels(params: { 
    id? : string | number;
    name?: string; 
    address?: string; 
    location?: string;
    page?: number;
    size?: number;
  }) {
    const { id, name, address, location } = params;
    const must = [];
    const filter = [];

    if(id){
      must.push({
        match: {
          id: id
        }
      });
    }

    if (name || address) {
      must.push({
        multi_match: {
          query: name || address,
          fields: ["name^3", "address^2"],
          fuzziness: "AUTO",
        },
      });
    } else {
      must.push({ match_all: {} }); // If no name or address, match all documents
    }

    // Filter by location if provided
    if (location) {
      filter.push({
        match: {
          location: {
            query: location,
            fuzziness: "AUTO",
          },
        },
      });
    }

  // Pagination control
  const page = Number(params.page) || 0;
  const size = Number(params.size) || 20;

    //  Construct ES query body
    const body = {
      query: { bool: { must, filter } },
      sort: [{ _score: { order: "desc" } }],
      from: page * size,
      size,
    };

    // Call repository search method
      const res = await this.esRepo.search(body);

      const hotels = res.hits.hits.map((hit: any) => hit._source);

      const total = typeof res.hits.total === "number" ? res.hits.total : res.hits.total?.value ?? 0;

      return {
        total,
        took: res?.took, // time taken by ES to execute the query
        page,
        size,
        hotels,
      };
  }
}
