package main

import (
	"log"

	"github.com/koromaru-tracker/koromaru/index/database"
	"github.com/koromaru-tracker/koromaru/index/model"
	"github.com/koromaru-tracker/koromaru/index/server"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "koromaru",
		Short: "BitTorrent Indexer",
		Long:  "A BitTorrent Indexer written in Go",
		Run:   startServer,
	}

	rootCmd.Flags().String("config", "/etc/koromaru.yaml", "location of configuration file")

	migrationCmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate database",
		Long:  "Migrate database with the latest schema",
		Run:   migrateDatabase,
	}
	migrationCmd.Flags().String("config", "/etc/koromaru.yaml", "location of configuration file")

	rootCmd.AddCommand(migrationCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal("failed when executing root cobra command: " + err.Error())
	}

}

func startServer(cmd *cobra.Command, args []string) {
	configFilePath, err := cmd.Flags().GetString("config")
	if err != nil {
		panic(err)
	}

	cfg, err := ParseConfig(configFilePath)
	if err != nil {
		panic(err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		panic(err)
	}

	server.Serve(cfg, db)
}

func migrateDatabase(cmd *cobra.Command, args []string) {
	log.Println("Migrating database...")
	configFilePath, err := cmd.Flags().GetString("config")
	if err != nil {
		panic(err)
	}

	cfg, err := ParseConfig(configFilePath)
	if err != nil {
		panic(err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&model.User{}, &model.Torrent{}); err != nil {
		panic(err)
	}
}
