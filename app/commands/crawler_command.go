package commands

import (
	"fmt"
	"gocrawler/app/crawler"
	"sync"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "crawler",
	Short: "Print hello message",

	Run: handle,
}

func handle(cmd *cobra.Command, args []string) {
	URLs := []string{
		"https://x.com/",
		"https://dojinja.com/",
		"https://eghtesadkhabar.com/",
		"http://qamarnews.com/",
	}

	maxWorker := 2
	bufferChannel := make(chan struct{}, maxWorker)
	var wg sync.WaitGroup

	for _, u := range URLs {

		bufferChannel <- struct{}{}
		wg.Add(1)

		go func(url string) {

			defer wg.Done()
			defer func() { <-bufferChannel }()

			result, err := crawler.Read(url)
			
			if err != nil {
				fmt.Println("❌ Error:", err)
				return
			}

			fmt.Println("====================================")
			fmt.Println("🌐 URL       :", result.URL)
			fmt.Println("📄 Title     :", result.Title)
			fmt.Println("📝 WordCount :", result.WordCount)
			fmt.Println("🔗 Links    :", result.LinksCount)
			fmt.Println("🛠 Status   :", result.StatusCode)
			fmt.Println("⏱ LoadTime :", result.LoadTime, "ms")
			fmt.Println("====================================\n")
		}(u)
	}

	wg.Wait()

	fmt.Println("All done!")
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
