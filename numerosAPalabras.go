// Package num2words implements numbers to words converter.
package numerosAPalabras

import "math"
import "strings"

// how many digit's groups to process
const groupsNumber int = 4

var _smallNumbers = []string{
	"cero", "uno", "dos", "tres", "cuatro",
	"cinco", "seis", "siete", "ocho", "nueve",
	"diez", "once", "doce", "trece", "catorce",
	"quince", "dieciséis", "diecisiete", "dieciocho", "diecinueve",
}
var _tens = []string{
	"", "", "veinti", "treinta", "cuarenta", "cincuenta",
	"sesenta", "setenta", "ochenta", "noventa",
}
var _scaleNumbers = []string{
	"", "mil", "millón", "billón",
}

type digitGroup int

// Convert converts number into the words representation.
func Convert(number float64) string {
	return convert(number, false)
}

// ConvertAnd converts number into the words representation
// with " and " added between number groups.
func ConvertAnd(number float64) string {
	return convert(number, true)
}

func convert(number float64, useAnd bool) string {
	entero := int(number)
	if entero == 100 {
		return "Cien"
	}
	if entero == 20000 {
		return "Veinte mil"
	}
	// Zero rule
	if number == 0.0 {
		return _smallNumbers[0]
	}

	// Divide into three-digits group
	var groups [groupsNumber]digitGroup
	positive := math.Abs(float64(number))

	// Form three-digit groups
	for i := 0; i < groupsNumber; i++ {
		groups[i] = digitGroup(math.Mod(positive, 1000))
		positive /= 1000
	}

	var textGroup [groupsNumber]string
	for i := 0; i < groupsNumber; i++ {
		textGroup[i] = digitGroup2Text(groups[i], useAnd)
	}
	combined := textGroup[0]
	and := useAnd && (groups[0] > 0 && groups[0] < 100)

	for i := 1; i < groupsNumber; i++ {
		if groups[i] != 0 {
			poner := textGroup[i]
			if textGroup[i] == "uno" {
				poner = ""
			}
			if _scaleNumbers[i] == "millón" || _scaleNumbers[i] == "millones" {
				if textGroup[i] != "uno" {
					_scaleNumbers[i] = "millones"
				} else {
					_scaleNumbers[i] = "millón"
					poner = "un"
				}
			}
			prefix := poner + " " + _scaleNumbers[i]

			if len(combined) != 0 {
				prefix += separator(and)
			}

			and = false

			combined = prefix + combined
		}
	}

	if number < 0 {
		combined = "menos " + combined
	}
	combined = strings.TrimSpace(combined)
	primeraLetra := combined[0:1]
	primeraLetra = strings.ToUpper(primeraLetra)
	combined = primeraLetra+combined[1:]
				 
	return combined
}

func intMod(x, y int) int {
	return int(math.Mod(float64(x), float64(y)))
}

func digitGroup2Text(group digitGroup, useAnd bool) (ret string) {
	hundreds := group / 100
	tensUnits := intMod(int(group), 100)

	if hundreds != 0 {
		if hundreds==1 {
			ret +=  " ciento"	
		} else {
			ret += _smallNumbers[hundreds] + "cientos"	
		}
		

		if tensUnits != 0 {
			ret += separator(useAnd)
		}
	}

	tens := tensUnits / 10
	units := intMod(tensUnits, 10)

	if tens >= 2 {
		ret += _tens[tens]

		if units != 0 {
			if tens == 2 {
				ret += "" + _smallNumbers[units]	
			} else {
				ret += " y " + _smallNumbers[units]	
			}
		}
	} else if tensUnits != 0 {
		ret += _smallNumbers[tensUnits]
	}

	return
}

// separator returns proper separator string between
// number groups.
func separator(useAnd bool) string {
	if useAnd {
		return " y "
	}
	return " "
}
