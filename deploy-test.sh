#!/bin/bash

eval "$(ssh-agent -s)" # Start ssh-agent cache
chmod 600 .travis/id_rsa # Allow read access to the private key
ssh-add .travis/id_rsa # Add the private key to SSH

ssh -tt -o StrictHostKeyChecking=no $REMOTE_USER@$SSH_KNOWNHOST <<EOF
docker pull ${TRAVIS_REPO_SLUG}:${IMAGE_VERSION} 
docker run -d ${TRAVIS_REPO_SLUG}:${IMAGE_VERSION}
EOF