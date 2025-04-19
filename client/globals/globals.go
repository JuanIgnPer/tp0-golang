package globals

type Config struct {
	Ip      string `json:"ip"`
	Puerto  int    `json:"puerto"`
	Mensaje string `json:"mensaje"`
}

// Defino variable global, puntero a UN Config. 
// Es decir la direccion de memoria no el objeto en si
var ClientConfig *Config
