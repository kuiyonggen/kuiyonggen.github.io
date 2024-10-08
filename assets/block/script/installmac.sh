#!/bin/bash
## install Command Line Tools(CLT) for Xcode
if [ -x "$(command -v git)" ]; then
    echo 'Command Line Tools(CLT) are already installed.'
else
    xcode-select --install
fi
## end of installing Command Line Tools(CLT) for Xcode

## ssh key
if test -r ~/.ssh/id_rsa_gh; then
    echo 'ssh key already existed.'
else
    echo 'Create ssh keys'
    #ssh-keygen -t RSA -C "dev@kyg" -f ~/.ssh/id_rsa_gh
    #ssh-keygen -t RSA -C "dev@huibunny" -f ~/.ssh/id_rsa_bb
    #ssh-keygen -t RSA -C "dev@local"
fi
## end ssh key 

## create .ssh config file
if test -r ~/.ssh/config; then
    echo '.ssh config file already existed.'
else
# Refer it: https://coderwall.com/p/7smjkq/multiple-ssh-keys-for-different-accounts-on-github-or-gitlab
    echo 'Create ~/.ssh/config'
    echo 'Host github.com
    HostName github.com
    PreferredAuthentications publickey
    IdentityFile ~/.ssh/id_rsa_gh
' > ~/.ssh/config
    # Delete cached keys
    ssh-add -D
    # Add keys
    ssh-add ~/.ssh/id_rsa_gh
    # List added keys
    ssh-add -l
    # Check connection
    ssh -T git@github.com
fi
## end of creating .ssh config file

## install Homebrew
if [ -x "$(command -v brew)" ]; then
    echo 'Homebrew is already installed.'
    source ~/.zshrc
else
    echo '
export HOMEBREW_INSTALL_FROM_API=1
export HOMEBREW_API_DOMAIN="https://mirrors.tuna.tsinghua.edu.cn/homebrew-bottles/api"
export HOMEBREW_BOTTLE_DOMAIN="https://mirrors.tuna.tsinghua.edu.cn/homebrew-bottles"
export HOMEBREW_BREW_GIT_REMOTE="https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/brew.git"
export HOMEBREW_CORE_GIT_REMOTE="https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/homebrew-core.git"
' > ~/.zshrc
    source ~/.zshrc
    echo 'Set Homebrew environment'
    echo $HOMEBREW_INSTALL_FROM_API
    echo $HOMEBREW_API_DOMAIN
    echo $HOMEBREW_BOTTLE_DOMAIN
    echo $HOMEBREW_BREW_GIT_REMOTE
    echo $HOMEBREW_CORE_GIT_REMOTE

    git clone --depth=1 https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/install.git brew-install

    /bin/bash brew-install/install.sh

    rm -rf brew-install

    if test -r ~/.zprofile;then echo 'eval "$(/opt/homebrew/bin/brew shellenv)"' >> ~/.zprofile;else touch ~/.zprofile && echo 'eval "$(/opt/homebrew/bin/brew shellenv)"' >> ~/.zprofile;fi
    source ~/.zprofile
    brew -v
fi
## end of installing Homebrew



## brew installation
brew install ruby@3.2
echo 'export PATH="/opt/homebrew/opt/ruby@3.2/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
ruby -v
#brew install ffmpeg
brew install bilix
if [ -x "$(command -v go)" ]; then
    go version
else
    brew install go
    go version
    go env -w GOPROXY=https://goproxy.cn,direct
    go env -w GO111MODULE=on
    go env
fi
brew install calibre
brew install vlc
#brew install docker
#brew install consul
brew install --cask google-chrome
## end brew installation
echo 'done'
