package lsp

import (
	"fmt"
)

type TipePembayaran interface {
	Bayar() string
}

type BayarDenganKartuKredit struct {
	Tujuan string
	Jumlah int
	Tenor  int
}

func (b *BayarDenganKartuKredit) Bayar() string {
	return fmt.Sprintf("Anda sudah membayar dengan CC senilai %d kepada %s dengan tenor %d", b.Jumlah, b.Tujuan, b.Tenor)
}

// Penamahan Kode baru
type BayarTunaiCOD struct {
	NamaPenerima string
	Jumlah       int
	BiayaOngkir  int
}

func (b *BayarTunaiCOD) Bayar() string {
	return fmt.Sprintf("Anda sudah membayar COD senilai %d kepada %s dengan Biaya Ongkir %d", b.Jumlah, b.NamaPenerima, b.BiayaOngkir)
}

// ==========================

func ProsesPembayaran(tipePembayaran TipePembayaran) string {
	return tipePembayaran.Bayar()
}

// func main() {

// 	cc := &BayarDenganKartuKredit{Tujuan: "Admin Fajrul", Jumlah: 300000, Tenor: 12}
// 	result := ProsesPembayaran(cc)
// 	log.Println("Kartu Kredit = ", result)

// 	cod := &BayarTunaiCOD{NamaPenerima: "Widia", Jumlah: 300000, BiayaOngkir: 10000}
// 	result = ProsesPembayaran(cod)
// 	log.Println("COD = ", result)

// }
