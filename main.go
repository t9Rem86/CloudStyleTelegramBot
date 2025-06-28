package main

import (
	"awesomeProject/methods"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// Загружаем переменные окружения
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Ошибка загрузки .env файла")
	}

	// Получаем URL базы данных
	databaseURL := os.Getenv("DATABASE_URL")
	fmt.Println("DATABASE_URL:", databaseURL)

	// Подключаемся к базе
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Не удалось подключиться к базе данных: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// Выполняем запрос всех строк
	rows, err := conn.Query(context.Background(), "SELECT id_type, name FROM type")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка выполнения запроса: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	// Перебираем все строки
	for rows.Next() {
		var id int
		var t string
		err := rows.Scan(&id, &t)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка чтения строки: %v\n", err)
			continue
		}
		fmt.Printf("ID: %d, Type: %s\n", id, t)
	}

	// Проверка на ошибки после итерации
	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при итерации по строкам: %v\n", rows.Err())
	}
	idBalance := 1
	var idType int
	var idBrand int
	err = conn.QueryRow(context.Background(), "SELECT id_type FROM type where name = 'Полузамок'").Scan(&idType)
	if err != nil {
		fmt.Fprintf(os.Stderr, "<UNK> <UNK> <UNK> <UNK> <UNK>: %v\n", err)
	}
	err = conn.QueryRow(context.Background(), "SELECT id_brand FROM brand where name = 'Nike'").Scan(&idBrand)
	err = methods.InsertDebit(*conn, idBalance, idType, 0, idBrand, 1600, 5, 300)
}
