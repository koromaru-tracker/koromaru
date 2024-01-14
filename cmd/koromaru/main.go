package main

import (
	"log"

	"github.com/google/uuid"
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

	installCmd := &cobra.Command{
		Use:   "install",
		Short: "Install database",
		Long:  "Install database with the latest schema and seed data",
		Run:   installDatabase,
	}
	installCmd.Flags().String("config", "/etc/koromaru.yaml", "location of configuration file")

	rootCmd.AddCommand(installCmd)
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

	if err := db.AutoMigrate(&model.User{}, &model.Torrent{}, &model.Role{}); err != nil {
		panic(err)
	}
}

func installDatabase(cmd *cobra.Command, args []string) {
	log.Println("Installing database...")
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

	// Create Roles
	roles := []model.Role{
		{
			ID:   uuid.Must(uuid.Parse("a87cc8f5-81e6-4702-a9c5-a80a5faf7ec4")),
			Name: "admin",
		},
		{
			ID:   uuid.Must(uuid.Parse("17ef6f73-cfc7-432e-b20b-7569b53d932d")),
			Name: "user",
		},
		{
			ID:   uuid.Must(uuid.Parse("bd34a62c-2051-4013-896d-82b76c4ee85a")),
			Name: "moderator",
		},
		{
			ID:   uuid.Must(uuid.Parse("524e410e-7618-4050-ab20-9014378d6b44")),
			Name: "uploader",
		},
		{
			ID:   uuid.Must(uuid.Parse("ed06650e-9321-43f8-a207-4f9109590874")),
			Name: "vip",
		},
		{
			ID:   uuid.Must(uuid.Parse("24f99b7d-3f3a-47e1-ad64-a68321f98d54")),
			Name: "system",
		},
	}

	for _, role := range roles {
		if err := db.Create(&role).Error; err != nil {
			panic(err)
		}
	}
}
