package domain

import (
	"fmt"
	"echofy_backend/src/core/errors"
)

func ShowInvalidFields(invalidFields []errors.InvalidField) {
	size := len(invalidFields)

	if size > 1 {
		fmt.Println("Houveram ", size, "campos inválidos, verifique abaixo quais foram")
	} else {
		fmt.Println("Houve apenas um campo inválido verifique abaixo qual foi")
	}

	for _, each := range invalidFields {
		fmt.Println(each.Name)
		fmt.Println(each.Description)
	}
}
