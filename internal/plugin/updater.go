// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The semrel Authors

// Package plugin updates .NET project version files in-place.
package plugin

import (
	"fmt"
	"os"
	"regexp"
)

var versionPattern = regexp.MustCompile(`(?s)(<Version>)[^<]*(</Version>)`)

// Updater updates .csproj and .nuspec versions.
type Updater struct{}

// NewUpdater creates an updater.
func NewUpdater() *Updater {
	return &Updater{}
}

// Update rewrites the Version element.
func (u *Updater) Update(path, version string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}
	if !versionPattern.Match(data) {
		return fmt.Errorf("Version element not found in %s", path)
	}
	updated := versionPattern.ReplaceAllString(string(data), `${1}`+version+`${2}`)
	if err := os.WriteFile(path, []byte(updated), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}
