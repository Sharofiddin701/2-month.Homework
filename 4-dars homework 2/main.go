package main

import (
	"database/sql"
	"fmt"
	"homework/2-oy/4-dars/homework/country"
	"homework/2-oy/4-dars/homework/storage"

	_ "github.com/lib/pq"
)

func main() {
	db, err := connectDB()
	if err != nil {
		fmt.Println("Ma'lumotlar bazasiga ulanishda xato:", err)
		return
	}
	defer db.Close()

	fmt.Println("Tanlang:")
	fmt.Println("1-Yaratish")
	fmt.Println("2-Yangilash")
	fmt.Println("3-O'chirish")
	fmt.Println("4-Hamma ma'lumotlarni olish")

	var userType int

	_, err = fmt.Scan(&userType)
	if err != nil {
		fmt.Println("Foydalanuvchi kiritishini o'qishda xato:", err)
		return
	}

	inv := storage.NewInventory(db)

	switch userType {
	case 1:
		var name string
		var code int
		fmt.Println("Mamlakat nomini kiriting:")
		_, err := fmt.Scan(&name)
		if err != nil {
			fmt.Println("Mamlakat nomini o'qishda xato:", err)
			return
		}
		fmt.Println("Mamlakat kodi kiriting:")
		_, err = fmt.Scan(&code)
		if err != nil {
			fmt.Println("Mamlakat kodini o'qishda xato:", err)
			return
		}
		country := country.Country{
			Name: name,
			Code: code,
		}
		err = inv.Create(country)
		if err != nil {
			fmt.Println("Mamlakat yaratishda xato:", err)
			return
		}
		fmt.Println("Mamlakat muvaffaqiyatli yaratildi")

	case 3:
		countries, err := inv.GetAll()
		if err != nil {
			fmt.Println("Mamlakatlar olishda xato:", err)
			return
		}
		fmt.Println("Mamlakatlar: ", countries)

		var id string
		fmt.Println("O'chiriladigan mamlakatning ID sini kiriting:")
		_, err = fmt.Scan(&id)
		if err != nil {
			fmt.Println("Mamlakat ID sini o'qishda xato:", err)
			return
		}
		err = inv.Delete(country.Country{}, id)
		if err != nil {
			fmt.Println("Mamlakatni o'chirishda xato:", err)
			return
		}
		fmt.Println("Mamlakat muvaffaqiyatli o'chirildi")

	case 2:
		countries, err := inv.GetAll()
		if err != nil {
			fmt.Println("Mamlakatlar olishda xato:", err)
			return
		}
		fmt.Println("Mamlakatlar: ", countries)

		var id, name string
		var code int
		fmt.Println("Yangilanadigan mamlakatning ID sini kiriting:")
		_, err = fmt.Scan(&id)
		if err != nil {
			fmt.Println("Mamlakat ID sini o'qishda xato:", err)
			return
		}
		fmt.Println("Yangilangan mamlakat nomini kiriting:")
		_, err = fmt.Scan(&name)
		if err != nil {
			fmt.Println("Yangilangan mamlakat nomini o'qishda xato:", err)
			return
		}
		fmt.Println("Yangilangan mamlakat kodi kiriting:")
		_, err = fmt.Scan(&code)
		if err != nil {
			fmt.Println("Yangilangan mamlakat kodini o'qishda xato:", err)
			return
		}
		err = inv.Update(country.Country{}, name, code, id)
		if err != nil {
			fmt.Println("Mamlakatni yangilashda xato:", err)
			return
		}
		fmt.Println("Mamlakat muvaffaqiyatli yangilandi")

	case 4:
		countries, err := inv.GetAll()
		if err != nil {
			fmt.Println("Mamlakatlar olishda xato:", err)
			return
		}
		fmt.Println("Mamlakatlar: ", countries)
	}
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=new password=1 database=name sslmode=disable")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
