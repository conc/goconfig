package goconfig

import (
	"testing"
)

func Test_Fortest(t *testing.T) {

	testConfig := LoadConfig("./testconfig/test.ini")

	if testConfig.GetString("user") != "hehehe" {
		t.Errorf("err GetString: user")
	}

	if testConfig.GetString("passwd") != "22222" {
		t.Errorf("err GetString: passwd")
	}

	if testConfig.GetString("new") != "" {
		t.Errorf("err GetString: new (is comment)")
	}

	if testConfig.GetInt("id") != 5 {
		t.Errorf("err GetInt: id")
	}

	if testConfig.GetString("key") != "alssdjflkasd==asd=fa-=sdfa=sf=asd=fa=s" {
		t.Errorf("err GetString: key")
	}
	if testConfig.GetString("key2") != "adjflasjlds = akljflasdjflasdfa333" {
		t.Errorf("err GetString: key2")
	}
	if testConfig.GetString("key3") != "#akdfjlkasjflksdjflkas" {
		t.Errorf("err GetString: key3")
	}
	if testConfig.GetString("k#ey4") != "jadslfk//asdfasd023o 0sdf9a0sdflaassd" {
		t.Errorf("err GetString: key4")
	}

	if testConfig.GetBool("key5") != false {
		t.Errorf("err GetBool: key5")
	}

}
