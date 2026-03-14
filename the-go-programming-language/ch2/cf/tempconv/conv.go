package tempconv

// CToF 把摄氏温度转化为华氏温度
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// CToK 把摄氏温度转化为开尔文温度
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

// FToC 把华氏温度转化为摄氏温度
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// KToC 把开尔文温度转化为摄氏温度
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}
