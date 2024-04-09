package config_manager

import "testing"

func TestCreatingInstance(t *testing.T) {
	conf := NewConfig()
	if conf == nil {
		t.Errorf("Want nil, Got %v\n", conf)
	}

	conf.set("Name", "John Doe")
	conf.set("Age", "31")

	if k, _ := conf.get("Name"); k != "John Doe" {
		t.Errorf("Want John Doe, Got %v\n", k)
	}

	conf2 := NewConfig()

	if _, err := conf2.get("Age"); err != nil {
		t.Errorf("Got Error: %v\n", err)
	}
}
