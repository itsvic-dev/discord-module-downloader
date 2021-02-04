#!/usr/bin/env python3

print(">>> Discord Module Downloader <<<\n--> reverse-engineered and coded by omame")

import requests
import os
import shutil

BASE_PATH = "https://discord.com/api"

def get_host_version(platform: str, release: str):
    req = requests.get(f"{BASE_PATH}/updates/{release}?platform={platform}")
    data = req.json()
    return data["name"]

def get_modules(platform: str, release: str, host_version: str):
    req = requests.get(f"{BASE_PATH}/modules/{release}/versions.json?host_version={host_version}&platform={platform}&_=0")
    data = req.json()
    return data

def get_module_path(release: str, module: str, version: str, platform: str, host_version: str):
    return f"{BASE_PATH}/modules/{release}/{module}/{version}?host_version={host_version}&platform={platform}"

release = input("Enter release candidate (valid values: stable, ptb, canary): ")
platform = input("Enter target platform (valid values: win, linux, osx): ")

print(f"==> Getting host version for {platform}/{release}...")
host_version = get_host_version(platform, release)
print(f"Host version: {host_version}")

print(f"==> Getting module info for {platform}/{release}/{host_version}...")
modules = get_modules(platform, release, host_version)
print(modules)

if os.path.exists("modules/"):
    shutil.rmtree("modules")

os.mkdir("modules")

for module in modules:
    module_name = module
    module_version = modules[module]
    print(f"--> Downloading module {module_name}@{module_version}...")
    req = requests.get(get_module_path(release, module_name, module_version, platform, host_version))
    with open(f"modules/{module_name}-{module_version}.zip", "wb") as file:
        file.write(req.content)

print("==> Done.")
