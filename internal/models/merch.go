package models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type MerchCategory struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MerchProduct struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	CategoryID     int       `json:"category_id"`
	PrimaryImageID int       `json:"primary_image_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type MerchSize struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ProductID int    `json:"product_id"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

type MerchColor struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ProductID int    `json:"product_id"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

type MerchPrice struct {
	ID        int     `json:"id"`
	Price     float64 `json:"price"`
	ProductID int     `json:"product_id"`
	SizeID    int     `json:"size_id"`
	ColorID   int     `json:"color_id"`
	CreatedAt int     `json:"created_at"`
	UpdatedAt int     `json:"updated_at"`
}

type MerchTransaction struct {
	ID             int       `json:"id"`
	UserID         uuid.UUID `json:"user_id"`
	ProductID      int       `json:"product_id"`
	CouponID       int       `json:"coupon_id"`
	ProofOfPayment string    `json:"proof_of_payment"`
	Status         string    `json:"status"`
	PaidAt         time.Time `json:"paid_at"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type MerchImage struct {
	ID        int    `json:"id"`
	Path      string `json:"path"`
	ProductID int    `json:"product_id"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

type MerchCoupon struct {
	ID            int            `json:"id"`
	Code          string         `json:"code"`
	Description   string         `json:"description"`
	DiscountType  string         `json:"discount_type"`
	DiscountValue pgtype.Numeric `json:"discount_value"`
	MaxUseCount   int            `json:"max_use_count"`
	UsedCount     int            `json:"used_count"`
	ValidFrom     time.Time      `json:"valid_from"`
	ValidUntil    time.Time      `json:"valid_until"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}
