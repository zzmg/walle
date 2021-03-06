FROM ccr.ccs.tencentyun.com/dhub.wallstcn.com/alpine:3.5
ENV CONFIG_ENV prod
ADD cron /
RUN cp /cron /etc/crontabs/root
RUN touch /var/log/cron.log
ADD run.sh /
RUN chmod +x /run.sh
ADD server /
ADD conf/ /conf
CMD ["bash","/run.sh"]