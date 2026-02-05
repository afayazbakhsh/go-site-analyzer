package commands

import (
	"context"
	"fmt"
	"gocrawler/app/crawler"
	"strconv"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var crawlerCmd = &cobra.Command{
	Use:   "crawler",
	Short: "Crawl the urls",
	Run:   handle,
}

var URLs = []string{
	"https://x.com/",
	"https://dojinja.com/",
	"https://eghtesadkhabar.com/",
	"https://qamarnews.com/",
}

func init() {
	rootCmd.AddCommand(crawlerCmd)

	// Flags
	crawlerCmd.Flags().Int("max-worker", 0, "max worker")
}

func handle(cmd *cobra.Command, args []string) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	bufferChannel := make(chan struct{}, setMaxWorker(cmd))
	resultChan := make(chan *crawler.ReadPage, 50)

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("⚠️ Context canceled before start db")
				return

			case id, ok := <-resultChan:
				if !ok {
					fmt.Println("✅ DB writer finished:", id)
					return
				}

				if err, pageData := crawler.Write(id); err != nil {
					fmt.Println("Failed to write:", err)
					return
				} else {
					fmt.Println("Success to write:", pageData)
				}

			case <-time.After(5 * time.Second):
				fmt.Println("waiting for results..")
			}
		}
	}()

	for _, u := range URLs {

		bufferChannel <- struct{}{}
		wg.Add(1)

		go func(ctx context.Context, url string) {

			defer wg.Done()
			defer func() { <-bufferChannel }()

			select {
			case <-ctx.Done():
				fmt.Println("⚠️ Context canceled before start:", url)
				return
			case <-time.After(5 * time.Second):
				fmt.Println("Timeout sending result: ", url)
			default:

				result, err := crawler.Read(ctx, url)
				resultChan <- result
				if err != nil {
					fmt.Println("❌ Error:", err)
					return
				}

				printResult(result)
			}
		}(ctx, u)
	}

	wg.Wait()

	fmt.Println("All done!")
}

func printResult(result *crawler.ReadPage) {
	fmt.Println("====================================")
	fmt.Println("🌐 URL       :", result.URL)
	fmt.Println("📄 Title     :", result.Title)
	fmt.Println("📝 WordCount :", result.WordCount)
	fmt.Println("🔗 Links    :", result.LinksCount)
	fmt.Println("🛠 Status   :", result.StatusCode)
	fmt.Println("⏱ LoadTime :", result.LoadTime, "ms")
	fmt.Println("====================================\n")
}

func setMaxWorker(cmd *cobra.Command) int {
	maxWorkerFlag := cmd.Flag("max-worker")

	if maxWorkerFlag != nil && maxWorkerFlag.Changed {
		v, err := strconv.Atoi(maxWorkerFlag.Value.String())
		if err != nil {
			fmt.Println("Invalid --max-worker flag value:", err)
			return viper.GetInt("crawler.maxWorker")
		}

		fmt.Println("Read from flag:", v)
		return v
	}

	v := viper.GetInt("crawler.maxWorker")
	fmt.Println("Read from config:", v)
	return v
}
