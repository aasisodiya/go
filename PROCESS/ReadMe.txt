# How to run any program as Daemon

---
> For code related details please refer below link
> 
> ***Reference:*** [Go Daemon](https://github.com/takama/daemon)

## Instructions on How to run any program as Daemon

---

1. Connect to your Ubuntu EC2 instance2.
2. Install Go Lang if not installed

    ```bash
    sudo add-apt-repository ppa:longsleep/golang-backports
    sudo apt-get update
    sudo apt-get install golang-go
    ```

    Also make sure that go version is **1.13 or above** using command `go version`

3. Create 2 New Folders - One each for logs and config file, and note down the folder location as we will need it again while writing the code. For my code I have created folders with respective addresses using following commands

    ```bash
    mkdir /home/ubuntu/logs
    mkdir /home/ubuntu/configs
    ```

4. Open your configs folder and create following config file. Here I am creating config file with name config.yml

    ```bash
    cd /home/ubuntu/configs
    cat > config.yml <<EOF
    service:
        name: testservicename
        port: 8000
        timer: 60
    logfile:
        location: /home/ubuntu/logs
        name: testservice.log
    EOF
    ```

    **Please make sure to make all necessary changes in config.yml based on your requirements and changes you need**
5. Now go to any folder you want to create your Go Lang Code, here in my case I am creating the sample program inside /home/ubuntu/sample. (Command used to create the directory `mkdir /home/ubuntu/sample`)
6. Create main.go file in your local editor before you move it to EC2. Take a copy of this code [Code](www.github.com) and now edit it according to your requirements.
    * First place that you will have to edit is constants defined in the code. Make sure the config file location is correct

        ```go
        const (
            description = "some description of the program"
            // This needs to be configured properly with a static file address (don't use relative address here)
            configFile = "/home/ubuntu/configs/config.yml"
        )
        ```

    * Second place that you will have to change is logic inside `yourMethod()`. Update the function name as per your requirement. `yourMethod()` contains the logic that you want to be executed in specified interval of time on repeating loop.



