
# Core features of HotelService

1. Hotel creation
2. Operation with existing rooms
3. Release occupied unConfirm booking/ghost bookings

4. Room generation
    * Room generation is divided into two parts
    (i) Bulk room creation
    (ii) Extend room availability to maintain a fixed booking window(like 90 days or 3 months) using the node-cron job
    


### 1. Hotel Creation(CRUD)

* Hotel creation is very simple, we follow router-->controller--->service--->repository stracture.

* take input according to dto and created hotel.
<pre>
    name: string;
    address: string;
    location: string;
    rating?: number;
    ratingCount?: number;
</pre>

### 2. Operation with existing rooms

**Here we have different routes and Functionality**

1. *checRoomAvailibility*

    <pre>
    (i). if you remember in the BookingService while we are creating a new booking then before creating a new booking first we check the room availability, by passing roomCategoryId and a specific date range. 
                if rooms are available as per given roomCategoryId and date range where bookingId is null in the room table then we return those date range rooms from the HotelService and then booking creation completed.
    </pre>

2. *updateBookingId in Rooms table*
    <pre>
            In the booking creation when booking creation becomes successful after that to book that room so that no other user can book it we update the bookingId inside the room table by making an api calls from bookingService to HotelService.
                            So here updateBookingId is nothing but that's exposed api who receive the request and update bookingId in the rooms table.
    </pre>

### 3. Release occupied unConfirm booking/ghost bookings

3. *release* occupied unConfirmed bookings/ghost bookings
    <pre>
            When we create new bookings then we are not confirming the booking in the same transaction, booking and confirmation happens in two transaction that why here we got this type of problem.

                                    Problem : The problem is when user creates the bookings
        but not confirmed that booking, since while create booking we update the bookingId inside the rooms table so that rooms is now not available for others because when we search rooms availability in the query we find where bookingId is null but here there is booking id so it will not come in the available rooms category.


        Solution : 
                To solve this problem we inserted two columns in our booking db 1.expiredAt, 2. releaseAt and status Enum "EXPIRED".

                when we create the booking and before assigning the bookingId to the room table we implemented updateExpiryTimeOfCreatedBooking in which we update expiredAt column of booking table from currentTime + 10 minutes;

                Cron implementation : 
                BookingService>src>cron>bookingCleanup.cron.ts

                Here we have implemented a cron job which runs every minute and from the booking tables we extract all those bookings whose status is pending and expiredAt is less then current time.
                                        Then we iterate on each booking and update the status from pending to expired and releaseAt time is current time. 


                After updating the booking table we make an api call to hotelService(already exposed an api) and remove the bookingId from the room table. In this way i solved the problem of ghost booking.

                so here release route is responsible for updating the bookingId inside the room table.

    </pre>


## 4. Room generation

### (i) Bulk room creation 

* Here room generation is happening using redis, there are 3 main components involve in this complete room generation process.
<pre>
(i) Producer --->> take payload from the postman(roomCategoryId, startDate, endDate)
(ii) Redis ---->> store producer payload
(iii) Processors(worker) ---->> took that payload from redis and start roomGeneration
</pre>
* starting of room generation from below : 
processors>roomGeneration.processors.ts ----->>> generateRoomsService(paylod)


***Core logic of room generation***

(i) Our function generateRoomsService start which have payload i.e roomCategoryId, startDate, endDate, priceOverride(optional), batchSize(default 100).

**roomCategoryId** : it signify the category of room like(SINGLE, DOUBLE, FAMILY, DELUX) etc.

**startDate** : the starting date from where you want to create room like(14-10-2025).

**endDate** : the last date till where you want to create room like(14-12-2025).
            here for two months i.e for 60 days we want to create room

**priceOverride** : this is optional because we have already decided price when we define the roomCategory like delux already price have 4500, if you want to override this you can pass priceOverride value here.

**batchSize** : the default batchSize is 100 i.e if we have to create 500 rooms then we will create in 5 batch not in one go, to avoid roomCreation failure or avoid unexpected issue, if you want to override the batch size you can.

```js
// Actual payload that we are sending for room generation : 
{
    "roomCategoryId": 4,
    "startDate": "2025-10-12T00:00:00.000Z",
    "endDate": "2025-11-15T00:00:00.000Z"
}

```

* In order to generate the room in batchSize first we check from the given roomCategoryId, is particular roomCategory exist or not.

* we wrap the incoming startDate and endDate with ``` new Date() ``` to make it object because the coming payload is in string on which we can't call the the method like **getTime(), setTime()**.

* calling the getTime() method on the date object we calculate the total days for which we want to generate the rooms.

* if batchSize is coming from the payload we extract that otherwise we keep defalt of 100 batchsize.

* In while loops it's checking that the clientSent startDate must be lte(less then equal) to endDate because the minimum batch size can be one day less then this not possible if in some case if start date and end date you send same then it will go in the infinite loop, to prevent this in the starting we have already put check.

* ``` batchEndDate.setDate(batchEndDate.getDate() + batchSize); ``` this part of the code set the batchEnd date, to understand this just log the vlaue you will understand this part.

* Then in next line we check if the batchEndDate that we calcute by adding the batchSize if that exceede the endDate it means the batchSize is less then 100(or whatever batchSize is!) so here we set the smaller batch and update the batchendDate.

* we call ```processDateBatch(roomCategory, startDate, batchEndDate)``` whatEver batchResult comes we update the totalRoomsCreated, totalDatesProcessed and increment the start date by one so that it create a fresh new batch and generate rooms for that batch as well.


***Understanding processDateBatch(roomCategory, startDate, batchEndDate)***

* since we have already roomCategoryId, startDate, endDate so instead of checking each day rooms availability, from the repository layer we make a db calls which returns us all the rooms in the particular given date range in array(it avoids the N+1 query problem).

* Now we iterate on that array and make a set from that dates of availibility which only keep the unique dates.

* In the while loop we are checking the currentDate must be less than the endDate and from the currentDate we extract the first date part and since we have the set so on that set we lookup with the currentDate if that set have no dates then we push that particular dateOfAvailability data to an array(roomsToCreate) and increment the currentDate, dateProcessed, and roomsCreated.

* when the while loops ends then we extract all data from that **roomsToCreate** array and create room in bulk in one go and return the roomsCreated and dateProcessed.


 

### (ii) Extend room availability to maintain a fixed booking window(like 90days or 3 months) using the node-cron job

* Hit startRoomSchedulerCronJob()

    (i) By hitting restAPI  ----> http://localhost:3001/api/v1/room-scheduler/start

    (ii) put your startRoomSchedulerCronJob() method in server.ts file


* we can hit  startRoomSchedulerCronJob() by two way either by the rest api calls or by simply put this method on server.ts file so that whenever our server starts our roomScheduler cron job also start running.

* First we check is the cronJob already running or not if it's already running then we simply return.

* Schedule job to run every minute in testing, every day at midnight in production decided from .env.

* we call extendRoomAvailability() function  inside this we call repository layer findLatestDatesForAllCategories which returns us the roomCategoryId and latestDate of the available rooms.

* Then we call extendCategoryAvailability with categoryData in which we have categoryId and latestDate. Here we calculated nextDate from the latestDate and by making a db query we find the roomCategory by the roomCategoryId.

* Again we make a db query to ```findRoomCategoryByIdAndDate(roomCategoryId, nextDate)``` here nextDate we have calculated from the latestDate. if rooms already available for that date then we skip room creation particular to that date.

* if rooms are not already available then we calculate the endDate form the nextDate(latestDate--->nextDate---->endDate) and from the available info like the roomCategoryId, startDate, endDate, priceOverrride, and batch we create a jobData.

* From the created jobData we schedule a cronJob in redis and whenever our worker is free then they took that job and create rooms like the bulk room creation but here we are creating only one rooms for each category with latest Date available.

* In this way our cronJob working fine and creating rooms depend on the cronJob setting.


## 5. Elastic search
- here we have main two features : 
i. index the hotel when new hotel is created, in the elastic search.
ii. hotel searching

### i. Indexing newly created hotel in the ES

* START --> routers--> v1---> hotel.router.ts--->createHotelHandler--->addHotelIndexInESJobToQueue

* addHotelIndexInESJobToQueue, it's a producer who takes the created hotelId and store it inside the redis, in a queue named "HOTEL_INDEXING_QUEUE" with a payload named "HOTEL_INDEXING_PAYLOAD", here producer works ended.

* Since inside the redis we have the hotelId, so now processor/consumer/worker takes that payload by matching the same queue name and the payload name.

* Now we find the hotel by that extracted id from the redis, once we get the hotel detail like their name, address and location then i call method transformHotelToESDoc(hotel).

* I keep transformHotelToESDoc(hotel) method because if in future we have to add more detail then we can transform the hotelData here like the rooms detail, geo location, price etc.

* we call indexHotel(doc) method from ElasticsearchRepository and here we have esClient which is coming from lib folder where i have defined my esClient liberary by importing {Client} from "@elastic/elasticsearch" npm package.

* In this way i have indexed all the hotel that are created and saved in the mysql db.

* In the similir fashion update and delete hotel is working in elastic search

### ii. Searching existing hotel from Elastic Search

* START : routers--> v1---> hotel.router.ts--->/search--->searchHotelHandler

* In searchHotelHandler whatever request is coming like id, name address, location accumulate all of this in a variable and pass into the searchHotels() methods.

* In searchHotels() methods we extract all those parameter and create two array [must] and [filter] in which we fill the coming parameter and create elastic search query and store this object in a variable name "body" and passs it to the esRepo.search(body) method.

* From search method we call our elasticSearch client esClient.search({ index: INDEX, body }) method in which we pass the index name and the search query stored in body, and it returns the matching hotels.

* In the returned hotels there are a lot of meta data so we filter it and only send the relevent data to the client.

