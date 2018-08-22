FROM ccr.ccs.tencentyun.com/dhub.wallstcn.com/alpine:3.5
ENV CONFIGOR_ENV ivktest
ADD cron /
RUN cp /cron /etc/crontabs/root
RUN touch /var/log/cron.log
ADD run.sh /
RUN chmod +x /run.sh
ADD server /
ADD conf/ /conf
CMD ["bash","/run.sh"]