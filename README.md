# Stalk | Chat with other devs over SSH in your Terminal

***

Where are the devs at? They're talking via SSH - or "Stalking"!

Stalk is a custom SSH server that takes you to a chat instead of a shell prompt. Because there's SSH apps on all platforms (even on phones) you can connect to Talk on any device!

## Usage

Try it out:

```sh
ssh talk.dasho.dev
```

If it's your first time logging in, you can choose your display name with the SSH username. For example, if you want to be called "curious", you can run:
```sh
ssh curious@talk.dasho.dev
```
If you want to change your display name after the first login, you should use the `nick` command.


If you're under a firewall, you can still join on port 443:
```sh
ssh talk.dasho.dev -p 443
```

If you add this to `~/.ssh/config`:
```ssh
Host talk
    HostName talk.dasho.dev
```

You'll be able to join with just:
```sh
ssh talk
```

Feel free to make a [new issue](https://github.com/d-a-s-h-o/stalk/issues) if something doesn't work.

### Want to host your own instance?

Quick start:
```shell
git clone https://github.com/d-a-s-h-o/stalk && cd stalk
go install # or build, if you want to keep things in the pwd
ssh-keygen -qN '' -f talk-sshkey # new ssh host key for the server
talk # run! the default config is used & written automatically
```
These commands download, build, setup and run a Talk server listening on port 2222, the default port (change by setting `$PORT`).

Check out the [Admin's Manual](Admin's%20Manual.md) for complete self-host documentation!

### Permission denied?

Talk uses public keys to identify users. If you are denied access: `foo@talk.dasho.dev: Permission denied (publickey)` try logging in on port 443, which does not require a key, using `ssh talk.dasho.dev -p 443`.

This error may happen because you do not have an SSH key pair. Generate one with the command `ssh-keygen` if this is the case. (you can usually check if you have a key pair by making sure a file of this form: `~/.ssh/id_*` exists)

### Help

```text
Welcome to Talk! Talk is a chat over SSH: github.com/d-a-s-h-o/talk
Because there's SSH apps on all platforms, even on mobile, you can join from anywhere.

Run `cmds` to see a list of commands.

Interesting features:
• Rooms! Run cd to see all rooms and use cd #foo to join a new room.
• Markdown support! Tables, headers, italics and everything. Just use \n in place of newlines.
• Code syntax highlighting. Use Markdown fences to send code. Run eg-code to see an example.
• Direct messages! Send a quick DM using =user <msg> or stay in DMs by running cd @user.
• Timezone support, use tz Continent/City to set your timezone.
• Built in Tic Tac Toe and Hangman! Run tic or hang <word> to start new games.
• Emoji replacements! :rocket: => 🚀  (like on Slack and Discord)

For replacing newlines, I often use bulkseotools.com/add-remove-line-breaks.php.
```
### Commands
```text
Commands
   =<user>   <msg>           DM <user> with <msg>
   users                     List users
   color     <color>         Change your name's color
   exit                      Leave the chat
   help                      Show help
   man       <cmd>           Get help for a specific command
   emojis                    See a list of emojis
   bell      on|off|all      ANSI bell on pings (on), never (off) or for every message (all)
   clear                     Clear the screen
   hang      <char|word>     Play hangman
   tic       <cell num>      Play tic tac toe!
   devmonk                   Test your typing speed
   cd        #room|user      Join #room, DM user or run cd to see a list
   tz        <zone> [24h]    Set your IANA timezone (like tz Asia/Dubai) and optionally set 24h
   nick      <name>          Change your username
   pronouns  @user|pronouns  Set your pronouns or get another user's
   theme     <theme>|list    Change the syntax highlighting theme
   rest                      Uncommon commands list
   cmds                      Show this message
```
```
The rest
   people                  See info about nice people who joined
   id       <user>         Get a unique ID for a user (hashed key)
   admins                  Print the ID (hashed key) for all admins
   eg-code  [big]          Example syntax-highlighted code
   lsbans                  List banned IDs
   ban      <user>         Ban <user> (admin)
   unban    <IP|ID> [dur]  Unban a person and optionally, for a duration (admin)
   kick     <user>         Kick <user> (admin)
   art                     Show some panda art
   pwd                     Show your current room
   shrug                   ¯\_(ツ)_/¯
```
tip: `kick` can help kick out an old session when rejoining if needed
## Integrations
See the [Admin's Manual](Admin's%20Manual.md) for more info.

Talk has a plugin API you can use to integrate your own services: [documentation](plugin/README.md). Feel free to add a plugin to the main instance. Just ask for a token on the server.


## Stargazers over time

[![Stargazers over time](https://starchart.cc/d-a-s-h-o/stalk.svg)](https://starchart.cc/d-a-s-h-o/stalk)
