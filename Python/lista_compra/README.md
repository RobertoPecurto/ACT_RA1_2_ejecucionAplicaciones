# 🛒 **Lista de la Compra en Python**

Este proyecto consiste en un programa sencillo que permite gestionar una **lista de la compra** con operaciones básicas (añadir, eliminar, mostrar y calcular el total).

---

## ⚙️ Instalación y configuración del entorno

### 🐍 Paso 1. Instalar Python

Descarga la versión más reciente de **Python** desde:
👉 [https://www.python.org/downloads/](https://www.python.org/downloads/)

Durante la instalación, asegúrate de marcar la opción **“Add Python to PATH”**.

Verifica la instalación con:

```bash
python --version
```

---

### 📁 Paso 2. Estructura del proyecto

Crea una carpeta de trabajo con los siguientes archivos:

```
📂 lista_compra/
 ├── lista_compra.py
 └── test_lista_compra.py
```

---

## 🧩 Código principal – `lista_compra.py`

El programa permite gestionar una lista de productos con sus precios.

### ✨ Funciones principales

| Función                            | Descripción                                   |
| ---------------------------------- | --------------------------------------------- |
| `agregar_producto(nombre, precio)` | Añade un producto y su precio al diccionario. |
| `eliminar_producto(nombre)`        | Elimina un producto si existe.                |
| `mostrar_lista()`                  | Muestra la lista con precios y el total.      |
| `calcular_total()`                 | Devuelve la suma total de todos los precios.  |

### ▶️ Ejecución del programa

El programa muestra un menú interactivo con las opciones:

```
1. Añadir producto
2. Ver lista y total
3. Eliminar producto
4. Salir
```

Ejemplo de ejecución:

```
=== LISTA DE COMPRA ===

1. Añadir producto
2. Ver lista y total
3. Eliminar producto
4. Salir
Opción: 1
Nombre del producto: pan
Precio del producto en €: 1.20
Producto agregado con éxito.

Opción: 2
Productos en la lista de compra:
- pan 1.2 €
El TOTAL es 1.2 €
```

Para ejecutar el programa:

```bash
python lista_compra.py
```

---

## 🧪 Tests unitarios – `test_lista_compra.py`

Se han creado **4 tests unitarios** usando la librería estándar `unittest`.

| Test                       | Descripción                                          |
| -------------------------- | ---------------------------------------------------- |
| ✅ `test_agregar_producto`  | Verifica que se añade correctamente un producto.     |
| ✅ `test_calcular_total`    | Comprueba que la suma de los precios es la esperada. |
| ✅ `test_eliminar_producto` | Verifica que se elimina un producto existente.       |
| ❌ `test_fallo`             | Test preparado para fallar intencionalmente.         |

### 🧭 Cómo ejecutar los tests

Desde la carpeta del proyecto:

```bash
python -m unittest test_lista_compra.py
```

Salida esperada:

```
...F
======================================================================
FAIL: test_fallo (test_lista_compra.TestListaCompra.test_fallo)
AssertionError: 3.3 != 5
----------------------------------------------------------------------
Ran 4 tests in 0.002s

FAILED (failures=1)
```

Los puntos (`...`) indican tests correctos.
La **F** corresponde al test **fallido intencionalmente** (para demostrar detección de errores).

---

## ✍️ Autor

**Roberto Pecurto Escrig**
Practica 2 - Ejecución de aplicaciones

