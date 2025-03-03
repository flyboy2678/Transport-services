package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Users interface {
		Create(context.Context, *User) error
		GetByID(context.Context, int64) (*User, error)
		GetByEmail(context.Context, string) (*User, error)
		UpdateByID(context.Context, *User) error
		DeleteByID(context.Context, int64) error
	}
	Trips interface {
		Create(context.Context, *Trip) error
		GetByID(context.Context, int64) (*Trip, error)
		GetByLocation(context.Context, string) ([]Trip, error)
		GetUpcoming(context.Context) ([]Trip, error)
		GetAll(context.Context) ([]Trip, error)
		UpdateByID(context.Context, *Trip) error
	}
	Bookings interface {
		Create(context.Context, *Booking) error
		GetByID(context.Context, int64) (*Booking, error)
		GetByTripID(context.Context, int64) ([]Booking, error)
		GetByUserID(context.Context, int64) ([]Booking, error)
		UpdateByID(context.Context, *Booking) error
	}
	Payments interface {
		Create(context.Context, *Payment) error
		GetByUserID(context.Context, int64) ([]Payment, error)
		UpdateByID(context.Context, *Payment) error
	}
	Subscriptions interface {
		Create(context.Context, *Subscription) error
		GetAll(context.Context) ([]Subscription, error)
		DeleteByEmail(context.Context, string) error
		DeleteByUserID(context.Context, int64) error
	}
	Invoices interface {
		Create(context.Context, *Invoice) error
		UpdateByInvoiceNumber(context.Context, *Invoice) error
		GetByInvoiceNumber(context.Context, string) (*Invoice, error)
	}
	Comments interface {
		Create(context.Context, *Comment) error
		GetByID(context.Context, int64) (*Comment, error)
		GetByTripID(context.Context, int64) ([]Comment, error)
		DeleteByID(context.Context, int64) error
		DeleteByTripID(context.Context, int64) error
	}
	Photos interface {
		Create(context.Context, *Photo) error
		GetByID(context.Context, int64) (*Photo, error)
		GetByTripID(context.Context, int64) ([]Photo, error)
		DeleteByID(context.Context, int64) error
		DeleteByTripID(context.Context, int64) error
	}
	Accomodations interface {
		Create(context.Context, *Accomodation) error
		GetByID(context.Context, int64) (*Accomodation, error)
		GetByTripID(context.Context, int64) ([]Accomodation, error)
		UpdateByID(context.Context, *Accomodation) error
	}
	AccomodationPhotos interface {
		Create(context.Context, *AccomodationPhoto) error
		GetById(context.Context, int64) (*AccomodationPhoto, error)
		GetByAccomodationId(context.Context, int64) ([]AccomodationPhoto, error)
		DeleteByID(context.Context, int64) error
		DeleteByAccomodationId(context.Context, int64) error
	}
	Activities interface {
		Create(context.Context, *Activity) error
		GetById(context.Context, int64) (*Activity, error)
		GetByTripId(context.Context, int64) ([]Activity, error)
		UpdateById(context.Context, *Activity) error
		DeleteById(context.Context, int64) error
	}
	ActivityPhotos interface {
		Create(context.Context, *ActivityPhoto) error
		GetById(context.Context, int64) (*ActivityPhoto, error)
		GetByActivityId(context.Context, int64) ([]ActivityPhoto, error)
		DeleteById(context.Context, int64) error
		DeleteByActivityId(context.Context, int64) error
	}
	//add more interface like based on the tables we are
	// on having in our database
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users:              &UserStore{db},
		Trips:              &TripStore{db},
		Bookings:           &BookingStore{db},
		Payments:           &PaymentStore{db},
		Subscriptions:      &SubscriptionStore{db},
		Invoices:           &InvoiceStore{db},
		Comments:           &CommentStore{db},
		Photos:             &PhotoStore{db},
		Accomodations:      &AccomodationStore{db},
		AccomodationPhotos: &AccomodationPhotoStore{db},
		Activities:         &ActivityStore{db},
		ActivityPhotos:     &ActivityPhotoStore{db},
	}
}
