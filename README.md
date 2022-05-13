# go-baby-bread

Verified Discord bot (Rewrite of it's C# brother)

## Requirements
- Golang
- config.json ( For utilities.go )

## /configs/config.json
```
{
    "secrets": 
    {
        "w2g": "api-key",
        "discord": "api-key"
    }
}
```

## Build and Run
### Local
```bash
> go get ./... ; go build ./... ; ./main
```
### Docker Compose
```yaml
version: '3'

services:
  go-baby-bread:
    image: ghcr.io/breadboi/go-baby-bread:latest
    volumes:
      - /path/to/go-baby-bread-configs:/app/configs:ro
```

## Major Features
* Watch2gether room generation
* Random Giveaways **(Undergoing rewrite from C#)**
* Mafia **(Undergoing rewrite from C#)**
* Team Randomizer
* Channel Lockdown **(Undergoing rewrite from C#)**


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[GPL-3.0](https://choosealicense.com/licenses/gpl-3.0/)
