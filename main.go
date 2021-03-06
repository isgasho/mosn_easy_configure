package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

func main() {
	p := prompt.New(nil, completer, prompt.OptionPrefix(">> "))
	for {
		cmd := p.Input()
		if cmd == "quit" || cmd == "exit" || cmd == "q" {
			os.Exit(0)
		}

		switch {
		case strings.HasPrefix(cmd, addClusterCmd+" "):
			fmt.Println(cmd)
			clusterName := strings.Split(cmd, " ")[1]
			enterCluster(clusterName)
		case strings.HasPrefix(cmd, addRouterCmd+" "):
			routerName := strings.TrimLeft(cmd, addRouterCmd+" ")
			enterRouter(routerName)
		case strings.HasPrefix(cmd, addListenerCmd+" "):
			listenerName := strings.TrimLeft(cmd, addListenerCmd+" ")
			enterListener(listenerName)
		default:
			fmt.Println("unsupported command")
		}
	}
}

const (
	addClusterCmd  = "addcluster"
	addRouterCmd   = "addrouter"
	addListenerCmd = "addlistener"
)

type clusterConfig struct{}

func enterCluster(clusterName string) {
	clusterCompleter := func(d prompt.Document) []prompt.Suggest {
		s := []prompt.Suggest{
			{Text: "hostadd", Description: "hostadd [ip:port]"},
			{Text: "lb_type LB_RANDOM", Description: "LB_RANDOM"},
			{Text: "lb_type LB_ROUNDROBIN", Description: "LB_ROUNDROBIN"},
		}
		return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
	}

	clusterPrompt := prompt.New(nil, clusterCompleter, prompt.OptionPrefix(fmt.Sprintf(`.. CLUSTER[%v]) `, clusterName)))

	for {
		cmd := clusterPrompt.Input()
		fmt.Println(cmd)
		if cmd == "exit" || cmd == "quit" || cmd == "q" {
			break
		}
	}
}

func enterRouter(routerName string) {
}

func enterListener(listenerName string) {
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: addListenerCmd, Description: addListenerCmd + " [listener_name]"},
		{Text: addRouterCmd, Description: addRouterCmd + " [router_name]"},
		{Text: addClusterCmd, Description: addClusterCmd + " [cluster_name]"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
