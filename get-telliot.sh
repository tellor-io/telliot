echo "Creating configs.."
mkdir ./configs
cd ./configs
wget https://raw.githubusercontent.com/tellor-io/telliot/master/configs/index.json
wget https://raw.githubusercontent.com/tellor-io/telliot/master/configs/manualData.json
wget https://raw.githubusercontent.com/tellor-io/telliot/master/configs/.env.example
mv .env.example .env
echo "Getting latest telliot release.."
cd ../
wget https://github.com/tellor-io/telliot/releases/latest/download/telliot
chmod +x telliot
echo "Finished."