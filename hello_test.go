package main

import "testing"

func TestHello(t *testing.T){
 if hello() != "Hello World" {
  t.Fatal("Test Fail")
 }

}
