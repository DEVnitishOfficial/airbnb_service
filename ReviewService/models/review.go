package models

type Review struct {
	Id        int64
	UserId    int64
	BookingId int64
	HotelId   int64
	Comment   string
	Rating    int
	CreatedAt string
	UpdatedAt string
	DeletedAt *string
	IsSynced  bool
}
type ReviewWithUser struct {
	UserId   int64
	Username string
	Email    string
	Comment  string
	Rating   int
	HotelId  int64
}
