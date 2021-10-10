package wallet

import (
	"testing"
)

func TestService_RegisterAccount_success(t *testing.T) {
	svc := Servise{}
	_, err := svc.RegisterAccount("+992000000001")
	if err != nil {
		t.Error(err)
	}
}
func TestService_FindAccountByID_success(t *testing.T) {
	svc := Servise{}
	_, err := svc.FindAccountByID(992000000001)
	if err == nil {
		t.Error(err)
	}
}