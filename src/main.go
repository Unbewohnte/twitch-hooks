package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
	"twitch-hooks/config"
	"twitch-hooks/discordhooks"
	"twitch-hooks/twitchhooks"
	"twitch-hooks/vkhooks"
)

var (
	logo string = `  _______       _ _       _           _                 _        
|__   __|     (_) |     | |         | |               | |       
   | |_      ___| |_ ___| |__ ______| |__   ___   ___ | | _____ 
   | \ \ /\ / / | __/ __| '_ \______| '_ \ / _ \ / _ \| |/ / __|
   | |\ V  V /| | || (__| | | |     | | | | (_) | (_) |   <\__ \
   |_| \_/\_/ |_|\__\___|_| |_|     |_| |_|\___/ \___/|_|\_\___/ by Unbewohnte`
	Config       config.Config
	pathToConfig *string = flag.String("config", config.DefaultConfigFilename, "Specifies path to a config in another directory")
	delay        *uint   = flag.Uint("delay", 5000, "Delay in seconds for each check cycle")
)

func init() {
	// process the config file
	flag.Parse()

	if *pathToConfig != config.DefaultConfigFilename {
		// config in another place
		config.ReadConfig(*pathToConfig)
	}

	if !config.ConfigExists(*pathToConfig) {
		// there is no existing config file;
		// create a new one and exit
		err := config.CreateConfig(filepath.Dir(*pathToConfig))
		if err != nil {
			panic(err)
		}

		fmt.Println("Created a new config file")
		os.Exit(0)
	}

	configContents, err := config.ReadConfig(*pathToConfig)
	if err != nil {
		panic(err)
	}

	Config = *configContents
}

func main() {
	fmt.Println(logo)

	if Config.Keys.Twitch.ClientID == "" || Config.Keys.Twitch.ClientSecret == "" {
		// no twitch api key used. Notify the user and check for the force-send flag
		fmt.Println("No Twitch API keys found")

		if !Config.ForceSend {
			// not forced to send messages. Exiting
			fmt.Println("Not forced to send. Exiting...")
			os.Exit(0)
		}
	}

	for {
		// retrieve access token
		tokenResp, err := twitchhooks.GetToken(&Config.Keys.Twitch)
		if err != nil {
			panic(err)
		}

		// check if live
		is_live, err := twitchhooks.IsLive(Config.TwitchName, &twitchhooks.RequestOptions{
			ApplicationKeys: Config.Keys.Twitch,
			AccessToken:     *tokenResp,
		})
		if err != nil {
			panic(err)
		}

		if is_live || Config.ForceSend {
			// live or forced to send -> send alerts
			if Config.Keys.Discord.WebhookUrl != "" {
				err := discordhooks.Post(Config.Keys.Discord.WebhookUrl, Config.Messages.DiscordMessage)
				if err != nil {
					panic(err)
				}
			}

			if Config.Keys.VK.Key != "" {
				vkhooks.Initialise(Config.Keys.VK.Key)
				err := vkhooks.Send(Config.Messages.VKmessage)
				if err != nil {
					panic(err)
				}
			}
			// alerted. Now exiting
			fmt.Println("Alerts has been sent ! My work is done here...")
			os.Exit(0)
		}
		duration, _ := time.ParseDuration(fmt.Sprintf("%ds", *delay))
		time.Sleep(duration)
	}

}
