package ocp

type OperasiMatematika interface {
	Kalkulasi(float64, float64) float64
}

type Pertambahan struct{}

func (a *Pertambahan) Kalkulasi(x, y float64) float64 {
	return x + y
}

// Tambah Kode Baru Tanpa Mengganggu Kode Lama, Misal pengurangan========
type Pengurangan struct{} // struct baru

func (a *Pengurangan) Kalkulasi(x, y float64) float64 {
	return x - y
} // implementasikan Kalkulasi yang Baru baru
// =================

// Export Fuction Kalkulasi, tanpa harus merubah ini
func Kalkulasikan(operasi OperasiMatematika, x, y float64) float64 {
	return operasi.Kalkulasi(x, y)
}
