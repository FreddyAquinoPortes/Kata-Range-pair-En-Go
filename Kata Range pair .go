package main

import "fmt"

type Range struct {
    start       int
    end         int
    startClosed bool
    endClosed   bool
}

// Crear un rango abierto o semicerrado
func NewRange(start int, startClosed bool, end int, endClosed bool) Range {
    return Range{start: start, end: end, startClosed: startClosed, endClosed: endClosed}
}

// Comprobar si un rango abierto o semicerrado contiene un punto específico
func (r Range) Contains(n int) bool {
    if r.startClosed {
        if r.endClosed {
            return r.start <= n && n <= r.end
        }
        return r.start <= n && n < r.end
    }
    if r.endClosed {
        return r.start < n && n <= r.end
    }
    return r.start < n && n < r.end
}

// Obtener todos los puntos en un rango abierto o semicerrado
func (r Range) AllPoints() []int {
    if r.startClosed && r.endClosed {
        points := make([]int, r.end-r.start+1)
        for i := 0; i <= r.end-r.start; i++ {
            points[i] = r.start + i
        }
        return points
    } else if r.startClosed && !r.endClosed {
        points := make([]int, r.end-r.start)
        for i := 0; i < r.end-r.start; i++ {
            points[i] = r.start + i
}
return points
} else if !r.startClosed && r.endClosed {
points := make([]int, r.end-r.start)
for i := 0; i < r.end-r.start; i++ {
points[i] = r.start + i + 1
}
return points
} else {
points := make([]int, r.end-r.start-1)
for i := 0; i < r.end-r.start-1; i++ {
points[i] = r.start + i + 1
}
return points
}
}

// Comprobar si un rango abierto o semicerrado contiene otro rango
func (r Range) ContainsRange(other Range) bool {
if r.startClosed {
if r.endClosed {
if other.startClosed {
if other.endClosed {
return r.start <= other.start && other.end <= r.end
}
return r.start <= other.start && other.end < r.end
}
if other.endClosed {
return r.start < other.start && other.end <= r.end
}
return r.start < other.start && other.end < r.end
}
if other.startClosed {
if other.endClosed {
return r.start <= other.start && other.end < r.end
}
return r.start <= other.start && other.end <= r.end
}
if other.endClosed {
return r.start <= other.start && other.end <= r.end
}
return r.start <= other.start && other.end < r.end
}
if r.endClosed {
if other.startClosed {
if other.endClosed {
return r.start < other.start && other.end <= r.end
}
return r.start < other.start && other.end < r.end
}
if other.endClosed {
return r.start <= other.start && other.end <= r.end
}
return r.start <= other.start && other.end < r.end
}
if other.startClosed {
if other.endClosed {
return r.start < other.start && other.end < r.end
}
return r.start < other.start && other.end <= r.end
}
if other.endClosed {
return r.start <= other.start && other.end < r.end
}
return r.start <= other.start && other.end <= r.end
}

// Obtener los puntos extremos de un rango abierto o semicerrado
func (r Range) EndPoints() (int, int) {
if r.startClosed {
if r.endClosed {
return r.start, r.end
}
return r.start, r.end - 1
}
if r.endClosed {
return r.start + 1, r.end
}
return r.start + 1, r.end - 1
}

// Comprobar si dos rangos abiertos o semicerrados se superponen
func (r Range) Overlaps(other Range) bool {
if r.startClosed {
if r.endClosed {
if other.startClosed {
if other.endClosed {
return r.start <= other.end && other.start <= r.end
}
return r.start <= other.end && other.start < r.end
}
if other.endClosed {
return r.start < other.end && other.start <= r.end
}
return r.start < other.end && other.start < r.end
}
if other.startClosed {
if other.endClosed {
return r.start <= other.end && other.start < r.end}
return r.start <= other.end && other.start <= r.end
}
if other.endClosed {
return r.start <= other.end && other.start <= r.end
}
return r.start <= other.end && other.start < r.end
}
if r.endClosed {
if other.startClosed {
if other.endClosed {
return r.start < other.end && other.start <= r.end
}
return r.start < other.end && other.start < r.end
}
if other.endClosed {
return r.start <= other.end && other.start <= r.end
}
return r.start <= other.end && other.start < r.end
}
if other.startClosed {
if other.endClosed {
return r.start < other.end && other.start < r.end
}
return r.start < other.end && other.start <= r.end
}
if other.endClosed {
return r.start <= other.end && other.start < r.end
}
return r.start <= other.end && other.start <= r.end
}

// Comprobar si dos rangos abiertos o semicerrados son iguales
func (r Range) Equals(other Range) bool {
    return r.start == other.start && r.end == other.end && r.startClosed == other.startClosed && r.endClosed == other.endClosed
}
func main() {
r1 := NewRange(2, true, 6, false)
r2 := NewRange(3, true, 5, true)
r3 := NewRange(7, true, 10, false)
fmt.Println("Contains:")
fmt.Println(r1.Contains(2)) // true
fmt.Println(r1.Contains(4)) // true
fmt.Println(r1.Contains(6)) // false

fmt.Println("AllPoints:")
fmt.Println(r1.AllPoints()) // [2, 3, 4, 5]

fmt.Println("ContainsRange:")
fmt.Println(r1.ContainsRange(r2)) // true
fmt.Println(r1.ContainsRange(r3)) // false

fmt.Println("EndPoints:")
start, end := r1.EndPoints()
fmt.Println(start, end) // 2 5

fmt.Println("Overlaps:")
fmt.Println(r1.Overlaps(r2)) // true
fmt.Println(r1.Overlaps(r3)) // false

fmt.Println("Equals:")
r4 := NewRange(2, true, 6, false)
fmt.Println(r1.Equals(r4)) // true