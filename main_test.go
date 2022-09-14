package main

import "testing"

func TestRomainToInteger(t *testing.T) {

	if val := romainToInteger("XII"); val != 12 {
		t.Errorf("Conversion incorrecte, obtenu: %d, attendu: %d.", val, 12)
	}

	if val := romainToInteger("XIX"); val != 19 {
		t.Errorf("Conversion incorrecte, obtenu: %d, attendu: %d.", val, 19)
	}

	if val := romainToInteger("XLIX"); val != 49 {
		t.Errorf("Conversion incorrecte, obtenu: %d, attendu: %d.", val, 49)
	}

	if val := romainToInteger("LXXVIII"); val != 78 {
		t.Errorf("Conversion incorrecte, obtenu: %d, attendu: %d.", val, 78)
	}

	if val := romainToInteger("LXXIX"); val != 79 {
		t.Errorf("Conversion incorrecte, obtenu: %d, attendu: %d.", val, 79)
	}

	if val := romainToInteger("CI"); val != 101 {
		t.Errorf("Conversion incorrecte, obtenu: %d, attendu: %d.", val, 101)
	}

	if val := romainToInteger("MCMLXXXIX"); val != 1989 {
		t.Errorf("Conversion incorrecte, obtenu: %d, attendu: %d.", val, 1989)
	}

	if val := romainToInteger("MMMMCMXXIV"); val != 4924 {
		t.Errorf("Conversion incorrecte, obtenu: %d, attendu: %d.", val, 4924)
	}

}
