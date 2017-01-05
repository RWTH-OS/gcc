#
# Regular cron jobs for the gcc package
#
0 4	* * *	root	[ -x /usr/bin/gcc_maintenance ] && /usr/bin/gcc_maintenance
