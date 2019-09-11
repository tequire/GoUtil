package auth

import (
	"github.com/coreos/go-oidc"
	"reflect"
	"testing"
	"unsafe"
)

func TestRequireScope(t *testing.T) {
	idToken := idTokenWithGoTalentScope()

	test1 := requireScope(idToken, and, "gotalent.full_access")
	test2 := requireScope(idToken, or, "gotalent.full_access", "gotalent.write_access")
	test3 := requireScope(idToken, and, "gotalent.full_access", "gotalent.write_access")
	test4 := GoTalentWritePolicy(idToken)
	test5 := requireScope(idToken, or, "identity.full_access", "identity.write_access")

	if !test1 {
		t.Errorf("Test1: requireScope returned %t, expected: %t", test1, true)
	}
	if !test2 {
		t.Errorf("Test2: requireScope returned %t, expected: %t", test2, true)
	}
	if test3 {
		t.Errorf("Test3: requireScope returned %t, expected: %t", test3, false)
	}
	if !test4 {
		t.Errorf("Test4: requireScope returned %t, expected: %t", test4, true)
	}
	if test5 {
		t.Errorf("Test5: requireScope returned %t, expected: %t", test5, false)
	}
}

func TestRequireRole(t *testing.T) {
	idToken := idTokenWithAdminRole()

	test1 := requireRole(idToken, and, "Admin")
	test2 := requireRole(idToken, or, "Admin", "HigheredEmployee")
	test3 := requireRole(idToken, and, "Admin", "HigheredEmployee")
	test4 := AdminPolicy(idToken)

	if !test1 {
		t.Errorf("Test1: requireRole returned %t, expected: %t", test1, true)
	}
	if !test2 {
		t.Errorf("Test1: requireRole returned %t, expected: %t", test2, true)
	}
	if test3 {
		t.Errorf("Test1: requireRole returned %t, expected: %t", test3, false)
	}
	if !test4 {
		t.Errorf("Test1: requireRole returned %t, expected: %t", test4, true)
	}
}

func idTokenWithAdminRole() *oidc.IDToken {
	// Initialize idToken with admin role
	idToken := &oidc.IDToken{}
	claims := []byte{123,34,110,98,102,34,58,49,53,54,56,50,49,52,54,51,48,44,34,101,120,112,34,58,49,53,54,56,50,49,56,50,51,48,44,34,105,115,115,34,58,34,104,116,116,112,115,58,47,47,105,100,101,110,116,105,116,121,45,100,101,118,46,104,105,103,104,101,114,101,100,46,103,108,111,98,97,108,34,44,34,97,117,100,34,58,91,34,104,116,116,112,115,58,47,47,105,100,101,110,116,105,116,121,45,100,101,118,46,104,105,103,104,101,114,101,100,46,103,108,111,98,97,108,47,114,101,115,111,117,114,99,101,115,34,44,34,85,115,101,114,65,80,73,34,93,44,34,99,108,105,101,110,116,95,105,100,34,58,34,99,97,102,54,102,57,100,51,45,57,52,50,49,45,52,102,50,102,45,57,49,57,51,45,57,55,98,100,56,56,56,55,100,54,52,57,34,44,34,115,117,98,34,58,34,50,54,49,100,53,101,52,100,45,52,56,97,49,45,52,48,55,101,45,97,50,55,55,45,48,102,48,98,51,49,97,48,50,100,50,55,34,44,34,97,117,116,104,95,116,105,109,101,34,58,49,53,54,56,50,49,52,54,51,48,44,34,105,100,112,34,58,34,108,111,99,97,108,34,44,34,111,114,103,46,115,116,117,100,101,110,116,34,58,34,57,52,54,34,44,34,114,111,108,101,34,58,34,65,100,109,105,110,34,44,34,101,109,97,105,108,34,58,34,97,110,100,101,114,115,64,114,101,99,114,117,116,46,110,111,34,44,34,115,99,111,112,101,34,58,91,34,111,112,101,110,105,100,34,44,34,65,80,73,95,70,85,76,76,95,85,83,69,82,95,65,67,67,69,83,83,34,44,34,111,102,102,108,105,110,101,95,97,99,99,101,115,115,34,93,44,34,97,109,114,34,58,91,34,112,119,100,34,93,125}
	setClaims(idToken, claims)
	return idToken
}

// idTokenWithGoTalentScope returns an idToken with "gotalent.full_access" scope
func idTokenWithGoTalentScope() *oidc.IDToken {
	// Initialize idToken with admin role
	idToken := &oidc.IDToken{}
	claims := []byte{123,34,110,98,102,34,58,49,53,54,56,50,49,51,50,49,52,44,34,101,120,112,34,58,49,53,54,56,51,56,54,48,49,52,44,34,105,115,115,34,58,34,104,116,116,112,115,58,47,47,105,100,101,110,116,105,116,121,46,104,105,103,104,101,114,101,100,46,103,108,111,98,97,108,34,44,34,97,117,100,34,58,91,34,104,116,116,112,115,58,47,47,105,100,101,110,116,105,116,121,46,104,105,103,104,101,114,101,100,46,103,108,111,98,97,108,47,114,101,115,111,117,114,99,101,115,34,44,34,103,111,116,97,108,101,110,116,34,93,44,34,99,108,105,101,110,116,95,105,100,34,58,34,74,111,98,65,100,115,65,112,105,34,44,34,115,99,111,112,101,34,58,91,34,103,111,116,97,108,101,110,116,46,102,117,108,108,95,97,99,99,101,115,115,34,93,125}
	setClaims(idToken, claims)
	return idToken
}

// A helper method for setting private fields in oidc.IDToken
func setClaims(f *oidc.IDToken, to []byte) {
	// Note, simply doing reflect.ValueOf(*f) won't work, need to do this
	pointerVal := reflect.ValueOf(f)
	val := reflect.Indirect(pointerVal)

	member := val.FieldByName("claims")
	ptrToField := unsafe.Pointer(member.UnsafeAddr())
	realPtrToField := (*[]byte)(ptrToField)
	*realPtrToField = to
}
