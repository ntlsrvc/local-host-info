package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
)

// Estrutura para armazenar as informações do computador
type Computador struct {
	Hostname   string
	IP         string
	MacAddress string
}

// Função principal
func main() {
	templates := template.Must(template.ParseGlob("*.html"))

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Rota para /home
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		// Obtém o nome do computador
		hostname, err := os.Hostname()
		if err != nil {
			log.Println("Erro ao obter o nome do computador:", err)
			hostname = "Desconhecido"
		}

		// Obtém o endereço IP
		ip, err := getIP()
		if err != nil {
			log.Println("Erro ao obter o IP:", err)
			ip = "Desconhecido"
		}

		// Obtém o endereço MAC
		mac, err := getMACAddress()
		if err != nil {
			log.Println("Erro ao obter o endereço MAC:", err)
			mac = "Desconhecido"
		}

		// Cria a estrutura Computador com os dados
		computador := Computador{
			Hostname:   hostname,
			IP:         ip,
			MacAddress: mac,
		}

		// Executa o template HTML e passa os dados
		templates.ExecuteTemplate(w, "index.html", computador)
	})

	// Inicia o servidor na porta 5000
	log.Fatal(http.ListenAndServe(":5000", nil))
}

// Função para obter o IP do computador
func getIP() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && ipNet.IP.To4() != nil {
				return ipNet.IP.String(), nil
			}
		}
	}
	return "", fmt.Errorf("não foi possível determinar o IP")
}

// Função para obter o endereço MAC
func getMACAddress() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		return iface.HardwareAddr.String(), nil
	}
	return "", fmt.Errorf("não foi possível determinar o endereço MAC")
}
