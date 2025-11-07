import unittest
import lista_compra

class TestListaCompra(unittest.TestCase):

    def setUp(self):
        lista_compra.lista_compra.clear()

    def test_agregar_producto(self):
        lista_compra.agregar_producto("pan", 1.20)
        self.assertIn("pan", lista_compra.lista_compra)
        self.assertEqual(lista_compra.lista_compra["pan"], 1.20)

    def test_calcular_total(self):
        lista_compra.agregar_producto("leche", 0.95)
        lista_compra.agregar_producto("huevos", 2.10)
        total = lista_compra.calcular_total()
        self.assertEqual(total, 3.05)

    def test_eliminar_producto(self):
        lista_compra.agregar_producto("pan", 1.20)
        resultado = lista_compra.eliminar_producto("pan")
        self.assertTrue(resultado)
        self.assertNotIn("pan", lista_compra.lista_compra)

    def test_fallo(self):
        lista_compra.agregar_producto("azucar", 3.30)
        total = lista_compra.calcular_total()
        self.assertEqual(total, 5)

    if __name__ == "__main__":
        unittest.main()