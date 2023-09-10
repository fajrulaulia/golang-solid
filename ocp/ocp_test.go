package ocp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTambah(t *testing.T) {

	x := 10.0
	y := 5.0

	tambah := new(Pertambahan)
	pertambahan := Kalkulasikan(tambah, x, y)

	assert.Equal(t, 15.0, pertambahan, fmt.Sprintf("%.2f + %.2f = %.2f\n", x, y, pertambahan))
}

// Kembangkan Kode Untuk Pengurangan tanpa mengacak Function Kalkulasikan
func TestKurang(t *testing.T) {

	x := 10.0
	y := 5.0

	kurang := new(Pengurangan)
	pengurangna := Kalkulasikan(kurang, x, y)

	assert.Equal(t, 5.0, pengurangna, fmt.Sprintf("%.2f - %.2f = %.2f\n", x, y, pengurangna))
}
