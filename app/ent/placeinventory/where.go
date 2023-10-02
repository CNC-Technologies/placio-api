// Code generated by ent, DO NOT EDIT.

package placeinventory

import (
	"placio-app/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldContainsFold(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldName, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldPrice, v))
}

// StockQuantity applies equality check predicate on the "stock_quantity" field. It's identical to StockQuantityEQ.
func StockQuantity(v int) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldStockQuantity, v))
}

// MinStockThreshold applies equality check predicate on the "min_stock_threshold" field. It's identical to MinStockThresholdEQ.
func MinStockThreshold(v int) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldMinStockThreshold, v))
}

// Sku applies equality check predicate on the "sku" field. It's identical to SkuEQ.
func Sku(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldSku, v))
}

// ExpiryDate applies equality check predicate on the "expiry_date" field. It's identical to ExpiryDateEQ.
func ExpiryDate(v time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldExpiryDate, v))
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldSize, v))
}

// Color applies equality check predicate on the "color" field. It's identical to ColorEQ.
func Color(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldColor, v))
}

// Brand applies equality check predicate on the "brand" field. It's identical to BrandEQ.
func Brand(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldBrand, v))
}

// PurchaseDate applies equality check predicate on the "purchase_date" field. It's identical to PurchaseDateEQ.
func PurchaseDate(v time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldPurchaseDate, v))
}

// LastUpdated applies equality check predicate on the "last_updated" field. It's identical to LastUpdatedEQ.
func LastUpdated(v time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldLastUpdated, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldContainsFold(FieldName, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLTE(FieldPrice, v))
}

// StockQuantityEQ applies the EQ predicate on the "stock_quantity" field.
func StockQuantityEQ(v int) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldStockQuantity, v))
}

// StockQuantityNEQ applies the NEQ predicate on the "stock_quantity" field.
func StockQuantityNEQ(v int) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNEQ(FieldStockQuantity, v))
}

// StockQuantityIn applies the In predicate on the "stock_quantity" field.
func StockQuantityIn(vs ...int) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldIn(FieldStockQuantity, vs...))
}

// StockQuantityNotIn applies the NotIn predicate on the "stock_quantity" field.
func StockQuantityNotIn(vs ...int) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNotIn(FieldStockQuantity, vs...))
}

// StockQuantityGT applies the GT predicate on the "stock_quantity" field.
func StockQuantityGT(v int) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGT(FieldStockQuantity, v))
}

// StockQuantityGTE applies the GTE predicate on the "stock_quantity" field.
func StockQuantityGTE(v int) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGTE(FieldStockQuantity, v))
}

// StockQuantityLT applies the LT predicate on the "stock_quantity" field.
func StockQuantityLT(v int) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLT(FieldStockQuantity, v))
}

// StockQuantityLTE applies the LTE predicate on the "stock_quantity" field.
func StockQuantityLTE(v int) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLTE(FieldStockQuantity, v))
}

// MinStockThresholdEQ applies the EQ predicate on the "min_stock_threshold" field.
func MinStockThresholdEQ(v int) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldMinStockThreshold, v))
}

// MinStockThresholdNEQ applies the NEQ predicate on the "min_stock_threshold" field.
func MinStockThresholdNEQ(v int) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNEQ(FieldMinStockThreshold, v))
}

// MinStockThresholdIn applies the In predicate on the "min_stock_threshold" field.
func MinStockThresholdIn(vs ...int) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldIn(FieldMinStockThreshold, vs...))
}

// MinStockThresholdNotIn applies the NotIn predicate on the "min_stock_threshold" field.
func MinStockThresholdNotIn(vs ...int) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNotIn(FieldMinStockThreshold, vs...))
}

// MinStockThresholdGT applies the GT predicate on the "min_stock_threshold" field.
func MinStockThresholdGT(v int) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGT(FieldMinStockThreshold, v))
}

// MinStockThresholdGTE applies the GTE predicate on the "min_stock_threshold" field.
func MinStockThresholdGTE(v int) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGTE(FieldMinStockThreshold, v))
}

// MinStockThresholdLT applies the LT predicate on the "min_stock_threshold" field.
func MinStockThresholdLT(v int) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLT(FieldMinStockThreshold, v))
}

// MinStockThresholdLTE applies the LTE predicate on the "min_stock_threshold" field.
func MinStockThresholdLTE(v int) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLTE(FieldMinStockThreshold, v))
}

// MinStockThresholdIsNil applies the IsNil predicate on the "min_stock_threshold" field.
func MinStockThresholdIsNil() predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldIsNull(FieldMinStockThreshold))
}

// MinStockThresholdNotNil applies the NotNil predicate on the "min_stock_threshold" field.
func MinStockThresholdNotNil() predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNotNull(FieldMinStockThreshold))
}

// SkuEQ applies the EQ predicate on the "sku" field.
func SkuEQ(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldSku, v))
}

// SkuNEQ applies the NEQ predicate on the "sku" field.
func SkuNEQ(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNEQ(FieldSku, v))
}

// SkuIn applies the In predicate on the "sku" field.
func SkuIn(vs ...string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldIn(FieldSku, vs...))
}

// SkuNotIn applies the NotIn predicate on the "sku" field.
func SkuNotIn(vs ...string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNotIn(FieldSku, vs...))
}

// SkuGT applies the GT predicate on the "sku" field.
func SkuGT(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGT(FieldSku, v))
}

// SkuGTE applies the GTE predicate on the "sku" field.
func SkuGTE(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGTE(FieldSku, v))
}

// SkuLT applies the LT predicate on the "sku" field.
func SkuLT(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLT(FieldSku, v))
}

// SkuLTE applies the LTE predicate on the "sku" field.
func SkuLTE(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLTE(FieldSku, v))
}

// SkuContains applies the Contains predicate on the "sku" field.
func SkuContains(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldContains(FieldSku, v))
}

// SkuHasPrefix applies the HasPrefix predicate on the "sku" field.
func SkuHasPrefix(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldHasPrefix(FieldSku, v))
}

// SkuHasSuffix applies the HasSuffix predicate on the "sku" field.
func SkuHasSuffix(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldHasSuffix(FieldSku, v))
}

// SkuIsNil applies the IsNil predicate on the "sku" field.
func SkuIsNil() predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldIsNull(FieldSku))
}

// SkuNotNil applies the NotNil predicate on the "sku" field.
func SkuNotNil() predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNotNull(FieldSku))
}

// SkuEqualFold applies the EqualFold predicate on the "sku" field.
func SkuEqualFold(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEqualFold(FieldSku, v))
}

// SkuContainsFold applies the ContainsFold predicate on the "sku" field.
func SkuContainsFold(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldContainsFold(FieldSku, v))
}

// ExpiryDateEQ applies the EQ predicate on the "expiry_date" field.
func ExpiryDateEQ(v time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldExpiryDate, v))
}

// ExpiryDateNEQ applies the NEQ predicate on the "expiry_date" field.
func ExpiryDateNEQ(v time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNEQ(FieldExpiryDate, v))
}

// ExpiryDateIn applies the In predicate on the "expiry_date" field.
func ExpiryDateIn(vs ...time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldIn(FieldExpiryDate, vs...))
}

// ExpiryDateNotIn applies the NotIn predicate on the "expiry_date" field.
func ExpiryDateNotIn(vs ...time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNotIn(FieldExpiryDate, vs...))
}

// ExpiryDateGT applies the GT predicate on the "expiry_date" field.
func ExpiryDateGT(v time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGT(FieldExpiryDate, v))
}

// ExpiryDateGTE applies the GTE predicate on the "expiry_date" field.
func ExpiryDateGTE(v time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGTE(FieldExpiryDate, v))
}

// ExpiryDateLT applies the LT predicate on the "expiry_date" field.
func ExpiryDateLT(v time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLT(FieldExpiryDate, v))
}

// ExpiryDateLTE applies the LTE predicate on the "expiry_date" field.
func ExpiryDateLTE(v time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLTE(FieldExpiryDate, v))
}

// ExpiryDateIsNil applies the IsNil predicate on the "expiry_date" field.
func ExpiryDateIsNil() predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldIsNull(FieldExpiryDate))
}

// ExpiryDateNotNil applies the NotNil predicate on the "expiry_date" field.
func ExpiryDateNotNil() predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNotNull(FieldExpiryDate))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLTE(FieldSize, v))
}

// SizeContains applies the Contains predicate on the "size" field.
func SizeContains(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldContains(FieldSize, v))
}

// SizeHasPrefix applies the HasPrefix predicate on the "size" field.
func SizeHasPrefix(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldHasPrefix(FieldSize, v))
}

// SizeHasSuffix applies the HasSuffix predicate on the "size" field.
func SizeHasSuffix(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldHasSuffix(FieldSize, v))
}

// SizeIsNil applies the IsNil predicate on the "size" field.
func SizeIsNil() predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldIsNull(FieldSize))
}

// SizeNotNil applies the NotNil predicate on the "size" field.
func SizeNotNil() predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNotNull(FieldSize))
}

// SizeEqualFold applies the EqualFold predicate on the "size" field.
func SizeEqualFold(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEqualFold(FieldSize, v))
}

// SizeContainsFold applies the ContainsFold predicate on the "size" field.
func SizeContainsFold(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldContainsFold(FieldSize, v))
}

// ColorEQ applies the EQ predicate on the "color" field.
func ColorEQ(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldColor, v))
}

// ColorNEQ applies the NEQ predicate on the "color" field.
func ColorNEQ(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNEQ(FieldColor, v))
}

// ColorIn applies the In predicate on the "color" field.
func ColorIn(vs ...string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldIn(FieldColor, vs...))
}

// ColorNotIn applies the NotIn predicate on the "color" field.
func ColorNotIn(vs ...string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNotIn(FieldColor, vs...))
}

// ColorGT applies the GT predicate on the "color" field.
func ColorGT(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGT(FieldColor, v))
}

// ColorGTE applies the GTE predicate on the "color" field.
func ColorGTE(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGTE(FieldColor, v))
}

// ColorLT applies the LT predicate on the "color" field.
func ColorLT(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLT(FieldColor, v))
}

// ColorLTE applies the LTE predicate on the "color" field.
func ColorLTE(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLTE(FieldColor, v))
}

// ColorContains applies the Contains predicate on the "color" field.
func ColorContains(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldContains(FieldColor, v))
}

// ColorHasPrefix applies the HasPrefix predicate on the "color" field.
func ColorHasPrefix(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldHasPrefix(FieldColor, v))
}

// ColorHasSuffix applies the HasSuffix predicate on the "color" field.
func ColorHasSuffix(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldHasSuffix(FieldColor, v))
}

// ColorIsNil applies the IsNil predicate on the "color" field.
func ColorIsNil() predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldIsNull(FieldColor))
}

// ColorNotNil applies the NotNil predicate on the "color" field.
func ColorNotNil() predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNotNull(FieldColor))
}

// ColorEqualFold applies the EqualFold predicate on the "color" field.
func ColorEqualFold(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEqualFold(FieldColor, v))
}

// ColorContainsFold applies the ContainsFold predicate on the "color" field.
func ColorContainsFold(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldContainsFold(FieldColor, v))
}

// BrandEQ applies the EQ predicate on the "brand" field.
func BrandEQ(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldBrand, v))
}

// BrandNEQ applies the NEQ predicate on the "brand" field.
func BrandNEQ(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNEQ(FieldBrand, v))
}

// BrandIn applies the In predicate on the "brand" field.
func BrandIn(vs ...string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldIn(FieldBrand, vs...))
}

// BrandNotIn applies the NotIn predicate on the "brand" field.
func BrandNotIn(vs ...string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNotIn(FieldBrand, vs...))
}

// BrandGT applies the GT predicate on the "brand" field.
func BrandGT(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGT(FieldBrand, v))
}

// BrandGTE applies the GTE predicate on the "brand" field.
func BrandGTE(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGTE(FieldBrand, v))
}

// BrandLT applies the LT predicate on the "brand" field.
func BrandLT(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLT(FieldBrand, v))
}

// BrandLTE applies the LTE predicate on the "brand" field.
func BrandLTE(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLTE(FieldBrand, v))
}

// BrandContains applies the Contains predicate on the "brand" field.
func BrandContains(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldContains(FieldBrand, v))
}

// BrandHasPrefix applies the HasPrefix predicate on the "brand" field.
func BrandHasPrefix(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldHasPrefix(FieldBrand, v))
}

// BrandHasSuffix applies the HasSuffix predicate on the "brand" field.
func BrandHasSuffix(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldHasSuffix(FieldBrand, v))
}

// BrandIsNil applies the IsNil predicate on the "brand" field.
func BrandIsNil() predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldIsNull(FieldBrand))
}

// BrandNotNil applies the NotNil predicate on the "brand" field.
func BrandNotNil() predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNotNull(FieldBrand))
}

// BrandEqualFold applies the EqualFold predicate on the "brand" field.
func BrandEqualFold(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEqualFold(FieldBrand, v))
}

// BrandContainsFold applies the ContainsFold predicate on the "brand" field.
func BrandContainsFold(v string) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldContainsFold(FieldBrand, v))
}

// PurchaseDateEQ applies the EQ predicate on the "purchase_date" field.
func PurchaseDateEQ(v time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldPurchaseDate, v))
}

// PurchaseDateNEQ applies the NEQ predicate on the "purchase_date" field.
func PurchaseDateNEQ(v time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNEQ(FieldPurchaseDate, v))
}

// PurchaseDateIn applies the In predicate on the "purchase_date" field.
func PurchaseDateIn(vs ...time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldIn(FieldPurchaseDate, vs...))
}

// PurchaseDateNotIn applies the NotIn predicate on the "purchase_date" field.
func PurchaseDateNotIn(vs ...time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNotIn(FieldPurchaseDate, vs...))
}

// PurchaseDateGT applies the GT predicate on the "purchase_date" field.
func PurchaseDateGT(v time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGT(FieldPurchaseDate, v))
}

// PurchaseDateGTE applies the GTE predicate on the "purchase_date" field.
func PurchaseDateGTE(v time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGTE(FieldPurchaseDate, v))
}

// PurchaseDateLT applies the LT predicate on the "purchase_date" field.
func PurchaseDateLT(v time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLT(FieldPurchaseDate, v))
}

// PurchaseDateLTE applies the LTE predicate on the "purchase_date" field.
func PurchaseDateLTE(v time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLTE(FieldPurchaseDate, v))
}

// PurchaseDateIsNil applies the IsNil predicate on the "purchase_date" field.
func PurchaseDateIsNil() predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldIsNull(FieldPurchaseDate))
}

// PurchaseDateNotNil applies the NotNil predicate on the "purchase_date" field.
func PurchaseDateNotNil() predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNotNull(FieldPurchaseDate))
}

// LastUpdatedEQ applies the EQ predicate on the "last_updated" field.
func LastUpdatedEQ(v time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldEQ(FieldLastUpdated, v))
}

// LastUpdatedNEQ applies the NEQ predicate on the "last_updated" field.
func LastUpdatedNEQ(v time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNEQ(FieldLastUpdated, v))
}

// LastUpdatedIn applies the In predicate on the "last_updated" field.
func LastUpdatedIn(vs ...time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldIn(FieldLastUpdated, vs...))
}

// LastUpdatedNotIn applies the NotIn predicate on the "last_updated" field.
func LastUpdatedNotIn(vs ...time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldNotIn(FieldLastUpdated, vs...))
}

// LastUpdatedGT applies the GT predicate on the "last_updated" field.
func LastUpdatedGT(v time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGT(FieldLastUpdated, v))
}

// LastUpdatedGTE applies the GTE predicate on the "last_updated" field.
func LastUpdatedGTE(v time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldGTE(FieldLastUpdated, v))
}

// LastUpdatedLT applies the LT predicate on the "last_updated" field.
func LastUpdatedLT(v time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLT(FieldLastUpdated, v))
}

// LastUpdatedLTE applies the LTE predicate on the "last_updated" field.
func LastUpdatedLTE(v time.Time) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.FieldLTE(FieldLastUpdated, v))
}

// HasPlace applies the HasEdge predicate on the "place" edge.
func HasPlace() predicate.PlaceInventory {
	return predicate.PlaceInventory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PlaceTable, PlaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlaceWith applies the HasEdge predicate on the "place" edge with a given conditions (other predicates).
func HasPlaceWith(preds ...predicate.Place) predicate.PlaceInventory {
	return predicate.PlaceInventory(func(s *sql.Selector) {
		step := newPlaceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInventoryType applies the HasEdge predicate on the "inventory_type" edge.
func HasInventoryType() predicate.PlaceInventory {
	return predicate.PlaceInventory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InventoryTypeTable, InventoryTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInventoryTypeWith applies the HasEdge predicate on the "inventory_type" edge with a given conditions (other predicates).
func HasInventoryTypeWith(preds ...predicate.InventoryType) predicate.PlaceInventory {
	return predicate.PlaceInventory(func(s *sql.Selector) {
		step := newInventoryTypeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAttributes applies the HasEdge predicate on the "attributes" edge.
func HasAttributes() predicate.PlaceInventory {
	return predicate.PlaceInventory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AttributesTable, AttributesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAttributesWith applies the HasEdge predicate on the "attributes" edge with a given conditions (other predicates).
func HasAttributesWith(preds ...predicate.PlaceInventoryAttribute) predicate.PlaceInventory {
	return predicate.PlaceInventory(func(s *sql.Selector) {
		step := newAttributesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMedia applies the HasEdge predicate on the "media" edge.
func HasMedia() predicate.PlaceInventory {
	return predicate.PlaceInventory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, MediaTable, MediaPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMediaWith applies the HasEdge predicate on the "media" edge with a given conditions (other predicates).
func HasMediaWith(preds ...predicate.Media) predicate.PlaceInventory {
	return predicate.PlaceInventory(func(s *sql.Selector) {
		step := newMediaStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTransactionHistories applies the HasEdge predicate on the "transaction_histories" edge.
func HasTransactionHistories() predicate.PlaceInventory {
	return predicate.PlaceInventory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TransactionHistoriesTable, TransactionHistoriesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTransactionHistoriesWith applies the HasEdge predicate on the "transaction_histories" edge with a given conditions (other predicates).
func HasTransactionHistoriesWith(preds ...predicate.TransactionHistory) predicate.PlaceInventory {
	return predicate.PlaceInventory(func(s *sql.Selector) {
		step := newTransactionHistoriesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReservationBlocks applies the HasEdge predicate on the "reservation_blocks" edge.
func HasReservationBlocks() predicate.PlaceInventory {
	return predicate.PlaceInventory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReservationBlocksTable, ReservationBlocksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReservationBlocksWith applies the HasEdge predicate on the "reservation_blocks" edge with a given conditions (other predicates).
func HasReservationBlocksWith(preds ...predicate.ReservationBlock) predicate.PlaceInventory {
	return predicate.PlaceInventory(func(s *sql.Selector) {
		step := newReservationBlocksStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBusiness applies the HasEdge predicate on the "business" edge.
func HasBusiness() predicate.PlaceInventory {
	return predicate.PlaceInventory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BusinessTable, BusinessColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessWith applies the HasEdge predicate on the "business" edge with a given conditions (other predicates).
func HasBusinessWith(preds ...predicate.Business) predicate.PlaceInventory {
	return predicate.PlaceInventory(func(s *sql.Selector) {
		step := newBusinessStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCategory applies the HasEdge predicate on the "category" edge.
func HasCategory() predicate.PlaceInventory {
	return predicate.PlaceInventory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CategoryTable, CategoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoryWith applies the HasEdge predicate on the "category" edge with a given conditions (other predicates).
func HasCategoryWith(preds ...predicate.Category) predicate.PlaceInventory {
	return predicate.PlaceInventory(func(s *sql.Selector) {
		step := newCategoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PlaceInventory) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PlaceInventory) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PlaceInventory) predicate.PlaceInventory {
	return predicate.PlaceInventory(sql.NotPredicates(p))
}
