```
      ___           ___           ___                       ___     
     /__/\         /  /\         /__/|        ___          /__/\    
     \  \:\       /  /::\       |  |:|       /  /\        |  |::\   
      \__\:\     /  /:/\:\      |  |:|      /  /:/        |  |:|:\  
  ___ /  /::\   /  /:/~/::\   __|  |:|     /__/::\      __|__|:|\:\ 
 /__/\  /:/\:\ /__/:/ /:/\:\ /__/\_|:|____ \__\/\:\__  /__/::::| \:\
 \  \:\/:/__\/ \  \:\/:/__\/ \  \:\/:::::/    \  \:\/\ \  \:\~~\__\/
  \  \::/       \  \::/       \  \::/~~~~      \__\::/  \  \:\      
   \  \:\        \  \:\        \  \:\          /__/:/    \  \:\     
    \  \:\        \  \:\        \  \:\         \__\/      \  \:\    
     \__\/         \__\/         \__\/                     \__\/    

```

A doctor for diagnosing your garden.

## Usage

On your Windows cell, run `hakim.exe`. It will error (hopefully helpfully) if there are errors, and just exit 0 if everything we check for is fine. You can optionally provide the `-gardenAddr` flag to target a Garden running somewhere other than `localhost:9241`.

## Download

The latest Hakim binaries can be found on the [releases page](https://github.com/cloudfoundry-incubator/diego-windows-release/releases) of [diego-windows-release](https://github.com/cloudfoundry-incubator/diego-windows-release/).

## Current checks

- Both [diego-windows-release](https://github.com/cloudfoundry-incubator/diego-windows-release/) services are running (metron.exe, rep.exe)
- Both [garden-windows-release](https://github.com/cloudfoundry-incubator/garden-windows-release/) services are running (garden-windows.exe, containerizer.exe)
- A container can be created through the garden API
