package lenconv

func MToI(m Meter) Inch {
	return Inch(m / 0.254)
}

func IToM(i Inch) Meter {
	return Meter(i * 0.254)
}
