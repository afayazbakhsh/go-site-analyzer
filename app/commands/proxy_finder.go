package commands

import (
	"context"
	"encoding/base64"
	"fmt"
	"gocrawler/app/crawler"
	"net"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var proxyFinder = &cobra.Command{
	Use:   "proxy",
	Short: "find best proxy",
	Run:   do_find_proxy,
}

var pingWG sync.WaitGroup

type ProxyConfig struct {
	Type   string // vmess | vless | trojan
	Raw    string
	Source string
}

var proxies []ProxyConfig

func init() {
	rootCmd.AddCommand(proxyFinder)
}

func do_find_proxy(cmd *cobra.Command, args []string) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	PROXY_GROUPS := []string{
		"https://raw.githubusercontent.com/Epodonios/v2ray-configs/refs/heads/main/All_Configs_Sub.txt",
		"https://raw.githubusercontent.com/barry-far/V2ray-config/main/All_Configs_Sub.txt",
		"https://raw.githubusercontent.com/shabane/kamaji/master/hub/tested/merged.txt",
		"https://raw.githubusercontent.com/MatinGhanbari/v2ray-configs/main/subscriptions/v2ray/super-sub.txt",
	}

	for _, link := range PROXY_GROUPS {

		response := crawler.Fetch(ctx, link)
		content := maybeBase64Decode(string(response.Body))

		lines := strings.Split(content, "\n")

		for _, line := range lines {
			line = strings.TrimSpace(line)

			switch {
			case strings.HasPrefix(line, "vmess://"):
				proxies = append(proxies, ProxyConfig{
					Type: "vmess",
					Raw:  line,
				})

			case strings.HasPrefix(line, "vless://"):
				proxies = append(proxies, ProxyConfig{
					Type: "vless",
					Raw:  line,
				})

			case strings.HasPrefix(line, "trojan://"):
				proxies = append(proxies, ProxyConfig{
					Type: "trojan",
					Raw:  line,
				})
			}
		}
	}

	proxies = deduplicate(proxies)

	pingWG.Add(1)
	go ping(proxies)
	pingWG.Wait()
}

func ping(proxies []ProxyConfig) {
	defer pingWG.Done()

	for _, v := range proxies {

		cmd := exec.Command("ping", "-c", "4", v.Raw)

		output, err := cmd.CombinedOutput()

		fmt.Println(output)

		if err != nil {
			fmt.Println("خطا:", err)
			return
		}
	}
}

func tcpPing(host, port string) time.Duration {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), 3*time.Second)
	if err != nil {
		return -1
	}
	conn.Close()
	return time.Since(start)
}

func maybeBase64Decode(s string) string {
	decoded, err := base64.StdEncoding.DecodeString(strings.TrimSpace(s))
	if err != nil {
		return s // base64 نبود
	}
	return string(decoded)
}

func deduplicate(proxies []ProxyConfig) []ProxyConfig {
	seen := make(map[string]struct{})
	var result []ProxyConfig

	for _, p := range proxies {
		key := p.Type + "|" + p.Raw

		if _, exists := seen[key]; exists {
			continue
		}

		seen[key] = struct{}{}
		result = append(result, p)
	}

	return result
}
