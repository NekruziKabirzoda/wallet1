package wallet

import (
	"testing"
)

func TestService_RegisterAccount_success(t *testing.T) {
	svc := Service{}
	_, err := svc.RegisterAccount("+992000000001")
	if err != nil {
		t.Error(err)
	}
}
func TestService_FindAccountByID_success(t *testing.T) {
	svc := Service{}
	_, err := svc.FindAccountByID(992000000001)
	if err == nil {
		t.Error(err)
	}
}

// func TestService_FindAccountByID_error(t *testing.T) {
// 	svc := Service{}
// 	_, err := svc.FindAccountByID(992000000001)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	expected := "account not found"

// 	result := svc.FindAccountByID(992000000001)

// 	if (expected != result) {
// 		t.Errorf("invalid result, expected %v, actual: %v", expected, result)
// 	}
// }