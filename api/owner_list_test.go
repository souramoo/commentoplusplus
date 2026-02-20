package main

import (
	"testing"
)

func TestOwnerListBasics(t *testing.T) {
	failTestOnError(t, setupTestEnv())

	ownerHex1, _ := ownerNew("test1@example.com", "Test1", "hunter2", false)
	ownerHex2, _ := ownerNew("test2@example.com", "Test2", "hunter2", false)

	owners, err := ownerList()
	if err != nil {
		t.Errorf("unexpected error listing owners: %v", err)
		return
	}

	// We expect three owners here because one owner is already created in setupTestEnv.
	if len(owners) != 3 {
		t.Errorf("expected number of owners to be 3 got %d", len(owners))
		return
	}

	foundOwner1 := false
	foundOwner2 := false
	for _, o := range owners {
		if o.OwnerHex == ownerHex1 {
			foundOwner1 = true
			if o.Email != "test1@example.com" {
				t.Errorf("expected email=test1@example.com got email=%s", o.Email)
			}
			if o.Name != "Test1" {
				t.Errorf("expected name=Test1 got name=%s", o.Name)
			}
		}
		if o.OwnerHex == ownerHex2 {
			foundOwner2 = true
			if o.Email != "test2@example.com" {
				t.Errorf("expected email=test2@example.com got email=%s", o.Email)
			}
			if o.Name != "Test2" {
				t.Errorf("expected name=Test2 got name=%s", o.Name)
			}
		}
	}

	if !foundOwner1 {
		t.Errorf("owner1 not found in list")
	}
	if !foundOwner2 {
		t.Errorf("owner2 not found in list")
	}
}

func TestOwnerListEmpty(t *testing.T) {
	failTestOnError(t, setupTestEnv())

	owners, err := ownerList()
	if err != nil {
		t.Errorf("unexpected error listing owners: %v", err)
		return
	}

	// We expect one owner here because one owner is already created in setupTestEnv.
	if len(owners) != 1 {
		t.Errorf("expected number of owners to be 1 got %d", len(owners))
		return
	}
}
