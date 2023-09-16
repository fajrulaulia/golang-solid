package lsp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPembayaranKartuKredit(t *testing.T) {
	cc := &BayarDenganKartuKredit{Tujuan: "Admin Fajrul", Jumlah: 300000, Tenor: 12}
	result := ProsesPembayaran(cc)
	assert.Equal(t, result, "Anda sudah membayar dengan CC senilai 300000 kepada Admin Fajrul dengan tenor 12")

	cod := &BayarTunaiCOD{NamaPenerima: "Widia", Jumlah: 300000, BiayaOngkir: 10000}
	result = ProsesPembayaran(cod)
	assert.Equal(t, result, "Anda sudah membayar COD senilai 300000 kepada Widia dengan Biaya Ongkir 10000")

}
