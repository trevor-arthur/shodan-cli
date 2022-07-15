# Shodan API CLI
*Made by: Trevor Arthur*

### How to Use:

	git clone https://github.com/trevor-arthur/shodan-cli

	echo "<your_shodan_api_key>" > ./shodan-cli/cmd/shodan/.env

	cd ./shodan-cli/cmd/shodan && go build -o shodan-cli

	./shodan-cli <searchterm>

Example:

	./shodan-cli tomcat
	
