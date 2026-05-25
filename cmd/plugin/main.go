// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The semrel Authors

package main

import (
	"log"

	plugin "github.com/SemRels/updater-nuget/internal/plugin"
)

func main() {
	publisher := plugin.NewPublisher(plugin.Config{})
	log.Printf("updater-nuget plugin ready: updates NuGet package versions and publishes packages (%T)", publisher)
}
