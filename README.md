<p align="center">
    <img src="assets/logo/cloud-computing.svg" width="250">
</p>

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting started](#getting-started)
  - [Relay Server](#relay-server)
  - [Game server](#game-server)
- [Bot Commands](#bot-commands)
- [Troubleshooting](#troubleshooting)
  - [Game server is not connecting to the relay server?](#game-server-is-not-connecting-to-the-relay-server)
  - [I'm getting errno 3 from the plugin](#im-getting-errno-3-from-the-plugin)
- [Credits](#credits)
- [License](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Introduction

<p>
    <a href="https://travis-ci.com/rumblefrog/source-chat-relay">
        <img src="https://img.shields.io/travis/com/rumblefrog/source-chat-relay.svg?style=for-the-badge">
    </a>
    <a href="https://discord.gg/TZ4BsrQ">
        <img src="https://img.shields.io/discord/443915420324331521.svg?style=for-the-badge">
    </a>
    <a href="https://github.com/rumblefrog/source-chat-relay/issues">
        <img src="https://img.shields.io/github/issues/rumblefrog/source-chat-relay.svg?style=for-the-badge">
    </a>
    <a href="https://github.com/rumblefrog/source-chat-relay/blob/master/LICENSE">
        <img src="https://img.shields.io/github/license/rumblefrog/source-chat-relay.svg?style=for-the-badge">
    </a>
</p>

<img src="assets/preview.gif">

Communicate between Discord & In-Game, monitor server without being in-game, control the flow of messages and user base engagement!

## Features
 - Receive and send messages bidrectionally
 - Channel configuration for powerful setups
 - Setup is incrediblily easy with Discord bot commands and simple config files
 - Upon disconnect, game servers will attempt to reconnect at a fixed interval

## Prerequisites
 - Server to host the relay binary on (with MySQL if not external)
 - Source game server with the [socket](https://forums.alliedmods.net/showthread.php?t=67640) extension
 - A Discord bot token (https://discordapp.com/developers/applications/)

## Getting started
 1. Download the latest release from [releases](https://github.com/rumblefrog/source-chat-relay/releases) for your operating system
 
#### Relay Server

1. Upload the binary (server[.exe]) to the server
2. Configure `config.toml.example` and rename it to `config.toml`
3. Start it by running `./server`

#### Game server

1. Upload the plugin to `addons/sourcemod/plugins`
2. Load the plugin via `sm plugins load Source-Chat-Relay`
3. Configure the config at `cfg/sourcemod/Source-Chat-Relay.cfg`
4. Reload the plugin via `sm plugins reload Source-Chat-Relay`

If all goes correctly, a new entity will appear when you use `r/entities` command

Following that, you may use `r/receivechannel` and `r/sendchannel` to set receive and send respectively

## Bot Commands

Before any clients can send messages, you must set the receive/send channels on them

 - r/receivechannel #TextChannel/EntityID Channel
    - #TextChannel/EntityID - Either a text channel mention via #channel-name or a game server entity ID obtainable via `r/entities`
    - Channel - List of channels to receive at. If more than one, you may use comma to separate them

 - r/sendchannel #TextChannel/EntityID Channel
    - #TextChannel/EntityID - Either a text channel mention via #channel-name or a game server entity ID obtainable via `r/entities`
    - Channel - List of channels to send to. If more than one, you may use comma to separate them

 - r/entities ?channel/server
    - ?channel/server - Optional argument, allows you filter the return results by either channel or server type

## Troubleshooting

Common problem & steps to troubleshoot

#### Game server is not connecting to the relay server?

This can be numerous amount of things, check the game server log for errors emitted by the plugin

#### I'm getting errno 3 from the plugin

Either the server is not started or your firewall is not configured to allow connections to it

## Credits
 - Ron for being a human linter

## License

This project is licensed under GPL 3.0

Icons made by <a href="https://www.flaticon.com/authors/itim2101" title="itim2101">itim2101</a> from <a href="https://www.flaticon.com/" title="Flaticon">www.flaticon.com</a> is licensed by <a href="http://creativecommons.org/licenses/by/3.0/" title="Creative Commons BY 3.0" target="_blank">CC 3.0 BY</a>