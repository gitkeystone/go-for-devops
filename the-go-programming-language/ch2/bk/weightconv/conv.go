package weightconv

func BToK(bl Pound) Kilogram {
	return Kilogram(bl * 0.45359237)
}

func KToB(k Kilogram) Pound {
	return Pound(k / 0.45359237)
}
