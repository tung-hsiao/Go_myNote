# Go_myNote


在 Linux 安裝 go

```
從官網下載壓縮檔

# 解壓縮剛剛下載的壓縮檔
sudo tar -xvf go1.15.linux-amd64.tar.gz
sudo mv go /usr/local

# 最後是設定環境變數，在~/.profile加入以下內容
export PATH=$PATH:/usr/local/go/bin
```
