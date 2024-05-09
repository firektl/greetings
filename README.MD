# Saludos en Go
Este paquete proporciona una forma simple de obtener saludos personalizados en Go

## Instalaciones 
Ejecuta el siguiente comando para instalar el paquete:
```bash
go get -u github.com/firektl/greetings
```

## Uso
Aquí tienes un ejemplo de cómo utilizar el paquete en tu código:

```go
package greetings

import (
	"fmt"
	"github.com/firektl/greetings"
)

func main(){
    message, err := greetings.Hello("Fiorela")

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(message)
}
```

Este ejemplo importa el paquete github.com/firektl/greetings y llama la función Hello