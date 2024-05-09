package greetings

import (
	"regexp"
	"testing"
)

// testing.T = Es un paquete para testear
// El nombre de la función debe iniciar con la palabra Test seguida del nombre
// Las funciones reciben un parámetro de tipo testing
// El objeto de tipo testing que se pasa como parámetro sirve para reportar el resultado de la prueba
func TestHelloName(t *testing.T) {
	name := "Juan"
	// La siguiente línea va a buscar en la variable name una coincidencia exacta con la expresión regular
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Juan")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Juan") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}

func TestSuma(t *testing.T) {
	valor := 5

	result := Suma(2, 3)

	if valor != result {
		t.Fatalf(`El resultado esperado debe devolver %v pero está devolviendo %v`, valor, result)
	}
}
