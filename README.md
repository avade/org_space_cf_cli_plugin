## Org-Space CLI Plugin
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

### Pre-requisites

```
git clone https://github.com/avade/org_space_cf_cli_plugin.git
cd org_space_cf_cli_plugin/main
go build org_space.go
cp org_space $GOPATH/bin
```

### Installation
```
cf install-plugin org_space
```

### Usage

```
cf org-space <orgname> <spacename>
```


