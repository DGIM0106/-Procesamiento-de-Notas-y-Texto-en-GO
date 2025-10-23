# Sistema de Análisis de Notas y Texto

Un programa en Go que procesa notas de estudiantes y analiza texto de entrada, demostrando el procesamiento básico de datos y manipulación de cadenas en Go.

## Descripción

Este programa consta de dos ejercicios principales:
1. **Análisis de Notas de Estudiantes**: Procesa un array bidimensional de notas para calcular notas más altas/más bajas, promedios individuales y promedio general del curso.
2. **Análisis de Texto**: Toma texto ingresado por el usuario y realiza análisis de frecuencia de palabras y listado alfabético.

## Características

### Ejercicio 1: Procesamiento de Notas
- Calcula la nota más alta y más baja para cada estudiante
- Computa promedios individuales de estudiantes
- Calcula el promedio general del curso
- Procesa una matriz de 6x4 de notas

### Ejercicio 2: Análisis de Texto
- Lee texto ingresado por el usuario
- Separa el texto en palabras
- Cuenta la frecuencia de palabras
- Muestra palabras en orden alfabético

## Estructura del Código

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    // Ejercicio 1: Procesamiento de notas
    // Ejercicio 2: Análisis de texto
}
```

## Instalación

1. Asegúrate de tener Go instalado en tu sistema (versión 1.16+ recomendada)
2. Clona o descarga el repositorio
3. Navega al directorio del proyecto

## Uso

### Ejecutar el Programa

```bash
go run main.go
```

### Ejemplo de Ejecución

**Salida del Ejercicio 1:**
```
******Ejercicio 1******
******Nota mas alta y baja del estudiante******
La nota mas alta del estudiante 1 es: 9.5
La nota mas baja del estudiante 1 es: 8.6
La nota mas alta del estudiante 2 es: 9.2
La nota mas baja del estudiante 2 es: 6.8
...
******Promedio por estudiante******
La nota promediada del estudiante 1 es de: 9.0
La nota promediada del estudiante 2 es de: 7.875
...
******Promedio general******
Promedio general del curso es de: 8.541666666666666
```

**Salida del Ejercicio 2:**
```
******Ejercicio 2******
Ingrese el texto: hola mundo hola otra vez mundo
Lista de palabras ordenadas
hola
mundo
hola
otra
vez
mundo
La palabra "hola" se repite: 2
La palabra "mundo" se repite: 2
La palabra "otra" se repite: 1
La palabra "vez" se repite: 1
```

## Formato de Entrada

### Datos de Notas
El programa usa una matriz 6x4 precargada que representa:
- 6 estudiantes
- 4 notas por estudiante

### Entrada de Texto
Cuando se solicite, ingresa cualquier texto. El programa:
- Separará el texto en palabras (usando espacios como separadores)
- Contará la frecuencia de palabras
- Mostrará palabras en el orden en que aparecen

## Funcionalidades

### Procesamiento de Notas
- **Iteración de Matriz**: Procesa las notas de cada estudiante
- **Cálculo Min/Max**: Encuentra notas más altas y más bajas por estudiante
- **Cálculo de Promedio**: Computa promedios individuales y general

### Análisis de Texto
- **Lectura de Entrada**: Usa `bufio.Scanner` para entrada del usuario
- **División de Cadenas**: Separa texto en palabras
- **Conteo de Frecuencia**: Usa mapas de Go para contar ocurrencias de palabras
- **Visualización de Salida**: Muestra tanto lista ordenada como conteos de frecuencia

## Estructuras de Datos

- **Array Bidimensional**: `[6][4]float64` para almacenar notas
- **Slices**: `[]float64` para almacenar sumas y promedios
- **Mapa**: `map[string]int` para conteo de frecuencia de palabras

## Dependencias

- Solo librerías estándar de Go:
  - `bufio`
  - `fmt`
  - `os`
  - `strings`

## Personalización

Para modificar los datos de notas, edita el array `bidi` en la función principal:

```go
bidi := [6][4]float64{
    {9.5, 8.9, 8.6, 9.0},
    {7.5, 6.8, 8.0, 9.2},
    // ... agregar más datos de estudiantes
}
```

## Limitaciones

- Los datos de notas están precargados (no se cargan desde archivos externos)
- La separación de texto usa separación simple por espacios
- La ordenación de palabras mantiene el orden de entrada en lugar de una ordenación alfabética verdadera

## Valor Educativo

Este programa demuestra:
- Manipulación de arrays bidimensionales en Go
- Cálculos estadísticos básicos
- Procesamiento de cadenas y uso de mapas
- Manejo de entrada del usuario
- Operaciones con slices y arrays dinámicos
