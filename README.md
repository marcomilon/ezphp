### EzPHP

```
 ______     _____  _    _ _____ 
|  ____|   |  __ \| |  | |  __ \
| |__   ___| |__) | |__| | |__) |
|  __| |_  /  ___/|  __  |  ___/
| |____ / /| |    | |  | | |
|______/___|_|    |_|  |_|_|
```

EzPHP is an alternative to Xamp/Wamp. EzPHP is the easiest way to setup a PHP development environment for learning PHP programming on Windows.

The scope of this project is to provide a single .exe file that will get you a PHP developing server.

![EzPHP](https://raw.githubusercontent.com/marcomilon/ezphp/master/ezphp.png)

### Installation

1. Download [ezphp.exe](https://github.com/marcomilon/ezphp/releases/download/1.1.1/ezphp.exe).
2. Create a new folder for your project and copy ezphp.exe.
3. Run ezphp.exe. If PHP is not installed locally ezphp will try to download and install PHP.
4. Open your browser in http://localhost:8080. 

Advanced user execute `ezphp.exe -h` to view all options.

### How it works?

After launching ezphp.exe you will get a PHP web server on port 8080. 
To start programming with PHP just copy your PHP files to the **Document root** directory and then open the url **http://localhost:8080** on your web browser.

### Why i created EzPHP?

XAMP or equivalent environments are difficult to set up for users who are learning to code. With EzPHP you have a development environment with just one click.
EzPHP is available only for windows.

### Requirements

PHP binaries required to have *Visual C++ Redistributable for Visual Studio 2017* installed on your computer.
In case you need it you can download it from https://www.microsoft.com/en-us/download/details.aspx?id=48145

### Contribution

Feel free to contribute! Just create a new issue or a new pull request.

### License

This library is released under the MIT License.

