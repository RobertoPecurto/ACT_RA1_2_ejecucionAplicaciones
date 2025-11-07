# diccionario dónde se almacenan los productos y sus precios
lista_compra = {}

# funcione para agregar un producto a la lista de compra
def agregar_producto(nombre, precio):
    lista_compra[nombre] = precio
    return lista_compra

# función para eliminar un producto de la lista de compra
def eliminar_producto(nombre):
    if nombre in lista_compra:
        del lista_compra[nombre]
        return True
    else:
        return False
    
# función para mostrar la lista de compra
def mostrar_lista():
    if lista_compra:
        print("Productos en la lista de compra:")
        for producto, precio in lista_compra.items():
            print("-", producto, precio, "€")
        print("El TOTAL es", calcular_total(), "€")
    else:
        print("La lista de compra está vacía.")

def calcular_total():
    return sum(lista_compra.values())

if __name__ == "__main__":
    print("=== LISTA DE COMPRA ===\n")

    while True:
        print("1. Añadir producto")
        print("2. Ver lista y total")
        print("3. Eliminar producto")
        print("4. Salir")

        opcion = input("Opción: ")

        if opcion == "1":
            nombre = input("Nombre del producto: ")
            try:
                precio = float(input("Precio del producto en €: "))
                agregar_producto(nombre, precio)
                print("Producto agregado con éxito.")
            except ValueError:
                print("Error: el precio debe ser un número.")

        elif opcion == "2":
            mostrar_lista()
        
        elif opcion == "3":
            nombre = input("Nombre del producto a eliminar: ")
            if eliminar_producto(nombre):
                print("Producto eliminado con éxito.")
            else:
                print("Error: no existe el producto.")

        elif opcion == "4":
            print("Saliendo del programa...")
            break

        else:
            print("Opción no válida. Inténtalo de nuevo.")