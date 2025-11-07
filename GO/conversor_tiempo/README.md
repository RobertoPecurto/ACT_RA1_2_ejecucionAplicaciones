# ⏱️ **Conversor de Tiempo en Go**

Este proyecto consiste en un programa sencillo que permite **convertir entre horas, minutos y segundos** de forma interactiva.

---

## ⚙️ Instalación y configuración del entorno

### 🟢 Paso 1. Instalar Go

Descarga la versión más reciente de **Go** desde:
👉 [https://go.dev/dl/](https://go.dev/dl/)

Durante la instalación, asegúrate de que **Go esté en el PATH**.

Verifica la instalación con:

```bash
go version
```

---

### 📁 Paso 2. Estructura del proyecto

Crea una carpeta de trabajo con los siguientes archivos:

```
📂 conversor_tiempo/
 ├── conversor.go
 ├── conversor_test.go
 └── go.mod
```

---

## 🧩 Código principal – `conversor.go`

El programa permite convertir entre horas, minutos y segundos mediante un menú interactivo.

### ✨ Funciones principales

| Función                        | Descripción                           |
| ------------------------------ | ------------------------------------- |
| `convertirHoras(h float64)`    | Convierte horas a minutos y segundos. |
| `convertirMinutos(m float64)`  | Convierte minutos a horas y segundos. |
| `convertirSegundos(s float64)` | Convierte segundos a horas y minutos. |

### ▶️ Ejecución del programa

El programa muestra un menú interactivo con las opciones:

```
1. Convertir horas a minutos y segundos
2. Convertir minutos a horas y segundos
3. Convertir segundos a horas y minutos
4. Salir
```

Ejemplo de ejecución:

```
=== CONVERSOR DE TIEMPO ===
1. Convertir horas a minutos y segundos
2. Convertir minutos a horas y segundos
3. Convertir segundos a horas y minutos
4. Salir
Selecciona una opción: 1
Introduce las horas: 2
2.00 horas son 120.00 minutos o 7200.00 segundos.
```

Para ejecutar el programa:

```bash
go run .
```

---

## 🧪 Tests unitarios – `conversor_test.go`

Se han creado **3 tests unitarios** usando la librería estándar `testing`.

| Test                             | Descripción                                                         |
| -------------------------------- | ------------------------------------------------------------------- |
| ✅ `TestConvertirHoras`           | Verifica que la conversión de horas a minutos/segundos es correcta. |
| ✅ `TestConvertirMinutos`         | Verifica que la conversión de minutos a horas/segundos es correcta. |
| ❌ `TestConvertirSegundosFallido` | Test preparado para fallar intencionalmente.                        |

### 🧭 Cómo ejecutar los tests

Desde la carpeta del proyecto:

```bash
go test
```

Salida esperada:

```
--- FAIL: TestConvertirSegundosFallido (0.00s)
    conversor_test.go:23: convertirSegundos(3600) horas esperado 2, obtenido 1.00
FAIL
exit status 1
FAIL    conversor_tiempo  0.003s
```

Los tests correctos aparecen como **PASS**, y el test preparado para fallar mostrará **FAIL**.

---

## ✍️ Autor

**Roberto Pecurto Escrig**
Practica 2 - Ejecución de aplicaciones

