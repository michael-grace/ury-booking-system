package config

type configuration struct {
	LogFile string
	Port    int
}

// Config contains the parameters for the booking system
var Config configuration = configuration{
	LogFile: "./bookings.log",
	Port:    80,
}
