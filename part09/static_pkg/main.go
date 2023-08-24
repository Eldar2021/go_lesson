package main

/// Statik Kütüphanesi ile Dosyaları Uygulamaya Gömme
/*
* Golang’ın müthiş yanlarından biri de bir uygulamayı build ettiğimizde bize
* tek çalıştırılabilir dosya şeklinde sunmasıdır. Fakat bu özellik sadece
* `.go` dosyalarını birleştirilmesinde kullanılıyor. Harici dosyalar programa
* gömülmüyor. Fakat Statik isimli kütüphane ile bu işlem mümkün kılınıyor.
*
* Kütüphanenin mantığından kısaca bahsedeyim. Belirlediğiniz bir dizindeki
* dosyaları bir kodlamaya çevirerek programın içine dosya gömmek yerine kod gömüyor.
* Ve bu kodu sanki dosyaymışcasına kullanabiliyoruz. Tabi ki sadece sunucu
* işlemlerinde işe yarar olduğunu belirtelim.
*
* Bu yöntemin güzel artı yönleri var.
*   - Programımız tek dosya halinde kalıyor.
*   - Programımız kapalı kaynak oluyor.
*
* go get github.com/rakyll/statik
 */

/*
* Kodlamaya dönüştürülmesini istediğimiz klasör ile işlem yapıyoruz. Yani
* public klasörü ile. Aşağıdaki komutu Proje klasörümüz içerisindeyken yazıyoruz.
*   - statik -src=/{publicpath} -f
 */

//  ?????
