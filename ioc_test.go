package ioc

import "testing"

type objTest struct {
	name        string
	description string
}

func TestIsRegistered(t *testing.T) {
	Bind("tester", func() interface{} {
		return &objTest{
			"test",
			"default",
		}
	})

	registered := IsRegistered("tester")

	if !registered {
		t.Error("Should be registered")
	}

	registered = IsRegistered("tester_2")

	if registered {
		t.Error("Should not be registered")
	}
}

func TestBind(t *testing.T) {
	Bind("tester", func() interface{} {
		return &objTest{
			"test",
			"default",
		}
	})

	obj, ok := Make("tester").(*objTest)

	if !ok {
		t.Error("Cannot get correct instance")
	}

	if obj.name != "test" {
		t.Errorf("Expected name: %s, actual: %s", "test", obj.name)
	}

	if obj.description != "default" {
		t.Errorf("Expected description: %s, actual: %s", "default", obj.description)
	}

	obj.name = "new-test"
	obj.description = "new-desc"

	obj2, ok := Make("tester").(*objTest)

	if obj.name == obj2.name {
		t.Error("Name should not be equal")
	}

	if obj.description == obj2.description {
		t.Error("Description should not be equal")
	}
}

func TestSingleton(t *testing.T) {
	Singleton("tester", func() interface{} {
		return &objTest{
			"test",
			"default",
		}
	})

	obj, ok := Make("tester").(*objTest)

	if !ok {
		t.Error("Cannot get correct instance")
	}

	if obj.name != "test" {
		t.Errorf("Expected name: %s, actual: %s", "test", obj.name)
	}

	if obj.description != "default" {
		t.Errorf("Expected description: %s, actual: %s", "default", obj.description)
	}

	obj.name = "new-test"
	obj.description = "new-desc"

	obj2, ok := Make("tester").(*objTest)

	if obj.name != obj2.name {
		t.Error("Name should be equal")
	}

	if obj.description != obj2.description {
		t.Error("Description should be equal")
	}
}
