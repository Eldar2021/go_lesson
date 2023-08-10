package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
)

/// Buffer
/*
* Buffer in Go, Go programlarında veri depolamak ve işlemek için
* kullanılan bir yapıdır. Buffer, bir dizi bayttan oluşur ve bu baytlar,
* okuma ve yazma işlemleri için bir tampon olarak kullanılabilir.
* Buffer, veri üzerinde çalışmak için daha verimli bir yol sunar,
* çünkü verileri doğrudan bellekte depolayabilir ve böylece diske
* erişim gerekmez.
 */

/*
* Buffer, Go programlarında aşağıdaki gibi kullanılabilir:
*
* 1) Verileri okumak ve yazmak için: Buffer, bir dizi bayttan veri
*    okumak ve yazmak için kullanılabilir. Bu, verileri doğrudan bellekten
*    okuyup yazabildiğiniz için diske erişimden daha verimlidir.
*
* 2) Verileri manipüle etmek için: Buffer, veriler üzerinde çeşitli
*    manipülasyonlar gerçekleştirmek için kullanılabilir. Örneğin,
*    verileri sıralamak, filtrelemek veya aralamak için bir buffer
*    kullanabilirsiniz.
*
* 3) Verileri iletmek için: Buffer, verileri bir kanaldan veya bir
*    soket üzerinden iletmek için kullanılabilir. Bu, verileri doğrudan
*    bellekten ilettiğiniz için diske erişimden daha verimlidir.
*
* Buffer, Go programlarında veri işleme için çok kullanışlı bir yapıdır.
* Verileri okumak, yazmak, manipüle etmek ve iletmek için bir buffer
* kullanabilirsiniz. Bu, diske erişimden daha verimli bir yol sunar ve
* programlarınızın daha hızlı çalışmasına yardımcı olur.
 */

// Verileri okumak ve yazmak için:
func main() {
	// ----------------------------------------------------
	// Verileri okumak ve yazmak için:

	// Create a buffer with a capacity of 100 bytes.
	a := bytes.NewBuffer(make([]byte, 100))

	// Write some data to the buffer.
	a.WriteString("Eldi Salam :)")

	// Read the data back from the buffer.
	data01 := a.String()

	// Use data
	fmt.Println(data01) // Eldi Salam :)

	// ----------------------------------------------------
	// Verileri manipüle etmek için:
	b := bytes.NewBufferString("Eldi Salam :)")

	// Reverse the data in the buffer.
	for i, j := 0, len(b.Bytes())-1; i < j; i, j = i+1, j-1 {
		b.Bytes()[i], b.Bytes()[j] = b.Bytes()[j], b.Bytes()[i]
	}

	// Read the data back from the buffer.
	data02 := b.String()

	// Use data
	fmt.Println(data02) // malaS idlE

	// ----------------------------------------------------
	// Verileri iletmek için:
	c := bytes.NewBufferString("Eldi Salam :)")

	// Connect to a TCP server.
	conn, errConnect := net.Dial("tcp", "localhost:8080")

	if errConnect != nil {
		log.Fatal("Error: Connect: ", errConnect)
		return
	}

	_, errWrite := conn.Write(c.Bytes())

	if errWrite != nil {
		log.Fatal("Error: Write: ", errConnect)
		return
	}

	// Close the connection.
	conn.Close()
}
