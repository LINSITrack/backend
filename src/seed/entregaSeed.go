package seed

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/LINSITrack/backend/src/models"
	"gorm.io/gorm"
)

// createSeedFile crea un archivo físico y su registro en la base de datos
func createSeedFile(db *gorm.DB, entregaID int, originalName, contentType, content, seedDir string) error {
	// Verificar si ya existe un archivo con el mismo nombre original para esta entrega
	var existingArchivo models.Archivo
	result := db.Where("artefact_id = ? AND original_name = ?", entregaID, originalName).First(&existingArchivo)

	if result.Error == nil {
		// El archivo ya existe en la BD, verificar si el archivo físico también existe
		if _, err := os.Stat(existingArchivo.FilePath); err == nil {
			// Archivo físico existe, no hacer nada
			log.Printf("Archivo '%s' already exists for entrega ID %d, skipping creation", originalName, entregaID)
			return nil
		} else {
			// Archivo físico no existe, pero registro en BD sí - eliminar registro huérfano
			log.Printf("Found orphaned file record for '%s', removing from database", originalName)
			db.Delete(&existingArchivo)
			// Continuar para crear el archivo nuevamente
		}
	}

	// Generar nombre único para el archivo
	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("%d_%d_%s", entregaID, timestamp, originalName)
	filePath := filepath.Join(seedDir, filename)

	// Crear archivo físico
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating physical file: %v", err)
	}
	defer file.Close()

	// Escribir contenido al archivo
	if _, err := file.WriteString(content); err != nil {
		os.Remove(filePath) // Limpiar archivo si falla la escritura
		return fmt.Errorf("error writing to file: %v", err)
	}

	// Obtener información del archivo
	fileInfo, err := file.Stat()
	if err != nil {
		os.Remove(filePath)
		return fmt.Errorf("error getting file info: %v", err)
	}

	// Crear registro en base de datos
	archivo := models.Archivo{
		ArtefactID:   entregaID,
		Filename:     filename,
		OriginalName: originalName,
		FilePath:     filePath,
		ContentType:  contentType,
		Size:         fileInfo.Size(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := db.Create(&archivo).Error; err != nil {
		os.Remove(filePath) // Limpiar archivo físico si falla la BD
		return fmt.Errorf("error saving to database: %v", err)
	}

	return nil
}

func EntregaSeed(db *gorm.DB) {
	// Recupera alumnos existentes
	var alumnos []models.Alumno
	if err := db.Find(&alumnos).Error; err != nil {
		log.Printf("Failed to get alumnos for entrega seed: %v", err)
		return
	}

	// Recupera TPs existentes
	var tps []models.TpModel
	if err := db.Find(&tps).Error; err != nil {
		log.Printf("Failed to get TPs for entrega seed: %v", err)
		return
	}

	if len(alumnos) == 0 || len(tps) == 0 {
		log.Println("No alumnos or TPs found. Run alumno and tp seeds first.")
		return
	}

	// Crear directorio para archivos de seed si no existe
	seedDir := "uploads/entregas"
	if err := os.MkdirAll(seedDir, 0755); err != nil {
		log.Printf("Failed to create seed directory: %v", err)
		return
	}

	entregas := []struct {
		FechaHora   string
		AlumnoIndex int
		TpIndex     int
		Archivos    []struct {
			OriginalName string
			ContentType  string
			Content      string
		}
	}{
		// Entregas para TP de Algoritmos - Comisión S11 (TP ID: 1)
		{
			FechaHora:   time.Date(2024, 12, 14, 18, 30, 0, 0, time.UTC).Format("2006-01-02 15:04:05"),
			AlumnoIndex: 0, // Juan Pérez
			TpIndex:     0, // TP Algoritmos de ordenamiento
			Archivos: []struct {
				OriginalName string
				ContentType  string
				Content      string
			}{
				{
					OriginalName: "algoritmos_ordenamiento.py",
					ContentType:  "text/x-python",
					Content: `# TP Algoritmos de Ordenamiento - Juan Pérez
def quicksort(arr):
    if len(arr) <= 1:
        return arr
    pivot = arr[len(arr) // 2]
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]
    return quicksort(left) + middle + quicksort(right)

def mergesort(arr):
    if len(arr) <= 1:
        return arr
    mid = len(arr) // 2
    left = mergesort(arr[:mid])
    right = mergesort(arr[mid:])
    return merge(left, right)

def merge(left, right):
    result = []
    i = j = 0
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1
    result.extend(left[i:])
    result.extend(right[j:])
    return result

# Pruebas de rendimiento
import time
import random

def test_algorithms():
    sizes = [1000, 5000, 10000]
    for size in sizes:
        arr = [random.randint(1, 1000) for _ in range(size)]
        
        # Test QuickSort
        start = time.time()
        quicksort(arr.copy())
        quick_time = time.time() - start
        
        # Test MergeSort
        start = time.time()
        mergesort(arr.copy())
        merge_time = time.time() - start
        
        print(f"Size {size}: QuickSort={quick_time:.4f}s, MergeSort={merge_time:.4f}s")

if __name__ == "__main__":
    test_algorithms()
`,
				},
				{
					OriginalName: "informe_rendimiento.pdf",
					ContentType:  "application/pdf",
					Content:      "PDF_CONTENT_PLACEHOLDER", // En un caso real, aquí iría contenido binario
				},
			},
		},
		{
			FechaHora:   time.Date(2024, 12, 15, 22, 45, 0, 0, time.UTC).Format("2006-01-02 15:04:05"),
			AlumnoIndex: 1, // María García
			TpIndex:     0, // TP Algoritmos de ordenamiento
			Archivos: []struct {
				OriginalName string
				ContentType  string
				Content      string
			}{
				{
					OriginalName: "ordenamiento_garcia.cpp",
					ContentType:  "text/x-c++src",
					Content: `// TP Algoritmos de Ordenamiento - María García
#include <iostream>
#include <vector>
#include <chrono>
#include <algorithm>
#include <random>

using namespace std;
using namespace std::chrono;

class SortingAlgorithms {
public:
    static void quickSort(vector<int>& arr, int low, int high) {
        if (low < high) {
            int pi = partition(arr, low, high);
            quickSort(arr, low, pi - 1);
            quickSort(arr, pi + 1, high);
        }
    }
    
    static void heapSort(vector<int>& arr) {
        int n = arr.size();
        for (int i = n / 2 - 1; i >= 0; i--)
            heapify(arr, n, i);
        
        for (int i = n - 1; i >= 0; i--) {
            swap(arr[0], arr[i]);
            heapify(arr, i, 0);
        }
    }

private:
    static int partition(vector<int>& arr, int low, int high) {
        int pivot = arr[high];
        int i = (low - 1);
        
        for (int j = low; j <= high - 1; j++) {
            if (arr[j] < pivot) {
                i++;
                swap(arr[i], arr[j]);
            }
        }
        swap(arr[i + 1], arr[high]);
        return (i + 1);
    }
    
    static void heapify(vector<int>& arr, int n, int i) {
        int largest = i;
        int l = 2 * i + 1;
        int r = 2 * i + 2;
        
        if (l < n && arr[l] > arr[largest])
            largest = l;
        
        if (r < n && arr[r] > arr[largest])
            largest = r;
        
        if (largest != i) {
            swap(arr[i], arr[largest]);
            heapify(arr, n, largest);
        }
    }
};

int main() {
    vector<int> sizes = {1000, 5000, 10000, 50000};
    
    for (int size : sizes) {
        vector<int> data(size);
        random_device rd;
        mt19937 gen(rd());
        uniform_int_distribution<> dis(1, 1000);
        
        for (int i = 0; i < size; i++) {
            data[i] = dis(gen);
        }
        
        // Test QuickSort
        vector<int> quickData = data;
        auto start = high_resolution_clock::now();
        SortingAlgorithms::quickSort(quickData, 0, size - 1);
        auto end = high_resolution_clock::now();
        auto quickTime = duration_cast<microseconds>(end - start);
        
        // Test HeapSort
        vector<int> heapData = data;
        start = high_resolution_clock::now();
        SortingAlgorithms::heapSort(heapData);
        end = high_resolution_clock::now();
        auto heapTime = duration_cast<microseconds>(end - start);
        
        cout << "Size " << size << ": QuickSort=" << quickTime.count() 
             << "μs, HeapSort=" << heapTime.count() << "μs" << endl;
    }
    
    return 0;
}
`,
				},
			},
		},

		// Entregas para TP de Estructuras de datos - Comisión S12 (TP ID: 2)
		{
			FechaHora:   time.Date(2024, 12, 19, 20, 15, 0, 0, time.UTC).Format("2006-01-02 15:04:05"),
			AlumnoIndex: 2, // Carlos López
			TpIndex:     1, // TP Árbol binario de búsqueda
			Archivos: []struct {
				OriginalName string
				ContentType  string
				Content      string
			}{
				{
					OriginalName: "arbol_binario_busqueda.java",
					ContentType:  "text/x-java-source",
					Content: `// TP Árbol Binario de Búsqueda - Carlos López
class NodoArbol {
    int dato;
    NodoArbol izquierdo;
    NodoArbol derecho;
    
    public NodoArbol(int dato) {
        this.dato = dato;
        this.izquierdo = null;
        this.derecho = null;
    }
}

public class ArbolBinarioBusqueda {
    private NodoArbol raiz;
    
    public ArbolBinarioBusqueda() {
        this.raiz = null;
    }
    
    public void insertar(int dato) {
        raiz = insertarRecursivo(raiz, dato);
    }
    
    private NodoArbol insertarRecursivo(NodoArbol nodo, int dato) {
        if (nodo == null) {
            return new NodoArbol(dato);
        }
        
        if (dato < nodo.dato) {
            nodo.izquierdo = insertarRecursivo(nodo.izquierdo, dato);
        } else if (dato > nodo.dato) {
            nodo.derecho = insertarRecursivo(nodo.derecho, dato);
        }
        
        return nodo;
    }
    
    public boolean buscar(int dato) {
        return buscarRecursivo(raiz, dato);
    }
    
    private boolean buscarRecursivo(NodoArbol nodo, int dato) {
        if (nodo == null) {
            return false;
        }
        
        if (dato == nodo.dato) {
            return true;
        }
        
        if (dato < nodo.dato) {
            return buscarRecursivo(nodo.izquierdo, dato);
        } else {
            return buscarRecursivo(nodo.derecho, dato);
        }
    }
    
    public void eliminar(int dato) {
        raiz = eliminarRecursivo(raiz, dato);
    }
    
    private NodoArbol eliminarRecursivo(NodoArbol nodo, int dato) {
        if (nodo == null) {
            return null;
        }
        
        if (dato < nodo.dato) {
            nodo.izquierdo = eliminarRecursivo(nodo.izquierdo, dato);
        } else if (dato > nodo.dato) {
            nodo.derecho = eliminarRecursivo(nodo.derecho, dato);
        } else {
            // Nodo a eliminar encontrado
            if (nodo.izquierdo == null) {
                return nodo.derecho;
            } else if (nodo.derecho == null) {
                return nodo.izquierdo;
            }
            
            nodo.dato = encontrarMinimo(nodo.derecho);
            nodo.derecho = eliminarRecursivo(nodo.derecho, nodo.dato);
        }
        
        return nodo;
    }
    
    private int encontrarMinimo(NodoArbol nodo) {
        int minimo = nodo.dato;
        while (nodo.izquierdo != null) {
            minimo = nodo.izquierdo.dato;
            nodo = nodo.izquierdo;
        }
        return minimo;
    }
    
    public void imprimirEnOrden() {
        imprimirEnOrdenRecursivo(raiz);
        System.out.println();
    }
    
    private void imprimirEnOrdenRecursivo(NodoArbol nodo) {
        if (nodo != null) {
            imprimirEnOrdenRecursivo(nodo.izquierdo);
            System.out.print(nodo.dato + " ");
            imprimirEnOrdenRecursivo(nodo.derecho);
        }
    }
    
    public static void main(String[] args) {
        ArbolBinarioBusqueda arbol = new ArbolBinarioBusqueda();
        
        // Insertar elementos
        int[] valores = {50, 30, 70, 20, 40, 60, 80};
        for (int valor : valores) {
            arbol.insertar(valor);
        }
        
        System.out.println("Árbol en orden:");
        arbol.imprimirEnOrden();
        
        // Buscar elementos
        System.out.println("Buscar 40: " + arbol.buscar(40));
        System.out.println("Buscar 25: " + arbol.buscar(25));
        
        // Eliminar elemento
        arbol.eliminar(30);
        System.out.println("Después de eliminar 30:");
        arbol.imprimirEnOrden();
    }
}
`,
				},
			},
		},
		{
			FechaHora:   time.Date(2024, 12, 20, 23, 30, 0, 0, time.UTC).Format("2006-01-02 15:04:05"),
			AlumnoIndex: 3, // Ana Martínez
			TpIndex:     1, // TP Árbol binario de búsqueda
			Archivos: []struct {
				OriginalName string
				ContentType  string
				Content      string
			}{
				{
					OriginalName: "bst_martinez.py",
					ContentType:  "text/x-python",
					Content: `# TP Árbol Binario de Búsqueda - Ana Martínez
class Nodo:
    def __init__(self, dato):
        self.dato = dato
        self.izquierdo = None
        self.derecho = None

class ArbolBinarioBusqueda:
    def __init__(self):
        self.raiz = None
    
    def insertar(self, dato):
        if self.raiz is None:
            self.raiz = Nodo(dato)
        else:
            self._insertar_recursivo(self.raiz, dato)
    
    def _insertar_recursivo(self, nodo, dato):
        if dato < nodo.dato:
            if nodo.izquierdo is None:
                nodo.izquierdo = Nodo(dato)
            else:
                self._insertar_recursivo(nodo.izquierdo, dato)
        elif dato > nodo.dato:
            if nodo.derecho is None:
                nodo.derecho = Nodo(dato)
            else:
                self._insertar_recursivo(nodo.derecho, dato)
    
    def buscar(self, dato):
        return self._buscar_recursivo(self.raiz, dato)
    
    def _buscar_recursivo(self, nodo, dato):
        if nodo is None or nodo.dato == dato:
            return nodo is not None
        
        if dato < nodo.dato:
            return self._buscar_recursivo(nodo.izquierdo, dato)
        else:
            return self._buscar_recursivo(nodo.derecho, dato)
    
    def eliminar(self, dato):
        self.raiz = self._eliminar_recursivo(self.raiz, dato)
    
    def _eliminar_recursivo(self, nodo, dato):
        if nodo is None:
            return nodo
        
        if dato < nodo.dato:
            nodo.izquierdo = self._eliminar_recursivo(nodo.izquierdo, dato)
        elif dato > nodo.dato:
            nodo.derecho = self._eliminar_recursivo(nodo.derecho, dato)
        else:
            # Nodo a eliminar encontrado
            if nodo.izquierdo is None:
                return nodo.derecho
            elif nodo.derecho is None:
                return nodo.izquierdo
            
            # Nodo con dos hijos
            nodo.dato = self._encontrar_minimo(nodo.derecho)
            nodo.derecho = self._eliminar_recursivo(nodo.derecho, nodo.dato)
        
        return nodo
    
    def _encontrar_minimo(self, nodo):
        while nodo.izquierdo is not None:
            nodo = nodo.izquierdo
        return nodo.dato
    
    def recorrido_inorden(self):
        resultado = []
        self._inorden_recursivo(self.raiz, resultado)
        return resultado
    
    def _inorden_recursivo(self, nodo, resultado):
        if nodo:
            self._inorden_recursivo(nodo.izquierdo, resultado)
            resultado.append(nodo.dato)
            self._inorden_recursivo(nodo.derecho, resultado)
    
    def altura(self):
        return self._altura_recursiva(self.raiz)
    
    def _altura_recursiva(self, nodo):
        if nodo is None:
            return 0
        return 1 + max(self._altura_recursiva(nodo.izquierdo), 
                      self._altura_recursiva(nodo.derecho))

if __name__ == "__main__":
    # Pruebas del árbol
    arbol = ArbolBinarioBusqueda()
    
    # Insertar elementos
    elementos = [50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45]
    for elemento in elementos:
        arbol.insertar(elemento)
    
    print("Árbol en orden:", arbol.recorrido_inorden())
    print("Altura del árbol:", arbol.altura())
    
    # Buscar elementos
    print("Buscar 40:", arbol.buscar(40))
    print("Buscar 100:", arbol.buscar(100))
    
    # Eliminar elemento
    arbol.eliminar(30)
    print("Después de eliminar 30:", arbol.recorrido_inorden())
`,
				},
				{
					OriginalName: "documentacion.md",
					ContentType:  "text/markdown",
					Content: `# Documentación - Árbol Binario de Búsqueda

## Autor: Ana Martínez

### Descripción
Implementación de un Árbol Binario de Búsqueda (BST) en Python con las operaciones básicas:
- Inserción
- Búsqueda  
- Eliminación
- Recorridos

### Complejidad Temporal
- **Búsqueda**: O(log n) promedio, O(n) peor caso
- **Inserción**: O(log n) promedio, O(n) peor caso  
- **Eliminación**: O(log n) promedio, O(n) peor caso

### Casos de Prueba
- Inserción de 11 elementos
- Verificación del orden tras recorrido inorden
- Cálculo de altura del árbol
- Búsqueda de elementos existentes y no existentes
- Eliminación de nodo con dos hijos

### Conclusiones
El BST implementado mantiene correctamente la propiedad de ordenamiento y maneja eficientemente los casos edge de eliminación.
`,
				},
			},
		},

		// Entregas para TP de Bases de Datos - Comisión S31 (TP ID: 4)
		{
			FechaHora:   time.Date(2024, 12, 9, 14, 10, 0, 0, time.UTC).Format("2006-01-02 15:04:05"),
			AlumnoIndex: 0, // Juan Pérez (entrega otro TP)
			TpIndex:     3, // TP Sistema bibliotecario
			Archivos: []struct {
				OriginalName string
				ContentType  string
				Content      string
			}{
				{
					OriginalName: "biblioteca_schema.sql",
					ContentType:  "application/sql",
					Content: `-- TP Sistema Bibliotecario - Juan Pérez
-- Base de datos para gestión bibliotecaria

CREATE DATABASE biblioteca_db;
USE biblioteca_db;

-- Tabla de autores
CREATE TABLE autores (
    id_autor INT PRIMARY KEY AUTO_INCREMENT,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    nacionalidad VARCHAR(50),
    fecha_nacimiento DATE
);

-- Tabla de géneros
CREATE TABLE generos (
    id_genero INT PRIMARY KEY AUTO_INCREMENT,
    nombre_genero VARCHAR(50) NOT NULL
);

-- Tabla de libros
CREATE TABLE libros (
    id_libro INT PRIMARY KEY AUTO_INCREMENT,
    isbn VARCHAR(13) UNIQUE NOT NULL,
    titulo VARCHAR(200) NOT NULL,
    id_autor INT,
    id_genero INT,
    fecha_publicacion DATE,
    editorial VARCHAR(100),
    paginas INT,
    disponibles INT DEFAULT 1,
    total_copias INT DEFAULT 1,
    FOREIGN KEY (id_autor) REFERENCES autores(id_autor),
    FOREIGN KEY (id_genero) REFERENCES generos(id_genero)
);

-- Tabla de usuarios
CREATE TABLE usuarios (
    id_usuario INT PRIMARY KEY AUTO_INCREMENT,
    numero_carnet VARCHAR(20) UNIQUE NOT NULL,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    email VARCHAR(150),
    telefono VARCHAR(20),
    fecha_registro DATE DEFAULT CURRENT_DATE,
    activo BOOLEAN DEFAULT TRUE
);

-- Tabla de préstamos
CREATE TABLE prestamos (
    id_prestamo INT PRIMARY KEY AUTO_INCREMENT,
    id_libro INT,
    id_usuario INT,
    fecha_prestamo DATE DEFAULT CURRENT_DATE,
    fecha_devolucion_esperada DATE,
    fecha_devolucion_real DATE NULL,
    multa DECIMAL(10,2) DEFAULT 0.00,
    estado ENUM('ACTIVO', 'DEVUELTO', 'VENCIDO') DEFAULT 'ACTIVO',
    FOREIGN KEY (id_libro) REFERENCES libros(id_libro),
    FOREIGN KEY (id_usuario) REFERENCES usuarios(id_usuario)
);

-- Insertar datos de prueba
INSERT INTO autores (nombre, apellido, nacionalidad, fecha_nacimiento) VALUES
('Gabriel', 'García Márquez', 'Colombiana', '1927-03-06'),
('Mario', 'Vargas Llosa', 'Peruana', '1936-03-28'),
('Isabel', 'Allende', 'Chilena', '1942-08-02');

INSERT INTO generos (nombre_genero) VALUES
('Realismo Mágico'),
('Novela'),
('Ensayo'),
('Biografía');

INSERT INTO libros (isbn, titulo, id_autor, id_genero, fecha_publicacion, editorial, paginas, total_copias, disponibles) VALUES
('9780060883287', 'Cien años de soledad', 1, 1, '1967-06-05', 'Sudamericana', 471, 3, 2),
('9788420471839', 'La ciudad y los perros', 2, 2, '1963-01-01', 'Seix Barral', 413, 2, 1),
('9788401242298', 'La casa de los espíritus', 3, 1, '1982-01-01', 'Plaza & Janés', 462, 2, 2);

-- Consultas complejas requeridas
-- 1. Libros más prestados
SELECT l.titulo, COUNT(p.id_prestamo) as total_prestamos
FROM libros l
LEFT JOIN prestamos p ON l.id_libro = p.id_libro
GROUP BY l.id_libro, l.titulo
ORDER BY total_prestamos DESC;

-- 2. Usuarios con multas pendientes
SELECT u.nombre, u.apellido, SUM(p.multa) as total_multa
FROM usuarios u
INNER JOIN prestamos p ON u.id_usuario = p.id_usuario
WHERE p.multa > 0 AND p.estado != 'DEVUELTO'
GROUP BY u.id_usuario
HAVING total_multa > 0;

-- 3. Libros disponibles por género
SELECT g.nombre_genero, COUNT(l.id_libro) as libros_disponibles
FROM generos g
LEFT JOIN libros l ON g.id_genero = l.id_genero AND l.disponibles > 0
GROUP BY g.id_genero, g.nombre_genero;
`,
				},
				{
					OriginalName: "analisis_normalizacion.pdf",
					ContentType:  "application/pdf",
					Content:      "PDF_CONTENT_PLACEHOLDER",
				},
			},
		},

		// Entregas restantes sin archivos para no hacer el seed demasiado largo
		{
			FechaHora:   time.Date(2024, 12, 10, 19, 45, 0, 0, time.UTC).Format("2006-01-02 15:04:05"),
			AlumnoIndex: 1, // María García
			TpIndex:     3, // TP Sistema bibliotecario
		},
		{
			FechaHora:   time.Date(2024, 12, 11, 17, 25, 0, 0, time.UTC).Format("2006-01-02 15:04:05"),
			AlumnoIndex: 2, // Carlos López
			TpIndex:     4, // TP Procedimientos almacenados
		},
		{
			FechaHora:   time.Date(2025, 1, 14, 12, 0, 0, 0, time.UTC).Format("2006-01-02 15:04:05"),
			AlumnoIndex: 3, // Ana Martínez
			TpIndex:     5, // TP Aplicación web MVC
		},
		{
			FechaHora:   time.Date(2025, 1, 15, 20, 30, 0, 0, time.UTC).Format("2006-01-02 15:04:05"),
			AlumnoIndex: 4, // Luis Rodríguez
			TpIndex:     5, // TP Aplicación web MVC
		},
	}

	for _, entregaData := range entregas {
		// Validar índices
		if entregaData.AlumnoIndex >= len(alumnos) {
			log.Printf("Alumno index %d out of range, skipping entrega", entregaData.AlumnoIndex)
			continue
		}
		if entregaData.TpIndex >= len(tps) {
			log.Printf("TP index %d out of range, skipping entrega", entregaData.TpIndex)
			continue
		}

		alumnoID := alumnos[entregaData.AlumnoIndex].ID
		tpID := tps[entregaData.TpIndex].ID

		// Verificar si ya existe una entrega para este alumno y TP
		var existingEntrega models.Entrega
		result := db.Where("alumno_id = ? AND tp_id = ?", alumnoID, tpID).First(&existingEntrega)

		if result.Error == nil {
			// La entrega ya existe, verificar si tiene archivos
			entregaExistenteID := existingEntrega.ID

			// Verificar si ya existen archivos para esta entrega
			var archivoCount int64
			db.Model(&models.Archivo{}).Where("artefact_id = ?", entregaExistenteID).Count(&archivoCount)

			if archivoCount == 0 && len(entregaData.Archivos) > 0 {
				// Verificar si ya existen archivos físicos para esta entrega antes de crearlos
				existingPhysicalFiles := false
				for _, archivoData := range entregaData.Archivos {
					// Buscar archivos que empiecen con el patrón entregaID_*_originalName
					pattern := fmt.Sprintf("%d_*_%s", entregaExistenteID, archivoData.OriginalName)
					matches, _ := filepath.Glob(filepath.Join(seedDir, pattern))
					if len(matches) > 0 {
						existingPhysicalFiles = true
						break
					}
				}

				if existingPhysicalFiles {
					log.Printf("Entrega ID %d has existing physical files in directory, skipping file creation", entregaExistenteID)
				} else {
					log.Printf("Entrega ID %d exists but has no files, creating seed files...", entregaExistenteID)
					// Crear archivos para la entrega existente
					for _, archivoData := range entregaData.Archivos {
						if err := createSeedFile(db, entregaExistenteID, archivoData.OriginalName, archivoData.ContentType, archivoData.Content, seedDir); err != nil {
							log.Printf("Failed to create archivo '%s' for existing entrega ID %d: %v", archivoData.OriginalName, entregaExistenteID, err)
						} else {
							log.Printf("Archivo '%s' created for existing entrega ID %d", archivoData.OriginalName, entregaExistenteID)
						}
					}
				}
			} else if archivoCount > 0 {
				log.Printf("Entrega for alumno ID %d, TP ID %d already exists with %d files, skipping", alumnoID, tpID, archivoCount)
			} else {
				log.Printf("Entrega for alumno ID %d, TP ID %d already exists (no files defined in seed)", alumnoID, tpID)
			}
		} else {
			newEntrega := models.Entrega{
				FechaHora: entregaData.FechaHora,
				AlumnoID:  alumnoID,
				TpID:      tpID,
			}

			if err := db.Create(&newEntrega).Error; err != nil {
				log.Printf("Failed to create entrega for alumno ID %d, TP ID %d: %v", alumnoID, tpID, err)
			} else {
				log.Printf("Entrega for alumno ID %d, TP ID %d created successfully (Entrega ID: %d)", alumnoID, tpID, newEntrega.ID)

				// Verificar si ya existen archivos físicos antes de crearlos
				hasExistingFiles := false
				for _, archivoData := range entregaData.Archivos {
					// Buscar archivos que empiecen con el patrón entregaID_*_originalName
					pattern := fmt.Sprintf("%d_*_%s", newEntrega.ID, archivoData.OriginalName)
					matches, _ := filepath.Glob(filepath.Join(seedDir, pattern))
					if len(matches) > 0 {
						hasExistingFiles = true
						break
					}
				}

				if hasExistingFiles {
					log.Printf("Physical files already exist for new entrega ID %d, skipping file creation", newEntrega.ID)
				} else {
					// Crear archivos asociados si existen
					for _, archivoData := range entregaData.Archivos {
						if err := createSeedFile(db, newEntrega.ID, archivoData.OriginalName, archivoData.ContentType, archivoData.Content, seedDir); err != nil {
							log.Printf("Failed to create archivo '%s' for entrega ID %d: %v", archivoData.OriginalName, newEntrega.ID, err)
						} else {
							log.Printf("Archivo '%s' created for entrega ID %d", archivoData.OriginalName, newEntrega.ID)
						}
					}
				}
			}
		}
	}

	// Log final
	log.Println("Entrega seed completed successfully")
}
