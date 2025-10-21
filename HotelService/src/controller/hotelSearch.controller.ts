// controller/hotelSearch.controller.ts
import { Request, Response, NextFunction } from "express";
import { StatusCodes } from "http-status-codes";
import { SearchService } from "../services/search.service";

export async function searchHotelHandler(req: Request, res: Response, next: NextFunction) {

  const hotelSearchService = new SearchService();
  const searchPayload = {
    name: req.query.name as string,
    address: req.query.address as string,
    location: req.query.location as string,
  };

  const hotels = await hotelSearchService.searchHotels(searchPayload);

  

  res.status(StatusCodes.OK).json({
    success: true,
    message: "Hotels fetched successfully",
    data: hotels,
  });
}
