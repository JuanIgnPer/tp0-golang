package utils

import (
	"bufio"
	"bytes"
	"client/globals"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Mensaje struct {
	Mensaje string `json:"mensaje"`
}

type Paquete struct {
	Valores []string `json:"valores"`
}

func IniciarConfiguracion(filePath string) *globals.Config {
	var config *globals.Config
	// configFile: Archivo abierto
	// err: error en caso de existir
	configFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	// Cerrar el archivo automaticamente cuando la funcion termine
	defer configFile.Close()

	// Decodificador de Json que va a leer el archivo abierto
	jsonParser := json.NewDecoder(configFile)
	// Inicializa el Config que esta en la direccion de memoria de config
	// Guarda el resultado del parseo en esa direccion
	jsonParser.Decode(&config)

	return config
}

func LeerConsola() {
	// Leer de la consola
	reader := bufio.NewReader(os.Stdin)
	log.Println("Ingrese los mensajes")
	text, _ := reader.ReadString('\n')
	log.Print(text)
}

func GenerarYEnviarPaquete(ip string, puerto int){
	paquete := Paquete{}
	// Leemos y cargamos el paquete
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		if text == "\r\n" || text == "\n" {
			// Corta cuando lee un enter
			break
		}
		paquete.Valores = append(paquete.Valores, text)
	}
	log.Printf("paquete a enviar: %+v", paquete)
	// Enviamos el paquete
	EnviarPaquete(ip, puerto, paquete)
}

func EnviarMensaje(ip string, puerto int, mensajeTxt string) {
	mensaje := Mensaje{Mensaje: mensajeTxt}
	body, err := json.Marshal(mensaje)
	if err != nil {
		log.Printf("error codificando mensaje: %s", err.Error())
	}

	url := fmt.Sprintf("http://%s:%d/mensaje", ip, puerto)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("error enviando mensaje a ip:%s puerto:%d", ip, puerto)
	}

	log.Printf("respuesta del servidor: %s", resp.Status)
}

func EnviarPaquete(ip string, puerto int, paquete Paquete) {
	body, err := json.Marshal(paquete)
	if err != nil {
		log.Printf("error codificando mensajes: %s", err.Error())
	}

	url := fmt.Sprintf("http://%s:%d/paquetes", ip, puerto)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("error enviando mensajes a ip:%s puerto:%d", ip, puerto)
	}

	log.Printf("respuesta del servidor: %s", resp.Status)
}

func ConfigurarLogger() {
	logFile, err := os.OpenFile("tp0.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}
