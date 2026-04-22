/*
 * SPDX-License-Identifier: GPL-3.0
 * Vencord Installer, a cross platform gui/cli app for installing Vencord
 * Copyright (c) 2023 Vendicated and Vencord contributors
 */

package main

import (
	"image/color"
	"vencord/buildinfo"
)

// https://zfrxncesck1.github.io/trashcord-fallback/releases/trashcord.json
const ReleaseUrl = "https://api.github.com/repos/zFrxncesck1/TrashCord/releases/latest"
const ReleaseUrlFallback = ""
const InstallerReleaseUrl = "https://api.github.com/repos/zFrxncesck1/TrashCordInstaller/releases/latest"
const InstallerReleaseUrlFallback = ""

var UserAgent = "TrashCordInstaller/" + buildinfo.InstallerGitHash + " (https://github.com/zFrxncesck1/TrashCordInstaller)"

var (
	DiscordGreen  = color.RGBA{R: 0x2D, G: 0x7C, B: 0x46, A: 0xFF}
	DiscordRed    = color.RGBA{R: 0xEC, G: 0x41, B: 0x44, A: 0xFF}
	DiscordBlue   = color.RGBA{R: 0x58, G: 0x65, B: 0xF2, A: 0xFF}
	DiscordYellow = color.RGBA{R: 0xfe, G: 0xe7, B: 0x5c, A: 0xff}
)

var LinuxDiscordNames = []string{
	"Discord",
	"DiscordPTB",
	"DiscordCanary",
	"DiscordDevelopment",
	"discord",
	"discordptb",
	"discordcanary",
	"discorddevelopment",
	"discord-ptb",
	"discord-canary",
	"discord-development",
	// Flatpak
	"com.discordapp.Discord",
	"com.discordapp.DiscordPTB",
	"com.discordapp.DiscordCanary",
	"com.discordapp.DiscordDevelopment",
}
