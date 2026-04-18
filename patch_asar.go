package main

import (
    "bytes"
    "os"
)

func patchAsar(asarPath string) error {
    data, err := os.ReadFile(asarPath)
    if err != nil {
        return err
    }

    old := []byte("https://github.com/Equicord/Equicord")
    new := []byte("https://github.com/zFrxncesck1/TrashCord")
    newData := bytes.ReplaceAll(data, old, new)

    oldRepo := []byte("Equicord/Equicord")
    newRepo := []byte("zFrxncesck1/TrashCord")
    newData = bytes.ReplaceAll(newData, oldRepo, newRepo)

    err = os.WriteFile(asarPath, newData, 0644)
    if err != nil {
        return err
    }

    Log.Info("Asar file patched with correct repository URL")
    return nil
}
