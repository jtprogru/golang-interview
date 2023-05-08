package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	err := runTests("./internal/...")
	if err != nil {
		fmt.Println("Тестирование завершилось с ошибками:", err)
		os.Exit(1)
	}

	fmt.Println("Тестирование успешно завершено")
	// Здесь ваш код для основной логики программы
	fmt.Println("Для более подробного вывода запускать:")
	fmt.Println("\tmake testalltasks")
}

func runTests(pattern string) error {
	cmd := exec.Command("go", "test", pattern)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("ошибка при запуске тестов: %w", err)
	}
	return nil
}
