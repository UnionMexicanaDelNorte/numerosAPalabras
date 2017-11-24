# numerosAPalabras
Convierte números a palabras en Golang, basado en: https://github.com/divan/num2words

## Uso

Primero, haz el import del package numerosAPalabras

```import github.com/UnionMexicanaDelNorte/numerosAPalabras```

Convertir el número
```go
  str := numerosAPalabras.Convert(17) // outputs "diecisiete"
  
  	strImporte := FloatToString(importe)
	puntos := strings.Split(strImporte, ".")
	centavos := puntos[1]
	str := numerosAPalabras.Convert(importe)+" "+centavos+"/100 M.N."
		
```

Cualquier ayuda será apreciada!
