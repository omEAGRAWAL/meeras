package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// User model
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Mobile   string             `bson:"mobile_number"`
	Password string             `bson:"password"`
	Role     string             `bson:"role"`
}

// Venue model
type Venue struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Location    string             `bson:"location"`
	Rating      string             `bson:"rating"`
	Description string             `bson:"description"`
	MapURL      string             `bson:"map_url"`
	ManagerID   primitive.ObjectID `bson:"manager_id"`
	Packages    []Package          `bson:"packages"`
}

// Package model
type Package struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	Price      float64            `bson:"price"`
	Decoration string             `bson:"decoration"`
	MaxPeople  int                `bson:"max_people"`
}

// Booking model
type Booking struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id"`
	VenueID   primitive.ObjectID `bson:"venue_id"`
	PackageID primitive.ObjectID `bson:"package_id"`
	Date      time.Time          `bson:"date"`
	TimeSlot  string             `bson:"time_slot"`
	Status    string             `bson:"status"`
}

// Role model
type Role struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	RoleName    string             `bson:"role_name"`
	Permissions []string           `bson:"permissions"`
}
