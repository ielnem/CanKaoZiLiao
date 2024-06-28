# Linux nfs客户端对于同时发起的NFS请求数量进行了控制，若该参数配置较小会导致IO性能较差。

```
cat /proc/sys/sunrpc/tcp_slot_table_entries
```

默认编译的内核该参数最大值为256，可以设置为128来提高性能

```
echo "options sunrpc tcp_slot_table_entries=128" >> /etc/modprobe.d/sunrpc.conf
echo "options sunrpc tcp_max_slot_table_entries=128" >>  /etc/modprobe.d/sunrpc.conf
sysctl -w sunrpc.tcp_slot_table_entries=128
```

修改完成后，需要重新挂载文件系统或重启机器。
