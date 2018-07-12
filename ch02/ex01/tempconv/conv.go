package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func KToC(k Kelvin) Celsius { return Celsius(k) + AbsoluteZeroC }

func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

func FToK(f Fahrenheit) Kelvin { return Kelvin((f - AbsoluteZeroF) * 5 / 9) }

func KToF(k Kelvin) Fahrenheit { return Fahrenheit(k*9/5) + AbsoluteZeroF }
