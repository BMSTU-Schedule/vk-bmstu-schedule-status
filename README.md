# vk-bmstu-schedule-status
Display BMSTU week info in VK chat title.

## Installation
```shell
go get -u github.com/trubitsyn/vk-bmstu-schedule-status/cmd/vk-bmstu-schedule-status
```

## Usage
```
Usage of vk-bmstu-schedule-status:
  -chat uint
        Chat ID
  -login string
        VK login
  -password string
        VK password
  -prefix string
        Chat title prefix
```

## Testing
```shell
go get -t github.com/trubitsyn/vk-bmstu-schedule-status
go test github.com/trubitsyn/vk-bmstu-schedule-status
```

## LICENSE
```
vk-bmstu-schedule-status
Copyright (C) 2019 Nikola Trubitsyn

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
```