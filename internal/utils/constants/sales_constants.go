package constants

const SaleStatusCreated = "created"
const SaleStatusPending = "pending"
const SaleStatusApproved = "approved"
const SaleStatusRejected = "rejected"

const SaleStatusCreatedVal = 1
const SaleStatusPendingVal = 2
const SaleStatusApprovedVal = 3
const SaleStatusRejectedVal = 4

var SalesStatus = map[int]string{
	SaleStatusCreatedVal:  SaleStatusCreated,
	SaleStatusPendingVal:  SaleStatusPending,
	SaleStatusApprovedVal: SaleStatusApproved,
	SaleStatusRejectedVal: SaleStatusRejected,
}
