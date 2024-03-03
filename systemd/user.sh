# 创建通用步骤
# alertmanager 、 victoriametrics、vmalert、webhook

useradd alertmanager
groupadd alertmanager
usermod -a -G  alertmanager alertmanager
mv ./alertmanager /usr/local/bin/
chown alertmanager:alertmanager /usr/local/bin/alertmanager

# Create systemd file

systemctl daemon-reload
systemctl start alertmanager.service
systemctl status alertmanager.servicex


# other
# vmalert
chmod -R 644 /etc/ruler/
chmod -R 644 /etc/alertmanager/