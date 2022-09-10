#!/bin/sh
set -e


echo "================================================================"
echo "===========================DEPLOYMENT==========================="
echo "================================================================"

echo "Update codebase..."
cd ~/skripsi/final-project-university
git fetch origin main
git reset --hard origin/main

# echo "Moving to backend folder"
# cd ~/skripsi/final-project-university/backend

# echo "Installing dependencies 🛠"
# go mod tidy

# echo "Restart pm2 service backend 🔥"
# pm2 restart deploy.json

# echo "Deploying Backend Application Successfully Yeayyyy ......."

echo "========================= FRONT END ========================="

echo "Moving to frontend folder"
cd ~/skripsi/final-project-university/frontend

echo "Installing dependencies 🛠"
yarn install

echo "Building application ⚙"
yarn build

echo "Restart pm2 service 🔥"
pm2 restart pm2.json

echo "Deploying Frontend Application Successfully Yeayyyy ........"