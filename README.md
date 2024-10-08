# dlevel

<h4 align="center"><b>Filter list of subdomains by level.</b></h4><br>

![image](https://user-images.githubusercontent.com/72344025/234942102-74d5cf10-b582-4278-9645-81d5de770980.png)

## Installation
```
git clone https://github.com/rix4uni/dlevel.git && cd dlevel && chmod +x dlevel && mv dlevel /usr/bin/
```

## Via wget
```
wget https://raw.githubusercontent.com/rix4uni/dlevel/refs/heads/main/dlevel && chmod +x dlevel && mv dlevel /usr/bin/
```

## Usage
```bash
# Level 1, all subdomains
$ cat subs.txt | dlevel -l 1
mtnairborn.com
mtnbusiness.com

# Level 2, all subdomains
$ cat subs.txt | dlevel -l 2
woolpack.mtnfootball.com
wcu-csm.mtnfootball.com
weirdnewsday.mtnfootball.com

# Level 8, 2 subdomains
$ cat subs.txt | dlevel -l 8 -f 2
h2a4.n1.ips.mtn.co.ugmut-pts35k-1.mtn.co.ug
h1bbd.n2.ips.mtn.co.ugmut-pts35k-1.mtn.co.ug

# All Highest Level subdomains -t not work with other options
$ cat subs.txt | dlevel -t
autodiscover.areeba.com.gnwimax-dns5.areeba.com.gn
wimax-dns0.areeba.com.gnequitygroup.africadnsmaster.hotmtnpromo.com
cnpr0-loc.lonestarcell.commtn-eschool.commtn-global.netexchange.mtn-ns.net
```
