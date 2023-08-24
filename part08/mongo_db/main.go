package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/// MongoDB
/*
* MongoDB, NoSQL bir veritabanı yönetim sistemidir. NoSQL veritabanları,
* ilişkisel veritabanlarından farklı olarak, verileri tablolar ve sütunlar yerine
* belgeler halinde depolarlar. MongoDB, belgelerin akıcı yapısından dolayı,
* özellikle büyük veri uygulamaları için uygundur.
*
* Go dili, Google tarafından geliştirilen açık kaynaklı bir programlama dilidir.
* Go, basit, açık ve verimli bir dildir. C ve C++ dillerinden esinlenen
* Go, yüksek performanslı uygulamalar oluşturmak için idealdir.
*
* Go ile MongoDB'yi kullanmak için, Go'nun MongoDB sürücüsünü kullanmanız gerekir.
* Go'nun MongoDB sürücüsü, MongoDB'ye erişmenizi ve verileri depolamanızı,
* okumanızı, güncellemenizi ve silmenizi sağlar.
*
* `go get go.mongodb.org/mongo-driver/mongo`
 */

func main() {
	// Bağlantı oluşturma
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	controError(err)

	collection := client.Database("mydb").Collection("mycollection")

	// Veri ekleme
	_, err = collection.InsertOne(context.TODO(), bson.D{{"key", "value"}})
	if err != nil {
		log.Fatal(err)
	}

	// Veri sorgulama
	var result bson.M
	err = collection.FindOne(context.TODO(), bson.D{{"key", "value"}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	// Veri Güncelleme (Update):
	// filter := bson.D{{"key", "value"}}
	// update := bson.D{{"$set", bson.D{{"newField", "newValue"}}}},

	// r,e := collection.UpdateOne(context.TODO(), filter, update)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Veri Silme (Delete):
	// filter := bson.D{{"key", "value"}}

    // result, err := collection.DeleteOne(context.TODO(), filter)
    // if err != nil {
    //     log.Fatal(err)
    // }

}

func controError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
