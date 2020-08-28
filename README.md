# Taskiwi

Taskiwi is a tool that works in the local web environment to easily view the task achievements.

In particular, we are targeting data exported by [org-clock-csv](https://github.com/atheriel/org-clock-csv) (Emacs extension).

## Installation

golang is required.

```bash
$ go get github.com/zeroclock/taskiwi
$ go install github.com/zeroclock/taskiwi
# Add a path to PATH if not
$ export PATH=$PATH:~/go/bin >> ~/.bash_profile
$ source ~/.bash_profile
```

## Usage

In your Emacs, `M-x` `org-clock-to-csv-file` and save the csv file to any directory.

```bash
$ taskiwi --path {your csv path}
```

Open `http://localhost:8080` in your browser.

## tests

```bash
$ go test github.com/zeroclock/taskiwi/...
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
