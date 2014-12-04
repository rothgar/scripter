scriptr
======

Scriptr is a go wrapper for running arbitrary commands. It was created out of the need to run an executable as a different user without `sudo` or `su`.

I wrote this tool to learn a little bit of go and solve a very specific problem. If you want something that probably works better you should check out how to do it in C. An example can be found [here](http://www.tuxation.com/setuid-on-shell-scripts.html).

Usage
-----

You can use scriptr by downloading the scriptr executable and running executables in your path.

```
./scriptr whoami
rothgar
```

If you want to run a program as a different user, you, or someone with access, can change the owner of the executable and set the sticky execution bit of scripter.

```
chmod +s scriptr
sudo chown root scriptr
```

Now scripter will fork proccesses as the owner.

```
./scriptr whoami
root
```

Scriptr will also accept flags to the called executable for options.

```
./scriptr id -un
root
```

Warning
-----

Scriptr is incredibly dangerous. Seriously. It should not replace built in security features of your operating system.

Caveats
-----

Scriptr will not work with shell scripts. For an example. If you have a script

```
#!/bin/bash
# ./whoami.sh

whoami
```

You can run the script with

```
chmod +x ./whoami.sh
./scriptr ./whoami.sh
rothgar
```
