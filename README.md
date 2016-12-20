This is a fork of the systray sources from
https://github.com/getlantern/systray

Fork was created in order to make modifications such that Mac OS code does not inhibit other
Mac-specific logic from working. In previous implementation, AppDelegate was both declared
and defined in implementation file, making categories difficult. Logic requiring access 
to earliest Application Delegate lifecycle could not be implemented. 

This fork aims to change that. Below are the original contents of the readme as written
by the original author.

- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

Package systray is a cross platfrom Go library to place an icon and menu in the notification area.
Tested on Windows 8, Mac OSX, Ubuntu 14.10 and Debian 7.6.

## Usage
```go
func main() {
	// Should be called at the very beginning of main().
	systray.Run(onReady)
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Pretty awesome超级棒")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
}
```
Menu item can be checked and / or disabled. Methods except `Run()` can be invoked from any goroutine. See demo code under `example` folder.

## Platform specific concerns

### Linux

```sh
sudo apt-get install libgtk-3-dev libappindicator3-dev
```
Checked menu item not implemented on Linux yet.

### Windows

Use the Visual Studio solution to build systray.dll. Make sure to target Windows
XP and build for Release (not Debug). Put the resulting dll in
`dll/systray_unsigned.dll` and then run `./signdll.bash` to sign it. Check the
resulting dll into git at Git at dll/systray.dll and run ./embeddll.bash to
generate the systraydll_windows.go file.

The solution is configured to build with platform toolset v90 and dynamic
linking to save on size and support Windows XP.  To get platform toolset v90,
you need to install Visual Studio 2008 (express edition is okay). You can
build with a more recent Visual Studio, you just need the old one for the
toolset.

## Try

Under `example` folder.
Place tray icon under `icon`, and use `make_icon.bat` or `make_icon.sh`, whichever suit for your os, to convert the icon to byte array.
Your icon should be .ico file under Windows, whereas .ico, .jpg and .png is supported on other platform.

```sh
go get
go run main.go
```

## Credits

- https://github.com/getlantern/systray
- https://github.com/xilp/systray
- https://github.com/cratonica/trayhost
