package utils_test

import (
	"Go-QuickStart/utils"
	"fmt"
	"testing"
)

func TestNaming(t *testing.T) {
	res := utils.Naming("Helo world", "/", "Kontol")
	fmt.Println(res)

}
