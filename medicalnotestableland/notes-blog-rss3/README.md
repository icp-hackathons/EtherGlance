We are extending and adapting the RSS3 powered blog platform and integrating it with medical notes application powered by tableland in the EtherGlance project. Please find the demo at EtherGlance directory at https://drive.google.com/drive/u/3/folders/196Mdygxm-ayOp34myifY2LjxoBb6kG84

# Project Setup

```bash
# 1. install the deps, make sure you have pnpm installed already
pnpm i 

# 2. init a env file to save your sensitive config
cat env-template > .env

# 3. rupdate the config file accordingly

# 4. you are ready to go
pnpm dev
```

# Utils 

## Log

__usage:__
```typescript
// saying u have VUE_APP_DEVELOPER = test

// log in default namespace //
log('What ever u want')
// test:default What ever u want


// log in sub namespace, i.e. logic1 //
const subLog = log.sub('logic1')
subLog('What ever u want')
// test:logic1 What ever u want
```
