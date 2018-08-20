package main

import (
	"github.com/1dustindavis/gorilla/pkg/catalog"
	"github.com/1dustindavis/gorilla/pkg/config"
	"github.com/1dustindavis/gorilla/pkg/installer"
	"github.com/1dustindavis/gorilla/pkg/manifest"
)

func main() {
	// Get the configuration
	config.Get()

	// Get the catalog
	catalog := catalog.Get()

	// Get the manifests
	manifests := manifest.Get()

	// Compile all of the installs, uninstalls, and updates into arrays
	var installs, uninstalls, updates []string
	for _, manifestItem := range manifests {
		// Installs
		for _, item := range manifestItem.Installs {
			if item != "" && catalog[item].InstallerItemLocation != "" {
				installs = append(installs, item)
			}
		}
		// Uninstalls
		for _, item := range manifestItem.Uninstalls {
			if item != "" {
				uninstalls = append(uninstalls, item)
			}
		}
		// Updates
		for _, item := range manifestItem.Updates {
			if item != "" {
				updates = append(updates, item)
			}
		}
	}

	// Iterate through the installs array, install dependencies, and then the item itself
	for _, item := range installs {
		// Check for dependencies and install if found
		if len(catalog[item].Dependencies) > 0 {
			for _, dependency := range catalog[item].Dependencies {
				installer.Install(catalog[dependency])
			}
		}
		// Install the item
		installer.Install(catalog[item])
	}

	// Iterate through the uninstalls array and uninstall the item
	for _, item := range uninstalls {
		// Install the item
		installer.Uninstall(catalog[item])
	}

	// Iterate through the updates array and update the item **if it is already installed**
	for _, item := range updates {
		// Install the item
		installer.Update(catalog[item])
	}
}
