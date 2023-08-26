/// urfave/cli
/*
urfave/cli, Go için bir komut satırı arayüzü (CLI) çerçevesidir. Bu çerçeve, 
CLI uygulamaları geliştirmeyi kolaylaştırmak ve hızlandırmak için tasarlanmıştır.

urfave/cli, aşağıdaki özellikleri sunar:
  - Basit ve kullanımı kolay bir API
  - Geniş özellik yelpazesi
  - Esnek ve özelleştirilebilir

urfave/cli ile bir CLI uygulaması geliştirmek için aşağıdaki adımları izleyin:
  1) Yeni bir Go projesi oluşturun.
  2) import komutunu kullanarak urfave/cli paketini projenize ekleyin.
  3) `cli.App` yapısını kullanarak bir uygulama oluşturun.
  4) Uygulamanıza komutlar ve bayraklar ekleyin.
  5) Uygulamanızın ana işlevini Run metodunda belirtin.

run istall pkg: `go get github.com/urfave/cli/v2`  
*/

// Flags
/*
Flags, komut satırı arayüzü (CLI) uygulamalarında yaygın olarak kullanılan bir özelliktir. 
Flags, kullanıcıların uygulamanın davranışını komut satırından değiştirmesine olanak tanır.

Flags, aşağıdakiler dahil olmak üzere çeşitli nedenlerle kullanılabilir:
  - Uygulama davranışını özelleştirmek: Flags, kullanıcıların uygulamanın davranışını kendi 
    ihtiyaçlarına göre özelleştirmesine olanak tanır. Örneğin, bir dosya dönüştürücü 
    uygulamasında, kullanıcılar dönüştürülecek dosyaların uzantısını veya dönüştürme 
  
  - Uygulama seçeneklerini etkinleştirmek veya devre dışı bırakmak: Flags, kullanıcıların 
    uygulamanın belirli seçeneklerini etkinleştirmesine veya devre dışı bırakmasına olanak tanır. 
    Örneğin, bir hata ayıklama aracında, kullanıcılar hata ayıklama modunu etkinleştirmek 
    veya devre dışı bırakmak için bir bayrak kullanabilir. formatını belirtmek için bir 
    bayrak kullanabilir.

  - Uygulama hakkında bilgi almak: Flags, kullanıcıların uygulama hakkında bilgi edinmesine 
    olanak tanır. Örneğin, bir web tarayıcısında, kullanıcılar uygulamanın sürümü veya mevcut 
    ayarları hakkında bilgi almak için bir bayrak kullanabilir.  

Flags, CLI uygulamalarının kullanıcı deneyimini iyileştirmenin güçlü bir yoludur. Flags, 
kullanıcıların uygulamayı daha kolay ve daha verimli bir şekilde kullanmalarına yardımcı olabilir.    
*/

// Subcommands 
/*
Daha git benzeri bir komut satırı uygulaması için alt komutlar tanımlanabilir.
*/
