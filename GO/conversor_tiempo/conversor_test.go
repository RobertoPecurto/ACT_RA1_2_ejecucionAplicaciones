package main

import (
	"testing"
)

// Test convertir horas a minutos y segundos
func TestConvertirHoras(t *testing.T) {
	min, seg := convertirHoras(2)
	if min != 120 {
		t.Errorf("convertirHoras(2) min esperado 120, obtenido %.2f", min)
	}
	if seg != 7200 {
		t.Errorf("convertirHoras(2) seg esperado 7200, obtenido %.2f", seg)
	}
}

// Test convertir minutos a horas y segundos
func TestConvertirMinutos(t *testing.T) {
	horas, seg := convertirMinutos(90)
	if horas != 1.5 {
		t.Errorf("convertirMinutos(90) horas esperado 1.5, obtenido %.2f", horas)
	}
	if seg != 5400 {
		t.Errorf("convertirMinutos(90) seg esperado 5400, obtenido %.2f", seg)
	}
}

// Test convertir segundos a horas y minutos (Preparado para fallar)
func TestConvertirSegundosFallido(t *testing.T) {
	horas, min := convertirSegundos(3600)
	if horas != 2 {
		t.Errorf("convertirSegundos(3600) horas esperado 2, obtenido %.2f", horas)
	}
	if min != 60 {
		t.Errorf("convertirSegundos(3600) minutos esperado 60, obtenido %.2f", min)
	}
}
