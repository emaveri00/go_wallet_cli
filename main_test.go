package main

import "testing"

func TestDeposit(t *testing.T) {
    p := Portafoglio{}
    p.Deposit(100.0)
    
    expected := 100.0
    if p.getBalance() != expected {
        t.Errorf("Errore: aspettavo %.2f ma ho ottenuto %.2f", expected, p.getBalance())
    }
}

func TestSpend(t *testing.T) {
    p := Portafoglio{balance: 100.0}
    
    // Test spesa valida
    success := p.Spend(50.0)
    if !success || p.getBalance() != 50.0 {
        t.Errorf("Errore nella spesa valida")
    }

    // Test spesa impossibile
    success = p.Spend(200.0)
    if success {
        t.Errorf("Errore: la spesa doveva fallire!")
    }
}