# ssh-rm

## What it does

Removes ssh keys and certs from the local ssh-agent using the fingerprint.

## Why?

The OpenSSH `ssh-add` command allows you to remove an ssh key from your local
ssh-agent by including the path to the file.

Example:
```
$ ssh-add -l
4096 SHA256:XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX /path/to/homedir/.ssh/id_rsa (RSA)
$ ssh-add -d /path/to/homedir/.ssh/id_rsa
Identity removed: /path/to/homedir/.ssh/id_rsa
```

## How?

Sometimes this may be inconvenient. This command will allow you to remove an
ssh key from the local agent using the key's fingerprint.

Example:
```
$ ssh-add -l
4096 SHA256:XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX /path/to/homedir/.ssh/id_rsa (RSA)
$ ssh-rm -fingerprint SHA256:XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
