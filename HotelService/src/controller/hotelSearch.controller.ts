// controller/hotelSearch.controller.ts
import { Request, Response, NextFunction } from "express";
import { StatusCodes } from "http-status-codes";
import { SearchService } from "../services/search.service";

export async function searchHotelHandler(req: Request, res: Response, next: NextFunction) {

    const hotelSearchService = new SearchService();
  const searchPayload = {
    name: req.query.name as string,
    city: req.query.city as string,
    minPrice: req.query.minPrice ? Number(req.query.minPrice) : undefined,
    maxPrice: req.query.maxPrice ? Number(req.query.maxPrice) : undefined,
    category: req.query.category as string,
    lat: req.query.lat ? Number(req.query.lat) : undefined,
    lon: req.query.lon ? Number(req.query.lon) : undefined,
    distance: req.query.distance as string, // e.g., "10km"
  };

  const hotels = await hotelSearchService.searchHotels(searchPayload);

  

  res.status(StatusCodes.OK).json({
    success: true,
    message: "Hotels fetched successfully",
    data: hotels,
  });
}
