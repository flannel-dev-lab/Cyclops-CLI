package main

import (
	"CyclopsCLI/config"
	"CyclopsCLI/database"
	"CyclopsCLI/router"
	"flag"
	"fmt"
	"log"
)

func main() {
	bootstrapPath := flag.String("bootstrap", "", "cyclops-cli bootstrap {PATH OF PROJECT}")

	fmt.Println(*bootstrapPath)

	flag.Parse()
	if bootstrapPath != nil {
		if err := config.BuildConfig(*bootstrapPath, "yaml"); err != nil {
			log.Fatal(err)
		}

		if err := database.BuildDatabase(*bootstrapPath); err != nil {
			log.Fatal(err)
		}

		if err := router.BuildRouter(*bootstrapPath); err != nil {
			log.Fatal(err)
		}

		log.Print("Project has been bootstrapped with necessary files")
	}
}
