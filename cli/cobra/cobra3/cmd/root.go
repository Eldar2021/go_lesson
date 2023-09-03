package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var name string

// rootCmd, uygulama alt-komutsuz şekilde çağırıldığında
// gerçekleşecekleri temsil eder.
var rootCmd = &cobra.Command{
	Use:   "cobra-3",
	Short: "uygulamanın kısa bilgisi",
	Long:  `uygulamanın uzun bilgilendirme metnini bu kısma girebilirsiniz.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Salam Duyno, Salam " + name + "!")
	},
}

// Execute fonksiyonu tüm komut ve alt-komutları
// uygulama çalıştırıldığında yüklenmesini sağlar
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// init fonksiyonu içerisinde ise komuta bağlı olan
	// alt-komutlarını ve flag'lerini bağlayabilir ve bunlar
	// için özelleştirmelerde bulunabilirsiniz.
	rootCmd.Flags().BoolP("toggle", "t", false, "toggle için yardım mesajı")

	rootCmd.Flags().StringVarP(&name, "name", "n", "Eldiiar", "--name=<Eldiiar> or -n=<Eldiiar>")
}
