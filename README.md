# APP Demo

Since there are many Web apps developed, a common development template is put together here.

### New app
1、You need copy `appdemo` to your `GOPATH` and rename:
```
$ git clone git@github.com:deepzz0/appdemo.git <app name>
```

3、Enter your app, run:
```
$ cd <app name>
$ make _new
```

3、Push the code to new repo:
```
$ git add .
$ git commit -m "init repo"
$ git remote add origin <your repo>
$ git push -u origin master
```

4、`make run app=demo` you can start your web app.

### Development

**Step1**

Understand the directory.

```
├── CHANGELOG.md     # Record version change.
├── LICENSE          # Open source license
├── Makefile         # Makefile: call scripts
├── README.md        # Read me docs.
├── api              # Protocol file
├── assets           # Assets
├── build            # Packaging and Continuous Integration.
├── cmd              # Main applications for this app.
├── conf             # Static configuration file.
├── docs             # Design and user documents.
├── examples         # Examples
├── go.mod           # Go mod file.
├── go.sum           # Go mod lock file
├── init             # Init scripts
├── pkg              # Library code that's ok to use by external applications.
├── scripts          # Scripts to perform various build, install, analysis, etc operations.
└── website          # APP's website data.
```



**Step2**

Code in pkg and cmd or website.


