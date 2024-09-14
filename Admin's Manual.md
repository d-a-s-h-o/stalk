# talk Admin's Manual

This document is for those who want to manage a self-hosted Talk server.

Feel free to make a [new issue](https://github.com/d-a-s-h-o/talk/issues) if something doesn't work.

## Installation
```shell
git clone https://github.com/d-a-s-h-o/talk
cd talk
```
To compile talk, you will need Go installed with a minimum version of 1.17.

Now run `go install` to install the talk binary globally, or run `go build` to build and keep the binary in the working directory.

You may need to generate a new key pair for your server using the `ssh-keygen` command. When prompted, save as `talk-sshkey` since this is the default location (it can be changed in the config).
While you can use the same key pair that your user account has, it is recommended to use a new key pair.

## Usage

```shell
./talk # use without "./" for a global binary
```

talk listens on port 2222 for new SSH connections by default. Users can now join using `ssh -p 2222 <server-hostname>`.

Set the environment variable `PORT` to a different port number or edit your config to change what port talk listens for SSH connections on. Users would then run `ssh -p <port> <server-hostname>` to join.

## Configuration

talk writes the default config file if one isn't found, so you do not need to make one before using talk. 

The default location talk looks for a config file is `talk.yml` in the current directory. Alternatively, it uses the path set in the `talk_CONFIG` environment variable.

An example config file:
```yaml
# what port to host a server on ($PORT overrides this)
port: 2222
# an alternate port to avoid firewalls
alt_port: 443
# what port to host profiling on (unimportant)
profile_port: 5555
# where to store data such as bans and logs
data_dir: talk-data
# where the SSH private key is stored
key_file: talk-sshkey
# whether to censor messages (optional)
censor: true
# a list of admin IDs and notes about them
admins:
  daba388eaea39428b7f6a287bade7a59ae428a53e210689c7642961ab68b41c2: 'dasho (sysAdmin)'
```

### Using admin power

As an admin, you can ban, unban and kick users. When logged into the chat, you can run commands like these:
```shell
ban <user>
ban <user> 1h10m
ban <user> "Reason for ban" 10m
unban <user ID or IP>
kick <user>
```

If running these commands makes sokka[bot] complain about authorization, you need to add your ID under the `admins` key in your config file (`talk.yml` by default).

### Enabling a user allowlist

Talk can use be used as a private chatroom. Add this to your config:

```yaml
private: true # enable allowlist checking
allowlist: 
  daba388eaea39428b7f6a287bade7a59ae428a53e210689c7642961ab68b41c2: 'dasho (sysAdmin)'
  ...
```

The `allowlist` has the same format as the `admins` list. Add the IDs of the allowed users and info about that user (this is to make IDs easier to identify when editing the config file, and isn't used by talk)

All admins are allowed even if their ID is not in the allowlist. So, if everyone on the private server is an admin, an allowlist isn't necessary, just enable private mode.

Message backlog on `#main` is disabled in private chats. Only those logged in at the same time as you can read your messages.

### Enabling integrations

Talk includes features that may not be needed by self-hosted instances. These are called integrations.

You can enable these integrations by setting the `integration_config` in your config file to some path:

```yaml
integration_config: talk-integrations.yml
```
Now make a new file at that path. This is your integration config file.

### Using the plugin API integration

Talk includes a built-in gRPC plugin API. This is useful for building your own integration or using a third-party one.

Documentation for using the gRPC API is available [here](plugin/README.md). This integration stores API tokens inside the data directory.

```yaml
rpc:
    port: 5556 # port to listen on for gRPC clients
```

Use the token issuing commands detailed in the [plugin documentation](plugin/README.md) to allow clients to authenticate.

You may also hard-code a key that can be used as an authentication token, but this is not recommended. This option can be useful for server owners trying to keep some API clients always online, since this key cannot be revoked by admins (unlike tokens).

```yaml
    key: "some string" # a string that gRPC clients authenticate with (optional)
```

You can use any number of integrations together.

There are 4 environment variables you can set to quickly disable integrations on the command line:
* `talk_OFFLINE_RPC=true` will disable the gRPC server
* `talk_OFFLINE=true` will disable all integrations.