package main

import (
	"core/models"
	"log"

	"gorm.io/gen"
)

func main() {
	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:        "./db/dao", // Output directory for generated code
		Mode:           gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
		FieldNullable:  true, // Generate pointer fields for nullable columns
		FieldCoverable: true, // Generate pointer fields for coverable columns
	})

	// Apply all your models for code generation
	g.ApplyBasic(
		models.Organization{},
		models.Team{},
		models.User{},
		models.TeamMember{},
		models.DataSource{},
		models.FileToMigrate{},
		models.Document{},
		models.MigrateJobs{},
	)

	log.Println("Starting code generation...")
	g.Execute()
	log.Println("Code generation completed successfully.")
}
