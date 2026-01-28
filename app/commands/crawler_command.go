package commands

import (
	"fmt"
	"gocrawler/app/crawler"
	"strconv"
	"sync"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var crawlerCmd = &cobra.Command{
	Use:   "crawler",
	Short: "Print hello message",
	Run:   handle,
}

func init() {
	rootCmd.AddCommand(crawlerCmd)
	crawlerCmd.Flags().Int("max-worker", 0, "max worker")
}

func handle(cmd *cobra.Command, args []string) {

	var maxWorker int

	URLs := []string{
		"https://x.com/",
		"https://dojinja.com/",
		"https://eghtesadkhabar.com/",
		"http://qamarnews.com/",
	}

	maxWorkerFlag := cmd.Flag("max-worker")

	if maxWorkerFlag != nil && maxWorkerFlag.Changed {
		
		v, err := strconv.Atoi(maxWorkerFlag.Value.String())
		if err != nil {
			fmt.Println("Invalid max-worker flag:", err)
			return
		}

		fmt.Println("Read from flags:", v)

		maxWorker = v

	} else {

		maxWorker = viper.GetInt("crawler.maxWorker")
	}

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
