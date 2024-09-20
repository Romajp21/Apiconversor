package main

//go get github.com/gin-gonic/gin correr este comando antes de ejecutar el proyecto
// paraejecutar el programa go run ./cmd/main.go

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Ruta para la raíz "/"
	r.GET("/", func(c *gin.Context) {
		html := `
		<h1>Bienvenido/a al API de Conversión</h1>
		<p>Ingresa el valor a convertir y selecciona una de las opciones:</p>
		<form action="/convert/celsius-to-fahrenheit" method="GET">
			<label>Convertir Celsius a Fahrenheit:</label>
			<input type="number" name="value" step="0.01" required>
			<button type="submit">Convertir</button>
		</form>
		<br>
		<form action="/convert/fahrenheit-to-celsius" method="GET">
			<label>Convertir Fahrenheit a Celsius:</label>
			<input type="number" name="value" step="0.01" required>
			<button type="submit">Convertir</button>
		</form>
		<br>
		<form action="/convert/km-to-miles" method="GET">
			<label>Convertir Kilómetros a Millas:</label>
			<input type="number" name="value" step="0.01" required>
			<button type="submit">Convertir</button>
		</form>
		<br>
		<form action="/convert/miles-to-km" method="GET">
			<label>Convertir Millas a Kilómetros:</label>
			<input type="number" name="value" step="0.01" required>
			<button type="submit">Convertir</button>
		</form>
		<br>
		<form action="/convert/inches-to-cm" method="GET">
			<label>Convertir Pulgadas a Centímetros:</label>
			<input type="number" name="value" step="0.01" required>
			<button type="submit">Convertir</button>
		</form>
		<br>
		<form action="/convert/cm-to-inches" method="GET">
			<label>Convertir Centímetros a Pulgadas:</label>
			<input type="number" name="value" step="0.01" required>
			<button type="submit">Convertir</button>
		</form>
		<br>
		<form action="/convert/feet-to-meters" method="GET">
			<label>Convertir Pies a Metros:</label>
			<input type="number" name="value" step="0.01" required>
			<button type="submit">Convertir</button>
		</form>
		<br>
		<form action="/convert/meters-to-feet" method="GET">
			<label>Convertir Metros a Pies:</label>
			<input type="number" name="value" step="0.01" required>
			<button type="submit">Convertir</button>
		</form>
		`
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, html)
	})

	// Rutas de la API
	r.GET("/convert/celsius-to-fahrenheit", celsiusToFahrenheit)
	r.GET("/convert/fahrenheit-to-celsius", fahrenheitToCelsius)
	r.GET("/convert/km-to-miles", kilometersToMiles)
	r.GET("/convert/miles-to-km", milesToKilometers)
	r.GET("/convert/inches-to-cm", inchesToCentimeters)
	r.GET("/convert/cm-to-inches", centimetersToInches)
	r.GET("/convert/feet-to-meters", feetToMeters)
	r.GET("/convert/meters-to-feet", metersToFeet)

	// Iniciar el servidor en el puerto 8080
	r.Run(":8080")
}

/// Funciones de conversión

func celsiusToFahrenheit(c *gin.Context) {
	value := c.Query("value")
	if val, err := strconv.ParseFloat(value, 64); err == nil {
		result := (val * 9 / 5) + 32
		c.JSON(http.StatusOK, gin.H{"Celsius": val, "Fahrenheit": result})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Valor no válido"})
	}
}

func fahrenheitToCelsius(c *gin.Context) {
	value := c.Query("value")
	if val, err := strconv.ParseFloat(value, 64); err == nil {
		result := (val - 32) * 5 / 9
		c.JSON(http.StatusOK, gin.H{"Fahrenheit": val, "Celsius": result})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Valor no válido"})
	}
}

func kilometersToMiles(c *gin.Context) {
	value := c.Query("value")
	if val, err := strconv.ParseFloat(value, 64); err == nil {
		result := val * 0.621371
		c.JSON(http.StatusOK, gin.H{"Kilómetros": val, "Millas": result})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Valor no válido"})
	}
}

func milesToKilometers(c *gin.Context) {
	value := c.Query("value")
	if val, err := strconv.ParseFloat(value, 64); err == nil {
		result := val / 0.621371
		c.JSON(http.StatusOK, gin.H{"Millas": val, "Kilómetros": result})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Valor no válido"})
	}
}

func inchesToCentimeters(c *gin.Context) {
	value := c.Query("value")
	if val, err := strconv.ParseFloat(value, 64); err == nil {
		result := val * 2.54
		c.JSON(http.StatusOK, gin.H{"Pulgadas": val, "Centímetros": result})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Valor no válido"})
	}
}

func centimetersToInches(c *gin.Context) {
	value := c.Query("value")
	if val, err := strconv.ParseFloat(value, 64); err == nil {
		result := val / 2.54
		c.JSON(http.StatusOK, gin.H{"Centímetros": val, "Pulgadas": result})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Valor no válido"})
	}
}

func feetToMeters(c *gin.Context) {
	value := c.Query("value")
	if val, err := strconv.ParseFloat(value, 64); err == nil {
		result := val * 0.3048
		c.JSON(http.StatusOK, gin.H{"Pies": val, "Metros": result})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Valor no válido"})
	}
}

func metersToFeet(c *gin.Context) {
	value := c.Query("value")
	if val, err := strconv.ParseFloat(value, 64); err == nil {
		result := val / 0.3048
		c.JSON(http.StatusOK, gin.H{"Metros": val, "Pies": result})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Valor no válido"})
	}
}
