## Usage

Clone the repository.

```
git clone git@github.com:daiki-onda/serverless-go-template.git
```

Installing all the dependencies of serverless.

```
yarn
```

Let's build.

```
$GOOS=linux go build -o bin/main
```

Deploy.

```
yarn run sls deploy
```
