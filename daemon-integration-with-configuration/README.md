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

    * Second place that you will have to change is logic inside function `yourMethod()`. Update the function name as per your requirement. Function `yourMethod()` contains the logic that you want to be executed in specified interval of time on repeating loop. Another way around is to directly call your intended function that you want to run, just simply by exporting it as public method. Then you can call it from the ticker section of `Manage()` function.
7. After you have finalised all changes to the code, copy the code to your ec2 instance. You can use vi editor to paste the code

    ```bash
    vi main.go
    #This will open an editor, now go to Insert mode by clicking 'i' then paste the code in main.go
    #Now save the file by pressing escape and then entering :wq! then press enter
    ```

    if you are using/referring your custom public method then make sure to copy the whole package to EC2
8. Now build the code using command `go build`. Now in case of following errors:
    * Package missing error use `go get package-name`
    * Config File/Log File related error, make sure to create/use correct location address in your code
    * In case of any other issues/error resolve them based on error message
9. If the build is successful you will get an executable file to run. In our code we have defined 5 types of arguments - install, start, status, stop and remove
    * install - Registers your program as a daemon service
    * start - Starts your program
    * status - Shows status of your process
    * stop - Stops the running process
    * remove - Removes/Uninstalls the registered service

    First thing you need to do is install the service using command

    ```bash
    ./testservicename install
    ```

    Then start the service using command

    ```bash
    ./testservicename start
    ```

    Now check the status for your service

    ```bash
    ./testservicename status
    ```

    Status should be shown as `running`, but if it is `stopped` then check logs for any issues and check your code properly. *For me most frequent reason for this was using unavailable port number*

    To stop your service use

    ```bash
    ./testservicename stop
    ```

    and to Uninstall your service use

    ```bash
    ./testservicename remove
    ```

10. Our service will log all our data to our specified logfile location. Go to logfile location and check the log file using command `cat testservice.log` or `tail -n num testservice.log` where num is the number of bottom last lines you want to be displayed.

> Important Note
>
> Value for timer in config file is by default defined in ns i.e Nano Seconds but in our code we have statically modified it to seconds. So...
>
> timer: 60 means 1 minute
>
> timer: 600 means 10 minutes
>
> timer: 3600 means 1 hour