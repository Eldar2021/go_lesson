package main

/// Declaring Models
/*
GORM, Go programlama dili için popüler bir ORM paketidir.
GORM kullanarak, Go struct'larını veritabanı tablolarına bağlayabilirsiniz.
Bu, veritabanı işlemlerini daha kolay ve verimli hale getirir.

GORM'da bir model tanımlamak için, type anahtar sözcüğünü kullanarak
bir struct oluşturursunuz. Struct'ın her bir alanı, veritabanı
tablosundaki bir sütuna karşılık gelir.

Örneğin, aşağıdaki kod bir User modelini tanımlar:
*/
type User1 struct {
	ID    int `gorm:"primary_key"`
	Name  string
	Email string `gorm:"unique"`
}

/*
Bu modelde, ID alanı veritabanı tablosundaki id sütununa karşılık gelir.
Name alanı name sütununa karşılık gelir ve Email alanı email
sütununa karşılık gelir. Email alanı ayrıca unique anahtar sözcüğünü
kullanır, bu da bu sütunda benzersiz değerlerin olması gerektiğini belirtir.

GORM, modelleri tanımlamak için bazı ek anahtar sözcükler de sağlar.
Bu anahtar sözcükler, sütun türlerini, sütun boyutlarını ve diğer
özellikleri belirtmek için kullanılabilir.

Örneğin, aşağıdaki kod User modelini değiştirir ve Name alanının
max_length değerini 255 olarak ayarlayarak sütun boyutunu belirtir:
*/
type User2 struct {
	ID    int    `gorm:"primary_key"`
	Name  string `gorm:"unique;max_length:255"`
	Email string
}

/*
GORM, modelleri tanımlamak için bir dizi varsayılan değer de kullanır.
Örneğin, GORM varsayılan olarak ID alanını birincil anahtar olarak kullanır.
GORM ayrıca CreatedAt ve UpdatedAt alanlarını otomatik olarak oluşturur ve günceller.

GORM modellerini kullanarak, veritabanı işlemleri için karmaşık SQL
sorguları yazmadan veritabanı tablolarına erişebilir ve bunları
değiştirebilirsiniz. GORM, veritabanı sorgularını otomatik olarak
oluşturur ve yürütür, bu da kodun daha okunaklı ve sürdürülebilir olmasını sağlar.
*/
/*
GORM modellerini kullanmanın bazı faydaları şunlardır:

  - Veritabanı işlemlerini basitleştirir: GORM, karmaşık SQL sorguları
    yazmadan veritabanı tablolarına erişebilir ve bunları değiştirebilirsiniz.

  - Veritabanı sorgularını otomatikleştirir: GORM, veritabanı
    sorgularını otomatik olarak oluşturur ve yürütür, bu da
    kodun daha okunaklı ve sürdürülebilir olmasını sağlar.

  - Veritabanı işlemleri üzerinde daha fazla kontrol sağlar: GORM,
    sütun türlerini, sütun boyutlarını ve diğer özellikleri
    belirtmek için anahtar sözcükler sağlar, bu da size veritabanı
    işlemleri üzerinde daha fazla kontrol sağlar.

GORM, Go programlama dilinde veritabanı işlemleri için kullanışlı
bir araçtır. GORM modellerini kullanarak, veritabanı işlemlerini
daha kolay ve verimli hale getirebilirsiniz.
*/
