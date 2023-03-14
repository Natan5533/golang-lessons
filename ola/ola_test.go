package main

import "testing"

func TestOla(t *testing.T) {
    verifyResponse := func(t *testing.T, response, expected string){
        t.Helper()
        if response != expected {
            t.Errorf("resultado '%s', esperado '%s'", response, expected)
            
        }
    }
    t.Run("Hello by user", func(t *testing.T) {
        response := Ola("Natan", "")
        expected := "Hello Natan"
        verifyResponse(t, response, expected)
    })
    t.Run("Hello world by default", func(t *testing.T){
        response := Ola("", "")
        expected := "Hello World"

        verifyResponse(t, response, expected)
    })
    t.Run("in portuguese", func(t *testing.T){
        response := Ola("Francisco", "portuguese")
        expected := "Eae, Francisco"

        verifyResponse(t, response, expected)
    })
    t.Run("In French", func(t *testing.T){
        response := Ola("Luffy", "french")
        expected := "Bonjour Luffy"

        verifyResponse(t, response, expected)
    })
}