#!/bin/zsh
ips=$(ifconfig|grep 'inet'|grep -v '127.0.0.1'|cut -d: -f2|awk '{ print $2}'|sed '/^$/d'|grep '192')
curl --location --request POST 'http://bot.alomerry.com:4376/report' \
--header 'Content-Type: application/json' \
--data-raw '{"workIP":"'$ips'"}'

# TODO 删除上一条消息/ip 未变化不发送
# crontab -e
# 0 */1 * * * /home/user/workspace/ip_reporter.sh 每小时上报一次