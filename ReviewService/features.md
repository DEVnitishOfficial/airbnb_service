
* crud part is very easy you can go through the code you will get that easily, here i will write those concept in which you can get confused or not able to get in one go.

# How async rating calculation work ?

**OVERVIEW**

*  *In the reviewService we have added a cron job which in procudtion run hourly and for testing purpose it's running in every 30sec, it fetch latest review created in last one hour or 30sec from review Service and previous hotelRating and count from hotel and finally after calculating the rating and ratingCount back we update the hotelService*

### Step By step working

* cronJob starts from app > application.go > cronjob.StartCron(svc, mode)

1.  StartCron method takes two parameter 1. svc 2. mode

* svc : it has the access of db connection, HotelClient instance, and ReviewRepositoryAggRating instance so calling the svc we can do any operation in order to complete the cron job.

* mode : it represents the "test" and "production" in which mode currently my server is running.

2. Inside the StartCron function we have initilized the cron job with UTC timezone and define a context of one minute and passed it to the ProcessPendingRatings(ctx) method i.e whatever goroutines is running to execute this function must have to complete within one minute otherwise when context is cancel then the running goroutine will be killed automatically.


3. In ProcessPendingRatings function we try to get application level or advisory lock so that we can ensure on a particular time there must be only running one cronJob in order to avoid any confilct.

4. After completing the cronjob we have release the lock in a very clean way by using the context.Background, this context ensure us doesn't matter how the ProcessPendingRatings function either successful or by any failure eventually it will be relaese the lock it's similir to the finally keyword in js(try, catch, finally).

5. we call FetchUnappliedAggregates methods with two parameter 1. ctx 2. cutoof

* inside this we have executed a sql query which calculte the rating count and sum of rating like if two user give a rating of 4 and 3 to the same hotel then it's rating count will be 2 and sum of rating will be 7 and return back in result array.

* we make a loop on the result(aggs) created transaction and passed the ctx.

* make an api call to hotelService(using GetHotelRating method) by providing the hotelId it returns us the rating and the ratingCount.

* Now we have oldAvg, oldCount, and newCount using this info we can calculate the avg rating using the below formula.----->>> newAvg = (oldAvg * oldCnt + a.Sum) / newCount

* newAvg = (oldAvg * oldCnt + currentIncomingRatingSumFromReview(like if two review with 3 and 4 sum will be 7)) / totalCountIncludingOldandComingNewFromREview(like old is 100 coming two new review then it will be 102)

* After calculating the newAvg we make an api call to the hotelService with new avg rating and the new total rating count and update the hotelService.

* Finally we call the MarkReviewsAsSynced function which update the revewService that the new rating is synced with overAll hotel rating.


# How you extracted userInfo with reviewInfo because both are in seperate db and in sperate microservices and how you communicate internally and authenticated the user in second internal api call and how api gateWay help here ?


### Features : 

* When we find hotelbyId then we only got the below info
 ```js
 {
            "Id": 1,
            "UserId": 23,
            "BookingId": 236,
            "HotelId": 20,
            "Comment": "I like the cleanness of the hotel and fooding part",
            "Rating": 4,
            "CreatedAt": "2025-10-07 15:50:46",
            "UpdatedAt": "2025-10-07 16:59:31",
            "DeletedAt": null,
            "IsSynced": true    
}
 ```

* In the above info you can see clearly we are getting the comment, rating, userid, bookingId, but there is no userName, email or their profile pic info, in actual review you will see the name of the reviewer, profile pic, rating and comment.

* So here my goal is to provide all the required information particular in this findByHotelId so that any client may request to this API endpoint and get all the required info and show them on the UI.

*  **overview** : we will hit the reviewService route getHotelById where i will get all the review of a particular hotel, from this response we have to accumulate all the userid and make an api call to the authService and get all userInfo and combine the result of both(review and auth service) api calls and show them to the client.

* **Start** : ReviewService---> router---> review_router.go---> r.Get("/reviews/hotel", rr.reviewController.GetReviewsByHotelId)

* controllers > review.go here we have extracted the hotelId(from req params) and authHeader from the request header(contain the token because the first req is gone through the apiGateway and inside the jwtAuth after verification of token again i set those token in the request header) and call service layer.

* The coming hotelId is string converted into the int and call the repo layer will hotelId and repo return us the array of review of that particular hotelId.

* iterate on each result and extract the userId and store in an array.

* In REviewService we have exposed an api named FetchUsersByIds(userIds, authHeader) with two paramete 1. array of userIds using which i will fetch the userInfo from AuthService and 2. authHeader(contain token, required because we are fetching the userInfo)

* FetchUsersByIds methods gives us the all the userInfo and the result came back into the service layer.

* In service layer i have combined both api data(reviewService(rating, comment, hotelId) and authService(name, email, userId)) and send back to the client.

* In this way i have extracted the complete review information which is ready to use on client side, for implementation detail you have to see the code.

