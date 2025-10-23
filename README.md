# -Procesamiento-de-Notas-y-Texto-en-GO
Descripción General
Este programa en Go realiza dos ejercicios principales:

Procesamiento de calificaciones estudiantiles

Análisis de texto ingresado por el usuario

Ejercicio 1: Procesamiento de Calificaciones
Estructura de Datos
go
bidi := [6][4]float64{
    {9.5, 8.9, 8.6, 9.0},
    {7.5, 6.8, 8.0, 9.2},
    {9.3, 9.7, 8.5, 8.4},
    {6.5, 7.8, 8.9, 9.1},
    {8.0, 8.3, 7.5, 6.9},
    {9.9, 10.0, 8.7, 9.8},
}
Funcionalidades Implementadas
1. Cálculo de Notas Extremas
Nota más alta: Para cada estudiante, encuentra la calificación máxima

Nota más baja: Para cada estudiante, encuentra la calificación mínima

2. Cálculo de Promedios
Promedio individual: Calcula el promedio de cada estudiante (suma de 4 notas / 4)

Promedio general: Calcula el promedio de todos los estudiantes

Salida Esperada
text
******Ejercicio 1******
******Nota mas alta y baja del estudiante******
La nota mas alta del estudiante 1 es: 9.5
La nota mas baja del estudiante 1 es: 8.6
...
******Promedio por estudiante******
La nota promediada del estudiante 1 es de: 9.0
...
******Promedio general******
Promedio general del curso es de: [valor calculado]
Ejercicio 2: Análisis de Texto
Proceso de Procesamiento
Entrada de texto: Usa bufio.NewScanner para leer entrada del usuario

Separación de palabras:

Primero divide por "/n" (posible error, debería ser "\n")

Luego usa strings.Fields para separar por espacios

Conteo de palabras: Utiliza un mapa para contar repeticiones

Salida ordenada: Muestra palabras en orden de aparición

Características
Almacenamiento: Slice txtsep para palabras individuales

Conteo: Mapa contador para frecuencia de palabras

Orden: Mantiene el orden original de las palabras

Posibles Mejoras
Corrección del separador de líneas (de "/n" a "\n")

Ordenamiento alfabético real de las palabras

Manejo de mayúsculas/minúsculas para conteo

Eliminación de caracteres especiales

Estructura del Código
Variables Principales
bidi: Matriz 6x4 de calificaciones

sumanot: Slice con sumas individuales

promedio: Slice con promedios individuales

contador: Mapa para frecuencia de palabras

Flujo de Ejecución
Procesamiento completo de calificaciones

Solicitud y procesamiento de texto

Análisis y presentación de resultados

Consideraciones Técnicas
Usa arrays estáticos para calificaciones

Emplea slices dinámicos para resultados

Utiliza maps para conteo eficiente

Maneja entrada estándar con bufio

El código demuestra conceptos fundamentales de Go como arrays, slices, maps, y procesamiento de strings.
