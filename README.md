
# Fix Crowdstrike BSOD

This little tool fixes the Crowdstrike Blue Screen of Death,
also known as the Crowdstrike Boot Loop Screen.


# Cause

On 2024-07-18, CrowdStrike deployed a defective update which
contains a lot of `NUL` bytes (`0x00`). This update caused
the Windows Kernel to be stuck in a boot loop because it crashes
and reboots.


# Solution

This small tool is written in `go` to be usable without external
dependencies and to ease up the task at hand. Download the prebuilt
binaries from the Releases section here on GitHub for your convenience,
and put it on a USB flash drive.

1. Prepare USB flash drive with the `fix-crowdstrike_amd64.exe` on it.
2. Boot the Windows system into `Safe Mode` or the `Windows Recovery Environment`.
3. Insert and Mount the USB flash drive, open the folder in the Explorer.
4. Right Click / Run As Administrator on `fix-crowdstrike_amd64.exe`.
5. Reboot the machine for the last time, it not crash now.


# How to boot the device into Windows Safe Mode

Microsoft Support Documentation: [Article](https://support.microsoft.com/en-us/windows/start-your-pc-in-safe-mode-in-windows-92c27cff-db89-8644-1ce4-b3e5e56fe234#WindowsVersion=Windows_10)

TL;DR:

1. Hold down the power button for 10 seconds to turn off your device.
2. Press the power button to turn on your device.
3. On the first sign that Windows has started (for example, the manufacturer's logo is shown), hold down the power button for 10 seconds to turn off your device.
4. Press the power button to turn on your device.
5. Again when Windows starts, hold down the power button for 10 seconds to turn off your device.
6. Again press the power button to turn on your device.
7. Allow your device to fully bootup. You will enter Windows Recovery Environment.
8. Select `Troubleshoot` / `Advanced Options` / `Startup Settings`.
9. After your device restarts, you'll see a list of options. Select option 5 (by pressing `F5`) for Safe Mode with Networking.


# License

GPL2
