package main

import (
	"os"
	"testing"

	"github.com/el-hadji-mamadou-sarr/gestion-de-livraison.git/pkg/utils"
	"github.com/stretchr/testify/assert"
)

// Test du nettoyage du fichier log
func TestClearLog(t *testing.T) {
	testFile := "test_status.log"
	os.WriteFile(testFile, []byte("log data"), 0644)

	utils.ClearLog(testFile)

	content, err := utils.ReadLog(testFile)
	assert.NoError(t, err)
	assert.Equal(t, "", content)

	// Suppression du fichier test
	os.Remove(testFile)
}

// Test de la lecture du fichier log
func TestReadStatus(t *testing.T) {
	testFile := "test_status.log"
	expectedContent := "Test delivery update"
	os.WriteFile(testFile, []byte(expectedContent), 0644)

	data, err := utils.ReadLog(testFile)
	assert.NoError(t, err)
	assert.Equal(t, expectedContent, data)

	// Nettoyage après test
	utils.ClearLog(testFile)
}

// Test du fonctionnement du système de tracking (simulation)
func TestTrackingSystem(t *testing.T) {
	statusFile := "test_status.log"
	expectedContent := "Delivery in progress..."
	os.WriteFile(statusFile, []byte(expectedContent), 0644)

	data, err := utils.ReadLog(statusFile)
	assert.NoError(t, err)
	assert.Equal(t, expectedContent, data)

	utils.ClearLog(statusFile)

	dataAfterClear, err := utils.ReadLog(statusFile)
	assert.NoError(t, err)
	assert.Equal(t, "", dataAfterClear)

	os.Remove(statusFile)
}
