Org-Space CLI Plugin
--------------------

Created by **Alex Ley & Ben Laplanche** for CF Hackday Winter 2014

![Dancing Panda](http://media.giphy.com/media/txsJLp7Z8zAic/giphy.gif)

### Purpose

Ever experience the frustration of having to run the following commands?

```
cf create-org awesomecompany
cf create-space awesomespace -o awesomecompany
cf target -o awesomecompany -s awesomespace
```

Then use our new plugin!

**Org-Space** will automatically detect whether the organisation and space exist, if not it will create them, and finally target them.

### Installation

```
go get github.com/avade/org_space_cf_cli_plugin
cf install-plugin $GOPATH/bin/org_space_cf_cli_plugin
```

### Usage

```
cf org-space <orgname> <spacename>
```
